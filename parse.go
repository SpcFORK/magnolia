package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type astNode interface {
	String() string
	pos() pos
}

func parseIntLiteral(payload string) (int64, error) {
	base := 10
	if strings.HasPrefix(payload, "0x") || strings.HasPrefix(payload, "0X") {
		base = 16
		payload = payload[2:]
	} else if strings.HasPrefix(payload, "0b") || strings.HasPrefix(payload, "0B") {
		base = 2
		payload = payload[2:]
	}

	return strconv.ParseInt(payload, base, 64)
}

type emptyNode struct {
	tok *token
}

func (n emptyNode) String() string {
	return "_"
}
func (n emptyNode) pos() pos {
	return n.tok.pos
}

type nullNode struct {
	tok *token
}

func (n nullNode) String() string {
	return "?"
}
func (n nullNode) pos() pos {
	return n.tok.pos
}

type stringNode struct {
	payload []byte
	tok     *token
}

func (n stringNode) String() string {
	return fmt.Sprintf("%s", strconv.Quote(string(n.payload)))
}
func (n stringNode) pos() pos {
	return n.tok.pos
}

type intNode struct {
	payload int64
	tok     *token
}

func (n intNode) String() string {
	return strconv.FormatInt(n.payload, 10)
}
func (n intNode) pos() pos {
	return n.tok.pos
}

type floatNode struct {
	payload float64
	tok     *token
}

func (n floatNode) String() string {
	return strconv.FormatFloat(n.payload, 'g', -1, 64)
}
func (n floatNode) pos() pos {
	return n.tok.pos
}

type boolNode struct {
	payload bool
	tok     *token
}

func (n boolNode) String() string {
	if n.payload {
		return "true"
	}
	return "false"
}
func (n boolNode) pos() pos {
	return n.tok.pos
}

type atomNode struct {
	payload string
	tok     *token
}

func (n atomNode) String() string {
	return ":" + n.payload
}
func (n atomNode) pos() pos {
	return n.tok.pos
}

type listNode struct {
	elems []astNode
	tok   *token
}

func (n listNode) String() string {
	elemStrings := make([]string, len(n.elems))
	for i, el := range n.elems {
		elemStrings[i] = el.String()
	}
	return "[" + strings.Join(elemStrings, ", ") + "]"
}
func (n listNode) pos() pos {
	return n.tok.pos
}

type objectEntry struct {
	key astNode
	val astNode
}

func (n objectEntry) String() string {
	return n.key.String() + ": " + n.val.String()
}

type objectNode struct {
	entries []objectEntry
	tok     *token
}

func (n objectNode) String() string {
	entryStrings := make([]string, len(n.entries))
	for i, ent := range n.entries {
		entryStrings[i] = ent.String()
	}
	return "{ " + strings.Join(entryStrings, ", ") + " }"
}
func (n objectNode) pos() pos {
	return n.tok.pos
}

type fnNode struct {
	name    string
	args    []string
	restArg string
	body    astNode
	tok     *token
}

func (n fnNode) String() string {
	var head string
	if n.name == "" {
		head = "fn"
	} else {
		head = "fn " + n.name
	}

	argStrings := make([]string, len(n.args))
	copy(argStrings, n.args)
	if n.restArg != "" {
		argStrings = append(argStrings, n.restArg+"...")
	}
	head += "(" + strings.Join(argStrings, ", ") + ")"

	return head + " " + n.body.String()
}
func (n fnNode) pos() pos {
	return n.tok.pos
}

type classNode struct {
	name        string
	args        []string
	restArg     string
	body        astNode
	staticExprs []astNode
	parents     []astNode
	tok         *token
}

func (n classNode) String() string {
	argStrings := make([]string, len(n.args))
	copy(argStrings, n.args)
	if n.restArg != "" {
		argStrings = append(argStrings, n.restArg+"...")
	}

	head := "cs " + n.name + "(" + strings.Join(argStrings, ", ") + ")"
	if len(n.parents) == 0 {
		return head + " " + n.body.String()
	}

	staticStrings := make([]string, len(n.staticExprs))
	for i, expr := range n.staticExprs {
		staticStrings[i] = expr.String()
	}

	parentStrings := make([]string, len(n.parents))
	for i, parent := range n.parents {
		parentStrings[i] = parent.String()
	}

	bodyParts := staticStrings
	bodyParts = append(bodyParts, "("+strings.Join(parentStrings, ", ")+") -> "+n.body.String())
	return head + " { " + strings.Join(bodyParts, ", ") + " }"
}
func (n classNode) pos() pos {
	return n.tok.pos
}

type identifierNode struct {
	payload string
	tok     *token
}

func (n identifierNode) String() string {
	return n.payload
}
func (n identifierNode) pos() pos {
	return n.tok.pos
}

type assignmentNode struct {
	isLocal bool
	left    astNode
	right   astNode
	tok     *token
}

func (n assignmentNode) String() string {
	if n.isLocal {
		return n.left.String() + " := " + n.right.String()
	}
	return n.left.String() + " <- " + n.right.String()
}
func (n assignmentNode) pos() pos {
	return n.tok.pos
}

type propertyAccessNode struct {
	left  astNode
	right astNode
	tok   *token
}

func (n propertyAccessNode) String() string {
	return "(" + n.left.String() + "." + n.right.String() + ")"
}
func (n propertyAccessNode) pos() pos {
	return n.tok.pos
}

type unaryNode struct {
	op    tokKind
	right astNode
	tok   *token
}

func (n unaryNode) String() string {
	opTok := token{kind: n.op}
	return opTok.String() + n.right.String()
}
func (n unaryNode) pos() pos {
	return n.tok.pos
}

type binaryNode struct {
	op    tokKind
	left  astNode
	right astNode
	tok   *token
}

func (n binaryNode) String() string {
	opTok := token{kind: n.op}
	return "(" + n.left.String() + " " + opTok.String() + " " + n.right.String() + ")"
}
func (n binaryNode) pos() pos {
	return n.tok.pos
}

type fnCallNode struct {
	fn      astNode
	args    []astNode
	restArg astNode
	tok     *token
}

func (n fnCallNode) String() string {
	argStrings := make([]string, len(n.args))
	for i, arg := range n.args {
		argStrings[i] = arg.String()
	}
	if n.restArg != nil {
		argStrings = append(argStrings, n.restArg.String()+"...")
	}
	return fmt.Sprintf("call[%s](%s)", n.fn, strings.Join(argStrings, ", "))
}
func (n fnCallNode) pos() pos {
	return n.tok.pos
}

type ifBranch struct {
	target astNode
	body   astNode
}

func (n ifBranch) String() string {
	return n.target.String() + " -> " + n.body.String()
}

type ifExprNode struct {
	cond     astNode
	branches []ifBranch
	tok      *token
}

func (n ifExprNode) String() string {
	branchStrings := make([]string, len(n.branches))
	for i, br := range n.branches {
		branchStrings[i] = br.String()
	}
	return "if " + n.cond.String() + " {" + strings.Join(branchStrings, ", ") + "}"
}
func (n ifExprNode) pos() pos {
	return n.tok.pos
}

type blockNode struct {
	exprs []astNode
	tok   *token
}

func (n blockNode) String() string {
	exprStrings := make([]string, len(n.exprs))
	for i, ex := range n.exprs {
		exprStrings[i] = ex.String()
	}
	return "{ " + strings.Join(exprStrings, ", ") + " }"
}

func (n blockNode) pos() pos {
	return n.tok.pos
}

type parser struct {
	tokens        []token
	index         int
	minBinaryPrec []int
}

func newParser(tokens []token) parser {
	return parser{
		tokens:        tokens,
		index:         0,
		minBinaryPrec: []int{0},
	}
}

func (p *parser) lastMinPrec() int {
	return p.minBinaryPrec[len(p.minBinaryPrec)-1]
}

func (p *parser) pushMinPrec(prec int) {
	p.minBinaryPrec = append(p.minBinaryPrec, prec)
}

func (p *parser) popMinPrec() {
	p.minBinaryPrec = p.minBinaryPrec[:len(p.minBinaryPrec)-1]
}

func (p *parser) isEOF() bool {
	return p.index == len(p.tokens)
}

func (p *parser) peek() token {
	return p.tokens[p.index]
}

func (p *parser) peekAhead(n int) token {
	if p.index+n > len(p.tokens) {
		// Use comma as "nothing is here" value
		return token{kind: comma}
	}
	return p.tokens[p.index+n]
}

func (p *parser) next() token {
	tok := p.tokens[p.index]

	if p.index < len(p.tokens) {
		p.index++
	}

	return tok
}

func (p *parser) back() {
	if p.index > 0 {
		p.index--
	}
}

func (p *parser) expect(kind tokKind) (token, error) {
	tok := token{kind: kind}

	if p.isEOF() {
		return token{kind: unknown}, parseError{
			reason: fmt.Sprintf("Unexpected end of input, expected %s", tok),
			pos:    tok.pos,
		}
	}

	next := p.next()
	if next.kind != kind {
		return token{kind: unknown}, parseError{
			reason: fmt.Sprintf("Unexpected token %s, expected %s", next, tok),
			pos:    next.pos,
		}
	}

	return next, nil
}

func (p *parser) readUntilTokenKind(kind tokKind) []token {
	tokens := []token{}
	for !p.isEOF() && p.peek().kind != kind {
		tokens = append(tokens, p.next())
	}
	return tokens
}

func (p *parser) restore(index int, minPrecLen int) {
	p.index = index
	p.minBinaryPrec = p.minBinaryPrec[:minPrecLen]
}

func cloneNameSet(src map[string]struct{}) map[string]struct{} {
	clone := make(map[string]struct{}, len(src))
	for key := range src {
		clone[key] = struct{}{}
	}
	return clone
}

func addPatternBindings(dst map[string]struct{}, node astNode) {
	switch n := node.(type) {
	case identifierNode:
		if n.payload != "" {
			dst[n.payload] = struct{}{}
		}
	case listNode:
		for _, elem := range n.elems {
			addPatternBindings(dst, elem)
		}
	case objectNode:
		for _, entry := range n.entries {
			addPatternBindings(dst, entry.val)
		}
	}
}

func makeClassSugarSelfNode(name string, tok *token) identifierNode {
	return identifierNode{payload: name, tok: tok}
}

func rewriteClassSugarAssignmentLeft(node astNode, visibleFields, allFields, shadowed map[string]struct{}, selfName string, isLocal bool) astNode {
	switch n := node.(type) {
	case identifierNode:
		if isLocal {
			return n
		}
		if _, ok := shadowed[n.payload]; ok {
			return n
		}
		if n.payload == "Self" {
			return identifierNode{payload: selfName, tok: n.tok}
		}
		if _, ok := visibleFields[n.payload]; ok {
			return propertyAccessNode{
				left:  makeClassSugarSelfNode(selfName, n.tok),
				right: n,
				tok:   n.tok,
			}
		}
		return n
	case listNode:
		elems := make([]astNode, len(n.elems))
		for i, elem := range n.elems {
			elems[i] = rewriteClassSugarAssignmentLeft(elem, visibleFields, allFields, shadowed, selfName, isLocal)
		}
		return listNode{elems: elems, tok: n.tok}
	case objectNode:
		entries := make([]objectEntry, len(n.entries))
		for i, entry := range n.entries {
			key := entry.key
			if _, ok := key.(identifierNode); !ok {
				key = rewriteClassSugarNode(key, visibleFields, allFields, shadowed, selfName)
			}
			entries[i] = objectEntry{
				key: key,
				val: rewriteClassSugarAssignmentLeft(entry.val, visibleFields, allFields, shadowed, selfName, isLocal),
			}
		}
		return objectNode{entries: entries, tok: n.tok}
	case propertyAccessNode:
		left := rewriteClassSugarNode(n.left, visibleFields, allFields, shadowed, selfName)
		right := n.right
		if _, ok := right.(identifierNode); !ok {
			right = rewriteClassSugarNode(right, visibleFields, allFields, shadowed, selfName)
		}
		return propertyAccessNode{left: left, right: right, tok: n.tok}
	default:
		return rewriteClassSugarNode(node, visibleFields, allFields, shadowed, selfName)
	}
}

func rewriteClassSugarNode(node astNode, visibleFields, allFields, shadowed map[string]struct{}, selfName string) astNode {
	switch n := node.(type) {
	case identifierNode:
		if n.payload == selfName {
			return n
		}
		if _, ok := shadowed[n.payload]; ok {
			return n
		}
		if n.payload == "Self" {
			return identifierNode{payload: selfName, tok: n.tok}
		}
		if _, ok := visibleFields[n.payload]; ok {
			return propertyAccessNode{
				left:  makeClassSugarSelfNode(selfName, n.tok),
				right: n,
				tok:   n.tok,
			}
		}
		return n
	case listNode:
		elems := make([]astNode, len(n.elems))
		for i, elem := range n.elems {
			elems[i] = rewriteClassSugarNode(elem, visibleFields, allFields, shadowed, selfName)
		}
		return listNode{elems: elems, tok: n.tok}
	case objectNode:
		entries := make([]objectEntry, len(n.entries))
		for i, entry := range n.entries {
			key := entry.key
			if _, ok := key.(identifierNode); !ok {
				key = rewriteClassSugarNode(key, visibleFields, allFields, shadowed, selfName)
			}
			entries[i] = objectEntry{
				key: key,
				val: rewriteClassSugarNode(entry.val, visibleFields, allFields, shadowed, selfName),
			}
		}
		return objectNode{entries: entries, tok: n.tok}
	case fnNode:
		fnShadowed := cloneNameSet(shadowed)
		if n.name != "" {
			fnShadowed[n.name] = struct{}{}
		}
		for _, arg := range n.args {
			if arg != "" {
				fnShadowed[arg] = struct{}{}
			}
		}
		if n.restArg != "" {
			fnShadowed[n.restArg] = struct{}{}
		}
		return fnNode{
			name:    n.name,
			args:    append([]string{}, n.args...),
			restArg: n.restArg,
			body:    rewriteClassSugarNode(n.body, allFields, allFields, fnShadowed, selfName),
			tok:     n.tok,
		}
	case classNode:
		classShadowed := cloneNameSet(shadowed)
		classShadowed[n.name] = struct{}{}
		for _, arg := range n.args {
			if arg != "" {
				classShadowed[arg] = struct{}{}
			}
		}
		if n.restArg != "" {
			classShadowed[n.restArg] = struct{}{}
		}
		staticExprs := make([]astNode, len(n.staticExprs))
		for i, expr := range n.staticExprs {
			staticExprs[i] = rewriteClassSugarNode(expr, visibleFields, allFields, classShadowed, selfName)
		}
		parents := make([]astNode, len(n.parents))
		for i, parent := range n.parents {
			parents[i] = rewriteClassSugarNode(parent, allFields, allFields, classShadowed, selfName)
		}
		return classNode{
			name:        n.name,
			args:        append([]string{}, n.args...),
			restArg:     n.restArg,
			body:        rewriteClassSugarNode(n.body, allFields, allFields, classShadowed, selfName),
			staticExprs: staticExprs,
			parents:     parents,
			tok:         n.tok,
		}
	case assignmentNode:
		return assignmentNode{
			isLocal: n.isLocal,
			left:    rewriteClassSugarAssignmentLeft(n.left, visibleFields, allFields, shadowed, selfName, n.isLocal),
			right:   rewriteClassSugarNode(n.right, visibleFields, allFields, shadowed, selfName),
			tok:     n.tok,
		}
	case propertyAccessNode:
		left := rewriteClassSugarNode(n.left, visibleFields, allFields, shadowed, selfName)
		right := n.right
		if _, ok := right.(identifierNode); !ok {
			right = rewriteClassSugarNode(right, visibleFields, allFields, shadowed, selfName)
		}
		return propertyAccessNode{left: left, right: right, tok: n.tok}
	case unaryNode:
		return unaryNode{op: n.op, right: rewriteClassSugarNode(n.right, visibleFields, allFields, shadowed, selfName), tok: n.tok}
	case binaryNode:
		return binaryNode{
			op:    n.op,
			left:  rewriteClassSugarNode(n.left, visibleFields, allFields, shadowed, selfName),
			right: rewriteClassSugarNode(n.right, visibleFields, allFields, shadowed, selfName),
			tok:   n.tok,
		}
	case fnCallNode:
		args := make([]astNode, len(n.args))
		for i, arg := range n.args {
			args[i] = rewriteClassSugarNode(arg, visibleFields, allFields, shadowed, selfName)
		}
		var restArg astNode
		if n.restArg != nil {
			restArg = rewriteClassSugarNode(n.restArg, visibleFields, allFields, shadowed, selfName)
		}
		return fnCallNode{
			fn:      rewriteClassSugarNode(n.fn, visibleFields, allFields, shadowed, selfName),
			args:    args,
			restArg: restArg,
			tok:     n.tok,
		}
	case ifExprNode:
		branches := make([]ifBranch, len(n.branches))
		for i, branch := range n.branches {
			branches[i] = ifBranch{
				target: rewriteClassSugarNode(branch.target, visibleFields, allFields, shadowed, selfName),
				body:   rewriteClassSugarNode(branch.body, visibleFields, allFields, shadowed, selfName),
			}
		}
		return ifExprNode{
			cond:     rewriteClassSugarNode(n.cond, visibleFields, allFields, shadowed, selfName),
			branches: branches,
			tok:      n.tok,
		}
	case blockNode:
		blockShadowed := cloneNameSet(shadowed)
		exprs := make([]astNode, len(n.exprs))
		for i, expr := range n.exprs {
			exprs[i] = rewriteClassSugarNode(expr, visibleFields, allFields, blockShadowed, selfName)
			if assign, ok := expr.(assignmentNode); ok && assign.isLocal {
				addPatternBindings(blockShadowed, assign.left)
			}
		}
		return blockNode{exprs: exprs, tok: n.tok}
	default:
		return node
	}
}

func classBodyFromAssignmentBlock(body astNode, reservedNames []string) (astNode, bool) {
	block, ok := body.(blockNode)
	if !ok || len(block.exprs) == 0 {
		return nil, false
	}

	fieldNames := make(map[string]struct{}, len(block.exprs))
	for _, expr := range block.exprs {
		assign, ok := expr.(assignmentNode)
		if !ok || !assign.isLocal {
			return nil, false
		}

		ident, ok := assign.left.(identifierNode)
		if !ok {
			return nil, false
		}

		fieldNames[ident.payload] = struct{}{}
	}

	selfName := "__oakClassSelf"
	reserved := make(map[string]struct{}, len(reservedNames)+len(fieldNames))
	for _, name := range reservedNames {
		if name != "" {
			reserved[name] = struct{}{}
		}
	}
	for {
		if _, used := fieldNames[selfName]; used {
			selfName += "_"
			continue
		}
		if _, used := reserved[selfName]; used {
			selfName += "_"
			continue
		}
		break
	}

	newExprs := make([]astNode, 0, len(block.exprs)+2)
	newExprs = append(newExprs, assignmentNode{
		isLocal: true,
		left:    identifierNode{payload: selfName, tok: block.tok},
		right:   objectNode{entries: []objectEntry{}, tok: block.tok},
		tok:     block.tok,
	})

	visibleFields := map[string]struct{}{}
	shadowed := map[string]struct{}{selfName: {}}
	for _, expr := range block.exprs {
		assign := expr.(assignmentNode)
		ident := assign.left.(identifierNode)
		newExprs = append(newExprs, assignmentNode{
			isLocal: false,
			left: propertyAccessNode{
				left:  identifierNode{payload: selfName, tok: ident.tok},
				right: ident,
				tok:   assign.tok,
			},
			right: rewriteClassSugarNode(assign.right, visibleFields, fieldNames, shadowed, selfName),
			tok:   assign.tok,
		})
		visibleFields[ident.payload] = struct{}{}
	}

	newExprs = append(newExprs, identifierNode{payload: selfName, tok: block.tok})
	return blockNode{exprs: newExprs, tok: block.tok}, true
}

func (p *parser) tryParseClassParentsClause() ([]astNode, astNode, bool, error) {
	startIndex := p.index
	startMinPrecLen := len(p.minBinaryPrec)

	restore := func() {
		p.restore(startIndex, startMinPrecLen)
	}

	if p.isEOF() || p.peek().kind != leftParen {
		return nil, nil, false, nil
	}

	p.pushMinPrec(0)
	popped := false
	pop := func() {
		if !popped {
			p.popMinPrec()
			popped = true
		}
	}

	p.next() // eat the leftParen

	parents := []astNode{}
	for !p.isEOF() && p.peek().kind != rightParen {
		parent, err := p.parseNode()
		if err != nil {
			pop()
			restore()
			return nil, nil, false, nil
		}

		parents = append(parents, parent)

		if p.peek().kind == comma {
			p.next()
			continue
		}

		if p.peek().kind != rightParen {
			pop()
			restore()
			return nil, nil, false, nil
		}
	}

	if _, err := p.expect(rightParen); err != nil {
		pop()
		restore()
		return nil, nil, false, nil
	}
	if _, err := p.expect(branchArrow); err != nil {
		pop()
		restore()
		return nil, nil, false, nil
	}

	body, err := p.parseNode()
	if err != nil {
		pop()
		restore()
		return nil, nil, false, nil
	}

	pop()
	return parents, body, true, nil
}

func (p *parser) tryParseInheritedClassBody() ([]astNode, []astNode, astNode, bool, error) {
	startIndex := p.index
	startMinPrecLen := len(p.minBinaryPrec)

	restore := func() {
		p.restore(startIndex, startMinPrecLen)
	}

	if p.isEOF() || p.peek().kind != leftBrace {
		return nil, nil, nil, false, nil
	}

	if _, err := p.expect(leftBrace); err != nil {
		restore()
		return nil, nil, nil, false, nil
	}

	staticExprs := []astNode{}
	for !p.isEOF() && p.peek().kind != rightBrace {
		parents, body, ok, err := p.tryParseClassParentsClause()
		if err != nil {
			restore()
			return nil, nil, nil, false, err
		}
		if ok {
			if p.peek().kind == comma {
				p.next()
			}
			if _, err := p.expect(rightBrace); err != nil {
				restore()
				return nil, nil, nil, false, nil
			}
			return staticExprs, parents, body, true, nil
		}

		expr, err := p.parseNode()
		if err != nil {
			restore()
			return nil, nil, nil, false, nil
		}
		if _, err := p.expect(comma); err != nil {
			restore()
			return nil, nil, nil, false, nil
		}

		staticExprs = append(staticExprs, expr)
	}

	restore()
	return nil, nil, nil, false, nil
}

// concrete astNode parse functions

type parseError struct {
	reason string
	pos
}

func (e parseError) Error() string {
	return fmt.Sprintf("Parse error at %s: %s", e.pos.String(), e.reason)
}

func (p *parser) parseAssignment(left astNode) (astNode, error) {
	if p.peek().kind != assign &&
		p.peek().kind != nonlocalAssign {
		return left, nil
	}

	next := p.next()
	node := assignmentNode{
		isLocal: next.kind == assign,
		left:    left,
		tok:     &next,
	}

	right, err := p.parseNode()
	if err != nil {
		return nil, err
	}
	node.right = right

	return node, nil
}

// parseUnit is responsible for parsing the smallest complete syntactic "units"
// of Oak's syntax, like literals including function literals, grouped
// expressions in blocks, and if/with expressions.
func (p *parser) parseUnit() (astNode, error) {
	tok := p.next()
	switch tok.kind {
	case qmark:
		return nullNode{tok: &tok}, nil
	case stringLiteral:
		payloadBuilder := bytes.Buffer{}
		runes := []rune(tok.payload)
		for i := 0; i < len(runes); i++ {
			c := runes[i]

			if c == '\\' {
				if i+1 >= len(runes) {
					break
				}
				i++
				c = runes[i]

				switch c {
				case 't':
					_ = payloadBuilder.WriteByte('\t')
				case 'n':
					_ = payloadBuilder.WriteByte('\n')
				case 'r':
					_ = payloadBuilder.WriteByte('\r')
				case 'f':
					_ = payloadBuilder.WriteByte('\f')
				case 'x':
					if i+2 >= len(runes) {
						_ = payloadBuilder.WriteByte('x')
						continue
					}

					hexCode, err := strconv.ParseUint(string(runes[i+1])+string(runes[i+2]), 16, 8)
					if err == nil {
						i += 2
						_ = payloadBuilder.WriteByte(uint8(hexCode))
					} else {
						_ = payloadBuilder.WriteByte('x')
					}
				default:
					_, _ = payloadBuilder.WriteRune(c)
				}
			} else {
				_, _ = payloadBuilder.WriteRune(c)
			}
		}
		return stringNode{payload: payloadBuilder.Bytes(), tok: &tok}, nil
	case numberLiteral:
		// Parse as float if contains '.' or 'e'/'E' (scientific notation)
		// But not for hex (0x) or binary (0b) literals, where e/E are valid digits
		isHexOrBinary := strings.HasPrefix(tok.payload, "0x") || strings.HasPrefix(tok.payload, "0X") ||
			strings.HasPrefix(tok.payload, "0b") || strings.HasPrefix(tok.payload, "0B")
		if !isHexOrBinary && (strings.ContainsRune(tok.payload, '.') || strings.ContainsAny(tok.payload, "eE")) {
			f, err := strconv.ParseFloat(tok.payload, 64)
			if err != nil {
				return nil, parseError{reason: err.Error(), pos: tok.pos}
			}
			return floatNode{
				payload: f,
				tok:     &tok,
			}, nil
		}
		n, err := parseIntLiteral(tok.payload)
		if err != nil {
			return nil, parseError{reason: err.Error(), pos: tok.pos}
		}
		return intNode{
			payload: n,
			tok:     &tok,
		}, nil
	case trueLiteral:
		return boolNode{payload: true, tok: &tok}, nil
	case falseLiteral:
		return boolNode{payload: false, tok: &tok}, nil
	case colon:
		switch p.peek().kind {
		case identifier:
			return atomNode{payload: p.next().payload, tok: &tok}, nil
		case ifKeyword:
			p.next()
			return atomNode{payload: "if", tok: &tok}, nil
		case fnKeyword:
			p.next()
			return atomNode{payload: "fn", tok: &tok}, nil
		case withKeyword:
			p.next()
			return atomNode{payload: "with", tok: &tok}, nil
		case csKeyword:
			p.next()
			return atomNode{payload: "cs", tok: &tok}, nil
		case trueLiteral:
			p.next()
			return atomNode{payload: "true", tok: &tok}, nil
		case falseLiteral:
			p.next()
			return atomNode{payload: "false", tok: &tok}, nil
		}
		return nil, parseError{
			reason: fmt.Sprintf("Expected identifier after ':', got %s", p.peek()),
			pos:    tok.pos,
		}
	case leftBracket:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		itemNodes := make([]astNode, 0, 4)
		for !p.isEOF() && p.peek().kind != rightBracket {
			node, err := p.parseNode()
			if err != nil {
				return nil, err
			}

			itemNodes = append(itemNodes, node)

			// Comma is optional but if present, consume it
			// This allows both [a, b] and [a, b,]
			if !p.isEOF() && p.peek().kind == comma {
				p.next() // consume the comma
				// If comma is followed by right bracket, that's ok (trailing comma)
				if p.isEOF() || p.peek().kind == rightBracket {
					break
				}
			} else if !p.isEOF() && p.peek().kind != rightBracket {
				// If no comma and not at end, this is an error
				return nil, parseError{
					reason: fmt.Sprintf("Unexpected token %s, expected comma or ]", p.peek()),
					pos:    p.peek().pos,
				}
			}
		}
		if _, err := p.expect(rightBracket); err != nil {
			return nil, err
		}

		return listNode{elems: itemNodes, tok: &tok}, nil
	case leftBrace:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		// empty {} is always considered an object -- an empty block is illegal
		if p.peek().kind == rightBrace {
			p.next() // eat the rightBrace
			return objectNode{entries: []objectEntry{}, tok: &tok}, nil
		}

		firstExpr, err := p.parseNode()
		if err != nil {
			return nil, err
		}
		if p.isEOF() {
			return nil, parseError{
				reason: "Unexpected end of input inside block or object",
				pos:    tok.pos,
			}
		}

		if p.peek().kind == colon {
			// it's an object
			p.next() // eat the colon
			valExpr, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			if _, err := p.expect(comma); err != nil {
				return nil, err
			}

			entries := []objectEntry{
				{key: firstExpr, val: valExpr},
			}

			for !p.isEOF() && p.peek().kind != rightBrace {
				key, err := p.parseNode()
				if err != nil {
					return nil, err
				}
				if _, err := p.expect(colon); err != nil {
					return nil, err
				}

				val, err := p.parseNode()
				if err != nil {
					return nil, err
				}
				if _, err := p.expect(comma); err != nil {
					return nil, err
				}

				entries = append(entries, objectEntry{
					key: key,
					val: val,
				})
			}
			if _, err := p.expect(rightBrace); err != nil {
				return nil, err
			}

			return objectNode{entries: entries, tok: &tok}, nil
		}

		// it's a block
		exprs := []astNode{firstExpr}
		if _, err := p.expect(comma); err != nil {
			return nil, err
		}

		for !p.isEOF() && p.peek().kind != rightBrace {
			expr, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			if _, err := p.expect(comma); err != nil {
				return nil, err
			}

			exprs = append(exprs, expr)
		}
		if _, err := p.expect(rightBrace); err != nil {
			return nil, err
		}

		return blockNode{exprs: exprs, tok: &tok}, nil
	case fnKeyword:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		name := ""
		if p.peek().kind == identifier {
			// optional named fn
			name = p.next().payload
		}

		args := []string{}
		var restArg string
		if p.peek().kind == leftParen {
			// optional argument list
			p.next() // eat the leftParen
			for !p.isEOF() && p.peek().kind != rightParen {
				arg, err := p.expect(identifier)
				if err != nil {
					p.back() // try again

					_, err := p.expect(underscore)
					if err != nil {
						return nil, err
					}

					args = append(args, "")

					if _, err := p.expect(comma); err != nil {
						return nil, err
					}

					continue
				}

				// maybe this is a rest arg
				if p.peek().kind == ellipsis {
					restArg = arg.payload
					p.next() // eat the ellipsis

					_, err = p.expect(comma)
					if err != nil {
						return nil, err
					}
					break
				}

				args = append(args, arg.payload)

				if _, err := p.expect(comma); err != nil {
					return nil, err
				}
			}
			if _, err := p.expect(rightParen); err != nil {
				return nil, err
			}
		}

		body, err := p.parseNode()
		if err != nil {
			return nil, err
		}

		// Exception to the "{} is empty object" rule is that `fn {}` parses as
		// a function with an empty block as a body
		if objBody, ok := body.(objectNode); ok && len(objBody.entries) == 0 {
			body = blockNode{exprs: []astNode{}, tok: objBody.tok}
		}

		return fnNode{
			name:    name,
			args:    args,
			restArg: restArg,
			body:    body,
			tok:     &tok,
		}, nil
	case csKeyword:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		nameTok, err := p.expect(identifier)
		if err != nil {
			return nil, err
		}

		args := []string{}
		var restArg string
		if p.peek().kind == leftParen {
			p.next() // eat the leftParen
			for !p.isEOF() && p.peek().kind != rightParen {
				arg, err := p.expect(identifier)
				if err != nil {
					p.back() // try again

					_, err := p.expect(underscore)
					if err != nil {
						return nil, err
					}

					args = append(args, "")

					if _, err := p.expect(comma); err != nil {
						return nil, err
					}

					continue
				}

				if p.peek().kind == ellipsis {
					restArg = arg.payload
					p.next() // eat the ellipsis

					_, err = p.expect(comma)
					if err != nil {
						return nil, err
					}
					break
				}

				args = append(args, arg.payload)

				if _, err := p.expect(comma); err != nil {
					return nil, err
				}
			}
			if _, err := p.expect(rightParen); err != nil {
				return nil, err
			}
		}

		if p.peek().kind == leftBrace {
			staticExprs, parents, body, ok, err := p.tryParseInheritedClassBody()
			if err != nil {
				return nil, err
			}
			if ok {
				return classNode{
					name:        nameTok.payload,
					args:        args,
					restArg:     restArg,
					body:        body,
					staticExprs: staticExprs,
					parents:     parents,
					tok:         &tok,
				}, nil
			}
		}

		body, err := p.parseNode()
		if err != nil {
			return nil, err
		}

		// Exception to the "{} is empty object" rule is that `cs Name {}`
		// should parse like `Name := fn Name {}` with an empty block body.
		if objBody, ok := body.(objectNode); ok && len(objBody.entries) == 0 {
			body = blockNode{exprs: []astNode{}, tok: objBody.tok}
		}

		reservedNames := append([]string{}, args...)
		if restArg != "" {
			reservedNames = append(reservedNames, restArg)
		}
		if desugaredBody, ok := classBodyFromAssignmentBlock(body, reservedNames); ok {
			body = desugaredBody
		}

		return classNode{
			name:    nameTok.payload,
			args:    args,
			restArg: restArg,
			body:    body,
			tok:     &tok,
		}, nil
	case underscore:
		return emptyNode{tok: &tok}, nil
	case identifier:
		return identifierNode{payload: tok.payload, tok: &tok}, nil
	case minus, exclam, tilde:
		right, err := p.parseSubNode()
		if err != nil {
			return nil, err
		}

		return unaryNode{
			op:    tok.kind,
			right: right,
			tok:   &tok,
		}, nil
	case ifKeyword:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		var condNode astNode
		branches := make([]ifBranch, 0, 4)

		// if no explicit condition is provided (i.e. if the keyword is
		// followed by a { ... }), we assume the condition is "true" to allow
		// for the useful `if { case, case ... }` pattern.
		var err error
		if p.peek().kind == leftBrace {
			condNode = boolNode{
				payload: true,
				tok:     &tok,
			}
		} else {
			condNode, err = p.parseNode()
			if err != nil {
				return nil, err
			}
		}

		// `if cond -> body` desugars to `if cond { true -> body }`. Note that
		// in this form, there can only be one condition expression; `if a, b,
		// c -> body` is not legal. However, `if a | b | c -> body` is
		// equivalent and valid.
		if p.peek().kind == branchArrow {
			arrowTok := p.next()

			body, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			// comma here marks end of the ifExpr, not end of branch, so we do
			// not consume it here.

			branches = append(branches, ifBranch{
				target: boolNode{
					payload: true,
					tok:     &arrowTok,
				},
				body: body,
			})
			return ifExprNode{
				cond:     condNode,
				branches: branches,
				tok:      &tok,
			}, nil
		}

		if _, err = p.expect(leftBrace); err != nil {
			return nil, err
		}

		for !p.isEOF() && p.peek().kind != rightBrace {
			targets := []astNode{}
			for !p.isEOF() && p.peek().kind != branchArrow {
				target, err := p.parseNode()
				if err != nil {
					return nil, err
				}
				if p.peek().kind != branchArrow {
					if _, err := p.expect(comma); err != nil {
						return nil, err
					}
				}

				targets = append(targets, target)
			}
			if _, err := p.expect(branchArrow); err != nil {
				return nil, err
			}

			body, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			if _, err := p.expect(comma); err != nil {
				return nil, err
			}

			// We want to support multi-target branches, but don't want to
			// incur the performance overhead in the interpreter/evaluator of
			// keeping every single target as a Go slice, when the vast
			// majority of targets will be single-value, which requires just a
			// pointer to an astNode.
			//
			// So instead of doing that, we penalize the multi-value case by
			// essentially considering it syntax sugar and splitting such
			// branches into multiple AST branches, each with one target value.
			for _, target := range targets {
				branches = append(branches, ifBranch{
					target: target,
					body:   body,
				})
			}
		}
		if _, err := p.expect(rightBrace); err != nil {
			return nil, err
		}

		return ifExprNode{
			cond:     condNode,
			branches: branches,
			tok:      &tok,
		}, nil
	case withKeyword:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		withExprBase, err := p.parseNode()
		if err != nil {
			return nil, err
		}

		withExprBaseCall, ok := withExprBase.(fnCallNode)
		if !ok {
			return nil, parseError{
				reason: fmt.Sprintf("with keyword should be followed by a function call, found %s", withExprBase),
				pos:    tok.pos,
			}
		}

		withExprLastArg, err := p.parseNode()
		if err != nil {
			return nil, err
		}

		withExprBaseCall.args = append(withExprBaseCall.args, withExprLastArg)
		return withExprBaseCall, nil
	case leftParen:
		p.pushMinPrec(0)
		defer p.popMinPrec()

		exprs := make([]astNode, 0, 4)
		for !p.isEOF() && p.peek().kind != rightParen {
			expr, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			if _, err := p.expect(comma); err != nil {
				return nil, err
			}

			exprs = append(exprs, expr)
		}
		if _, err := p.expect(rightParen); err != nil {
			return nil, err
		}
		// TODO: If only one body expr and body expr is identifier or literal,
		// unwrap the blockNode and just return the bare child
		return blockNode{exprs: exprs, tok: &tok}, nil
	}
	return nil, parseError{
		reason: fmt.Sprintf("Unexpected token %s at start of unit", tok),
		pos:    tok.pos,
	}
}

func infixOpPrecedence(op tokKind) int {
	switch op {
	case plus, minus:
		return 40
	case times, divide:
		return 50
	case power:
		return 75
	case modulus:
		return 80
	case eq, deepEq, greater, less, geq, leq, neq:
		return 30
	case and:
		return 20
	case xor:
		return 15
	case or:
		return 10
	case pushArrow, rshift:
		// both list/string concat and bitwise shifts use low precedence
		return 1
	default:
		return -1
	}
}

// parseSubNode is responsible for parsing independent "terms" in the Oak
// syntax, like terms in unary and binary expressions and in pipelines. It is
// in between parseUnit and parseNode.
func (p *parser) parseSubNode() (astNode, error) {
	p.pushMinPrec(0)
	defer p.popMinPrec()

	node, err := p.parseUnit()
	if err != nil {
		return nil, err
	}

	for !p.isEOF() {
		switch p.peek().kind {
		case dot:
			next := p.next() // eat the dot
			right, err := p.parseUnit()
			if err != nil {
				return nil, err
			}

			node = propertyAccessNode{
				left:  node,
				right: right,
				tok:   &next,
			}
		case leftParen:
			next := p.next() // eat the leftParen

			args := make([]astNode, 0, 4)
			var restArg astNode = nil
			for !p.isEOF() && p.peek().kind != rightParen {
				arg, err := p.parseNode()
				if err != nil {
					return nil, err
				}
				if p.peek().kind == ellipsis {
					p.next() // eat the ellipsis

					if _, err = p.expect(comma); err != nil {
						return nil, err
					}

					restArg = arg

					break
				} else {
					args = append(args, arg)
				}

				if _, err = p.expect(comma); err != nil {
					return nil, err
				}
			}
			if _, err := p.expect(rightParen); err != nil {
				return nil, err
			}

			node = fnCallNode{
				fn:      node,
				args:    args,
				restArg: restArg,
				tok:     &next,
			}
		default:
			return node, nil
		}
	}

	return node, nil
}

// parseNode returns the next top-level astNode from the parser
func (p *parser) parseNode() (astNode, error) {
	node, err := p.parseSubNode()
	if err != nil {
		return nil, err
	}

	for !p.isEOF() && p.peek().kind != comma {
		switch p.peek().kind {
		case assign, nonlocalAssign:
			// whatever follows an assignment expr cannot bind to the
			// assignment expression itself by syntax rule, so we simply return
			return p.parseAssignment(node)
		case plus, minus, times, divide, modulus, power,
			xor, and, or, pushArrow, rshift,
			greater, less, eq, deepEq, geq, leq, neq:
			// this case implements a mini Pratt parser threaded through the
			// larger Oak syntax parser, using the parser struct itself to keep
			// track of the power / precedence stack since other forms may be
			// parsed in between, as in 1 + f(g(x := y)) + 2
			minPrec := p.lastMinPrec()

			for {
				if p.isEOF() {
					return nil, parseError{
						reason: "Incomplete binary expression",
						pos:    p.peek().pos,
					}
				}

				peeked := p.peek()
				op := peeked.kind
				prec := infixOpPrecedence(op)
				if prec <= minPrec {
					break
				}
				p.next() // eat the operator

				if p.isEOF() {
					return nil, parseError{
						reason: fmt.Sprintf("Incomplete binary expression with %s", token{kind: op}),
						pos:    p.peek().pos,
					}
				}

				p.pushMinPrec(prec)
				right, err := p.parseNode()
				if err != nil {
					return nil, err
				}
				p.popMinPrec()

				node = binaryNode{
					op:    op,
					left:  node,
					right: right,
					tok:   &peeked,
				}
			}

			// whatever follows a binary expr cannot bind to the binary
			// expression by syntax rule, so we simply return
			return node, nil
		case pipeArrow:
			pipe := p.next() // eat the pipe

			pipeRight, err := p.parseSubNode()
			if err != nil {
				return nil, err
			}
			pipedFnCall, ok := pipeRight.(fnCallNode)
			if !ok {
				return nil, parseError{
					reason: fmt.Sprintf("Expected function call after |>, got %s", pipeRight),
					pos:    pipe.pos,
				}
			}

			pipedFnCall.args = append([]astNode{node}, pipedFnCall.args...)
			node = pipedFnCall
		default:
			return node, nil
		}
	}
	// the trailing comma is handled as necessary in callers of parseNode

	return node, nil
}

func (p *parser) parse() ([]astNode, error) {
	nodes := make([]astNode, 0, 8)

	for !p.isEOF() {
		node, err := p.parseNode()
		if err != nil {
			return nodes, err
		}

		if _, err = p.expect(comma); err != nil {
			return nodes, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
