package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ---------------------------------------------------------------------------
// Opcodes — matches the wasm-vm.oak opcode set for consistency
// ---------------------------------------------------------------------------

type opcode byte

const (
	opHalt       opcode = 0
	opNop        opcode = 1
	opConstNull  opcode = 2
	opConstEmpty opcode = 3
	opConstTrue  opcode = 4
	opConstFalse opcode = 5
	opConstInt   opcode = 6  // i32 operand (little-endian)
	opConstFloat opcode = 7  // u16 constant pool index
	opConstStr   opcode = 8  // u16 constant pool index
	opConstAtom  opcode = 9  // u16 constant pool index
	opPop        opcode = 10 // discard TOS
	opDup        opcode = 11 // duplicate TOS
	opLoadLocal  opcode = 12 // u16 slot
	opStoreLocal opcode = 13 // u16 slot
	opLoadUpval  opcode = 14 // u16 name constant index (name-based scope chain lookup)
	opStoreUpval opcode = 15 // u16 name constant index (name-based scope chain lookup)
	opAdd        opcode = 16
	opSub        opcode = 17
	opMul        opcode = 18
	opDiv        opcode = 19
	opMod        opcode = 20
	opPow        opcode = 21
	opNeg        opcode = 22
	opBAnd       opcode = 23
	opBOr        opcode = 24
	opBXor       opcode = 25
	opBRShift    opcode = 26
	opEq         opcode = 27
	opNeq        opcode = 28
	opGt         opcode = 29
	opLt         opcode = 30
	opGeq        opcode = 31
	opLeq        opcode = 32
	opNot        opcode = 33
	opConcat     opcode = 34 // << (push/append)
	opMakeList   opcode = 35 // u16 count
	opMakeObj    opcode = 36 // u16 pair count
	opGetProp    opcode = 37 // pop key, pop obj, push result
	opSetProp    opcode = 38 // pop val, pop key, pop obj → set, push val
	opJump       opcode = 39 // i32 absolute offset
	opJumpFalse  opcode = 40 // i32 absolute offset (pops condition)
	opClosure    opcode = 41 // u16 function template index
	opCall       opcode = 42 // u8 arity
	opReturn     opcode = 43
	opTailCall   opcode = 44 // u8 arity
	opBuiltin    opcode = 45 // u16 builtin index, u8 arity
	opImport     opcode = 46 // u16 constant pool index (module name)
	opImportDyn  opcode = 47 // pop string from stack, import that module
	opDeepEq     opcode = 48
	opSwap       opcode = 49
	opMatchJump  opcode = 50 // pop & compare TOS, jump i32 if no match
	opScopePush  opcode = 51
	opScopePop   opcode = 52
	opCallSpread opcode = 53 // u8 arity — last arg is a list to spread

	// Lexical-coordinate upvalue ops (M1 TODO-09, §7.4). Operand: u16 depth
	// (1 = immediate lexical parent function, counting up to the chunk
	// top-level), u16 slot (the ancestor's local slot), u16 name constant
	// index kept solely for the defensive name-based fallback.
	opLoadUpvalFast  opcode = 54
	opStoreUpvalFast opcode = 55

	// In-frame higher-order-function loops (M1 TODO-12, §7.8). Operand:
	// u16 function-template index of the literal fnNode callback. The VM
	// runs the callback as a plain frame between iterations (no child VM,
	// no closure allocation); opReturn drives the iteration state machine.
	opHOFMap    opcode = 56 // pop subject; push list/str of f(elem, i)
	opHOFFilter opcode = 57 // pop subject; push elements where f(elem, i) == true
	opHOFEach   opcode = 58 // pop subject; call f(elem, i); push null
	opHOFReduce opcode = 59 // pop seed, pop subject; push f(acc, elem, i) fold

	// Superinstruction (M1 TODO-10, §7.5): fuses the ubiquitous comparison-
	// match prologue `LOAD_LOCAL s; CONST_INT imm; LT; DUP; CONST_TRUE; EQ;
	// JUMP_FALSE; POP(+POP)` (9 dispatches) into one. Operand: u16 local
	// slot, i32 immediate, i32 else-target. Emitted only inside function
	// bodies (top-level phantom-local semantics stay on the generic path).
	// Falls through when locals[slot] < imm (Eq-true), else jumps after
	// consuming the condition; the else path pushes null for no-match
	// fall-through, matching the tree-walker (fixes the cond-leak bug).
	opLtLocalJumpTrue opcode = 60

	// Inline-cached property access (M1 TODO-11, §7.7). Operand: u16
	// constant-pool index of the interned property name. Avoids pushing a
	// key *StringValue per access (zero alloc) and skips the key pop.
	opGetPropConst opcode = 61 // pop obj; push obj.<name>
	opSetPropConst opcode = 62 // pop val, pop obj; obj.<name> = val; push val
)

var opcodeNames = [...]string{
	"HALT", "NOP", "CONST_NULL", "CONST_EMPTY",
	"CONST_TRUE", "CONST_FALSE", "CONST_INT", "CONST_FLOAT",
	"CONST_STRING", "CONST_ATOM", "POP", "DUP",
	"LOAD_LOCAL", "STORE_LOCAL", "LOAD_UPVAL", "STORE_UPVAL",
	"ADD", "SUB", "MUL", "DIV", "MOD", "POW", "NEG",
	"BAND", "BOR", "BXOR", "BRSHIFT",
	"EQ", "NEQ", "GT", "LT", "GEQ", "LEQ",
	"NOT", "CONCAT", "MAKE_LIST", "MAKE_OBJECT",
	"GET_PROP", "SET_PROP",
	"JUMP", "JUMP_FALSE",
	"CLOSURE", "CALL", "RETURN", "TAIL_CALL",
	"BUILTIN", "IMPORT", "IMPORT_DYN", "DEEP_EQ",
	"SWAP", "MATCH_JUMP",
	"SCOPE_PUSH", "SCOPE_POP",
	"CALL_SPREAD",
	"LOAD_UPVAL_FAST", "STORE_UPVAL_FAST",
	"HOF_MAP", "HOF_FILTER", "HOF_EACH", "HOF_REDUCE",
	"LT_LOCAL_JUMP_TRUE",
	"GET_PROP_CONST", "SET_PROP_CONST",
}

// ---------------------------------------------------------------------------
// Constant pool entry
// ---------------------------------------------------------------------------

type constKind byte

const (
	constString constKind = 0
	constAtom   constKind = 1
	constFloat  constKind = 2
)

type constEntry struct {
	kind constKind
	str  string  // for constString and constAtom
	f    float64 // for constFloat
}

// ---------------------------------------------------------------------------
// Function template (one per fn definition in source)
// ---------------------------------------------------------------------------

type funcTemplate struct {
	offset     int    // bytecode offset where body starts
	arity      int    // number of named parameters
	localCount int    // total local slots needed
	name       string // function name (may be empty)
	hasRestArg bool
	localNames []string // name for each local slot (for scope chain construction)
	defn       *fnNode  // preserved AST for interpreter() engine switching
	// needsHeapScope is true only when this function's body contains an inner
	// fnNode literal (at any depth). Inner closures capture the enclosing
	// scope chain, so frames of such functions must allocate a heap vmScope;
	// all other frames run on the VM's contiguous locals slab with zero
	// per-call allocations (M1 TODO-08, §7.2).
	needsHeapScope bool
}

// ---------------------------------------------------------------------------
// Bytecode chunk — output of compilation
// ---------------------------------------------------------------------------

// sourceMapEntry maps a bytecode offset to a source position.
type sourceMapEntry struct {
	offset int
	pos    pos
}

type bytecodeChunk struct {
	code          []byte
	constants     []constEntry
	functions     []funcTemplate
	topLevelNames []string // local slot names for top-level code
	sourceMap     []sourceMapEntry
}

// ---------------------------------------------------------------------------
// .mgb binary format loader
// ---------------------------------------------------------------------------

const mgbMagic = "MGbc"
const mgbVersion = 2

// loadBytecodeChunk deserializes a .mgb binary file into a bytecodeChunk.
// Format: "MGbc" <version:u16LE> <bcSection> <cpSection> <ftSection>
// Each section: <length:u32LE> <data:N bytes>
func loadBytecodeChunk(data []byte) (*bytecodeChunk, error) {
	if len(data) < 6 {
		return nil, fmt.Errorf("mgb: file too short")
	}
	if string(data[0:4]) != mgbMagic {
		return nil, fmt.Errorf("mgb: invalid magic (expected %q)", mgbMagic)
	}
	version := uint16(data[4]) | uint16(data[5])<<8
	if version != mgbVersion {
		return nil, fmt.Errorf("mgb: unsupported version %d", version)
	}

	off := 6

	// --- Bytecode section ---
	if off+4 > len(data) {
		return nil, fmt.Errorf("mgb: truncated bytecode section header")
	}
	bcLen := int(u32le(data, off))
	off += 4
	if off+bcLen > len(data) {
		return nil, fmt.Errorf("mgb: truncated bytecode section")
	}
	code := make([]byte, bcLen)
	copy(code, data[off:off+bcLen])
	off += bcLen

	// --- Constant pool section ---
	if off+4 > len(data) {
		return nil, fmt.Errorf("mgb: truncated constant pool section header")
	}
	cpLen := int(u32le(data, off))
	off += 4
	cpEnd := off + cpLen
	if cpEnd > len(data) {
		return nil, fmt.Errorf("mgb: truncated constant pool section")
	}
	constants, err := deserializeConstantPool(data[off:cpEnd])
	if err != nil {
		return nil, err
	}
	off = cpEnd

	// --- Function table section ---
	if off+4 > len(data) {
		return nil, fmt.Errorf("mgb: truncated function table section header")
	}
	ftLen := int(u32le(data, off))
	off += 4
	ftEnd := off + ftLen
	if ftEnd > len(data) {
		return nil, fmt.Errorf("mgb: truncated function table section")
	}
	functions, err := deserializeFunctionTable(data[off:ftEnd])
	if err != nil {
		return nil, err
	}
	off = ftEnd

	// --- Top-level names section ---
	var topLevelNames []string
	if off+4 <= len(data) {
		tlLen := int(u32le(data, off))
		off += 4
		tlEnd := off + tlLen
		if tlEnd > len(data) {
			return nil, fmt.Errorf("mgb: truncated top-level names section")
		}
		topLevelNames, err = deserializeNameList(data[off:tlEnd])
		if err != nil {
			return nil, err
		}
	}

	return &bytecodeChunk{
		code:          code,
		constants:     constants,
		functions:     functions,
		topLevelNames: topLevelNames,
	}, nil
}

func u32le(data []byte, off int) uint32 {
	return uint32(data[off]) | uint32(data[off+1])<<8 | uint32(data[off+2])<<16 | uint32(data[off+3])<<24
}

func deserializeConstantPool(data []byte) ([]constEntry, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("mgb: constant pool too short")
	}
	count := int(u32le(data, 0))
	off := 4
	entries := make([]constEntry, 0, count)

	for i := 0; i < count; i++ {
		if off >= len(data) {
			return nil, fmt.Errorf("mgb: constant pool truncated at entry %d", i)
		}
		kind := constKind(data[off])
		off++

		if off+4 > len(data) {
			return nil, fmt.Errorf("mgb: constant pool truncated at entry %d length", i)
		}
		slen := int(u32le(data, off))
		off += 4

		if off+slen > len(data) {
			return nil, fmt.Errorf("mgb: constant pool truncated at entry %d data", i)
		}
		raw := string(data[off : off+slen])
		off += slen

		switch kind {
		case constString, constAtom:
			entries = append(entries, constEntry{kind: kind, str: raw})
		case constFloat:
			f, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				return nil, fmt.Errorf("mgb: bad float constant %q: %w", raw, err)
			}
			entries = append(entries, constEntry{kind: kind, f: f})
		default:
			return nil, fmt.Errorf("mgb: unknown constant kind %d", kind)
		}
	}
	return entries, nil
}

func deserializeFunctionTable(data []byte) ([]funcTemplate, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("mgb: function table too short")
	}
	count := int(u32le(data, 0))
	off := 4
	fns := make([]funcTemplate, 0, count)

	for i := 0; i < count; i++ {
		if off+9 > len(data) {
			return nil, fmt.Errorf("mgb: function table truncated at entry %d", i)
		}
		offset := int(u32le(data, off))
		off += 4
		arity := int(uint16(data[off]) | uint16(data[off+1])<<8)
		off += 2
		localCount := int(uint16(data[off]) | uint16(data[off+1])<<8)
		off += 2
		hasRestArg := data[off] != 0
		off++

		// Read local names
		var localNames []string
		if off+2 <= len(data) {
			nameCount := int(uint16(data[off]) | uint16(data[off+1])<<8)
			off += 2
			localNames = make([]string, 0, nameCount)
			for j := 0; j < nameCount; j++ {
				if off+2 > len(data) {
					return nil, fmt.Errorf("mgb: function %d local name %d truncated", i, j)
				}
				nl := int(uint16(data[off]) | uint16(data[off+1])<<8)
				off += 2
				if off+nl > len(data) {
					return nil, fmt.Errorf("mgb: function %d local name %d data truncated", i, j)
				}
				localNames = append(localNames, string(data[off:off+nl]))
				off += nl
			}
		}

		fns = append(fns, funcTemplate{
			offset:     offset,
			arity:      arity,
			localCount: localCount,
			hasRestArg: hasRestArg,
			localNames: localNames,
		})
	}
	return fns, nil
}

func deserializeNameList(data []byte) ([]string, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("mgb: name list too short")
	}
	count := int(u32le(data, 0))
	off := 4
	names := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if off+2 > len(data) {
			return nil, fmt.Errorf("mgb: name list truncated at entry %d", i)
		}
		nl := int(uint16(data[off]) | uint16(data[off+1])<<8)
		off += 2
		if off+nl > len(data) {
			return nil, fmt.Errorf("mgb: name list data truncated at entry %d", i)
		}
		names = append(names, string(data[off:off+nl]))
		off += nl
	}
	return names, nil
}

// ---------------------------------------------------------------------------
// VM scope chain (for upvalue resolution)
// ---------------------------------------------------------------------------

// vmScope maps local variable names to values via shared slices.
// The values slice is shared with the callFrame's locals array, so
// STORE_LOCAL updates are visible through the scope chain.
type vmScope struct {
	names  []string // name at each slot index
	values []Value  // value at each slot index (shared with frame locals)
	parent *vmScope
}

// get walks the scope chain to find a variable by name.
func (s *vmScope) get(name string) (Value, bool) {
	for scope := s; scope != nil; scope = scope.parent {
		for i, n := range scope.names {
			if n == name && i < len(scope.values) {
				return scope.values[i], true
			}
		}
	}
	return nil, false
}

// set walks the scope chain to update a variable by name.
func (s *vmScope) set(name string, val Value) bool {
	for scope := s; scope != nil; scope = scope.parent {
		for i, n := range scope.names {
			if n == name && i < len(scope.values) {
				scope.values[i] = val
				return true
			}
		}
	}
	return false
}

// ---------------------------------------------------------------------------
// Compiler
// ---------------------------------------------------------------------------

type compiler struct {
	code         []byte
	constants    []constEntry
	functions    []funcTemplate
	locals       []string   // current scope local names
	scopeStack   [][]string // saved locals for outer scopes
	parentLocals [][]string // stack of parent scope local name lists (for upvalue resolution)
	sourceMap    []sourceMapEntry
	fnScopeFlags []bool // per-function: body contains an inner fnNode (needsHeapScope)
}

func newCompiler() *compiler {
	return &compiler{
		code:      make([]byte, 0, 256),
		constants: make([]constEntry, 0, 16),
		functions: make([]funcTemplate, 0, 8),
		locals:    make([]string, 0, 8),
		sourceMap: make([]sourceMapEntry, 0, 32),
	}
}

// notePos records the source position for the current bytecode offset.
func (co *compiler) notePos(p pos) {
	co.sourceMap = append(co.sourceMap, sourceMapEntry{offset: co.offset(), pos: p})
}

// --- Emission helpers ---

func (co *compiler) emit(b byte) {
	co.code = append(co.code, b)
}

func (co *compiler) emitOp(op opcode) {
	co.code = append(co.code, byte(op))
}

func (co *compiler) emitU16(v int) {
	co.code = append(co.code, byte(v&0xFF), byte((v>>8)&0xFF))
}

func (co *compiler) emitI32(v int) {
	co.code = append(co.code,
		byte(v&0xFF),
		byte((v>>8)&0xFF),
		byte((v>>16)&0xFF),
		byte((v>>24)&0xFF),
	)
}

func (co *compiler) patchI32(offset int, v int) {
	co.code[offset] = byte(v & 0xFF)
	co.code[offset+1] = byte((v >> 8) & 0xFF)
	co.code[offset+2] = byte((v >> 16) & 0xFF)
	co.code[offset+3] = byte((v >> 24) & 0xFF)
}

func (co *compiler) offset() int {
	return len(co.code)
}

// --- Constant pool ---

func (co *compiler) addString(s string) int {
	for i, c := range co.constants {
		if c.kind == constString && c.str == s {
			return i
		}
	}
	idx := len(co.constants)
	co.constants = append(co.constants, constEntry{kind: constString, str: s})
	return idx
}

func (co *compiler) addAtom(name string) int {
	for i, c := range co.constants {
		if c.kind == constAtom && c.str == name {
			return i
		}
	}
	idx := len(co.constants)
	co.constants = append(co.constants, constEntry{kind: constAtom, str: name})
	return idx
}

func (co *compiler) addFloat(v float64) int {
	for i, c := range co.constants {
		if c.kind == constFloat && c.f == v {
			return i
		}
	}
	idx := len(co.constants)
	co.constants = append(co.constants, constEntry{kind: constFloat, f: v})
	return idx
}

// --- Local variable tracking ---

func (co *compiler) resolveLocal(name string) int {
	for i := len(co.locals) - 1; i >= 0; i-- {
		if co.locals[i] == name {
			return i
		}
	}
	return -1
}

func (co *compiler) declareLocal(name string) int {
	idx := co.resolveLocal(name)
	if idx >= 0 {
		return idx
	}
	idx = len(co.locals)
	co.locals = append(co.locals, name)
	return idx
}

// pushScope saves the current locals for restoration via popScope.
// Currently unused but retained for future nested-scope compilation.
var _ = (*compiler).pushScope //nolint:unused

func (co *compiler) pushScope() {
	saved := make([]string, len(co.locals))
	copy(saved, co.locals)
	co.scopeStack = append(co.scopeStack, saved)
}

// popScope restores locals saved by pushScope.
var _ = (*compiler).popScope //nolint:unused

func (co *compiler) popScope() {
	n := len(co.scopeStack)
	if n > 0 {
		co.locals = co.scopeStack[n-1]
		co.scopeStack = co.scopeStack[:n-1]
	}
}

// resolveUpvalueCoord locates name in an ancestor function's local list and
// returns its lexical coordinates: depth 1 = the immediate lexical parent
// function, depth len(parentLocals) = the chunk top-level; slot is the
// ancestor's local slot index. Runtime resolution walks the persistent heap
// scope chain (every function on a closure's lexical chain is heap-scoped by
// TODO-08's needsHeapScope induction), giving O(depth) direct indexing with
// no name hashing (M1 TODO-09, §7.4).
func (co *compiler) resolveUpvalueCoord(name string) (depth int, slot int, ok bool) {
	for i := len(co.parentLocals) - 1; i >= 0; i-- {
		for j, n := range co.parentLocals[i] {
			if n == name {
				return len(co.parentLocals) - i, j, true
			}
		}
	}
	return 0, 0, false
}

// resolveUpvalue checks if name exists in any parent scope.
func (co *compiler) resolveUpvalue(name string) bool {
	for i := len(co.parentLocals) - 1; i >= 0; i-- {
		for _, n := range co.parentLocals[i] {
			if n == name {
				return true
			}
		}
	}
	return false
}

// ---------------------------------------------------------------------------
// AST → Bytecode compilation
// ---------------------------------------------------------------------------

func (co *compiler) compileNode(node astNode) {
	co.compileNodeTail(node, false)
}

func (co *compiler) compileNodeTail(node astNode, isTail bool) {
	if node == nil {
		co.emitOp(opConstNull)
		return
	}

	// Record source position for this node's bytecode
	co.notePos(node.pos())

	switch n := node.(type) {
	case nullNode:
		co.emitOp(opConstNull)

	case emptyNode:
		co.emitOp(opConstEmpty)

	case boolNode:
		if n.payload {
			co.emitOp(opConstTrue)
		} else {
			co.emitOp(opConstFalse)
		}

	case intNode:
		co.emitOp(opConstInt)
		co.emitI32(int(n.payload))

	case floatNode:
		idx := co.addFloat(n.payload)
		co.emitOp(opConstFloat)
		co.emitU16(idx)

	case *stringNode:
		idx := co.addString(string(n.payload))
		co.emitOp(opConstStr)
		co.emitU16(idx)

	case stringNode:
		idx := co.addString(string(n.payload))
		co.emitOp(opConstStr)
		co.emitU16(idx)

	case atomNode:
		idx := co.addAtom(n.payload)
		co.emitOp(opConstAtom)
		co.emitU16(idx)

	case identifierNode:
		// First check current locals
		localIdx := co.resolveLocal(n.payload)
		if localIdx >= 0 {
			co.emitOp(opLoadLocal)
			co.emitU16(localIdx)
		} else if depth, slot, ok := co.resolveUpvalueCoord(n.payload); ok {
			// Found in a parent scope with known coordinates — emit direct
			// lexical-coordinate lookup (name kept for the defensive fallback).
			nameIdx := co.addString(n.payload)
			co.emitOp(opLoadUpvalFast)
			co.emitU16(depth)
			co.emitU16(slot)
			co.emitU16(nameIdx)
		} else if co.resolveUpvalue(n.payload) {
			// Found in a parent scope — emit name-based upvalue lookup
			nameIdx := co.addString(n.payload)
			co.emitOp(opLoadUpval)
			co.emitU16(nameIdx)
		} else if len(co.parentLocals) > 0 {
			// Inside a function body: unknown identifier is an outer variable
			// (e.g., forward reference to a sibling function defined later).
			// Use runtime scope chain lookup.
			nameIdx := co.addString(n.payload)
			co.emitOp(opLoadUpval)
			co.emitU16(nameIdx)
		} else {
			// Top-level: forward reference to a local declared later
			localIdx = co.declareLocal(n.payload)
			co.emitOp(opLoadLocal)
			co.emitU16(localIdx)
		}

	case unaryNode:
		co.compileNode(n.right)
		switch n.op {
		case minus:
			co.emitOp(opNeg)
		case exclam:
			co.emitOp(opNot)
		case tilde:
			co.emitOp(opConstInt)
			co.emitI32(-1)
			co.emitOp(opBXor)
		}

	case binaryNode:
		// NOTE: the former `n.op == pipeArrow` branch was dead code and has
		// been removed (M1 TODO-05, §7.6). The parser desugars `x |> f(a, b)`
		// into fnCallNode{fn: f, args: [x, a, b]} at parse time (parse.go),
		// so a binaryNode carrying pipeArrow can never reach the compiler
		// from parsed Oak programs.
		co.compileNode(n.left)
		co.compileNode(n.right)
		switch n.op {
		case plus:
			co.emitOp(opAdd)
		case minus:
			co.emitOp(opSub)
		case times:
			co.emitOp(opMul)
		case divide:
			co.emitOp(opDiv)
		case modulus:
			co.emitOp(opMod)
		case power:
			co.emitOp(opPow)
		case and:
			co.emitOp(opBAnd)
		case or:
			co.emitOp(opBOr)
		case xor:
			co.emitOp(opBXor)
		case rshift:
			co.emitOp(opBRShift)
		case eq:
			co.emitOp(opEq)
		case deepEq:
			co.emitOp(opDeepEq)
		case neq:
			co.emitOp(opNeq)
		case greater:
			co.emitOp(opGt)
		case less:
			co.emitOp(opLt)
		case geq:
			co.emitOp(opGeq)
		case leq:
			co.emitOp(opLeq)
		case pushArrow:
			co.emitOp(opConcat)
		}

	case assignmentNode:
		co.compileAssignment(n)

	case propertyAccessNode:
		co.compileNode(n.left)
		if ident, ok := n.right.(identifierNode); ok {
			// Inline cache (M1 TODO-11): constant key fused into the op.
			idx := co.addString(ident.payload)
			co.emitOp(opGetPropConst)
			co.emitU16(idx)
		} else {
			co.compileNode(n.right)
			co.emitOp(opGetProp)
		}

	case listNode:
		for _, elem := range n.elems {
			co.compileNode(elem)
		}
		co.emitOp(opMakeList)
		co.emitU16(len(n.elems))

	case objectNode:
		for _, entry := range n.entries {
			// Object keys: identifiers are string keys, not variable references
			if ident, ok := entry.key.(identifierNode); ok {
				idx := co.addString(ident.payload)
				co.emitOp(opConstStr)
				co.emitU16(idx)
			} else {
				co.compileNode(entry.key)
			}
			co.compileNode(entry.val)
		}
		co.emitOp(opMakeObj)
		co.emitU16(len(n.entries))

	case blockNode:
		if len(n.exprs) == 0 {
			co.emitOp(opConstNull)
		} else {
			for i, expr := range n.exprs {
				isLast := i == len(n.exprs)-1
				if isLast {
					co.compileNodeTail(expr, isTail)
				} else {
					co.compileNode(expr)
					co.emitOp(opPop)
				}
			}
		}

	case ifExprNode:
		co.compileIfTail(n, isTail)

	case fnNode:
		// An inner fnNode captures the enclosing scope chain: mark the
		// immediately-enclosing function as needing a heap scope (M1 TODO-08).
		if len(co.fnScopeFlags) > 0 {
			co.fnScopeFlags[len(co.fnScopeFlags)-1] = true
		}
		co.compileFunction(&n)
	case *fnNode:
		if len(co.fnScopeFlags) > 0 {
			co.fnScopeFlags[len(co.fnScopeFlags)-1] = true
		}
		co.compileFunction(n)

	case classNode:
		co.compileClass(n)

	case fnCallNode:
		co.compileFnCallTail(n, isTail)

	default:
		co.emitOp(opConstNull)
	}
}

// compileAssignment handles := and <- for identifiers, property access, and destructuring.
func (co *compiler) compileAssignment(n assignmentNode) {
	switch left := n.left.(type) {
	case identifierNode:
		co.compileNode(n.right)
		// Check if the target is a local or an upvalue
		localIdx := co.resolveLocal(left.payload)
		if localIdx >= 0 {
			co.emitOp(opDup)
			co.emitOp(opStoreLocal)
			co.emitU16(localIdx)
		} else if depth, slot, ok := co.resolveUpvalueCoord(left.payload); ok {
			// Assignment to an outer variable with known coordinates
			nameIdx := co.addString(left.payload)
			co.emitOp(opDup)
			co.emitOp(opStoreUpvalFast)
			co.emitU16(depth)
			co.emitU16(slot)
			co.emitU16(nameIdx)
		} else if co.resolveUpvalue(left.payload) {
			// Assignment to an outer variable
			nameIdx := co.addString(left.payload)
			co.emitOp(opDup)
			co.emitOp(opStoreUpval)
			co.emitU16(nameIdx)
		} else {
			// New local variable
			localIdx = co.declareLocal(left.payload)
			co.emitOp(opDup)
			co.emitOp(opStoreLocal)
			co.emitU16(localIdx)
		}

	case propertyAccessNode:
		co.compileNode(left.left)
		if ident, ok := left.right.(identifierNode); ok {
			// Inline cache (M1 TODO-11): constant key fused into the op.
			idx := co.addString(ident.payload)
			co.compileNode(n.right)
			co.emitOp(opSetPropConst)
			co.emitU16(idx)
		} else {
			co.compileNode(left.right)
			co.compileNode(n.right)
			co.emitOp(opSetProp)
		}

	case listNode:
		co.compileNode(n.right)
		for i, elem := range left.elems {
			if _, isEmpty := elem.(emptyNode); isEmpty {
				continue
			}
			if ident, ok := elem.(identifierNode); ok {
				co.emitOp(opDup)
				co.emitOp(opConstInt)
				co.emitI32(i)
				co.emitOp(opGetProp)
				idx := co.declareLocal(ident.payload)
				co.emitOp(opStoreLocal)
				co.emitU16(idx)
			}
		}

	case objectNode:
		co.compileNode(n.right)
		for _, entry := range left.entries {
			co.emitOp(opDup)
			// Keys in destructuring patterns are property names, not variable refs
			if ident, ok := entry.key.(identifierNode); ok {
				idx := co.addString(ident.payload)
				co.emitOp(opGetPropConst)
				co.emitU16(idx)
			} else {
				co.compileNode(entry.key)
				co.emitOp(opGetProp)
			}
			if ident, ok := entry.val.(identifierNode); ok {
				idx := co.declareLocal(ident.payload)
				co.emitOp(opStoreLocal)
				co.emitU16(idx)
			} else {
				co.emitOp(opPop)
			}
		}

	default:
		co.compileNode(n.right)
	}
}

// compileIf compiles an if-expression (pattern matching).
// Convenience wrapper; callers use compileIfTail directly.
var _ = (*compiler).compileIf //nolint:unused

func (co *compiler) compileIf(n ifExprNode) {
	co.compileIfTail(n, false)
}

func (co *compiler) compileIfTail(n ifExprNode, isTail bool) {
	// Superinstruction fast path (M1 TODO-10, §7.5): `if local < imm` with a
	// literal-true arm (+ optional wildcard) is the hottest pattern in the
	// corpus (fib, loop counters, AVM1) — 9 dispatches fuse into 1.
	if co.parentLocals != nil {
		if sl, imm, ok := co.matchLtLocalImmTrue(n); ok {
			co.compileLtLocalImmTrue(n, sl, imm, isTail)
			return
		}
	}
	co.compileNode(n.cond)
	var endJumps []int

	// M1 TODO-05 (§7.6): group consecutive branches that share an identical
	// body AST. The parser's multi-target sugar `a, b -> body` splits one arm
	// into multiple ifBranch entries referencing the SAME body node, so the
	// naive path compiled the body once per target. Here every matching
	// target jumps directly to a single compiled copy of the body, shrinking
	// emitted bytecode for idioms like `'Z', 'B', 'C', 'S', 'I' -> 0`.
	//
	// Body identity is keyed by the body's rendered form: identical render
	// implies identical compiled code, so coalescing is always safe.
	type bodyGroup struct {
		targets     []astNode
		body        astNode
		bodyKey     string
		hasWildcard bool
	}
	groups := make([]*bodyGroup, 0, len(n.branches))
	for _, br := range n.branches {
		if _, isEmpty := br.target.(emptyNode); isEmpty {
			// Wildcard matches unconditionally and terminates its group.
			groups = append(groups, &bodyGroup{body: br.body, hasWildcard: true})
			continue
		}
		key := br.body.String()
		if len(groups) > 0 {
			last := groups[len(groups)-1]
			if !last.hasWildcard && last.bodyKey == key {
				last.targets = append(last.targets, br.target)
				continue
			}
		}
		groups = append(groups, &bodyGroup{
			targets: []astNode{br.target},
			body:    br.body,
			bodyKey: key,
		})
	}

	for _, g := range groups {
		var groupSkips []int
		if g.hasWildcard {
			// Default/wildcard branch: pop condition, execute body
			co.emitOp(opPop)
			co.compileNodeTail(g.body, isTail)
		} else {
			// For each target: test condition, and on match jump to the
			// single shared body emitted after all tests.
			var toBody []int
			for _, t := range g.targets {
				co.emitOp(opDup)
				co.compileNode(t)
				co.emitOp(opEq)
				co.emitOp(opJumpFalse)
				skipOffset := co.offset()
				co.emitI32(0)
				groupSkips = append(groupSkips, skipOffset)
				// Match: pop condition, jump to shared body
				co.emitOp(opPop)
				co.emitOp(opJump)
				toBody = append(toBody, co.offset())
				co.emitI32(0)
			}
			bodyStart := co.offset()
			for _, off := range toBody {
				co.patchI32(off, bodyStart)
			}
			co.compileNodeTail(g.body, isTail)
		}
		co.emitOp(opJump)
		endJumps = append(endJumps, co.offset())
		co.emitI32(0)
		// Non-matching targets skip past this group's shared body and the
		// group's end-jump; that position is only known now, after the body
		// has been emitted.
		afterGroup := co.offset()
		for _, s := range groupSkips {
			co.patchI32(s, afterGroup)
		}
	}

	// No-match fall-through parity (tree-walker returns null): the last
	// non-wildcard group's false path reaches endTarget with the condition
	// still on the stack — pop it and push null. With a trailing wildcard,
	// the wildcard always matches and consumes the condition itself.
	if len(groups) > 0 && !groups[len(groups)-1].hasWildcard {
		co.emitOp(opPop)
		co.emitOp(opConstNull)
	}
	endTarget := co.offset()
	for _, off := range endJumps {
		co.patchI32(off, endTarget)
	}
}

// matchLtLocalImmTrue reports whether the if-expression has the superinstr
// shape: cond = `local < imm-literal`, first arm targets literal true, and
// the final arm (when present) is a wildcard. Returns the local slot and
// immediate.
func (co *compiler) matchLtLocalImmTrue(n ifExprNode) (int, int64, bool) {
	bin, ok := n.cond.(binaryNode)
	if !ok || bin.op != less {
		return 0, 0, false
	}
	id, ok := bin.left.(identifierNode)
	if !ok {
		return 0, 0, false
	}
	slot := co.resolveLocal(id.payload)
	if slot < 0 {
		return 0, 0, false
	}
	lit, ok := bin.right.(intNode)
	if !ok {
		return 0, 0, false
	}
	if len(n.branches) == 0 {
		return 0, 0, false
	}
	if tb, ok := n.branches[0].target.(boolNode); !ok || !tb.payload {
		return 0, 0, false
	}
	if len(n.branches) == 1 {
		return slot, lit.payload, true
	}
	if len(n.branches) == 2 {
		if _, ok := n.branches[1].target.(emptyNode); ok {
			return slot, lit.payload, true
		}
	}
	return 0, 0, false
}

// compileLtLocalImmTrue emits the fused comparison-match superinstruction.
func (co *compiler) compileLtLocalImmTrue(n ifExprNode, slot int, imm int64, isTail bool) {
	co.emitOp(opLtLocalJumpTrue)
	co.emitU16(slot)
	co.emitI32(int(imm))
	elseSite := co.offset()
	co.emitI32(0)

	co.compileNodeTail(n.branches[0].body, isTail)
	if len(n.branches) == 2 {
		// Else target = start of the wildcard body (its value is the result).
		co.emitOp(opJump)
		endSite := co.offset()
		co.emitI32(0)
		co.patchI32(elseSite, co.offset())
		co.compileNodeTail(n.branches[1].body, isTail)
		co.patchI32(endSite, co.offset())
	} else {
		// No wildcard: else jumps to an emitted null (no-match parity with
		// the tree-walker); the match path skips it.
		co.emitOp(opJump)
		skipNull := co.offset()
		co.emitI32(0)
		co.patchI32(elseSite, co.offset())
		co.emitOp(opConstNull)
		co.patchI32(skipNull, co.offset())
	}
}

// compileFunction compiles a fn definition (registers the template, emits
// CLOSURE, and binds the name in the outer scope when named).
func (co *compiler) compileFunction(n *fnNode) {
	fnIdx := co.emitFunctionBody(n)

	// Emit CLOSURE in outer code
	co.emitOp(opClosure)
	co.emitU16(fnIdx)

	// If named, also bind as local in outer scope
	if n.name != "" {
		co.emitOp(opDup)
		co.emitOp(opStoreLocal)
		co.emitU16(co.resolveLocal(n.name))
	}
}

// hofCallbackFnIdx compiles a literal fnNode callback argument into a
// registered function template and returns its index, for opHOF* ops.
// Returns ok=false for non-fnNode arguments (fast path not applicable).
func (co *compiler) hofCallbackFnIdx(arg astNode) (int, bool) {
	switch cb := arg.(type) {
	case fnNode:
		return co.emitFunctionBody(&cb), true
	case *fnNode:
		return co.emitFunctionBody(cb), true
	}
	return 0, false
}

// emitFunctionBody compiles a fn definition body and registers its function
// template, returning the template index. Shared by inline closure literals
// (which then emit opClosure themselves) and HOF in-frame callbacks (whose
// templates are referenced directly by opHOF* ops — no closure allocation,
// M1 TODO-12 §7.8).
func (co *compiler) emitFunctionBody(n *fnNode) int {
	// NOTE: Do NOT capture fnIdx here — inner function compilations
	// will register their templates first, shifting the index.

	// Pre-declare the function name in the outer scope BEFORE compiling body.
	// This ensures the name is in the parent locals so the body can
	// reference it as an upvalue for self-recursion.
	if n.name != "" {
		co.declareLocal(n.name)
	}

	// Save outer locals (now includes the pre-declared function name)
	savedLocals := make([]string, len(co.locals))
	copy(savedLocals, co.locals)

	// Push outer locals onto parentLocals stack for upvalue resolution
	co.parentLocals = append(co.parentLocals, savedLocals)

	// Track whether this body contains inner fnNodes (heap-scope requirement)
	co.fnScopeFlags = append(co.fnScopeFlags, false)

	// Jump over inline function body
	co.emitOp(opJump)
	jumpOverOffset := co.offset()
	co.emitI32(0)

	fnBodyStart := co.offset()

	// Fresh locals for the function body
	co.locals = co.locals[:0]
	for _, arg := range n.args {
		if arg == "" {
			// Empty args (_) must each occupy their own slot even though
			// they share the same "" name. declareLocal would deduplicate
			// them, causing localCount < arity and an index panic.
			idx := len(co.locals)
			co.locals = append(co.locals, "")
			_ = idx
		} else {
			co.declareLocal(arg)
		}
	}
	if n.restArg != "" {
		co.declareLocal(n.restArg)
	}

	// Compile body (may register inner function templates!)
	// Mark body as tail position so recursive calls emit opTailCall
	co.compileNodeTail(n.body, true)
	co.emitOp(opReturn)

	fnBodyEnd := co.offset()
	localCount := len(co.locals)
	localNamesCopy := make([]string, localCount)
	copy(localNamesCopy, co.locals)

	// Pop parentLocals stack
	co.parentLocals = co.parentLocals[:len(co.parentLocals)-1]

	// Restore outer locals
	co.locals = savedLocals

	// Patch jump-over
	co.patchI32(jumpOverOffset, fnBodyEnd)

	// Register function template AFTER body compilation (correct index)
	needsHeapScope := co.fnScopeFlags[len(co.fnScopeFlags)-1]
	co.fnScopeFlags = co.fnScopeFlags[:len(co.fnScopeFlags)-1]
	fnIdx := len(co.functions)
	co.functions = append(co.functions, funcTemplate{
		offset:         fnBodyStart,
		arity:          len(n.args),
		localCount:     localCount,
		name:           n.name,
		hasRestArg:     n.restArg != "",
		localNames:     localNamesCopy,
		defn:           n,
		needsHeapScope: needsHeapScope,
	})
	return fnIdx
}

// compileClass compiles a class node (desugars to fn).
func (co *compiler) compileClass(n classNode) {
	fnNode := &fnNode{
		name:    n.name,
		args:    n.args,
		restArg: n.restArg,
		body:    n.body,
		tok:     n.tok,
	}
	co.compileFunction(fnNode)
}

// compileFnCall compiles a function call.
// Convenience wrapper; callers use compileFnCallTail directly.
var _ = (*compiler).compileFnCall //nolint:unused

func (co *compiler) compileFnCall(n fnCallNode) {
	co.compileFnCallTail(n, false)
}

func (co *compiler) compileFnCallTail(n fnCallNode, isTail bool) {
	// Check for import() calls
	if ident, ok := n.fn.(identifierNode); ok && ident.payload == "import" && len(n.args) == 1 && n.restArg == nil {
		if strArg, ok := n.args[0].(stringNode); ok {
			// import('literal') → opImport with constant index
			idx := co.addString(string(strArg.payload))
			co.emitOp(opImport)
			co.emitU16(idx)
			return
		}
		// import(expr) → evaluate expr, then opImportDyn
		co.compileNode(n.args[0])
		co.emitOp(opImportDyn)
		return
	}

	// In-frame HOF loops (M1 TODO-12, §7.8): std map/filter/each/reduce called
	// with a literal fnNode callback compile to dedicated ops that run the
	// callback as an in-frame loop — no closure allocation, no per-element
	// accumulator `<<` churn, no child VMs. The callee (identifier or
	// property access such as std.map) is evaluated first and left for the
	// op to pop, so both `map(...)` and `std.map(...)` shapes specialize;
	// shadowed identifiers keep the generic call.
	if n.restArg == nil {
		hofTarget := ""
		switch fnExpr := n.fn.(type) {
		case identifierNode:
			if co.resolveLocal(fnExpr.payload) < 0 && !co.resolveUpvalue(fnExpr.payload) {
				hofTarget = fnExpr.payload
			}
		case propertyAccessNode:
			if id, ok := fnExpr.right.(identifierNode); ok {
				if _, ok2 := fnExpr.left.(identifierNode); ok2 {
					hofTarget = id.payload
				}
			}
		}
		switch hofTarget {
		case "map", "filter", "each":
			if len(n.args) == 2 {
				if fnIdx, ok2 := co.hofCallbackFnIdx(n.args[1]); ok2 {
					// stack: [callee, subject]
					co.compileNode(n.fn)
					co.compileNode(n.args[0])
					switch hofTarget {
					case "map":
						co.emitOp(opHOFMap)
					case "filter":
						co.emitOp(opHOFFilter)
					default:
						co.emitOp(opHOFEach)
					}
					co.emitU16(fnIdx)
					return
				}
			}
		case "reduce":
			if len(n.args) == 3 {
				if fnIdx, ok2 := co.hofCallbackFnIdx(n.args[2]); ok2 {
					// stack: [callee, subject, seed]
					co.compileNode(n.fn)
					co.compileNode(n.args[0])
					co.compileNode(n.args[1])
					co.emitOp(opHOFReduce)
					co.emitU16(fnIdx)
					return
				}
			}
		}
	}

	// Check for built-in function calls
	if ident, ok := n.fn.(identifierNode); ok {
		if bi := resolveBuiltinIndex(ident.payload); bi >= 0 {
			for _, arg := range n.args {
				co.compileNode(arg)
			}
			if n.restArg != nil {
				co.compileNode(n.restArg)
			}
			co.emitOp(opBuiltin)
			co.emitU16(bi)
			arity := len(n.args)
			if n.restArg != nil {
				arity++
			}
			co.emit(byte(arity & 0xFF))
			return
		}
	}

	// Generic function call
	co.compileNode(n.fn)
	for _, arg := range n.args {
		co.compileNode(arg)
	}
	if n.restArg != nil {
		co.compileNode(n.restArg)
	}
	arity := len(n.args)
	if n.restArg != nil {
		arity++
		co.emitOp(opCallSpread)
	} else if isTail && len(co.parentLocals) > 0 {
		co.emitOp(opTailCall)
	} else {
		co.emitOp(opCall)
	}
	co.emit(byte(arity & 0xFF))
}

// ---------------------------------------------------------------------------
// Built-in function name → index resolution
// ---------------------------------------------------------------------------

var builtinNames = [...]string{
	"print", "len", "type", "string", "int", "float",
	"codepoint", "char", "keys", "values", "slice",
	"append", "wait", "exit",
}

func resolveBuiltinIndex(name string) int {
	for i, n := range builtinNames {
		if n == name {
			return i
		}
	}
	return -1
}

// ---------------------------------------------------------------------------
// Top-level compilation entry point
// ---------------------------------------------------------------------------

func compileToByteCode(nodes []astNode) *bytecodeChunk {
	co := newCompiler()
	for i, node := range nodes {
		co.compileNode(node)
		if i < len(nodes)-1 {
			co.emitOp(opPop)
		}
	}
	co.emitOp(opHalt)

	topNames := make([]string, len(co.locals))
	copy(topNames, co.locals)

	return &bytecodeChunk{
		code:          co.code,
		constants:     co.constants,
		functions:     co.functions,
		topLevelNames: topNames,
		sourceMap:     co.sourceMap,
	}
}

// ---------------------------------------------------------------------------
// Disassembler (for debugging)
// ---------------------------------------------------------------------------

// disassemble returns a human-readable dump of bytecode (for debugging).
var _ = disassemble //nolint:unused

func disassemble(chunk *bytecodeChunk) string {
	var sb strings.Builder
	bc := chunk.code
	i := 0
	for i < len(bc) {
		op := opcode(bc[i])
		name := "???"
		if int(op) < len(opcodeNames) {
			name = opcodeNames[op]
		}
		fmt.Fprintf(&sb, "%4d: %s", i, name)
		advance := 1
		switch op {
		case opConstInt:
			v := int(bc[i+1]) | int(bc[i+2])<<8 | int(bc[i+3])<<16 | int(bc[i+4])<<24
			fmt.Fprintf(&sb, " %d", v)
			advance = 5
		case opConstFloat, opConstStr, opConstAtom:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			if op == opConstStr && v < len(chunk.constants) {
				fmt.Fprintf(&sb, " #%d (%q)", v, chunk.constants[v].str)
			} else if op == opConstAtom && v < len(chunk.constants) {
				fmt.Fprintf(&sb, " #%d (:%s)", v, chunk.constants[v].str)
			} else if op == opConstFloat && v < len(chunk.constants) {
				fmt.Fprintf(&sb, " #%d (%g)", v, chunk.constants[v].f)
			} else {
				fmt.Fprintf(&sb, " #%d", v)
			}
			advance = 3
		case opLoadLocal, opStoreLocal:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			fmt.Fprintf(&sb, " @%d", v)
			advance = 3
		case opLoadUpval, opStoreUpval:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			if v < len(chunk.constants) {
				fmt.Fprintf(&sb, " '%s'", chunk.constants[v].str)
			} else {
				fmt.Fprintf(&sb, " #%d", v)
			}
			advance = 3
		case opHOFMap, opHOFFilter, opHOFEach, opHOFReduce:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			fmt.Fprintf(&sb, " fn#%d", v)
			advance = 3
		case opGetPropConst, opSetPropConst:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			if v < len(chunk.constants) {
				fmt.Fprintf(&sb, " '%s'", chunk.constants[v].str)
			} else {
				fmt.Fprintf(&sb, " #%d", v)
			}
			advance = 3
		case opLtLocalJumpTrue:
			s := int(bc[i+1]) | int(bc[i+2])<<8
			imm := int32(uint32(bc[i+3]) | uint32(bc[i+4])<<8 | uint32(bc[i+5])<<16 | uint32(bc[i+6])<<24)
			tgt := int32(uint32(bc[i+7]) | uint32(bc[i+8])<<8 | uint32(bc[i+9])<<16 | uint32(bc[i+10])<<24)
			fmt.Fprintf(&sb, " @%d <%d else->%d", s, imm, tgt)
			advance = 11
		case opLoadUpvalFast, opStoreUpvalFast:
			d := int(bc[i+1]) | int(bc[i+2])<<8
			s := int(bc[i+3]) | int(bc[i+4])<<8
			v := int(bc[i+5]) | int(bc[i+6])<<8
			if v < len(chunk.constants) {
				fmt.Fprintf(&sb, " d%d s%d '%s'", d, s, chunk.constants[v].str)
			} else {
				fmt.Fprintf(&sb, " d%d s%d #%d", d, s, v)
			}
			advance = 7
		case opJump, opJumpFalse, opMatchJump:
			v := int(bc[i+1]) | int(bc[i+2])<<8 | int(bc[i+3])<<16 | int(bc[i+4])<<24
			fmt.Fprintf(&sb, " ->%d", v)
			advance = 5
		case opClosure:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			fmt.Fprintf(&sb, " fn#%d", v)
			advance = 3
		case opCall, opTailCall, opCallSpread:
			fmt.Fprintf(&sb, " (%d args)", bc[i+1])
			advance = 2
		case opMakeList, opMakeObj:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			fmt.Fprintf(&sb, " [%d]", v)
			advance = 3
		case opBuiltin:
			bi := int(bc[i+1]) | int(bc[i+2])<<8
			ar := bc[i+3]
			bname := "?"
			if bi < len(builtinNames) {
				bname = builtinNames[bi]
			}
			fmt.Fprintf(&sb, " %s(%d)", bname, ar)
			advance = 4
		case opImport:
			v := int(bc[i+1]) | int(bc[i+2])<<8
			fmt.Fprintf(&sb, " #%d", v)
			advance = 3
		}
		sb.WriteByte('\n')
		i += advance
	}
	return sb.String()
}

// ===========================================================================
// Virtual Machine
// ===========================================================================

// closureVal represents a function closure on the VM stack.
type closureVal struct {
	fnIdx       int
	parentScope *vmScope                                  // captured scope (shared reference, not a copy)
	call        func(args []Value) (Value, *runtimeError) // set for interop with tree-walker
	defn        *fnNode                                   // preserved AST for interpreter() engine switching
}

// callFrame represents a single function call on the call stack.
type callFrame struct {
	returnPC int     // where to resume after RETURN
	baseSlot int     // base index on value stack for this frame's locals
	fnIdx    int     // function template index (-1 for top-level)
	locals   []Value // local variable slots
	scope    *vmScope
	baseLocal int // slab window base (localSP at call entry); -1 for heap frames
	// parentScope is the closure's captured scope for slab frames (whose own
	// scope is nil). Name-based upvalue reads/writes resolve against it.
	parentScope *vmScope
}

// VM is the bytecode virtual machine.
type VM struct {
	chunk      *bytecodeChunk
	pc         int
	stack      []Value
	sp         int // stack pointer (next free slot)
	frames     []callFrame
	ctx        *Context // for built-in function dispatch + imports
	initLocals []Value  // pre-populated local slot values (for bytecode() builtin)
	outerScope *vmScope // enclosing scope for upvalue resolution in bytecode() builtin
	// localsSlab is a contiguous pre-allocated register file shared by all
	// non-escaping frames of this VM. Each call windows into it; returns
	// restore localSP in O(1). Zero heap allocations per slab call.
	localsSlab []Value
	localSP    int
	// hofStack drives in-frame higher-order-function loops (opHOF*). Each
	// entry is one active HOF whose callback frame is currently executing.
	hofStack []hofState
}

// hofKind enumerates the in-frame HOF loop shapes.
type hofKind byte

const (
	hofMap hofKind = iota
	hofFilter
	hofEach
	hofReduce
)

// hofState is the iteration state of one active in-frame HOF loop. The
// callback runs as a plain VM frame; opReturn detects the callback frame by
// index (frameIdx == len(frames) after pop), accumulates, and either pushes
// the next iteration's frame or finishes the loop.
type hofState struct {
	kind      hofKind
	fnIdx     int
	sub       *ListValue   // list subject (element source)
	sstr      *StringValue // string subject (element source)
	acc       *ListValue   // list result accumulator (map/filter)
	sacc      *StringValue // string result accumulator (map/filter)
	accNull   bool         // map over string saw a non-string/int result → overall null
	seed      Value        // reduce accumulator (current value)
	frame     callFrame    // prepared callback frame, re-pushed per iteration
	frameIdx  int          // vm.frames index the callback frame occupies
	i, n      int          // iteration index and subject length
	returnPC  int          // pc right after the opHOF* instruction
	baseLocal int          // slab window base of the callback frame; -1 heap
}

// elemAt returns the subject element at index i (list element or 1-char
// string), matching std's xs.(i) index access.
func (st *hofState) elemAt(i int) Value {
	if st.sub != nil {
		return (*st.sub)[i]
	}
	return MakeSingleByteString((*st.sstr)[i])
}

// hofFillArgs writes the callback's arguments into the frame's locals window:
// (elem, i) for map/filter/each, (acc, elem, i) for reduce. Slots beyond the
// callback's arity stay null (normalizeCallArgs padding semantics).
func (vm *VM) hofFillArgs(st *hofState) {
	locals := st.frame.locals
	for j := range locals {
		locals[j] = null
	}
	idxVal := makeIntValue(int64(st.i))
	switch st.kind {
	case hofReduce:
		if len(locals) > 0 {
			locals[0] = st.seed
		}
		if len(locals) > 1 {
			locals[1] = st.elemAt(st.i)
		}
		if len(locals) > 2 {
			locals[2] = idxVal
		}
	default:
		if len(locals) > 0 {
			locals[0] = st.elemAt(st.i)
		}
		if len(locals) > 1 {
			locals[1] = idxVal
		}
	}
}

// hofRefresh prepares the callback frame for the next iteration. Self-tail
// recursion and same-shape reuse keep the original locals slice; a non-self
// tail call inside the callback swaps frame.locals to a foreign slice that
// may be captured, so heap frames get a fresh locals + scope (slab windows
// are always private — opReturn freed them before we get here).
func (vm *VM) hofRefresh(st *hofState) {
	if st.baseLocal >= 0 {
		vm.localSP = st.baseLocal + len(st.frame.locals) + vmSlabHeadroom
	} else if st.frame.fnIdx != st.fnIdx {
		locals := make([]Value, len(st.frame.locals))
		for j := range locals {
			locals[j] = null
		}
		st.frame.locals = locals
		if st.frame.scope != nil {
			st.frame.scope = &vmScope{names: st.frame.scope.names, values: locals, parent: st.frame.scope.parent}
		}
	}
	vm.hofFillArgs(st)
}

// hofAccumulate folds one callback return value into the loop state.
func (st *hofState) hofAccumulate(retVal Value) {
	switch st.kind {
	case hofMap:
		if st.acc != nil {
			*st.acc = append(*st.acc, retVal)
		} else {
			switch r := retVal.(type) {
			case *StringValue:
				*st.sacc = append(*st.sacc, *r...)
			case IntValue:
				*st.sacc = append(*st.sacc, byte(r))
			default:
				// acc << x on a string with a non-string/int result yields
				// null, which poisons the rest of the std loop.
				st.accNull = true
			}
		}
	case hofFilter:
		if bv, ok := retVal.(BoolValue); ok && bool(bv) {
			if st.acc != nil {
				*st.acc = append(*st.acc, st.elemAt(st.i))
			} else {
				*st.sacc = append(*st.sacc, (*st.sstr)[st.i])
			}
		}
	case hofReduce:
		st.seed = retVal
	}
}

// hofResult computes the loop's final return value, mirroring std semantics:
// map/filter return the accumulator, each returns null, reduce returns the
// final accumulator.
func (st *hofState) hofResult() Value {
	switch st.kind {
	case hofMap:
		if st.acc != nil {
			return st.acc
		}
		if st.accNull {
			return null
		}
		return st.sacc
	case hofFilter:
		if st.acc != nil {
			return st.acc
		}
		return st.sacc
	case hofEach:
		return null
	default:
		return st.seed
	}
}

// hofStart classifies the subject and either (a) starts the in-frame loop by
// pushing the first callback frame (returns started=true), (b) finishes
// immediately for empty subjects by pushing the result (started=false), or
// (c) falls back to the generic std-lib call for unsupported subjects
// (started=false). Mirrors lib/std.oak exactly (M1 TODO-12, §7.8).
func (vm *VM) hofStart(kind hofKind, fnIdx int, subject, seed, callee Value, topScope *vmScope) (bool, *runtimeError) {
	ft := &vm.chunk.functions[fnIdx]
	curScope := vm.currentScope(topScope)
	st := hofState{kind: kind, fnIdx: fnIdx, seed: seed, baseLocal: -1}
	finish := func(result Value) (bool, *runtimeError) {
		vm.push(result)
		return false, nil
	}
	switch s := subject.(type) {
	case *ListValue:
		st.sub = s
		st.n = len(*s)
		if kind == hofMap || kind == hofFilter {
			acc := make(ListValue, 0, st.n)
			st.acc = &acc
		}
	case *StringValue:
		st.sstr = s
		st.n = len(*s)
		if kind == hofMap || kind == hofFilter {
			st.sacc = &StringValue{}
		}
	default:
		// Unsupported subject: exact std semantics via the generic path,
		// calling the same callee value the op popped off the stack.
		callback := &closureVal{fnIdx: fnIdx, parentScope: curScope, defn: ft.defn}
		var args []Value
		if kind == hofReduce {
			args = []Value{subject, seed, callback}
		} else {
			args = []Value{subject, callback}
		}
		res, err := vm.genericCall(callee, args)
		if err != nil {
			return false, err
		}
		return finish(res)
	}
	if st.n == 0 {
		return finish(st.hofResult())
	}
	st.returnPC = vm.pc
	// Build the first callback frame (slab window when possible).
	if !ft.needsHeapScope && !ft.hasRestArg {
		if base := vm.localSP; base+ft.localCount+vmSlabHeadroom <= len(vm.localsSlab) {
			window := vm.localsSlab[base : base+ft.localCount : base+ft.localCount+vmSlabHeadroom]
			st.baseLocal = base
			st.frame = callFrame{
				returnPC:    st.returnPC,
				baseSlot:    vm.sp,
				fnIdx:       fnIdx,
				locals:      window,
				scope:       nil,
				baseLocal:   base,
				parentScope: curScope,
			}
			vm.hofFillArgs(&st)
			vm.localSP = base + ft.localCount + vmSlabHeadroom
			st.frameIdx = len(vm.frames)
			vm.hofStack = append(vm.hofStack, st)
			vm.frames = append(vm.frames, vm.hofStack[len(vm.hofStack)-1].frame)
			vm.pc = ft.offset
			return true, nil
		}
	}
	locals := make([]Value, ft.localCount)
	for j := range locals {
		locals[j] = null
	}
	st.frame = callFrame{
		returnPC:    st.returnPC,
		baseSlot:    vm.sp,
		fnIdx:       fnIdx,
		locals:      locals,
		scope:       nil,
		baseLocal:   -1,
		parentScope: curScope,
	}
	if ft.needsHeapScope {
		st.frame.scope = &vmScope{names: ft.localNames, values: locals, parent: curScope}
	}
	vm.hofFillArgs(&st)
	st.frameIdx = len(vm.frames)
	vm.hofStack = append(vm.hofStack, st)
	vm.frames = append(vm.frames, vm.hofStack[len(vm.hofStack)-1].frame)
	vm.pc = ft.offset
	return true, nil
}

// vmSlabSlots is the per-VM locals slab size. Frames that do not fit fall
// back to heap-allocated locals (correct, just slower).
const vmSlabSlots = 65536

// vmSlabHeadroom reserves extra capacity per slab window so that the
// defensive append growth in opStoreLocal stays inside this frame's own
// reserved region instead of corrupting the next frame's window.
const vmSlabHeadroom = 8

func newVM(chunk *bytecodeChunk, ctx *Context) *VM {
	return &VM{
		chunk:      chunk,
		pc:         0,
		stack:      make([]Value, 1024),
		sp:         0,
		frames:     make([]callFrame, 0, 64),
		ctx:        ctx,
		localsSlab: make([]Value, vmSlabSlots),
	}
}

// vmPool recycles VM structs to reduce allocations for bytecode() calls.
var vmPool = sync.Pool{
	New: func() interface{} {
		return &VM{
			stack:      make([]Value, 1024),
			frames:     make([]callFrame, 0, 64),
			localsSlab: make([]Value, vmSlabSlots),
		}
	},
}

// acquireVM gets a VM from the pool and initializes it for the given chunk.
func acquireVM(chunk *bytecodeChunk, ctx *Context) *VM {
	vm := vmPool.Get().(*VM)
	vm.chunk = chunk
	vm.pc = 0
	vm.sp = 0
	vm.frames = vm.frames[:0]
	vm.ctx = ctx
	vm.initLocals = nil
	vm.outerScope = nil
	vm.localSP = 0
	vm.hofStack = vm.hofStack[:0]
	return vm
}

// releaseVM returns a VM to the pool after clearing references.
func releaseVM(vm *VM) {
	// Clear stack references to avoid holding onto GC-pinned values
	for i := 0; i < vm.sp; i++ {
		vm.stack[i] = nil
	}
	// Clear consumed slab windows for the same reason.
	for i := 0; i < vm.localSP; i++ {
		vm.localsSlab[i] = nil
	}
	vm.localSP = 0
	// Clear HOF loop states (hold Value references).
	for i := range vm.hofStack {
		vm.hofStack[i].frame.locals = nil
		vm.hofStack[i].sub = nil
		vm.hofStack[i].sstr = nil
		vm.hofStack[i].acc = nil
		vm.hofStack[i].sacc = nil
		vm.hofStack[i].seed = nil
	}
	vm.hofStack = vm.hofStack[:0]
	vm.chunk = nil
	vm.ctx = nil
	vm.initLocals = nil
	vm.outerScope = nil
	vmPool.Put(vm)
}

// genericCall invokes callee with args outside the fast-path frame machinery:
// bytecode closures run in a pooled child VM (bridged scopes), Oak AST
// closures and arbitrary callables dispatch through the tree-walker. Used by
// opHOF* for subjects the in-frame loops do not support, preserving exact
// std-lib semantics (M1 TODO-12 §7.8).
func (vm *VM) genericCall(callee Value, args []Value) (Value, *runtimeError) {
	switch fn := callee.(type) {
	case *closureVal:
		if fn.defn != nil {
			chunk := vm.ctx.getOrCompileChunk(fn.defn)
			initVals := normalizeCallArgs(len(fn.defn.args), args)
			if fn.defn.restArg != "" {
				var restList ListValue
				if len(args) > len(fn.defn.args) {
					restList = ListValue(args[len(fn.defn.args):])
				} else {
					restList = ListValue{}
				}
				initVals = append(initVals[:len(fn.defn.args)], &restList)
			}
			sub := acquireVM(chunk, vm.ctx)
			sub.initLocals = initVals
			sub.outerScope = fn.parentScope
			res, err := sub.run()
			releaseVM(sub)
			return res, err
		}
		if fn.call != nil {
			return fn.call(args)
		}
		return vm.ctx.evalFnCall(callee, false, args)
	case BuiltinFnValue:
		for i, a := range args {
			args[i] = vm.interopValue(a)
		}
		return fn.fn(args)
	default:
		return vm.ctx.evalFnCall(callee, false, args)
	}
}

func (vm *VM) push(v Value) {
	if vm.sp >= len(vm.stack) {
		vm.stack = append(vm.stack, make([]Value, len(vm.stack))...)
	}
	vm.stack[vm.sp] = v
	vm.sp++
}

func (vm *VM) pop() Value {
	vm.sp--
	return vm.stack[vm.sp]
}

func (vm *VM) peek() Value {
	return vm.stack[vm.sp-1]
}

func (vm *VM) fetchU8() byte {
	b := vm.chunk.code[vm.pc]
	vm.pc++
	return b
}

func (vm *VM) fetchU16() int {
	lo := vm.chunk.code[vm.pc]
	hi := vm.chunk.code[vm.pc+1]
	vm.pc += 2
	return int(lo) | int(hi)<<8
}

func (vm *VM) fetchI32() int {
	b := vm.chunk.code
	v := int(int32(b[vm.pc]) | int32(b[vm.pc+1])<<8 | int32(b[vm.pc+2])<<16 | int32(b[vm.pc+3])<<24)
	vm.pc += 4
	return v
}

// currentFrame returns the top call frame, or nil at top-level.
func (vm *VM) currentFrame() *callFrame {
	if len(vm.frames) == 0 {
		return nil
	}
	return &vm.frames[len(vm.frames)-1]
}

// currentScope returns the scope for upvalue resolution.
func (vm *VM) currentScope(topScope *vmScope) *vmScope {
	// Slab frames have nil scopes (their locals are slot-addressed and can
	// never be captured by closures). Walk down to the nearest heap scope so
	// name-based upvalue lookups from inside slab frames still resolve
	// against the enclosing lexical scope chain (M1 TODO-08).
	for i := len(vm.frames) - 1; i >= 0; i-- {
		if s := vm.frames[i].scope; s != nil {
			return s
		}
	}
	return topScope
}

// posAtPC returns the source position for a given program counter value
// by binary-searching the source map for the last entry at or before pc.
func (vm *VM) posAtPC(pc int) pos {
	sm := vm.chunk.sourceMap
	if len(sm) == 0 {
		return pos{}
	}
	// Binary search: find the last entry where offset <= pc
	lo, hi := 0, len(sm)-1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if sm[mid].offset <= pc {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if sm[lo].offset <= pc {
		return sm[lo].pos
	}
	return pos{}
}

// vmError creates a runtimeError with the current source position and a
// stack trace unwound from the call frame stack.
func (vm *VM) vmError(reason string) *runtimeError {
	p := vm.posAtPC(vm.pc - 1) // pc already advanced past the opcode
	err := &runtimeError{
		reason: reason,
		pos:    p,
	}
	// Build stack trace from call frames (most recent first)
	for i := len(vm.frames) - 1; i >= 0; i-- {
		f := &vm.frames[i]
		name := ""
		if f.fnIdx >= 0 && f.fnIdx < len(vm.chunk.functions) {
			name = vm.chunk.functions[f.fnIdx].name
		}
		framePos := vm.posAtPC(f.returnPC)
		err.stackTrace = append(err.stackTrace, stackEntry{
			name: name,
			pos:  framePos,
		})
	}
	return err
}

// wrapClosureForInterop wraps a bytecode closureVal as a BuiltinFnValue so
// it can be called by the tree-walking interpreter.  Each invocation spawns
// a small child VM that runs only the closure's body and returns.
var _ = (*VM).wrapClosureForInterop //nolint:unused

func (vm *VM) wrapClosureForInterop(cv *closureVal) BuiltinFnValue {
	chunk := vm.chunk
	ctx := vm.ctx
	return BuiltinFnValue{
		name: chunk.functions[cv.fnIdx].name,
		fn: func(args []Value) (Value, *runtimeError) {
			child := acquireVM(chunk, ctx)
			result, err := child.callClosure(cv, args)
			releaseVM(child)
			return result, err
		},
	}
}

// callClosure executes a closureVal with the given arguments in a fresh
// mini-run and returns the result.
func (vm *VM) callClosure(cv *closureVal, args []Value) (Value, *runtimeError) {
	ft := &vm.chunk.functions[cv.fnIdx]

	locals := make([]Value, ft.localCount)
	for i := range locals {
		locals[i] = null
	}
	for i := 0; i < ft.arity && i < len(args); i++ {
		locals[i] = args[i]
	}
	if ft.hasRestArg && ft.arity < len(locals) {
		var restList ListValue
		if len(args) > ft.arity {
			restList = ListValue(args[ft.arity:])
		} else {
			restList = ListValue{}
		}
		locals[ft.arity] = &restList
	}

	scope := &vmScope{
		names:  ft.localNames,
		values: locals,
		parent: cv.parentScope,
	}

	vm.frames = append(vm.frames, callFrame{
		returnPC:  len(vm.chunk.code), // halt after return
		baseSlot:  vm.sp,
		fnIdx:     cv.fnIdx,
		locals:    locals,
		scope:     scope,
		baseLocal: -1,
	})
	vm.pc = ft.offset
	return vm.run()
}

// interopValue prepares a value for safe use by the tree-walking interpreter.
// closureVal gets its call field populated so evalFnCall can invoke it.
// Objects/lists are walked in-place (not copied) to find nested closureVals.
func (vm *VM) interopValue(v Value) Value {
	switch val := v.(type) {
	case *closureVal:
		vm.ensureClosureCallable(val)
	case ObjectValue:
		for _, item := range val {
			vm.interopValue(item)
		}
	case *ListValue:
		for _, item := range *val {
			vm.interopValue(item)
		}
	}
	return v
}

// ensureClosureCallable populates the call field on a closureVal so that
// the tree-walking interpreter can invoke it through evalFnCall.
// A dedicated child VM is allocated per closure to avoid pool Get/Put overhead
// on every callback invocation (e.g., 10k map iterations).
func (vm *VM) ensureClosureCallable(cv *closureVal) {
	if cv.call != nil {
		return
	}
	chunk := vm.chunk
	ctx := vm.ctx
	childVM := newVM(chunk, ctx)
	cv.call = func(args []Value) (Value, *runtimeError) {
		childVM.sp = 0
		childVM.frames = childVM.frames[:0]
		return childVM.callClosure(cv, args)
	}
}

// ---------------------------------------------------------------------------
// VM execution — single flat loop, no recursive run() calls
// ---------------------------------------------------------------------------

func (vm *VM) run() (Value, *runtimeError) {
	code := vm.chunk.code

	// Skip expensive top-level allocation when we're already inside a frame
	// (e.g., callClosure for callback dispatch). The frame has its own locals
	// and scope, so topLocals/topScope are never accessed.
	var topLocals []Value
	var topScope *vmScope

	if len(vm.frames) == 0 {
		// Top-level execution — allocate locals
		topLocalCount := len(vm.chunk.topLevelNames)
		if topLocalCount < 256 {
			topLocalCount = 256 // ensure enough slots for forward-declared locals
		}
		topLocals = make([]Value, topLocalCount)
		for i := range topLocals {
			topLocals[i] = null
		}

		// Apply pre-populated initial locals (from bytecode() builtin)
		for i, v := range vm.initLocals {
			if i < len(topLocals) {
				topLocals[i] = v
			}
		}

		// Build top-level scope
		topScope = &vmScope{
			names:  vm.chunk.topLevelNames,
			values: topLocals,
			parent: vm.outerScope,
		}
	}

	for {
		if vm.pc >= len(code) {
			break
		}

		op := opcode(code[vm.pc])
		vm.pc++

		switch op {
		case opHalt:
			if vm.sp > 0 {
				return vm.pop(), nil
			}
			return null, nil

		case opNop:
			// do nothing

		case opConstNull:
			vm.push(null)

		case opConstEmpty:
			vm.push(empty)

		case opConstTrue:
			vm.push(BoolValue(true))

		case opConstFalse:
			vm.push(BoolValue(false))

		case opConstInt:
			v := vm.fetchI32()
			vm.push(makeIntValue(int64(v)))

		case opConstFloat:
			idx := vm.fetchU16()
			vm.push(FloatValue(vm.chunk.constants[idx].f))

		case opConstStr:
			idx := vm.fetchU16()
			sv := StringValue(vm.chunk.constants[idx].str)
			vm.push(&sv)

		case opConstAtom:
			idx := vm.fetchU16()
			vm.push(AtomValue(vm.chunk.constants[idx].str))

		case opPop:
			vm.sp--

		case opDup:
			vm.push(vm.peek())

		case opSwap:
			vm.stack[vm.sp-1], vm.stack[vm.sp-2] = vm.stack[vm.sp-2], vm.stack[vm.sp-1]

		case opLoadLocal:
			slot := vm.fetchU16()
			frame := vm.currentFrame()
			if frame != nil {
				if slot < len(frame.locals) {
					vm.push(frame.locals[slot])
				} else {
					vm.push(null)
				}
			} else {
				if slot < len(topLocals) {
					v := topLocals[slot]
					if v == null && slot < len(vm.chunk.topLevelNames) {
						// Phantom local not yet assigned — check Context scope (builtins)
						if cv, _ := vm.ctx.scope.get(vm.chunk.topLevelNames[slot]); cv != nil && cv != null {
							topLocals[slot] = cv // cache so upvalue lookups also find it
							vm.push(cv)
						} else {
							vm.push(null)
						}
					} else {
						vm.push(v)
					}
				} else {
					vm.push(null)
				}
			}

		case opStoreLocal:
			slot := vm.fetchU16()
			val := vm.pop()
			frame := vm.currentFrame()
			if frame != nil {
				for slot >= len(frame.locals) {
					frame.locals = append(frame.locals, null)
				}
				frame.locals[slot] = val
			} else {
				for slot >= len(topLocals) {
					topLocals = append(topLocals, null)
					// Also extend topScope.values to keep it in sync
					topScope.values = topLocals
				}
				topLocals[slot] = val
			}

		case opLtLocalJumpTrue:
			slot := vm.fetchU16()
			imm := vm.fetchI32()
			target := vm.fetchI32()
			var cond Value
			frame := vm.currentFrame()
			if frame != nil {
				var lv Value = null
				if slot < len(frame.locals) {
					lv = frame.locals[slot]
				}
				if li, ok := lv.(IntValue); ok {
					cond = BoolValue(li < IntValue(imm))
				} else {
					cond = vmLt(lv, makeIntValue(int64(imm)))
				}
			} else {
				// Top level: mirror opLoadLocal's topLocals read.
				var lv Value = null
				if slot < len(topLocals) {
					lv = topLocals[slot]
				}
				cond = vmLt(lv, makeIntValue(int64(imm)))
			}
			if bv, ok := cond.(BoolValue); ok && bool(bv) {
				// Match: fall through to the true arm (cond consumed).
			} else {
				// No match (incl. non-bool cond): jump to the else stream —
				// wildcard body, or an emitted null for the 1-arm shape.
				vm.pc = target
			}

		case opHOFMap, opHOFFilter, opHOFEach, opHOFReduce:
			fnIdx := vm.fetchU16()
			kind := hofKind(op - opHOFMap)
			var seed Value
			if kind == hofReduce {
				// stack: [callee, subject, seed] (seed emitted last)
				seed = vm.pop()
			}
			subject := vm.pop()
			callee := vm.pop()
			started, hErr := vm.hofStart(kind, fnIdx, subject, seed, callee, topScope)
			if hErr != nil {
				return nil, hErr
			}
			if started {
				continue
			}

		case opLoadUpvalFast, opStoreUpvalFast:
			depth := int(vm.fetchU16())
			slot := int(vm.fetchU16())
			nameIdx := vm.fetchU16()
			name := vm.chunk.constants[nameIdx].str
			frame := vm.currentFrame()
			var sc *vmScope
			if frame != nil {
				if frame.scope != nil {
					sc = frame.scope.parent
				} else {
					sc = frame.parentScope
				}
			} else {
				sc = topScope
			}
			// Walk (depth-1) lexical hops up the persistent heap scope chain.
			for hop := 1; hop < depth && sc != nil; hop++ {
				sc = sc.parent
			}
			if sc != nil && slot < len(sc.values) {
				if op == opLoadUpvalFast {
					v := sc.values[slot]
					if v == nil {
						v = null
					}
					vm.push(v)
					continue
				}
				// opStoreUpvalFast: frame.locals aliases sc.values, so the
				// write lands in the ancestor's slot directly.
				sc.values[slot] = vm.pop()
				continue
			}
			// Defensive fallback: resolve by name (bridges, extended slots).
			if op == opLoadUpvalFast {
				found := false
				if sc != nil {
					if val, ok := sc.get(name); ok && val != nil && val != null {
						vm.push(val)
						found = true
					}
				}
				if !found {
					if val, _ := vm.ctx.scope.get(name); val != nil && val != null {
						vm.push(val)
					} else {
						vm.push(null)
					}
				}
				continue
			}
			val := vm.pop()
			if sc != nil {
				sc.set(name, val)
			}

		case opLoadUpval:
			nameIdx := vm.fetchU16()
			name := vm.chunk.constants[nameIdx].str
			scope := vm.currentScope(topScope)
			found := false
			if scope != nil {
				// For a function frame, start searching from the parent scope
				// (the closure's captured scope), not the current function's scope.
				// Slab frames carry the captured scope in frame.parentScope.
				frame := vm.currentFrame()
				var searchScope *vmScope
				if frame != nil {
					if frame.scope != nil {
						searchScope = frame.scope.parent
					} else {
						searchScope = frame.parentScope
					}
				}
				if searchScope == nil {
					searchScope = scope
				}
				if val, ok := searchScope.get(name); ok && val != nil && val != null {
					vm.push(val)
					found = true
				}
			}
			if !found {
				// Fall back to the tree-walker scope (builtins, globals)
				if val, _ := vm.ctx.scope.get(name); val != nil && val != null {
					vm.push(val)
				} else {
					vm.push(null)
				}
			}

		case opStoreUpval:
			nameIdx := vm.fetchU16()
			name := vm.chunk.constants[nameIdx].str
			val := vm.pop()
			scope := vm.currentScope(topScope)
			if scope != nil {
				frame := vm.currentFrame()
				var searchScope *vmScope
				if frame != nil {
					if frame.scope != nil {
						searchScope = frame.scope.parent
					} else {
						searchScope = frame.parentScope
					}
				}
				if searchScope == nil {
					searchScope = scope
				}
				searchScope.set(name, val)
			}

		case opAdd:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(li + ri)))
					break
				}
			}
			// Fast path: string + string (common in string-heavy code)
			if ls, ok := left.(*StringValue); ok {
				if rs, ok := right.(*StringValue); ok {
					buf := make([]byte, 0, len(*ls)+len(*rs))
					buf = append(buf, *ls...)
					buf = append(buf, *rs...)
					sv := StringValue(buf)
					vm.push(&sv)
					break
				}
			}
			vm.push(vmAdd(left, right))

		case opSub:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(li - ri)))
					break
				}
			}
			vm.push(vmSub(left, right))

		case opMul:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(li * ri)))
					break
				}
			}
			vm.push(vmMul(left, right))

		case opDiv:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok && ri != 0 {
					if li%ri == 0 {
						vm.push(makeIntValue(int64(li / ri)))
					} else {
						vm.push(FloatValue(float64(li) / float64(ri)))
					}
					break
				}
			}
			result, errReason := vmDiv(left, right)
			if errReason != "" {
				return nil, vm.vmError(errReason)
			}
			vm.push(result)

		case opMod:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok && ri != 0 {
					vm.push(makeIntValue(int64(li % ri)))
					break
				}
			}
			result, errReason := vmMod(left, right)
			if errReason != "" {
				return nil, vm.vmError(errReason)
			}
			vm.push(result)

		case opPow:
			right := vm.pop()
			left := vm.pop()
			vm.push(vmPow(left, right))

		case opNeg:
			v := vm.pop()
			switch val := v.(type) {
			case IntValue:
				vm.push(makeIntValue(int64(-val)))
			case FloatValue:
				vm.push(FloatValue(-val))
			default:
				vm.push(null)
			}

		case opBAnd:
			right := vm.pop()
			left := vm.pop()
			switch lv := left.(type) {
			case IntValue:
				if rv, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(lv & rv)))
				} else {
					vm.push(null)
				}
			case BoolValue:
				if rv, ok := right.(BoolValue); ok {
					vm.push(BoolValue(bool(lv) && bool(rv)))
				} else {
					vm.push(null)
				}
			case *StringValue:
				if rv, ok := right.(*StringValue); ok {
					maxLen := len(*lv)
					if len(*rv) > maxLen {
						maxLen = len(*rv)
					}
					res := make([]byte, maxLen)
					for i := 0; i < maxLen; i++ {
						var lb, rb byte
						if i < len(*lv) {
							lb = (*lv)[i]
						}
						if i < len(*rv) {
							rb = (*rv)[i]
						}
						res[i] = lb & rb
					}
					sv := StringValue(res)
					vm.push(&sv)
				} else {
					vm.push(null)
				}
			default:
				vm.push(null)
			}

		case opBOr:
			right := vm.pop()
			left := vm.pop()
			switch lv := left.(type) {
			case IntValue:
				if rv, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(lv | rv)))
				} else {
					vm.push(null)
				}
			case BoolValue:
				if rv, ok := right.(BoolValue); ok {
					vm.push(BoolValue(bool(lv) || bool(rv)))
				} else {
					vm.push(null)
				}
			case *StringValue:
				if rv, ok := right.(*StringValue); ok {
					maxLen := len(*lv)
					if len(*rv) > maxLen {
						maxLen = len(*rv)
					}
					res := make([]byte, maxLen)
					for i := 0; i < maxLen; i++ {
						var lb, rb byte
						if i < len(*lv) {
							lb = (*lv)[i]
						}
						if i < len(*rv) {
							rb = (*rv)[i]
						}
						res[i] = lb | rb
					}
					sv := StringValue(res)
					vm.push(&sv)
				} else {
					vm.push(null)
				}
			default:
				vm.push(null)
			}

		case opBXor:
			right := vm.pop()
			left := vm.pop()
			switch lv := left.(type) {
			case IntValue:
				if rv, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(lv ^ rv)))
				} else {
					vm.push(null)
				}
			case BoolValue:
				if rv, ok := right.(BoolValue); ok {
					vm.push(BoolValue(bool(lv) != bool(rv)))
				} else {
					vm.push(null)
				}
			case *StringValue:
				if rv, ok := right.(*StringValue); ok {
					maxLen := len(*lv)
					if len(*rv) > maxLen {
						maxLen = len(*rv)
					}
					res := make([]byte, maxLen)
					for i := 0; i < maxLen; i++ {
						var lb, rb byte
						if i < len(*lv) {
							lb = (*lv)[i]
						}
						if i < len(*rv) {
							rb = (*rv)[i]
						}
						res[i] = lb ^ rb
					}
					sv := StringValue(res)
					vm.push(&sv)
				} else {
					vm.push(null)
				}
			default:
				vm.push(null)
			}

		case opBRShift:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(makeIntValue(int64(li >> uint(ri))))
				} else {
					vm.push(null)
				}
			} else {
				vm.push(null)
			}

		case opEq:
			right := vm.pop()
			left := vm.pop()
			// Fast path for int comparison (dominant case in pattern matching)
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li == ri))
					break
				}
			}
			// Object identity parity with the tree-walker's eq (evalBinaryOp):
			// `=` on object records compares map identity. Deep Eq here would
			// recurse infinitely on cyclic graphs (lru nodes, parent links).
			if lo, ok := left.(ObjectValue); ok {
				if ro, ok2 := right.(ObjectValue); ok2 {
					vm.push(BoolValue(ifaceDataPtr(lo) == ifaceDataPtr(ro)))
				} else {
					vm.push(BoolValue(false))
				}
				break
			}
			// Fast path for single-byte string comparison (dominant in char loops)
			if ls, ok := left.(*StringValue); ok {
				if rs, ok := right.(*StringValue); ok {
					if len(*ls) == 1 && len(*rs) == 1 {
						vm.push(BoolValue((*ls)[0] == (*rs)[0]))
						break
					}
				}
			}
			vm.push(BoolValue(left.Eq(right)))

		case opNeq:
			right := vm.pop()
			left := vm.pop()
			if lo, ok := left.(ObjectValue); ok {
				if ro, ok2 := right.(ObjectValue); ok2 {
					vm.push(BoolValue(ifaceDataPtr(lo) != ifaceDataPtr(ro)))
				} else {
					vm.push(BoolValue(true))
				}
				break
			}
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li != ri))
					break
				}
			}
			vm.push(BoolValue(!left.Eq(right)))

		case opDeepEq:
			right := vm.pop()
			left := vm.pop()
			vm.push(BoolValue(vmDeepEq(left, right)))

		case opGt:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li > ri))
					break
				}
			}
			vm.push(vmGt(left, right))

		case opLt:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li < ri))
					break
				}
			}
			vm.push(vmLt(left, right))

		case opGeq:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li >= ri))
					break
				}
			}
			vm.push(vmGeq(left, right))

		case opLeq:
			right := vm.pop()
			left := vm.pop()
			if li, ok := left.(IntValue); ok {
				if ri, ok := right.(IntValue); ok {
					vm.push(BoolValue(li <= ri))
					break
				}
			}
			vm.push(vmLeq(left, right))

		case opNot:
			v := vm.pop()
			if bv, ok := v.(BoolValue); ok {
				vm.push(BoolValue(!bv))
			} else {
				vm.push(BoolValue(true))
			}

		case opConcat:
			right := vm.pop()
			left := vm.pop()
			vm.push(vmConcat(left, right))

		case opMakeList:
			count := vm.fetchU16()
			elems := make([]Value, count)
			for j := count - 1; j >= 0; j-- {
				elems[j] = vm.pop()
			}
			lv := ListValue(elems)
			vm.push(&lv)

		case opMakeObj:
			count := vm.fetchU16()
			obj := make(ObjectValue, count)
			for j := 0; j < count; j++ {
				val := vm.pop()
				key := vm.pop()
				var keyStr string
				switch k := key.(type) {
				case *StringValue:
					keyStr = string(*k)
				case AtomValue:
					keyStr = string(k)
				default:
					keyStr = key.String()
				}
				obj[keyStr] = val
			}
			vm.push(obj)

		case opGetProp:
			key := vm.pop()
			obj := vm.pop()
			result, errReason := vmGetProp(obj, key)
			if errReason != "" {
				return nil, vm.vmError(errReason)
			}
			vm.push(result)

		case opSetProp:
			val := vm.pop()
			key := vm.pop()
			obj := vm.pop()
			vmSetProp(obj, key, val)
			vm.push(val)

		case opGetPropConst:
			idx := vm.fetchU16()
			keyStr := vm.chunk.constants[idx].str
			obj := vm.pop()
			switch o := obj.(type) {
			case ObjectValue:
				if v, ok := o[keyStr]; ok {
					vm.push(v)
				} else {
					vm.push(null)
				}
			case ClassValue:
				if v, ok := o.static[keyStr]; ok {
					vm.push(v)
				} else {
					vm.push(null)
				}
			default:
				sv := StringValue(keyStr)
				res, errMsg := vmGetProp(obj, &sv)
				if errMsg != "" {
					return nil, vm.vmError(errMsg)
				}
				vm.push(res)
			}

		case opSetPropConst:
			idx := vm.fetchU16()
			keyStr := vm.chunk.constants[idx].str
			val := vm.pop()
			obj := vm.pop()
			switch o := obj.(type) {
			case ObjectValue:
				mu := objLock(o)
				o[keyStr] = val
				mu.Unlock()
			default:
				sv := StringValue(keyStr)
				vmSetProp(obj, &sv, val)
			}
			vm.push(val)

		case opJump:
			target := vm.fetchI32()
			vm.pc = target

		case opJumpFalse:
			target := vm.fetchI32()
			cond := vm.pop()
			if !isTruthy(cond) {
				vm.pc = target
			}

		case opClosure:
			fnIdx := vm.fetchU16()
			// Capture the current scope as a shared reference (NOT a copy).
			// This allows closures to see updates made after the closure is created
			// (e.g., named functions storing themselves in the outer scope).
			curScope := vm.currentScope(topScope)
			ft := &vm.chunk.functions[fnIdx]
			vm.push(&closureVal{fnIdx: fnIdx, parentScope: curScope, defn: ft.defn})

		case opCall:
			arity := int(vm.fetchU8())

			// Peek at callee (below args on stack)
			callee := vm.stack[vm.sp-arity-1]

			if fn, ok := callee.(*closureVal); ok {
				// Fast path for closure calls: copy args directly from
				// stack to locals, avoiding intermediate args slice allocation.
				ft := &vm.chunk.functions[fn.fnIdx]

				// Slab fast path (M1 TODO-08, §7.2): frames of functions that
				// contain no inner fnNode run entirely on the VM's contiguous
				// locals slab — 0 heap allocations for locals, scope, or frame.
				if !ft.needsHeapScope && !ft.hasRestArg {
					if base := vm.localSP; base+ft.localCount+vmSlabHeadroom <= len(vm.localsSlab) {
						window := vm.localsSlab[base : base+ft.localCount : base+ft.localCount+vmSlabHeadroom]
						for i := range window {
							window[i] = null
						}
						argBase := vm.sp - arity
						n := ft.arity
						if arity < n {
							n = arity
						}
						copy(window[:n], vm.stack[argBase:argBase+n])
						vm.localSP = base + ft.localCount + vmSlabHeadroom
						vm.sp = argBase - 1 // pop callee + args
						vm.frames = append(vm.frames, callFrame{
							returnPC:    vm.pc,
							baseSlot:    vm.sp,
							fnIdx:       fn.fnIdx,
							locals:      window,
							scope:       nil,
							baseLocal:   base,
							parentScope: fn.parentScope,
						})
						vm.pc = ft.offset
						continue
					}
				}

				locals := make([]Value, ft.localCount)
				for i := range locals {
					locals[i] = null
				}
				paramCount := ft.arity
				// Copy args from stack directly to locals
				argBase := vm.sp - arity
				for i := 0; i < paramCount && i < arity; i++ {
					locals[i] = vm.stack[argBase+i]
				}
				if ft.hasRestArg && paramCount < len(locals) {
					var restList ListValue
					if arity > paramCount {
						restList = make(ListValue, arity-paramCount)
						copy(restList, vm.stack[argBase+paramCount:vm.sp])
					} else {
						restList = ListValue{}
					}
					locals[paramCount] = &restList
				}

				// Pop callee + args
				vm.sp = argBase - 1

				scope := &vmScope{
					names:  ft.localNames,
					values: locals,
					parent: fn.parentScope,
				}

				vm.frames = append(vm.frames, callFrame{
					returnPC: vm.pc,
					baseSlot: vm.sp,
					fnIdx:    fn.fnIdx,
					locals:   locals,
					scope:    scope,
				})
				vm.pc = ft.offset
			} else {
				// General path for builtins, FnValue, ClassValue
				args := make([]Value, arity)
				for j := arity - 1; j >= 0; j-- {
					args[j] = vm.pop()
				}
				vm.pop() // pop callee

				switch fn := callee.(type) {
				case BuiltinFnValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, bErr := fn.fn(args)
					if bErr != nil {
						return nil, bErr
					}
					vm.push(result)

				case FnValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, fErr := vm.ctx.evalFnCall(fn, false, args)
					if fErr != nil {
						return nil, fErr
					}
					vm.push(result)

				case ClassValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, cErr := vm.ctx.evalFnCall(fn, false, args)
					if cErr != nil {
						return nil, cErr
					}
					vm.push(result)

				default:
					return nil, vm.vmError(fmt.Sprintf("%s is not a function and cannot be called", callee))
				}
			}

		case opReturn:
			retVal := vm.pop()
			if len(vm.frames) == 0 {
				return retVal, nil
			}
			frame := vm.frames[len(vm.frames)-1]
			vm.frames = vm.frames[:len(vm.frames)-1]
			vm.pc = frame.returnPC
			vm.sp = frame.baseSlot
			if frame.baseLocal >= 0 {
				// Slab frame: O(1) bulk "free" by restoring the slab pointer.
				vm.localSP = frame.baseLocal
			}
			// In-frame HOF loop step (M1 TODO-12, §7.8): if the frame that
			// just returned was an opHOF* callback, accumulate and iterate
			// (or finish) instead of returning to the call site.
			if hs := len(vm.hofStack); hs > 0 && vm.hofStack[hs-1].frameIdx == len(vm.frames) {
				st := &vm.hofStack[hs-1]
				st.hofAccumulate(retVal)
				st.i++
				if st.i < st.n {
					// Re-push the same prepared callback frame for the next
					// element; refresh its locals window in place.
					vm.hofRefresh(st)
					vm.frames = append(vm.frames, st.frame)
					vm.pc = vm.chunk.functions[st.fnIdx].offset
					continue
				}
				// Loop finished: free the slab window and push the result.
				if st.baseLocal >= 0 {
					vm.localSP = st.baseLocal
				}
				result := st.hofResult()
				vm.hofStack = vm.hofStack[:hs-1]
				vm.push(result)
				vm.pc = st.returnPC
				continue
			}
			vm.push(retVal)

		case opCallSpread:
			arity := int(vm.fetchU8())
			raw := make([]Value, arity)
			for j := arity - 1; j >= 0; j-- {
				raw[j] = vm.pop()
			}
			callee := vm.pop()

			// Spread the last argument if it's a list
			var args []Value
			if arity > 0 {
				last := raw[arity-1]
				if lv, ok := last.(*ListValue); ok {
					args = make([]Value, 0, arity-1+len(*lv))
					args = append(args, raw[:arity-1]...)
					args = append(args, *lv...)
				} else {
					args = raw
				}
			} else {
				args = raw
			}

			switch fn := callee.(type) {
			case *closureVal:
				ft := &vm.chunk.functions[fn.fnIdx]
				locals := make([]Value, ft.localCount)
				for i := range locals {
					locals[i] = null
				}
				paramCount := ft.arity
				for i := 0; i < paramCount && i < len(args); i++ {
					locals[i] = args[i]
				}
				if ft.hasRestArg && paramCount < len(locals) {
					var restList ListValue
					if len(args) > paramCount {
						restList = ListValue(args[paramCount:])
					} else {
						restList = ListValue{}
					}
					locals[paramCount] = &restList
				}
				scope := &vmScope{
					names:  ft.localNames,
					values: locals,
					parent: fn.parentScope,
				}
				vm.frames = append(vm.frames, callFrame{
					returnPC:  vm.pc,
					baseSlot:  vm.sp,
					fnIdx:     fn.fnIdx,
					locals:    locals,
					scope:     scope,
					baseLocal: -1,
				})
				vm.pc = ft.offset

			case BuiltinFnValue:
				for i, a := range args {
					args[i] = vm.interopValue(a)
				}
				result, bErr := fn.fn(args)
				if bErr != nil {
					return nil, bErr
				}
				vm.push(result)

			case FnValue:
				for i, a := range args {
					args[i] = vm.interopValue(a)
				}
				result, fErr := vm.ctx.evalFnCall(fn, false, args)
				if fErr != nil {
					return nil, fErr
				}
				vm.push(result)

			case ClassValue:
				for i, a := range args {
					args[i] = vm.interopValue(a)
				}
				result, cErr := vm.ctx.evalFnCall(fn, false, args)
				if cErr != nil {
					return nil, cErr
				}
				vm.push(result)

			default:
				return nil, vm.vmError(fmt.Sprintf("%s is not a function and cannot be called", callee))
			}

		case opTailCall:
			arity := int(vm.fetchU8())
			argBase := vm.sp - arity
			callee := vm.stack[argBase-1]

			switch fn := callee.(type) {
			case *closureVal:
				ft := &vm.chunk.functions[fn.fnIdx]

				// Reuse the current frame for tail calls
				frame := vm.currentFrame()
				if frame != nil {
					// Fast path: self-recursion — overwrite the frame's locals
					// window in place directly from the operand stack (zero
					// allocations: no args slice, no locals, no scope).
					// Self-recursion where the callee frame is a slab frame is
					// handled by len(frame.locals) >= ft.localCount plus the
					// in-place copy below.
					if fn.fnIdx == frame.fnIdx && len(frame.locals) >= ft.localCount {
						locals := frame.locals
						paramCount := ft.arity
						n := paramCount
						if arity < n {
							n = arity
						}
						copy(locals[:n], vm.stack[argBase:argBase+n])
						for i := paramCount; i < ft.localCount; i++ {
							locals[i] = null
						}
						if ft.hasRestArg && paramCount < len(locals) {
							var restList ListValue
							if arity > paramCount {
								restList = ListValue(vm.stack[argBase+paramCount : vm.sp])
							} else {
								restList = ListValue{}
							}
							locals[paramCount] = &restList
						}
						vm.sp = argBase - 1 // pop callee + args
						if frame.scope != nil {
							frame.scope.values = locals
							frame.scope.parent = fn.parentScope
						}
						vm.pc = ft.offset
					} else {
						// Different function or different local count — allocate new
						args := make([]Value, arity)
						copy(args, vm.stack[argBase:vm.sp])
						vm.sp = argBase - 1 // pop callee + args
						locals := make([]Value, ft.localCount)
						for i := range locals {
							locals[i] = null
						}
						paramCount := ft.arity
						for i := 0; i < paramCount && i < len(args); i++ {
							locals[i] = args[i]
						}
						if ft.hasRestArg && paramCount < len(locals) {
							var restList ListValue
							if len(args) > paramCount {
								restList = ListValue(args[paramCount:])
							} else {
								restList = ListValue{}
							}
							locals[paramCount] = &restList
						}

						scope := &vmScope{
							names:  ft.localNames,
							values: locals,
							parent: fn.parentScope,
						}

						frame.locals = locals
						frame.scope = scope
						frame.fnIdx = fn.fnIdx
						vm.pc = ft.offset
					}
				} else {
					// Tail call at top level — same as regular call
					args := make([]Value, arity)
					copy(args, vm.stack[argBase:vm.sp])
					vm.sp = argBase - 1 // pop callee + args
					locals := make([]Value, ft.localCount)
					for i := range locals {
						locals[i] = null
					}
					paramCount := ft.arity
					for i := 0; i < paramCount && i < len(args); i++ {
						locals[i] = args[i]
					}
					scope := &vmScope{
						names:  ft.localNames,
						values: locals,
						parent: fn.parentScope,
					}
					vm.frames = append(vm.frames, callFrame{
						returnPC:  vm.pc,
						baseSlot:  vm.sp,
						fnIdx:     fn.fnIdx,
						locals:    locals,
						scope:     scope,
						baseLocal: -1,
					})
					vm.pc = ft.offset
				}

			default:
				// For non-closures, just do a regular call
				args := make([]Value, arity)
				copy(args, vm.stack[argBase:vm.sp])
				vm.sp = argBase - 1 // pop callee + args
				switch fn2 := callee.(type) {
				case BuiltinFnValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, bErr := fn2.fn(args)
					if bErr != nil {
						return nil, bErr
					}
					vm.push(result)
				case FnValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, fErr := vm.ctx.evalFnCall(fn2, false, args)
					if fErr != nil {
						return nil, fErr
					}
					vm.push(result)
				case ClassValue:
					for i, a := range args {
						args[i] = vm.interopValue(a)
					}
					result, cErr := vm.ctx.evalFnCall(fn2, false, args)
					if cErr != nil {
						return nil, cErr
					}
					vm.push(result)
				default:
					return nil, vm.vmError(fmt.Sprintf("%s is not a function and cannot be called", callee))
				}
			}

		case opBuiltin:
			bi := vm.fetchU16()
			arity := int(vm.fetchU8())
			args := make([]Value, arity)
			for j := arity - 1; j >= 0; j-- {
				args[j] = vm.pop()
			}
			result, err := vm.callBuiltin(bi, args)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case opImport:
			idx := vm.fetchU16()
			modName := vm.chunk.constants[idx].str
			result, err := vm.doImport(modName)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case opImportDyn:
			modVal := vm.pop()
			var modName string
			if s, ok := modVal.(*StringValue); ok {
				modName = string(*s)
			} else {
				return nil, vm.vmError(fmt.Sprintf("import() expects a string argument, got %s", modVal))
			}
			result, err := vm.doImport(modName)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case opScopePush:
			// No-op in frame-based approach

		case opScopePop:
			// No-op in frame-based approach

		case opMatchJump:
			target := vm.fetchI32()
			pattern := vm.pop()
			tos := vm.peek()
			if !tos.Eq(pattern) {
				vm.pc = target
			}

		default:
			return nil, vm.vmError(fmt.Sprintf("Unknown bytecode opcode: %d", op))
		}
	}

	if vm.sp > 0 {
		return vm.pop(), nil
	}
	return null, nil
}

// ---------------------------------------------------------------------------
// Built-in function dispatch
// ---------------------------------------------------------------------------

func (vm *VM) callBuiltin(idx int, args []Value) (Value, *runtimeError) {
	switch idx {
	case 0: // print
		if len(args) > 0 {
			if sv, ok := args[0].(*StringValue); ok {
				os.Stdout.Write(*sv)
			} else {
				fmt.Print(args[0])
			}
		}
		return null, nil
	case 1: // len
		if len(args) < 1 {
			return makeIntValue(int64(0)), nil
		}
		switch v := args[0].(type) {
		case *StringValue:
			return makeIntValue(int64(len(*v))), nil
		case *ListValue:
			return makeIntValue(int64(len(*v))), nil
		case ObjectValue:
			return makeIntValue(int64(len(v))), nil
		default:
			return makeIntValue(int64(0)), nil
		}
	case 2: // type
		if len(args) < 1 {
			return AtomValue("null"), nil
		}
		switch args[0].(type) {
		case NullValue:
			return AtomValue("null"), nil
		case EmptyValue:
			return AtomValue("empty"), nil
		case IntValue:
			return AtomValue("int"), nil
		case FloatValue:
			return AtomValue("float"), nil
		case BoolValue:
			return AtomValue("bool"), nil
		case *StringValue:
			return AtomValue("string"), nil
		case AtomValue:
			return AtomValue("atom"), nil
		case *ListValue:
			return AtomValue("list"), nil
		case ObjectValue:
			return AtomValue("object"), nil
		case *closureVal:
			return AtomValue("function"), nil
		case FnValue, BuiltinFnValue:
			return AtomValue("function"), nil
		default:
			return AtomValue("null"), nil
		}
	case 3: // string
		if len(args) < 1 {
			return MakeString(""), nil
		}
		switch v := args[0].(type) {
		case *StringValue:
			return args[0], nil
		case AtomValue:
			return MakeString(string(v)), nil
		default:
			_ = v
			return MakeString(args[0].String()), nil
		}
	case 4: // int
		if len(args) < 1 {
			return makeIntValue(int64(0)), nil
		}
		switch v := args[0].(type) {
		case IntValue:
			return v, nil
		case FloatValue:
			return makeIntValue(int64(v)), nil
		case *StringValue:
			if n, err := strconv.ParseInt(string(*v), 10, 64); err == nil {
				return makeIntValue(int64(n)), nil
			}
			return null, nil
		default:
			return makeIntValue(int64(0)), nil
		}
	case 5: // float
		if len(args) < 1 {
			return FloatValue(0), nil
		}
		switch v := args[0].(type) {
		case FloatValue:
			return v, nil
		case IntValue:
			return FloatValue(v), nil
		case *StringValue:
			if f, err := strconv.ParseFloat(string(*v), 64); err == nil {
				return FloatValue(f), nil
			}
			return null, nil
		default:
			return FloatValue(0), nil
		}
	case 6: // codepoint
		if len(args) < 1 {
			return null, nil
		}
		if sv, ok := args[0].(*StringValue); ok && len(*sv) > 0 {
			return makeIntValue(int64((*sv)[0])), nil
		}
		return null, nil
	case 7: // char
		if len(args) < 1 {
			return MakeString(""), nil
		}
		if iv, ok := args[0].(IntValue); ok {
			return MakeSingleByteString(byte(iv & 0xFF)), nil
		}
		return MakeString(""), nil
	case 8: // keys
		if len(args) < 1 {
			lv := ListValue{}
			return &lv, nil
		}
		if obj, ok := args[0].(ObjectValue); ok {
			keys := make([]Value, 0, len(obj))
			for k := range obj {
				keys = append(keys, MakeString(k))
			}
			lv := ListValue(keys)
			return &lv, nil
		}
		lv := ListValue{}
		return &lv, nil
	case 9: // values
		if len(args) < 1 {
			lv := ListValue{}
			return &lv, nil
		}
		if obj, ok := args[0].(ObjectValue); ok {
			vals := make([]Value, 0, len(obj))
			for _, v := range obj {
				vals = append(vals, v)
			}
			lv := ListValue(vals)
			return &lv, nil
		}
		lv := ListValue{}
		return &lv, nil
	case 10: // slice
		if len(args) < 1 {
			return null, nil
		}
		var startArg Value = null
		if len(args) >= 2 {
			startArg = args[1]
		}
		var endArg Value = null
		if len(args) >= 3 {
			endArg = args[2]
		}
		return vmSlice(args[0], startArg, endArg)
	case 11: // append
		if len(args) < 2 {
			return null, nil
		}
		return vmAppend(args[0], args[1])
	case 12: // wait
		return null, nil
	case 13: // exit
		code := 0
		if len(args) > 0 {
			if iv, ok := args[0].(IntValue); ok {
				code = int(iv)
			}
		}
		os.Exit(code)
		return null, nil
	default:
		return null, nil
	}
}

// doImport handles module imports. Falls back to the tree-walking interpreter.
func (vm *VM) doImport(modName string) (Value, *runtimeError) {
	val, err := vm.ctx.evalGo(strings.NewReader(fmt.Sprintf("import('%s')", modName)))
	if err != nil {
		return nil, vm.vmError(fmt.Sprintf("Error importing %s: %s", modName, err))
	}
	return val, nil
}

// ===========================================================================
// Helper functions for VM operations
// ===========================================================================

func isTruthy(v Value) bool {
	switch val := v.(type) {
	case BoolValue:
		return bool(val)
	case NullValue:
		return false
	case EmptyValue:
		return false
	case IntValue:
		return val != 0
	default:
		return true
	}
}

func vmDeepEq(a, b Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	switch av := a.(type) {
	case *ListValue:
		if bv, ok := b.(*ListValue); ok {
			if len(*av) != len(*bv) {
				return false
			}
			for i := range *av {
				if !vmDeepEq((*av)[i], (*bv)[i]) {
					return false
				}
			}
			return true
		}
		return false
	case ObjectValue:
		if bv, ok := b.(ObjectValue); ok {
			if len(av) != len(bv) {
				return false
			}
			for k, v := range av {
				bval, exists := bv[k]
				if !exists || !vmDeepEq(v, bval) {
					return false
				}
			}
			return true
		}
		return false
	default:
		return a.Eq(b)
	}
}

func vmAdd(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return makeIntValue(int64(av + bv))
		case FloatValue:
			return FloatValue(float64(av) + float64(bv))
		case PointerValue:
			return PointerValue(uintptr(bv) + uintptr(av))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return FloatValue(float64(av) + float64(bv))
		case FloatValue:
			return FloatValue(av + bv)
		}
	case PointerValue:
		if bv, ok := b.(IntValue); ok {
			return PointerValue(uintptr(av) + uintptr(bv))
		}
	case *StringValue:
		if bv, ok := b.(*StringValue); ok {
			result := make([]byte, len(*av)+len(*bv))
			copy(result, *av)
			copy(result[len(*av):], *bv)
			sv := StringValue(result)
			return &sv
		}
	case BoolValue:
		if bv, ok := b.(BoolValue); ok {
			return BoolValue(av || bv)
		}
	}
	return null
}

func vmSub(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return makeIntValue(int64(av - bv))
		case FloatValue:
			return FloatValue(float64(av) - float64(bv))
		case PointerValue:
			return PointerValue(uintptr(bv) - uintptr(av))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return FloatValue(float64(av) - float64(bv))
		case FloatValue:
			return FloatValue(av - bv)
		}
	case PointerValue:
		switch bv := b.(type) {
		case IntValue:
			return PointerValue(uintptr(av) - uintptr(bv))
		case PointerValue:
			return makeIntValue(int64(uintptr(av) - uintptr(bv)))
		}
	}
	return null
}

func vmMul(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return makeIntValue(int64(av * bv))
		case FloatValue:
			return FloatValue(float64(av) * float64(bv))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return FloatValue(float64(av) * float64(bv))
		case FloatValue:
			return FloatValue(av * bv)
		}
	case BoolValue:
		if bv, ok := b.(BoolValue); ok {
			return BoolValue(av && bv)
		}
	}
	return null
}

func vmDiv(a, b Value) (Value, string) {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			if av%bv == 0 {
				return makeIntValue(int64(av / bv)), ""
			}
			return FloatValue(float64(av) / float64(bv)), ""
		case FloatValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(float64(av) / float64(bv)), ""
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(float64(av) / float64(bv)), ""
		case FloatValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(av / bv), ""
		}
	}
	return null, ""
}

func vmMod(a, b Value) (Value, string) {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return makeIntValue(int64(av % bv)), ""
		case FloatValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(math.Mod(float64(av), float64(bv))), ""
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(math.Mod(float64(av), float64(bv))), ""
		case FloatValue:
			if bv == 0 {
				return nil, "Division by zero"
			}
			return FloatValue(math.Mod(float64(av), float64(bv))), ""
		}
	}
	return null, ""
}

func vmPow(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return FloatValue(math.Pow(float64(av), float64(bv)))
		case FloatValue:
			return FloatValue(math.Pow(float64(av), float64(bv)))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return FloatValue(math.Pow(float64(av), float64(bv)))
		case FloatValue:
			return FloatValue(math.Pow(float64(av), float64(bv)))
		}
	}
	return null
}

func vmGt(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(av > bv)
		case FloatValue:
			return BoolValue(float64(av) > float64(bv))
		case PointerValue:
			return BoolValue(uintptr(av) > uintptr(bv))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(float64(av) > float64(bv))
		case FloatValue:
			return BoolValue(av > bv)
		}
	case PointerValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(uintptr(av) > uintptr(bv))
		case PointerValue:
			return BoolValue(uintptr(av) > uintptr(bv))
		}
	case *StringValue:
		if bv, ok := b.(*StringValue); ok {
			return BoolValue(string(*av) > string(*bv))
		}
	}
	return BoolValue(false)
}

func vmLt(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(av < bv)
		case FloatValue:
			return BoolValue(float64(av) < float64(bv))
		case PointerValue:
			return BoolValue(uintptr(av) < uintptr(bv))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(float64(av) < float64(bv))
		case FloatValue:
			return BoolValue(av < bv)
		}
	case PointerValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(uintptr(av) < uintptr(bv))
		case PointerValue:
			return BoolValue(uintptr(av) < uintptr(bv))
		}
	case *StringValue:
		if bv, ok := b.(*StringValue); ok {
			return BoolValue(string(*av) < string(*bv))
		}
	}
	return BoolValue(false)
}

func vmGeq(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(av >= bv)
		case FloatValue:
			return BoolValue(float64(av) >= float64(bv))
		case PointerValue:
			return BoolValue(uintptr(av) >= uintptr(bv))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(float64(av) >= float64(bv))
		case FloatValue:
			return BoolValue(av >= bv)
		}
	case PointerValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(uintptr(av) >= uintptr(bv))
		case PointerValue:
			return BoolValue(uintptr(av) >= uintptr(bv))
		}
	case *StringValue:
		if bv, ok := b.(*StringValue); ok {
			return BoolValue(string(*av) >= string(*bv))
		}
	}
	return BoolValue(false)
}

func vmLeq(a, b Value) Value {
	switch av := a.(type) {
	case IntValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(av <= bv)
		case FloatValue:
			return BoolValue(float64(av) <= float64(bv))
		case PointerValue:
			return BoolValue(uintptr(av) <= uintptr(bv))
		}
	case FloatValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(float64(av) <= float64(bv))
		case FloatValue:
			return BoolValue(av <= bv)
		}
	case PointerValue:
		switch bv := b.(type) {
		case IntValue:
			return BoolValue(uintptr(av) <= uintptr(bv))
		case PointerValue:
			return BoolValue(uintptr(av) <= uintptr(bv))
		}
	case *StringValue:
		if bv, ok := b.(*StringValue); ok {
			return BoolValue(string(*av) <= string(*bv))
		}
	}
	return BoolValue(false)
}

func vmConcat(left, right Value) Value {
	switch lv := left.(type) {
	case IntValue:
		if rv, ok := right.(IntValue); ok {
			if rv < 0 {
				return null
			}
			return makeIntValue(int64(lv << uint(rv)))
		}
	case *StringValue:
		if rv, ok := right.(*StringValue); ok {
			mu := strLock(lv)
			*lv = append(*lv, *rv...)
			mu.Unlock()
			return lv
		}
		if rv, ok := right.(IntValue); ok {
			mu := strLock(lv)
			*lv = append(*lv, byte(rv))
			mu.Unlock()
			return lv
		}
	case *ListValue:
		mu := listLock(lv)
		*lv = append(*lv, right)
		mu.Unlock()
		return lv
	}
	return null
}

func vmGetProp(obj Value, key Value) (Value, string) {
	switch o := obj.(type) {
	case *StringValue:
		if idx, ok := key.(IntValue); ok {
			i := int(idx)
			if i < 0 || i >= len(*o) {
				return null, ""
			}
			return MakeSingleByteString((*o)[i]), ""
		}
		return nil, fmt.Sprintf("Cannot index into string with non-integer index %s", key)
	case *ListValue:
		if idx, ok := key.(IntValue); ok {
			i := int(idx)
			if i < 0 || i >= len(*o) {
				return null, ""
			}
			return (*o)[i], ""
		}
		return nil, fmt.Sprintf("Cannot index into list with non-integer index %s", key)
	case ObjectValue:
		var keyStr string
		switch k := key.(type) {
		case *StringValue:
			keyStr = string(*k)
		case AtomValue:
			keyStr = string(k)
		default:
			keyStr = key.String()
		}
		if v, ok := o[keyStr]; ok {
			return v, ""
		}
		return null, ""
	case ClassValue:
		var keyStr string
		switch k := key.(type) {
		case *StringValue:
			keyStr = string(*k)
		case AtomValue:
			keyStr = string(k)
		default:
			keyStr = key.String()
		}
		if v, ok := o.static[keyStr]; ok {
			return v, ""
		}
		return null, ""
	}
	return nil, fmt.Sprintf("Expected string, list, or object in left-hand side of property access, got %s", obj)
}

func vmSetProp(obj Value, key Value, val Value) {
	switch o := obj.(type) {
	case *ListValue:
		if idx, ok := key.(IntValue); ok {
			i := int(idx)
			if i >= 0 && i < len(*o) {
				mu := listLock(o)
				(*o)[i] = val
				mu.Unlock()
			}
		}
	case ObjectValue:
		var keyStr string
		switch k := key.(type) {
		case *StringValue:
			keyStr = string(*k)
		case AtomValue:
			keyStr = string(k)
		default:
			keyStr = key.String()
		}
		mu := objLock(o)
		o[keyStr] = val
		mu.Unlock()
	case *StringValue:
		if idx, ok := key.(IntValue); ok {
			i := int(idx)
			if i >= 0 && i < len(*o) {
				if bv, ok := val.(IntValue); ok {
					mu := strLock(o)
					(*o)[i] = byte(bv)
					mu.Unlock()
				}
			}
		}
	}
}

func vmSlice(collection, startVal, endVal Value) (Value, *runtimeError) {
	start := 0
	end := 0

	if sv, ok := startVal.(IntValue); ok {
		start = int(sv)
	}

	switch c := collection.(type) {
	case *StringValue:
		if ev, ok := endVal.(IntValue); ok {
			end = int(ev)
		} else {
			end = len(*c)
		}
		if start < 0 {
			start = 0
		}
		if end > len(*c) {
			end = len(*c)
		}
		if start > end {
			start = end
		}
		sliceLen := end - start
		// Fast path: single byte returns cached string
		if sliceLen == 1 {
			return MakeSingleByteString((*c)[start]), nil
		}
		// Fast path: empty slice
		if sliceLen == 0 {
			return MakeString(""), nil
		}
		result := make([]byte, sliceLen)
		copy(result, (*c)[start:end])
		sv := StringValue(result)
		return &sv, nil
	case *ListValue:
		if ev, ok := endVal.(IntValue); ok {
			end = int(ev)
		} else {
			end = len(*c)
		}
		if start < 0 {
			start = 0
		}
		if end > len(*c) {
			end = len(*c)
		}
		if start > end {
			start = end
		}
		result := make([]Value, end-start)
		copy(result, (*c)[start:end])
		lv := ListValue(result)
		return &lv, nil
	default:
		return null, nil
	}
}

func vmAppend(list, val Value) (Value, *runtimeError) {
	switch lv := list.(type) {
	case *ListValue:
		if rv, ok := val.(*ListValue); ok {
			result := make([]Value, len(*lv)+len(*rv))
			copy(result, *lv)
			copy(result[len(*lv):], *rv)
			newList := ListValue(result)
			return &newList, nil
		}
		result := make([]Value, len(*lv)+1)
		copy(result, *lv)
		result[len(*lv)] = val
		newList := ListValue(result)
		return &newList, nil
	case *StringValue:
		if rv, ok := val.(*StringValue); ok {
			result := make([]byte, len(*lv)+len(*rv))
			copy(result, *lv)
			copy(result[len(*lv):], *rv)
			sv := StringValue(result)
			return &sv, nil
		}
	}
	return null, nil
}

// closureVal implements the Value interface so it can live on the VM stack.
func (c *closureVal) String() string {
	if c.fnIdx < 0 {
		return "fn() { <closure> }"
	}
	return fmt.Sprintf("fn#%d() { <closure> }", c.fnIdx)
}

func (c *closureVal) Eq(other Value) bool {
	if _, ok := other.(EmptyValue); ok {
		return true
	}

	if oc, ok := other.(*closureVal); ok {
		return c.fnIdx == oc.fnIdx
	}
	return false
}
