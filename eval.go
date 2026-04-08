package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// byte slice helpers from the Ink interpreter source code,
// github.com/thesephist/ink

// zero-extend a slice of bytes to given length
func zeroExtend(s []byte, max int) []byte {
	if max <= len(s) {
		return s
	}

	extended := make([]byte, max)
	copy(extended, s)
	return extended
}

// return the max length of two slices
func maxLen(a, b []byte) int {
	if alen, blen := len(a), len(b); alen < blen {
		return blen
	} else {
		return alen
	}
}

type Value interface {
	String() string
	Eq(Value) bool
}

type EmptyValue byte

// interned "empty" value
const empty EmptyValue = 0

func (v EmptyValue) String() string {
	return "_"
}
func (v EmptyValue) Eq(u Value) bool {
	return true
}

// Null need not contain any data, so we use the most compact data
// representation we can.
type NullValue byte

// interned "null"
const null NullValue = 0

func (v NullValue) String() string {
	return "?"
}
func (v NullValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if _, ok := u.(NullValue); ok {
		return true
	}
	return false
}

type StringValue []byte

// MakeSingleByteString returns a new *StringValue for a single byte.
// Each call allocates a fresh value so that in-place mutation via << does not
// corrupt shared state.
func MakeSingleByteString(b byte) *StringValue {
	sv := StringValue([]byte{b})
	return &sv
}

func MakeString(s string) *StringValue {
	v := StringValue(s)
	return &v
}
func (v *StringValue) String() string {
	return fmt.Sprintf("'%s'", string(*v))
}
func (v *StringValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(*StringValue); ok {
		// Fast path: both are single-byte (dominant case in char-by-char loops)
		if len(*v) == 1 && len(*w) == 1 {
			return (*v)[0] == (*w)[0]
		}
		return bytes.Equal(*v, *w)
	}
	return false
}
func (v *StringValue) stringContent() string {
	return string(*v)
}

type IntValue int64

func (v IntValue) String() string {
	return strconv.FormatInt(int64(v), 10)
}
func (v IntValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(IntValue); ok {
		return v == w
	} else if w, ok := u.(FloatValue); ok {
		return FloatValue(v) == w
	}

	return false
}

type FloatValue float64

func (v FloatValue) String() string {
	return strconv.FormatFloat(float64(v), 'g', -1, 64)
}
func (v FloatValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(FloatValue); ok {
		return v == w
	} else if w, ok := u.(IntValue); ok {
		return v == FloatValue(w)
	}

	return false
}

// PointerValue represents a raw memory address. It is primarily used for
// interoperability with OS APIs and low-level memory operations. Unlike
// IntValue, pointers are treated as a separate type in the language to make
// intent explicit, but many arithmetic operations between ints and pointers
// are still permitted (e.g. pointer + int). Internally we just store a
// uintptr.
type PointerValue uintptr

func (v PointerValue) String() string {
	return fmt.Sprintf("ptr(0x%x)", uintptr(v))
}
func (v PointerValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	switch w := u.(type) {
	case PointerValue:
		return uintptr(v) == uintptr(w)
	case IntValue:
		return uintptr(v) == uintptr(w)
	}
	return false
}

type BoolValue bool

// interned bools
const oakTrue = BoolValue(true)
const oakFalse = BoolValue(false)

func (v BoolValue) String() string {
	if v {
		return "true"
	}
	return "false"
}
func (v BoolValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(BoolValue); ok {
		return v == w
	}

	return false
}

type AtomValue string

func (v AtomValue) String() string {
	return ":" + string(v)
}
func (v AtomValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(AtomValue); ok {
		return v == w
	}

	return false
}

type ListValue []Value

func MakeList(xs ...Value) *ListValue {
	v := ListValue(xs)
	return &v
}
func (v *ListValue) String() string {
	valStrings := make([]string, len(*v))
	for i, val := range *v {
		valStrings[i] = val.String()
	}
	return "[" + strings.Join(valStrings, ", ") + "]"
}
func (v *ListValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(*ListValue); ok {
		if len(*v) != len(*w) {
			return false
		}

		for i, el := range *v {
			if !el.Eq((*w)[i]) {
				return false
			}
		}
		return true
	}

	return false
}

type ObjectValue map[string]Value

// Striped RWMutex pool for concurrent ObjectValue map access.
// Each map is assigned a stripe based on its runtime pointer so that
// distinct objects rarely share a lock while the same object always
// maps to the same lock.
const objMuN = 32

var objMu [objMuN]sync.RWMutex

// objLock acquires an exclusive lock for the given ObjectValue and
// returns the mutex so the caller can Unlock it.
func objLock(m ObjectValue) *sync.RWMutex {
	mu := &objMu[reflect.ValueOf(m).Pointer()/8%objMuN]
	mu.Lock()
	return mu
}

// objRLock acquires a shared read lock for the given ObjectValue and
// returns the mutex so the caller can RUnlock it.
func objRLock(m ObjectValue) *sync.RWMutex {
	mu := &objMu[reflect.ValueOf(m).Pointer()/8%objMuN]
	mu.RLock()
	return mu
}

// Striped RWMutex pool for concurrent ListValue access.
const listMuN = 32

var listMu [listMuN]sync.RWMutex

// listLock acquires an exclusive lock for the given *ListValue and
// returns the mutex so the caller can Unlock it.
func listLock(l *ListValue) *sync.RWMutex {
	mu := &listMu[reflect.ValueOf(l).Pointer()/8%listMuN]
	mu.Lock()
	return mu
}

// listRLock acquires a shared read lock for the given *ListValue and
// returns the mutex so the caller can RUnlock it.
func listRLock(l *ListValue) *sync.RWMutex {
	mu := &listMu[reflect.ValueOf(l).Pointer()/8%listMuN]
	mu.RLock()
	return mu
}

// Striped RWMutex pool for concurrent StringValue access.
const strMuN = 32

var strMu [strMuN]sync.RWMutex

// strLock acquires an exclusive lock for the given *StringValue and
// returns the mutex so the caller can Unlock it.
func strLock(s *StringValue) *sync.RWMutex {
	mu := &strMu[reflect.ValueOf(s).Pointer()/8%strMuN]
	mu.Lock()
	return mu
}

// strRLock acquires a shared read lock for the given *StringValue and
// returns the mutex so the caller can RUnlock it.
func strRLock(s *StringValue) *sync.RWMutex {
	mu := &strMu[reflect.ValueOf(s).Pointer()/8%strMuN]
	mu.RLock()
	return mu
}

// only used for efficient serialization to string
type serializedObjEntry struct {
	key  string
	full string
}

func (v ObjectValue) String() string {
	return v.stringWithSeen(make(map[*ObjectValue]bool))
}

func (v ObjectValue) stringWithSeen(seen map[*ObjectValue]bool) string {
	if seen[&v] {
		return "{...}"
	}
	seen[&v] = true

	// Snapshot key-value pairs under read lock, then release before
	// recursing into child ObjectValues (avoids recursive RLock).
	mu := objRLock(v)
	type kv struct {
		k string
		v Value
	}
	pairs := make([]kv, 0, len(v))
	for key, val := range v {
		pairs = append(pairs, kv{key, val})
	}
	mu.RUnlock()

	entries := make([]serializedObjEntry, len(pairs))
	for i, p := range pairs {
		valStr := ""
		if ov, ok := p.v.(ObjectValue); ok {
			valStr = ov.stringWithSeen(seen)
		} else {
			valStr = p.v.String()
		}
		entries[i] = serializedObjEntry{
			key:  p.k,
			full: p.k + ": " + valStr,
		}
	}

	// sort entries lexicographically for easier debugging use
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].key < entries[j].key
	})

	sb := strings.Builder{}
	sb.WriteString("{")
	for i, entry := range entries {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(entry.full)
	}
	sb.WriteString("}")

	return sb.String()
}
func (v ObjectValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(ObjectValue); ok {
		muV := objRLock(v)
		vLen := len(v)
		muV.RUnlock()
		muW := objRLock(w)
		wLen := len(w)
		muW.RUnlock()
		if vLen != wLen {
			return false
		}

		// Snapshot v under read lock, then compare without holding it.
		muV = objRLock(v)
		type kv struct {
			k string
			v Value
		}
		pairs := make([]kv, 0, vLen)
		for key, val := range v {
			pairs = append(pairs, kv{key, val})
		}
		muV.RUnlock()

		for _, p := range pairs {
			muW = objRLock(w)
			wVal, ok := w[p.k]
			muW.RUnlock()
			if ok {
				if !p.v.Eq(wVal) {
					return false
				}
			} else {
				return false
			}
		}

		return true
	}

	return false
}

type FnValue struct {
	defn *fnNode
	scope
}

func (v FnValue) String() string {
	return v.defn.String()
}
func (v FnValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(FnValue); ok {
		return v.defn == w.defn
	}

	return false
}

type ClassValue struct {
	defn   *classNode
	scope  scope
	static ObjectValue
}

func (v ClassValue) String() string {
	return v.defn.String()
}
func (v ClassValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(ClassValue); ok {
		return v.defn == w.defn
	}

	return false
}

type thunkValue struct {
	defn *fnNode
	scope
}

func (v thunkValue) String() string {
	return fmt.Sprintf("thunk of fn %s: %s", v.defn.name, v.defn.body)
}
func (v thunkValue) Eq(u Value) bool {
	panic("Illegal to compare thunk values!")
}

// unwrapThunk iteratively evaluates a chain of tail-call thunks.
// Retained for reference; current eval loop inlines this logic.
var _ = (*Context).unwrapThunk //nolint:unused

func (c *Context) unwrapThunk(thunk thunkValue) (v Value, err *runtimeError) {
	for isThunk := true; isThunk; thunk, isThunk = v.(thunkValue) {
		v, err = c.evalExprWithOpt(thunk.defn.body, thunk.scope, true)
		if err != nil {
			err.stackTrace = append(err.stackTrace, stackEntry{
				name: thunk.defn.name,
				pos:  thunk.defn.pos(),
			})
			return
		}
	}

	return
}

type scope struct {
	parent *scope
	vars   map[string]Value
	mu     *sync.RWMutex
}

var muPool = sync.Pool{
	New: func() interface{} {
		return &sync.RWMutex{}
	},
}

func newScope(parent *scope) scope {
	return scope{
		parent: parent,
		vars:   make(map[string]Value),
		mu:     muPool.Get().(*sync.RWMutex),
	}
}

// newScopeN creates a scope with a pre-sized map for n expected variables.
func newScopeN(parent *scope, n int) scope {
	return scope{
		parent: parent,
		vars:   make(map[string]Value, n),
		mu:     muPool.Get().(*sync.RWMutex),
	}
}

// newScopeFast creates a scope without a mutex for single-threaded hot paths.
// When the engine is in concurrent mode (after the first go() call), it falls
// back to allocating a mutex to protect against cross-goroutine access.
func newScopeFast(parent *scope, n int) scope {
	return scope{
		parent: parent,
		vars:   make(map[string]Value, n),
	}
}

// newScopeFastConcurrent is like newScopeFast but always allocates a mutex.
// Used when the engine knows goroutines are active.
func newScopeFastConcurrent(parent *scope, n int) scope {
	return scope{
		parent: parent,
		vars:   make(map[string]Value, n),
		mu:     muPool.Get().(*sync.RWMutex),
	}
}

// bridgeVmScope converts a bytecode vmScope chain into a tree-walker scope,
// parented under the Context's top-level scope (which contains builtins and imports).
// This enables interpreter() to tree-walk closures originally compiled for the bytecode VM.
func (c *Context) bridgeVmScope(vs *vmScope) scope {
	sc := newScopeN(&c.scope, 16)
	for s := vs; s != nil; s = s.parent {
		for i, name := range s.names {
			if name != "" && i < len(s.values) && s.values[i] != nil {
				if _, ok := sc.vars[name]; !ok {
					sc.vars[name] = s.values[i]
				}
			}
		}
	}
	return sc
}

// bridgeToVmScope converts a tree-walker scope into a vmScope chain,
// enabling bytecode() to resolve upvalues from the enclosing interpreter scope.
func bridgeToVmScope(sc *scope) *vmScope {
	if sc == nil {
		return nil
	}
	// Flatten the scope chain into a single vmScope with all visible names
	var names []string
	var values []Value
	seen := make(map[string]bool)
	for s := sc; s != nil; s = s.parent {
		if s.mu != nil {
			s.mu.RLock()
		}
		for k, v := range s.vars {
			if !seen[k] {
				seen[k] = true
				names = append(names, k)
				values = append(values, v)
			}
		}
		if s.mu != nil {
			s.mu.RUnlock()
		}
	}
	return &vmScope{
		names:  names,
		values: values,
		parent: nil,
	}
}

func (sc *scope) lockRef() *sync.RWMutex {
	if sc.mu == nil {
		sc.mu = muPool.Get().(*sync.RWMutex)
	}
	return sc.mu
}

func (sc *scope) snapshotVars() map[string]Value {
	mu := sc.lockRef()
	mu.RLock()
	defer mu.RUnlock()

	clone := make(map[string]Value, len(sc.vars))
	for key, val := range sc.vars {
		clone[key] = val
	}
	return clone
}

func (sc *scope) get(name string) (Value, *runtimeError) {
	if sc.mu != nil {
		sc.mu.RLock()
		v, ok := sc.vars[name]
		sc.mu.RUnlock()
		if ok {
			return v, nil
		}
	} else {
		if v, ok := sc.vars[name]; ok {
			return v, nil
		}
	}
	if sc.parent != nil {
		return sc.parent.get(name)
	}
	return null, nil
}

func (sc *scope) put(name string, v Value) {
	if sc.mu != nil {
		sc.mu.Lock()
		sc.vars[name] = v
		sc.mu.Unlock()
	} else {
		sc.vars[name] = v
	}
}

func (sc *scope) update(name string, v Value) *runtimeError {
	if sc.mu != nil {
		sc.mu.Lock()
		if _, ok := sc.vars[name]; ok {
			sc.vars[name] = v
			sc.mu.Unlock()
			return nil
		}
		sc.mu.Unlock()
	} else {
		if _, ok := sc.vars[name]; ok {
			sc.vars[name] = v
			return nil
		}
	}
	if sc.parent != nil {
		return sc.parent.update(name, v)
	}
	return &runtimeError{
		reason: fmt.Sprintf("%s is undefined", name),
	}
}

type engine struct {
	// interpreter lock to ensure lack of data races
	sync.Mutex
	// serializes asynchronous callback evaluation spawned from builtins
	asyncEvalLock sync.Mutex
	// interpreter event loop waitgroup
	sync.WaitGroup
	// set to true when go() builtin is first used; scopes skip mutexes until then
	concurrent bool
	// for deduplicating imports
	importMap  map[string]scope
	importLock sync.RWMutex
	// file fd -> Go's File map
	fileMap map[uintptr]*os.File
	fdLock  sync.Mutex
	// stable storage for pointers created from atom names
	nameRefs    map[uintptr]*StringValue
	namePtrs    map[string]uintptr
	nameRefLock sync.Mutex
	// Go interop channels
	chanMap    map[int64]chan Value
	chanLock   sync.Mutex
	nextChanID int64
	// log async error streams through this
	reportErr func(error)
	// bytecode() builtin: cache compiled chunks keyed by *fnNode pointer
	bytecodeCache     map[*fnNode]*bytecodeChunk
	bytecodeCacheLock sync.RWMutex
}

type Context struct {
	// shared interpreter state
	eng *engine
	// directory containing the root file of this context, used for loading
	// other modules with relative paths / URLs
	rootPath string
	// current file name being executed (for error reporting)
	currentFile string
	// top level ("global") scope of this context
	scope
	// cached Oak VM for evaluation
	vm Value
}

func normalizeRootPath(rootPath string) string {
	if rootPath == "" {
		rootPath = "."
	}

	absRoot, err := filepath.Abs(rootPath)
	if err != nil {
		return filepath.Clean(rootPath)
	}

	return absRoot
}

func NewContext(rootPath string) Context {
	eng := engine{
		importMap:     map[string]scope{},
		fileMap:       map[uintptr]*os.File{},
		nameRefs:      map[uintptr]*StringValue{},
		namePtrs:      map[string]uintptr{},
		chanMap:       map[int64]chan Value{},
		bytecodeCache: map[*fnNode]*bytecodeChunk{},
		reportErr: func(err error) {
			fmt.Println(err)
		},
	}
	return Context{
		eng:      &eng,
		rootPath: normalizeRootPath(rootPath),
		scope:    newScope(nil),
		vm:       nil,
	}
}

func NewContextWithCwd() Context {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get working directory")
		os.Exit(1)
	}
	return NewContext(cwd)
}

func (c *Context) ChildContext(rootPath string) Context {
	return Context{
		eng:      c.eng,
		rootPath: normalizeRootPath(rootPath),
		scope:    newScope(nil),
	}
}

func (c *Context) subScope(parent *scope) {
	c.scope = newScope(parent)
}

func (c *Context) Lock() {
	// No-op: locking causes deadlock during recursive imports
	// c.eng.Lock()
}

func (c *Context) Unlock() {
	// No-op: locking causes deadlock during recursive imports
	// c.eng.Unlock()
}

func (c *Context) Wait() {
	c.eng.Wait()
}

type stackEntry struct {
	name string
	pos
}

func (e stackEntry) String() string {
	if e.name != "" {
		return fmt.Sprintf("  in fn %s %s", e.name, e.pos)
	}
	return fmt.Sprintf("  in anonymous fn %s", e.pos)
}

type runtimeError struct {
	reason string
	pos
	stackTrace []stackEntry
}

func (e *runtimeError) Error() string {
	trace := make([]string, len(e.stackTrace))
	for i, entry := range e.stackTrace {
		trace[i] = entry.String()
	}
	return fmt.Sprintf("Runtime error %s: %s\n%s", e.pos, e.reason, strings.Join(trace, "\n"))
}

func (c *Context) evalGo(programReader io.Reader) (Value, error) {
	c.Lock()
	defer c.Unlock()

	program, err := io.ReadAll(programReader)
	if err != nil {
		return nil, err
	}

	fileName := c.currentFile
	if fileName == "" {
		fileName = "(input)"
	}
	tokenizer := newTokenizer(string(program), fileName)
	tokens := tokenizer.tokenize()

	parser := newParser(tokens)
	nodes, err := parser.parse()
	if err != nil {
		return nil, err
	}

	nodes = optimizeAST(nodes)

	val, runtimeErr := c.evalNodes(nodes)
	if runtimeErr == nil {
		return val, nil
	}
	return val, runtimeErr

}

func (c *Context) Eval(programReader io.Reader) (Value, error) {
	c.Lock()
	defer c.Unlock()
	return c.evalGo(programReader)
}

// EvalBytecode parses, optimizes, compiles to bytecode, and executes via the
// stack-based VM instead of the tree-walking interpreter.
func (c *Context) EvalBytecode(programReader io.Reader) (Value, error) {
	c.Lock()
	defer c.Unlock()

	program, err := io.ReadAll(programReader)
	if err != nil {
		return nil, err
	}

	fileName := c.currentFile
	if fileName == "" {
		fileName = "(input)"
	}
	tokenizer := newTokenizer(string(program), fileName)
	tokens := tokenizer.tokenize()

	parser := newParser(tokens)
	nodes, parseErr := parser.parse()
	if parseErr != nil {
		return nil, parseErr
	}

	nodes = optimizeAST(nodes)

	chunk := compileToByteCode(nodes)
	vm := newVM(chunk, c)
	val, runtimeErr := vm.run()
	if runtimeErr == nil {
		return val, nil
	}
	return val, runtimeErr
}

// EvalBytecodeChunk runs a pre-compiled bytecodeChunk (e.g. loaded from a .mgb file).
func (c *Context) EvalBytecodeChunk(chunk *bytecodeChunk) (Value, error) {
	c.Lock()
	defer c.Unlock()

	vm := newVM(chunk, c)
	val, runtimeErr := vm.run()
	if runtimeErr == nil {
		return val, nil
	}
	return val, runtimeErr
}

func normalizeCallArgs(paramCount int, args []Value) []Value {
	if len(args) >= paramCount {
		return args
	}

	normalized := make([]Value, paramCount)
	copy(normalized, args)
	// remaining slots are already nil (zero value), fill with null
	for i := len(args); i < paramCount; i++ {
		normalized[i] = null
	}
	return normalized
}

func bindCallScope(parent *scope, params []string, restArg string, args []Value, concurrent bool) scope {
	n := len(params)
	if restArg != "" {
		n++
	}
	var callScope scope
	if concurrent {
		callScope = newScopeFastConcurrent(parent, n)
	} else {
		callScope = newScopeFast(parent, n)
	}

	for i, argName := range params {
		if argName != "" {
			callScope.vars[argName] = args[i]
		}
	}

	if restArg != "" {
		var restList ListValue
		if len(args) > len(params) {
			restList = ListValue(args[len(params):])
		} else {
			restList = ListValue{}
		}
		callScope.vars[restArg] = &restList
	}

	return callScope
}

func mergeObjectValue(dst ObjectValue, src ObjectValue) {
	// Snapshot src under read lock to avoid holding two locks (deadlock-safe).
	srcMu := objRLock(src)
	type kv struct {
		k string
		v Value
	}
	pairs := make([]kv, 0, len(src))
	for key, val := range src {
		pairs = append(pairs, kv{key, val})
	}
	srcMu.RUnlock()

	dstMu := objLock(dst)
	for _, p := range pairs {
		dst[p.k] = p.v
	}
	dstMu.Unlock()
}

func (c *Context) constructClassValue(class ClassValue, args ...Value) (Value, *runtimeError) {
	args = normalizeCallArgs(len(class.defn.args), args)
	constructorScope := bindCallScope(&class.scope, class.defn.args, class.defn.restArg, args, c.eng.concurrent)

	if len(class.defn.parents) == 0 {
		val, err := c.evalExpr(class.defn.body, constructorScope)
		if err != nil {
			err.stackTrace = append(err.stackTrace, stackEntry{
				name: class.defn.name,
				pos:  class.defn.pos(),
			})
		}
		return val, err
	}

	instance := ObjectValue{}
	for _, parentNode := range class.defn.parents {
		parentValue, err := c.evalExpr(parentNode, constructorScope)
		if err != nil {
			return nil, err
		}

		if parentObj, ok := parentValue.(ObjectValue); ok {
			mergeObjectValue(instance, parentObj)
			continue
		}

		parentInstance, err := c.evalFnCall(parentValue, false, args)
		if err != nil {
			err.stackTrace = append(err.stackTrace, stackEntry{
				name: class.defn.name,
				pos:  class.defn.pos(),
			})
			return nil, err
		}

		parentObj, ok := parentInstance.(ObjectValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Parent class %s must construct an object, got %s", parentNode, parentInstance),
				pos:    parentNode.pos(),
				stackTrace: []stackEntry{
					{
						name: class.defn.name,
						pos:  class.defn.pos(),
					},
				},
			}
		}

		mergeObjectValue(instance, parentObj)
	}

	ownValue, err := c.evalExpr(class.defn.body, constructorScope)
	if err != nil {
		err.stackTrace = append(err.stackTrace, stackEntry{
			name: class.defn.name,
			pos:  class.defn.pos(),
		})
		return nil, err
	}

	if ownValue == nil || ownValue.Eq(null) {
		return instance, nil
	}

	ownObj, ok := ownValue.(ObjectValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Inherited class body for %s must construct an object, got %s", class.defn.name, ownValue),
			pos:    class.defn.body.pos(),
			stackTrace: []stackEntry{
				{
					name: class.defn.name,
					pos:  class.defn.pos(),
				},
			},
		}
	}

	mergeObjectValue(instance, ownObj)
	return instance, nil
}

// EvalFnValue is the variadic convenience wrapper, used by env.go builtins.
func (c *Context) EvalFnValue(maybeFn Value, thunkable bool, args ...Value) (Value, *runtimeError) {
	return c.evalFnCall(maybeFn, thunkable, args)
}

// EvalFnValueAsync serializes callback evaluation across asynchronous Go
// goroutines to avoid concurrent writes to shared Oak object maps.
func (c *Context) EvalFnValueAsync(maybeFn Value, thunkable bool, args ...Value) (Value, *runtimeError) {
	c.eng.asyncEvalLock.Lock()
	defer c.eng.asyncEvalLock.Unlock()
	return c.evalFnCall(maybeFn, thunkable, args)
}

// EvalFnValueParallel evaluates a function without acquiring asyncEvalLock,
// allowing true parallel execution. The caller must ensure the function does
// not race on shared mutable objects. Designed for thread.parallel tasks that
// work on their own closure-captured data and return results via channels.
func (c *Context) EvalFnValueParallel(maybeFn Value, thunkable bool, args ...Value) (Value, *runtimeError) {
	return c.evalFnCall(maybeFn, thunkable, args)
}

// forkContext creates a lightweight Context that shares the engine (channels,
// imports, WaitGroup) with the parent but can evaluate code independently
// without competing for asyncEvalLock. Each forked context has its own scope
// chain starting from the parent's current scope, so variable reads are safe
// (scope.mu provides per-scope RWMutex protection) and local bindings stay
// isolated. Use for CPU-bound parallel tasks that don't mutate shared Objects
// or Lists.
func (c *Context) forkContext() *Context {
	return &Context{
		eng:         c.eng,
		rootPath:    c.rootPath,
		currentFile: c.currentFile,
		scope:       c.scope,
		vm:          nil, // forked contexts do not inherit cached VM
	}
}

// evalFnCall is the slice-based fast path for function evaluation.
func (c *Context) evalFnCall(maybeFn Value, thunkable bool, args []Value) (Value, *runtimeError) {
	if fn, ok := maybeFn.(FnValue); ok {
		args = normalizeCallArgs(len(fn.defn.args), args)
		fnScope := bindCallScope(&fn.scope, fn.defn.args, fn.defn.restArg, args, c.eng.concurrent)

		if thunkable {
			return thunkValue{
				defn:  fn.defn,
				scope: fnScope,
			}, nil
		}

		// Direct evaluation — skip thunk wrapper for non-thunkable calls
		v, err := c.evalExprWithOpt(fn.defn.body, fnScope, true)
		if err != nil {
			err.stackTrace = append(err.stackTrace, stackEntry{
				name: fn.defn.name,
				pos:  fn.defn.pos(),
			})
			return nil, err
		}
		// Unwrap any returned thunks (tail call optimization)
		for {
			thunk, isThunk := v.(thunkValue)
			if !isThunk {
				return v, nil
			}
			v, err = c.evalExprWithOpt(thunk.defn.body, thunk.scope, true)
			if err != nil {
				err.stackTrace = append(err.stackTrace, stackEntry{
					name: thunk.defn.name,
					pos:  thunk.defn.pos(),
				})
				return nil, err
			}
		}
	} else if fn, ok := maybeFn.(BuiltinFnValue); ok {
		return fn.fn(args)
	} else if class, ok := maybeFn.(ClassValue); ok {
		return c.constructClassValue(class, args...)
	} else if cv, ok := maybeFn.(*closureVal); ok {
		if cv.call != nil {
			return cv.call(args)
		}
		// closureVal with preserved AST: convert to FnValue and tree-walk
		if cv.defn != nil {
			fn := FnValue{
				defn:  cv.defn,
				scope: c.bridgeVmScope(cv.parentScope),
			}
			return c.evalFnCall(fn, thunkable, args)
		}
	}

	return nil, &runtimeError{
		reason: fmt.Sprintf("%s is not a function and cannot be called", maybeFn),
	}
}

func (c *Context) evalNodes(nodes []astNode) (Value, *runtimeError) {
	var returnVal Value = null
	var err *runtimeError
	for _, expr := range nodes {
		returnVal, err = c.evalExpr(expr, c.scope)
		if err != nil {
			return nil, err
		}
	}
	return returnVal, nil
}

var divisionByZeroErr = runtimeError{
	reason: "Division by zero",
}

func intBinaryOp(op tokKind, left, right IntValue) (Value, *runtimeError) {
	switch op {
	case plus:
		return IntValue(left + right), nil
	case minus:
		return IntValue(left - right), nil
	case times:
		return IntValue(left * right), nil
	case divide:
		if right == 0 {
			return nil, &divisionByZeroErr
		}
		if left%right == 0 {
			return IntValue(left / right), nil
		}
		return FloatValue(float64(left) / float64(right)), nil
	case modulus:
		if right == 0 {
			return nil, &divisionByZeroErr
		}
		return IntValue(left % right), nil
	case power:
		// int ** int can produce a float if exponent is negative
		if right < 0 {
			return FloatValue(math.Pow(float64(left), float64(right))), nil
		}
		// for positive exponents, try to keep as int if possible
		result := math.Pow(float64(left), float64(right))
		if result > math.MaxInt64 || result < math.MinInt64 {
			return FloatValue(result), nil
		}
		return IntValue(int64(result)), nil
	case xor:
		return IntValue(left ^ right), nil
	case and:
		return IntValue(left & right), nil
	case or:
		return IntValue(left | right), nil
	case pushArrow:
		// bitwise left shift
		if right < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Shift amount must be non-negative, got %d", right),
			}
		}
		return IntValue(left << uint(right)), nil
	case rshift:
		// bitwise right shift
		if right < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Shift amount must be non-negative, got %d", right),
			}
		}
		return IntValue(left >> uint(right)), nil
	case greater:
		return BoolValue(left > right), nil
	case less:
		return BoolValue(left < right), nil
	case geq:
		return BoolValue(left >= right), nil
	case leq:
		return BoolValue(left <= right), nil
	}
	return nil, &runtimeError{
		reason: fmt.Sprintf("Invalid binary operator %s for ints %s, %s", token{kind: op}, left, right),
	}
}

func floatBinaryOp(op tokKind, left, right FloatValue) (Value, *runtimeError) {
	switch op {
	case plus:
		return FloatValue(left + right), nil
	case minus:
		return FloatValue(left - right), nil
	case times:
		return FloatValue(left * right), nil
	case divide:
		if right == 0 {
			return nil, &divisionByZeroErr
		}
		return FloatValue(left / right), nil
	case modulus:
		if right == 0 {
			return nil, &divisionByZeroErr
		}
		return FloatValue(math.Mod(float64(left), float64(right))), nil
	case power:
		return FloatValue(math.Pow(float64(left), float64(right))), nil
	case greater:
		return BoolValue(left > right), nil
	case less:
		return BoolValue(left < right), nil
	case geq:
		return BoolValue(left >= right), nil
	case leq:
		return BoolValue(left <= right), nil
	}
	return nil, &runtimeError{
		reason: fmt.Sprintf("Invalid binary operator %s for floats %s, %s", token{kind: op}, left, right),
	}
}

func (c *Context) evalAsObjKey(node astNode, sc scope) (Value, *runtimeError) {
	if ident, ok := node.(identifierNode); ok {
		return MakeString(ident.payload), nil
	}

	return c.evalExpr(node, sc)
}

func (c *Context) evalExpr(node astNode, sc scope) (Value, *runtimeError) {
	return c.evalExprWithOpt(node, sc, false)
}

func incompatibleError(op tokKind, left, right Value, position pos) *runtimeError {
	return &runtimeError{
		reason: fmt.Sprintf("Cannot %s incompatible values %s, %s",
			token{kind: op}, left, right),
		pos: position,
	}
}

func shallowEqual(left, right Value) bool {
	switch l := left.(type) {
	case *ListValue:
		r, ok := right.(*ListValue)
		return ok && l == r
	case ObjectValue:
		r, ok := right.(ObjectValue)
		if !ok {
			return false
		}
		return reflect.ValueOf(l).Pointer() == reflect.ValueOf(r).Pointer()
	default:
		return left.Eq(right)
	}
}

func (c *Context) evalExprWithOpt(node astNode, sc scope, thunkable bool) (Value, *runtimeError) {
	switch n := node.(type) {
	case emptyNode:
		return empty, nil
	case nullNode:
		return null, nil
	case stringNode:
		payload := make([]byte, len(n.payload))
		copy(payload, n.payload)
		v := StringValue(payload)
		return &v, nil
	case intNode:
		return IntValue(n.payload), nil
	case floatNode:
		return FloatValue(n.payload), nil
	case boolNode:
		return BoolValue(n.payload), nil
	case atomNode:
		return AtomValue(n.payload), nil
	case listNode:
		var err *runtimeError
		elems := make([]Value, len(n.elems))
		for i, elNode := range n.elems {
			elems[i], err = c.evalExpr(elNode, sc)
			if err != nil {
				return nil, err
			}
		}
		list := ListValue(elems)
		return &list, nil
	case objectNode:
		obj := ObjectValue{}
		for _, entry := range n.entries {
			var keyString string

			if identKey, ok := entry.key.(identifierNode); ok {
				keyString = identKey.payload
			} else {
				key, err := c.evalExpr(entry.key, sc)
				if err != nil {
					return nil, err
				}
				switch typedKey := key.(type) {
				case *StringValue:
					keyString = string(*typedKey)
				case AtomValue:
					keyString = string(typedKey)
				case IntValue:
					keyString = typedKey.String()
				case FloatValue:
					keyString = typedKey.String()
				default:
					return nil, &runtimeError{reason: fmt.Sprintf("Expected a string, atom, or number as object key, got %s", key.String()),
						pos: entry.key.pos(),
					}
				}
			}

			val, err := c.evalExpr(entry.val, sc)
			if err != nil {
				return nil, err
			}

			obj[keyString] = val
		}
		return obj, nil
	case classNode:
		class := ClassValue{
			defn:   &n,
			scope:  sc,
			static: ObjectValue{},
		}
		sc.put(n.name, class)

		if len(n.staticExprs) != 0 {
			staticScope := newScope(&sc)
			staticScope.put(n.name, class)

			for _, expr := range n.staticExprs {
				if _, err := c.evalExpr(expr, staticScope); err != nil {
					return nil, err
				}
			}

			for key, val := range staticScope.snapshotVars() {
				if key != n.name {
					class.static[key] = val
				}
			}
		}

		sc.put(n.name, class)
		return class, nil
	case fnNode:
		fn := FnValue{
			defn:  &n,
			scope: sc,
		}
		if fn.defn.name != "" {
			sc.put(fn.defn.name, fn)
		}
		return fn, nil
	case identifierNode:
		// Inline fast path: check local scope first without method call overhead.
		// Most variable accesses hit the local scope (function params, block locals).
		if sc.mu == nil {
			if v, ok := sc.vars[n.payload]; ok {
				return v, nil
			}
		} else {
			sc.mu.RLock()
			v, ok := sc.vars[n.payload]
			sc.mu.RUnlock()
			if ok {
				return v, nil
			}
		}
		// Fall back to parent chain walk
		if sc.parent != nil {
			val, err := sc.parent.get(n.payload)
			if err != nil {
				err.pos = n.pos()
			}
			return val, err
		}
		return null, nil
	case assignmentNode:
		assignedValue, err := c.evalExpr(n.right, sc)
		if err != nil {
			return nil, err
		}

		switch left := n.left.(type) {
		case identifierNode:
			if n.isLocal {
				// Inline sc.put() to avoid method call overhead on hot path
				if sc.mu != nil {
					sc.mu.Lock()
					sc.vars[left.payload] = assignedValue
					sc.mu.Unlock()
				} else {
					sc.vars[left.payload] = assignedValue
				}
			} else {
				err := sc.update(left.payload, assignedValue)
				if err != nil {
					err.pos = n.pos()
					return nil, err
				}
			}
			return assignedValue, nil
		case listNode:
			assignedList, ok := assignedValue.(*ListValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("right side %s of list destructuring is not a list", n.right),
					pos:    n.pos(),
				}
			}

			for i, mustBeIdent := range left.elems {
				ident, ok := mustBeIdent.(identifierNode)
				if !ok {
					if _, ok = mustBeIdent.(emptyNode); ok {
						continue
					}

					return nil, &runtimeError{
						reason: fmt.Sprintf("element %s in destructured list %s is not an identifier", mustBeIdent, left),
						pos:    n.pos(),
					}
				}

				var destructuredEl Value
				if i < len(*assignedList) {
					destructuredEl = (*assignedList)[i]
				} else {
					destructuredEl = null
				}

				if n.isLocal {
					sc.put(ident.payload, destructuredEl)
				} else {
					err := sc.update(ident.payload, destructuredEl)
					if err != nil {
						return nil, err
					}
				}
			}
			return assignedValue, nil
		case objectNode:
			assignedObj, ok := assignedValue.(ObjectValue)
			if ok {
				for _, entryNode := range left.entries {
					key, err := c.evalAsObjKey(entryNode.key, sc)
					if err != nil {
						return nil, err
					}

					mustBeIdent := entryNode.val
					ident, ok := mustBeIdent.(identifierNode)
					if !ok {
						if _, ok = mustBeIdent.(emptyNode); ok {
							continue
						}

						return nil, &runtimeError{
							reason: fmt.Sprintf("value %s in destructured object %s is not an identifier", mustBeIdent, left),
							pos:    n.pos(),
						}
					}

					var keyString string
					if k, ok := key.(*StringValue); ok {
						keyString = string(*k)
					} else if k, ok := key.(AtomValue); ok {
						keyString = string(k)
					} else {
						keyString = key.String()
					}

					var destructuredEl Value
					if val, ok := assignedObj[keyString]; ok {
						destructuredEl = val
					} else {
						destructuredEl = null
					}

					if n.isLocal {
						sc.put(ident.payload, destructuredEl)
					} else {
						err := sc.update(ident.payload, destructuredEl)
						if err != nil {
							return nil, err
						}
					}
				}
				return assignedValue, nil
			}

			// Double destructuring: object pattern := list
			// Maps list elements by position to object keys, binds variables,
			// and returns the constructed object.
			assignedList, ok := assignedValue.(*ListValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("right side %s of object destructuring is not an object or list", n.right),
					pos:    n.pos(),
				}
			}

			constructedObj := ObjectValue{}
			for i, entryNode := range left.entries {
				key, err := c.evalAsObjKey(entryNode.key, sc)
				if err != nil {
					return nil, err
				}

				var keyString string
				if k, ok := key.(*StringValue); ok {
					keyString = string(*k)
				} else if k, ok := key.(AtomValue); ok {
					keyString = string(k)
				} else {
					keyString = key.String()
				}

				var destructuredEl Value
				if i < len(*assignedList) {
					destructuredEl = (*assignedList)[i]
				} else {
					destructuredEl = null
				}

				constructedObj[keyString] = destructuredEl

				mustBeIdent := entryNode.val
				ident, ok := mustBeIdent.(identifierNode)
				if !ok {
					if _, ok = mustBeIdent.(emptyNode); ok {
						continue
					}

					return nil, &runtimeError{
						reason: fmt.Sprintf("value %s in destructured object %s is not an identifier", mustBeIdent, left),
						pos:    n.pos(),
					}
				}

				if n.isLocal {
					sc.put(ident.payload, destructuredEl)
				} else {
					err := sc.update(ident.payload, destructuredEl)
					if err != nil {
						return nil, err
					}
				}
			}
			return constructedObj, nil
		case propertyAccessNode:
			assign := left

			assignLeft, err := c.evalExpr(assign.left, sc)
			if err != nil {
				return nil, err
			}

			// Fast path: obj.identifier assignment — avoid MakeString allocation
			if ident, isIdent := assign.right.(identifierNode); isIdent {
				switch target := assignLeft.(type) {
				case ObjectValue:
					mu := objLock(target)
					if _, ok := assignedValue.(EmptyValue); ok {
						delete(target, ident.payload)
					} else {
						target[ident.payload] = assignedValue
					}
					mu.Unlock()
					return assignLeft, nil
				case ClassValue:
					mu := objLock(target.static)
					if _, ok := assignedValue.(EmptyValue); ok {
						delete(target.static, ident.payload)
					} else {
						target.static[ident.payload] = assignedValue
					}
					mu.Unlock()
					return assignLeft, nil
				}
			}

			assignRight, err := c.evalAsObjKey(assign.right, sc)
			if err != nil {
				return nil, err
			}

			switch target := assignLeft.(type) {
			case *StringValue:
				assignedString, ok := assignedValue.(*StringValue)
				if !ok {
					return nil, &runtimeError{
						reason: fmt.Sprintf("Cannot assign non-string value %s to string in %s", assignedValue, assign),
						pos:    n.pos(),
					}
				}

				byteIndexVal, ok := assignRight.(IntValue)
				if !ok {
					return nil, &runtimeError{
						reason: fmt.Sprintf("Cannot index into string with non-integer index %s", assignRight),
						pos:    n.pos(),
					}
				}
				byteIndex := int(byteIndexVal)

				if byteIndex < 0 || byteIndex > len(*target) {
					return nil, &runtimeError{
						reason: fmt.Sprintf("String assignment index %d out of range in %s", byteIndex, n),
						pos:    n.pos(),
					}
				}

				mu := strLock(target)
				if byteIndex == len(*target) {
					// append
					*target = append(*target, *assignedString...)
				} else {
					for byteOffset, byteAtOffset := range *assignedString {
						if byteIndex+byteOffset < len(*target) {
							(*target)[byteIndex+byteOffset] = byteAtOffset
						} else {
							*target = append(*target, byteAtOffset)
						}
					}
				}
				mu.Unlock()
			case *ListValue:
				listIndexVal, ok := assignRight.(IntValue)
				if !ok {
					return nil, &runtimeError{
						reason: fmt.Sprintf("Cannot index into list with non-integer index %s", assignRight),
						pos:    n.pos(),
					}
				}
				listIndex := int(listIndexVal)

				if listIndex < 0 || listIndex > len(*target) {
					return nil, &runtimeError{
						reason: fmt.Sprintf("List assignment index %d out of range in %s", listIndex, n),
						pos:    n.pos(),
					}
				}

				mu := listLock(target)
				if listIndex == len(*target) {
					*target = append(*target, assignedValue)
				} else {
					(*target)[listIndex] = assignedValue
				}
				mu.Unlock()
			case ObjectValue:
				var objKeyString string
				if objKey, ok := assignRight.(*StringValue); ok {
					objKeyString = string(*objKey)
				} else if objKey, ok := assignRight.(AtomValue); ok {
					objKeyString = string(objKey)
				} else {
					objKeyString = assignRight.String()
				}

				mu := objLock(target)
				if _, ok := assignedValue.(EmptyValue); ok {
					delete(target, objKeyString)
				} else {
					target[objKeyString] = assignedValue
				}
				mu.Unlock()
			case ClassValue:
				var objKeyString string
				if objKey, ok := assignRight.(*StringValue); ok {
					objKeyString = string(*objKey)
				} else if objKey, ok := assignRight.(AtomValue); ok {
					objKeyString = string(objKey)
				} else {
					objKeyString = assignRight.String()
				}

				mu := objLock(target.static)
				if _, ok := assignedValue.(EmptyValue); ok {
					delete(target.static, objKeyString)
				} else {
					target.static[objKeyString] = assignedValue
				}
				mu.Unlock()
			case NullValue:
				return nil, &runtimeError{
					reason: fmt.Sprintf("Cannot assign to property of undefined value in %s", left.String()),
					pos:    n.pos(),
				}
			default:
				return nil, &runtimeError{
					reason: fmt.Sprintf("Expected string, list, or object in left-hand side of property assignment, got %s", left.String()),
					pos:    n.pos(),
				}
			}

			return assignLeft, nil
		default:
			return nil, &runtimeError{
				reason: fmt.Sprintf("Invalid assignment target %s", left.String()),
				pos:    n.pos(),
			}
		}
	case propertyAccessNode:
		left, err := c.evalExpr(n.left, sc)
		if err != nil {
			return nil, err
		}

		// Fast path: obj.identifier — the most common property access pattern.
		// Avoids MakeString allocation and type-switch round-trip.
		if ident, isIdent := n.right.(identifierNode); isIdent {
			switch target := left.(type) {
			case ObjectValue:
				mu := objRLock(target)
				val, ok := target[ident.payload]
				mu.RUnlock()
				if ok {
					return val, nil
				}
				return null, nil
			case ClassValue:
				mu := objRLock(target.static)
				val, ok := target.static[ident.payload]
				mu.RUnlock()
				if ok {
					return val, nil
				}
				return null, nil
			}
		}

		right, err := c.evalAsObjKey(n.right, sc)
		if err != nil {
			return nil, err
		}

		switch target := left.(type) {
		case *StringValue:
			byteIndex, ok := right.(IntValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("Cannot index into string with non-integer index %s", right),
					pos:    n.pos(),
				}
			}

			if byteIndex < 0 || int64(byteIndex) >= int64(len(*target)) {
				return null, nil
			}

			return MakeSingleByteString((*target)[byteIndex]), nil
		case *ListValue:
			listIndex, ok := right.(IntValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("Cannot index into list with non-integer index %s", right),
					pos:    n.pos(),
				}
			}

			if listIndex < 0 || int64(listIndex) >= int64(len(*target)) {
				return null, nil
			}

			return (*target)[listIndex], nil
		case ObjectValue:
			var objKeyString string
			if objKey, ok := right.(*StringValue); ok {
				objKeyString = string(*objKey)
			} else if objKey, ok := right.(AtomValue); ok {
				objKeyString = string(objKey)
			} else {
				objKeyString = right.String()
			}

			mu := objRLock(target)
			val, ok := target[objKeyString]
			mu.RUnlock()
			if ok {
				return val, nil
			}

			return null, nil
		case ClassValue:
			var objKeyString string
			if objKey, ok := right.(*StringValue); ok {
				objKeyString = string(*objKey)
			} else if objKey, ok := right.(AtomValue); ok {
				objKeyString = string(objKey)
			} else {
				objKeyString = right.String()
			}

			mu := objRLock(target.static)
			val, ok := target.static[objKeyString]
			mu.RUnlock()
			if ok {
				return val, nil
			}

			return null, nil
		}

		return nil, &runtimeError{
			reason: fmt.Sprintf("Expected string, list, or object in left-hand side of property access, got %s", left.String()),
			pos:    n.pos(),
		}
	case unaryNode:
		rightComputed, err := c.evalExpr(n.right, sc)
		if err != nil {
			return nil, err
		}

		switch right := rightComputed.(type) {
		case IntValue:
			switch n.op {
			case plus:
				return right, nil
			case minus:
				return -right, nil
			case tilde:
				return ^right, nil
			}
		case FloatValue:
			switch n.op {
			case plus:
				return right, nil
			case minus:
				return -right, nil
			}
		case BoolValue:
			switch n.op {
			case exclam:
				return !right, nil
			}
		}
		return nil, &runtimeError{
			reason: fmt.Sprintf("%s is not a valid unary operator for %s", token{kind: n.op}, rightComputed),
			pos:    n.pos(),
		}
	case binaryNode:
		leftComputed, err := c.evalExpr(n.left, sc)
		if err != nil {
			return nil, err
		}

		// short-circuit boolean comparisons
		if leftBool, ok := leftComputed.(BoolValue); ok {
			if leftBool && (n.op == or || n.op == plus) ||
				!leftBool && (n.op == and || n.op == times) {
				return leftBool, nil
			}
		}

		rightComputed, err := c.evalExpr(n.right, sc)
		if err != nil {
			return nil, err
		}

		switch n.op {
		case eq:
			return BoolValue(shallowEqual(leftComputed, rightComputed)), nil
		case deepEq:
			return BoolValue(leftComputed.Eq(rightComputed)), nil
		case neq:
			// Fast path: int != int (very common in loop guards)
			if li, ok := leftComputed.(IntValue); ok {
				if ri, ok := rightComputed.(IntValue); ok {
					return BoolValue(li != ri), nil
				}
			}
			return BoolValue(!leftComputed.Eq(rightComputed)), nil
		}

		// Fast path: int op int (dominant case in arithmetic-heavy code)
		if leftInt, ok := leftComputed.(IntValue); ok {
			if rightInt, ok := rightComputed.(IntValue); ok {
				switch n.op {
				case plus:
					return IntValue(leftInt + rightInt), nil
				case minus:
					return IntValue(leftInt - rightInt), nil
				case times:
					return IntValue(leftInt * rightInt), nil
				case less:
					return BoolValue(leftInt < rightInt), nil
				case greater:
					return BoolValue(leftInt > rightInt), nil
				case leq:
					return BoolValue(leftInt <= rightInt), nil
				case geq:
					return BoolValue(leftInt >= rightInt), nil
				case divide:
					if rightInt == 0 {
						return nil, &runtimeError{reason: "Division by zero", pos: n.pos()}
					}
					if leftInt%rightInt == 0 {
						return IntValue(leftInt / rightInt), nil
					}
					return FloatValue(float64(leftInt) / float64(rightInt)), nil
				case modulus:
					if rightInt == 0 {
						return nil, &runtimeError{reason: "Division by zero", pos: n.pos()}
					}
					return IntValue(leftInt % rightInt), nil
				default:
					val, err := intBinaryOp(n.op, leftInt, rightInt)
					if err != nil {
						err.pos = n.pos()
					}
					return val, err
				}
			}
		}

		// Fast path: float op float
		if leftFloat, ok := leftComputed.(FloatValue); ok {
			if rightFloat, ok := rightComputed.(FloatValue); ok {
				switch n.op {
				case plus:
					return FloatValue(leftFloat + rightFloat), nil
				case minus:
					return FloatValue(leftFloat - rightFloat), nil
				case times:
					return FloatValue(leftFloat * rightFloat), nil
				case divide:
					if rightFloat == 0 {
						return nil, &runtimeError{reason: "Division by zero", pos: n.pos()}
					}
					return FloatValue(leftFloat / rightFloat), nil
				case less:
					return BoolValue(leftFloat < rightFloat), nil
				case greater:
					return BoolValue(leftFloat > rightFloat), nil
				case leq:
					return BoolValue(leftFloat <= rightFloat), nil
				case geq:
					return BoolValue(leftFloat >= rightFloat), nil
				default:
					val, err := floatBinaryOp(n.op, leftFloat, rightFloat)
					if err != nil {
						err.pos = n.pos()
					}
					return val, err
				}
			}
		}

		// Fast path: string + string (common in template/concat-heavy code)
		if n.op == plus {
			if leftStr, ok := leftComputed.(*StringValue); ok {
				if rightStr, ok := rightComputed.(*StringValue); ok {
					base := make([]byte, 0, len(*leftStr)+len(*rightStr))
					base = append(base, *leftStr...)
					base = append(base, *rightStr...)
					baseStr := StringValue(base)
					return &baseStr, nil
				}
			}
		}

		switch left := leftComputed.(type) {
		case IntValue:
			// if the right side is a pointer we can perform pointer arithmetic or
			// comparisons; for simplicity we treat `int + ptr` as `ptr + int` and
			// `int - ptr` as `ptr - int`.
			if ptr, ok := rightComputed.(PointerValue); ok {
				switch n.op {
				case plus:
					return PointerValue(uintptr(ptr) + uintptr(left)), nil
				case minus:
					return PointerValue(uintptr(ptr) - uintptr(left)), nil
				case greater:
					return BoolValue(uintptr(left) > uintptr(ptr)), nil
				case less:
					return BoolValue(uintptr(left) < uintptr(ptr)), nil
				case geq:
					return BoolValue(uintptr(left) >= uintptr(ptr)), nil
				case leq:
					return BoolValue(uintptr(left) <= uintptr(ptr)), nil
				default:
					return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
				}
			}

			right, ok := rightComputed.(IntValue)
			if !ok {
				rightFloat, ok := rightComputed.(FloatValue)
				if !ok {
					return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
				}

				leftFloat := FloatValue(float64(int64(left)))
				val, err := floatBinaryOp(n.op, leftFloat, rightFloat)
				if err != nil {
					err.pos = n.pos()
				}
				return val, err
			}

			val, err := intBinaryOp(n.op, left, right)
			if err != nil {
				err.pos = n.pos()
			}
			return val, err
		case FloatValue:
			right, ok := rightComputed.(FloatValue)
			if !ok {
				rightInt, ok := rightComputed.(IntValue)
				if !ok {
					return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
				}

				right = FloatValue(float64(int64(rightInt)))
				val, err := floatBinaryOp(n.op, left, right)
				if err != nil {
					err.pos = n.pos()
				}
				return val, err
			}

			val, err := floatBinaryOp(n.op, left, right)
			if err != nil {
				err.pos = n.pos()
			}
			return val, err
		case *StringValue:
			right, ok := rightComputed.(*StringValue)
			if !ok {
				return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
			}

			switch n.op {
			case plus:
				base := make([]byte, 0, len(*left)+len(*right))
				base = append(base, *left...)
				base = append(base, *right...)
				baseStr := StringValue(base)
				return &baseStr, nil
			case xor:
				max := maxLen(*left, *right)

				ls, rs := zeroExtend(*left, max), zeroExtend(*right, max)
				res := make([]byte, max)
				for i := range res {
					res[i] = ls[i] ^ rs[i]
				}
				resStr := StringValue(res)
				return &resStr, nil
			case and:
				max := maxLen(*left, *right)

				ls, rs := zeroExtend(*left, max), zeroExtend(*right, max)
				res := make([]byte, max)
				for i := range res {
					res[i] = ls[i] & rs[i]
				}
				resStr := StringValue(res)
				return &resStr, nil
			case or:
				max := maxLen(*left, *right)

				ls, rs := zeroExtend(*left, max), zeroExtend(*right, max)
				res := make([]byte, max)
				for i := range res {
					res[i] = ls[i] | rs[i]
				}
				resStr := StringValue(res)
				return &resStr, nil
			case pushArrow:
				mu := strLock(left)
				*left = append(*left, *right...)
				mu.Unlock()
				return left, nil
			case greater:
				return BoolValue(bytes.Compare(*left, *right) > 0), nil
			case less:
				return BoolValue(bytes.Compare(*left, *right) < 0), nil
			case geq:
				return BoolValue(bytes.Compare(*left, *right) >= 0), nil
			case leq:
				return BoolValue(bytes.Compare(*left, *right) <= 0), nil
			}
			return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
		case BoolValue:
			right, ok := rightComputed.(BoolValue)
			if !ok {
				return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
			}

			switch n.op {
			case plus, or:
				return BoolValue(left || right), nil
			case times, and:
				return BoolValue(left && right), nil
			case xor:
				return BoolValue(left != right), nil
			}
		case *ListValue:
			switch n.op {
			case pushArrow:
				mu := listLock(left)
				*left = append(*left, rightComputed)
				mu.Unlock()
				return left, nil
			}
			return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
		case PointerValue:
			// pointer arithmetic and comparisons with integers or other pointers
			switch r := rightComputed.(type) {
			case IntValue:
				switch n.op {
				case plus:
					return PointerValue(uintptr(left) + uintptr(r)), nil
				case minus:
					return PointerValue(uintptr(left) - uintptr(r)), nil
				case greater:
					return BoolValue(uintptr(left) > uintptr(r)), nil
				case less:
					return BoolValue(uintptr(left) < uintptr(r)), nil
				case geq:
					return BoolValue(uintptr(left) >= uintptr(r)), nil
				case leq:
					return BoolValue(uintptr(left) <= uintptr(r)), nil
				default:
					return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
				}
			case PointerValue:
				switch n.op {
				case minus:
					return IntValue(int64(uintptr(left) - uintptr(r))), nil
				case greater:
					return BoolValue(uintptr(left) > uintptr(r)), nil
				case less:
					return BoolValue(uintptr(left) < uintptr(r)), nil
				case geq:
					return BoolValue(uintptr(left) >= uintptr(r)), nil
				case leq:
					return BoolValue(uintptr(left) <= uintptr(r)), nil
				default:
					return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
				}
			default:
				return nil, incompatibleError(n.op, leftComputed, rightComputed, n.pos())
			}
		}
		return nil, &runtimeError{
			reason: fmt.Sprintf("Binary operator %s is not defined for values %s, %s",
				token{kind: n.op}, leftComputed, rightComputed),
			pos: n.pos(),
		}
	case fnCallNode:
		// Super-fast path: identifier function call — inline scope lookup
		// to avoid the full evalExpr type switch for the function itself.
		var maybeFn Value
		if fnIdent, isFnIdent := n.fn.(identifierNode); isFnIdent {
			// Inline scope.get for the function identifier
			if sc.mu == nil {
				if v, ok := sc.vars[fnIdent.payload]; ok {
					maybeFn = v
				}
			} else {
				sc.mu.RLock()
				v, ok := sc.vars[fnIdent.payload]
				sc.mu.RUnlock()
				if ok {
					maybeFn = v
				}
			}
			if maybeFn == nil && sc.parent != nil {
				maybeFn, _ = sc.parent.get(fnIdent.payload)
			}
			if maybeFn == nil {
				maybeFn = null
			}
		} else {
			var err *runtimeError
			maybeFn, err = c.evalExpr(n.fn, sc)
			if err != nil {
				return nil, err
			}
		}

		// Fast path: direct FnValue call with exact args and no rest/spread
		if fn, isFn := maybeFn.(FnValue); isFn && n.restArg == nil &&
			fn.defn.restArg == "" && len(n.args) == len(fn.defn.args) {
			nParams := len(fn.defn.args)
			fnScope := newScopeFast(&fn.scope, nParams)
			for i, argNode := range n.args {
				argVal, err := c.evalExpr(argNode, sc)
				if err != nil {
					return nil, err
				}
				if fn.defn.args[i] != "" {
					fnScope.vars[fn.defn.args[i]] = argVal
				}
			}
			if thunkable {
				return thunkValue{defn: fn.defn, scope: fnScope}, nil
			}
			v, err := c.evalExprWithOpt(fn.defn.body, fnScope, true)
			if err != nil {
				err.stackTrace = append(err.stackTrace, stackEntry{
					name: fn.defn.name, pos: fn.defn.pos(),
				})
				return nil, err
			}
			for {
				thunk, isThunk := v.(thunkValue)
				if !isThunk {
					return v, nil
				}
				v, err = c.evalExprWithOpt(thunk.defn.body, thunk.scope, true)
				if err != nil {
					err.stackTrace = append(err.stackTrace, stackEntry{
						name: thunk.defn.name, pos: thunk.defn.pos(),
					})
					return nil, err
				}
			}
		}

		args := make([]Value, len(n.args))
		for i, argNode := range n.args {
			var err *runtimeError
			args[i], err = c.evalExpr(argNode, sc)
			if err != nil {
				return nil, err
			}
		}
		if n.restArg != nil {
			rest, err := c.evalExpr(n.restArg, sc)
			if err != nil {
				return nil, err
			}

			restList, ok := rest.(*ListValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("Cannot spread a non-list value %s in a function call %s", rest, n),
					pos:    n.pos(),
				}
			}

			args = append(args, *restList...)
		}

		val, err := c.evalFnCall(maybeFn, thunkable, args)
		// we only overwrite the error pos if it's nil (i.e. if it was a "nil
		// is not a function" error, where EvalFnValue can't correctly position
		// the error itself)
		if err != nil && err.pos.line == 0 {
			err.pos = n.pos()
		}
		return val, err
	case ifExprNode:
		cond, err := c.evalExpr(n.cond, sc)
		if err != nil {
			return nil, err
		}

		// Fast paths for common condition types to avoid Eq() interface dispatch
		switch cv := cond.(type) {
		case IntValue:
			for _, branch := range n.branches {
				target, err := c.evalExpr(branch.target, sc)
				if err != nil {
					return nil, err
				}
				if _, isEmpty := target.(EmptyValue); isEmpty {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
				if tv, ok := target.(IntValue); ok && cv == tv {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
			}
		case AtomValue:
			for _, branch := range n.branches {
				target, err := c.evalExpr(branch.target, sc)
				if err != nil {
					return nil, err
				}
				if _, isEmpty := target.(EmptyValue); isEmpty {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
				if tv, ok := target.(AtomValue); ok && cv == tv {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
			}
		case BoolValue:
			for _, branch := range n.branches {
				target, err := c.evalExpr(branch.target, sc)
				if err != nil {
					return nil, err
				}
				if _, isEmpty := target.(EmptyValue); isEmpty {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
				if tv, ok := target.(BoolValue); ok && cv == tv {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
			}
		default:
			for _, branch := range n.branches {
				target, err := c.evalExpr(branch.target, sc)
				if err != nil {
					return nil, err
				}
				if _, isEmpty := target.(EmptyValue); isEmpty {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
				if cond.Eq(target) {
					return c.evalExprWithOpt(branch.body, sc, thunkable)
				}
			}
		}
		return null, nil
	case blockNode:
		// empty block returns ? (null)
		if len(n.exprs) == 0 {
			return null, nil
		}

		// Fast path: blocks without local assignments don't need a new scope.
		// This avoids a map allocation + mutex pool acquisition per block.
		blockScope := sc
		if n.hasLocal {
			blockScope = newScope(&sc)
		}

		last := len(n.exprs) - 1
		for _, expr := range n.exprs[:last] {
			_, err := c.evalExprWithOpt(expr, blockScope, false)
			if err != nil {
				return nil, err
			}
		}

		return c.evalExprWithOpt(n.exprs[last], blockScope, thunkable)
	}

	panic(fmt.Sprintf("Unexpected astNode type: %s", node))
}
