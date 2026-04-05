package main

import (
	"bufio"
	"bytes"
	"context"
	crand "crypto/rand"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	utf16pkg "unicode/utf16"
	"unsafe"

	"github.com/gorilla/websocket"
)

func (c *Context) requireArgLen(fnName string, args []Value, count int) *runtimeError {
	if len(args) < count {
		return &runtimeError{
			reason: fmt.Sprintf("%s requires %d arguments, got %d", fnName, count, len(args)),
		}
	}

	return nil
}

type builtinFn func([]Value) (Value, *runtimeError)

type BuiltinFnValue struct {
	name string
	fn   builtinFn
}

func (v BuiltinFnValue) String() string {
	return fmt.Sprintf("fn %s { <native fn> }", v.name)
}
func (v BuiltinFnValue) Eq(u Value) bool {
	if _, ok := u.(EmptyValue); ok {
		return true
	}

	if w, ok := u.(BuiltinFnValue); ok {
		return v.name == w.name
	}
	return false
}

func (c *Context) LoadFunc(name string, fn builtinFn) {
	c.scope.put(name, BuiltinFnValue{
		name: name,
		fn:   fn,
	})
}

func (c *Context) LoadBuiltins() {
	// global initializations
	rand.Seed(time.Now().UnixNano())

	// core language and reflection
	c.LoadFunc("import", c.oakImport)
	c.LoadFunc("int", c.oakInt)
	c.LoadFunc("float", c.oakFloat)
	c.LoadFunc("atom", c.oakAtom)
	c.LoadFunc("string", c.oakString)
	c.LoadFunc("codepoint", c.oakCodepoint)
	c.LoadFunc("char", c.oakChar)
	c.LoadFunc("type", c.oakType)
	c.LoadFunc("name", c.oakName)
	c.LoadFunc("csof", c.oakClassMatch)
	c.LoadFunc("len", c.oakLen)
	c.LoadFunc("keys", c.oakKeys)
	c.LoadFunc("___stdlibs", c.oakStdlibs)

	// os interfaces
	c.LoadFunc("args", c.oakArgs)
	c.LoadFunc("env", c.oakEnv)
	c.LoadFunc("time", c.oakTime)
	c.LoadFunc("nanotime", c.oakNanotime)
	c.LoadFunc("rand", c.oakRand)
	c.LoadFunc("srand", c.oakSrand)
	c.LoadFunc("wait", c.callbackify(c.oakWait))
	c.LoadFunc("exit", c.oakExit)
	c.LoadFunc("exec", c.callbackify(c.oakExec))
	c.LoadFunc("sysproc", c.oakSysProc)
	c.LoadFunc("syscall", c.oakSyscall)
	c.LoadFunc("win_msg_loop", c.oakWinMsgLoop)
	c.LoadFunc("utf16", c.oakUTF16)
	c.LoadFunc("bits", c.oakBits)
	c.LoadFunc("addr", c.oakAddr)
	c.LoadFunc("pointer", c.oakPointer) // convert integer or string to a pointer value
	c.LoadFunc("memread", c.oakMemRead)
	c.LoadFunc("memwrite", c.oakMemWrite)
	c.LoadFunc("go", c.oakGo)
	c.LoadFunc("lock_thread", c.oakLockThread)
	c.LoadFunc("unlock_thread", c.oakUnlockThread)
	c.LoadFunc("make_chan", c.oakMakeChan)
	c.LoadFunc("chan_send", c.oakChanSend)
	c.LoadFunc("chan_recv", c.oakChanRecv)

	// i/o interfaces
	c.LoadFunc("input", c.callbackify(c.oakInput))
	c.LoadFunc("print", c.oakPrint)
	c.LoadFunc("ls", c.callbackify(c.oakLs))
	c.LoadFunc("rm", c.callbackify(c.oakRm))
	c.LoadFunc("mkdir", c.callbackify(c.oakMkdir))
	c.LoadFunc("stat", c.callbackify(c.oakStat))
	c.LoadFunc("open", c.callbackify(c.oakOpen))
	c.LoadFunc("close", c.callbackify(c.oakClose))
	c.LoadFunc("read", c.callbackify(c.oakRead))
	c.LoadFunc("write", c.callbackify(c.oakWrite))
	c.LoadFunc("listen", c.oakListen)
	c.LoadFunc("req", c.callbackify(c.oakReq))
	c.LoadFunc("socket_connect", c.callbackify(c.oakSocketConnect))
	c.LoadFunc("socket_send", c.callbackify(c.oakSocketSend))
	c.LoadFunc("socket_recv", c.callbackify(c.oakSocketRecv))
	c.LoadFunc("socket_recv_exact", c.callbackify(c.oakSocketRecvExact))
	c.LoadFunc("socket_recv_line", c.callbackify(c.oakSocketRecvLine))
	c.LoadFunc("socket_starttls", c.callbackify(c.oakSocketStartTLS))
	c.LoadFunc("socket_close", c.callbackify(c.oakSocketClose))
	c.LoadFunc("socket_listen", c.oakSocketListen)
	c.LoadFunc("ws_dial", c.callbackify(c.oakWSDial))
	c.LoadFunc("ws_send", c.callbackify(c.oakWSSend))
	c.LoadFunc("ws_recv", c.callbackify(c.oakWSRecv))
	c.LoadFunc("ws_close", c.callbackify(c.oakWSClose))
	c.LoadFunc("ws_listen", c.oakWSListen)

	// math
	c.LoadFunc("sin", c.oakSin)
	c.LoadFunc("cos", c.oakCos)
	c.LoadFunc("tan", c.oakTan)
	c.LoadFunc("asin", c.oakAsin)
	c.LoadFunc("acos", c.oakAcos)
	c.LoadFunc("atan", c.oakAtan)
	c.LoadFunc("pow", c.oakPow)
	c.LoadFunc("log", c.oakLog)

	// language and runtime APIs
	c.LoadFunc("___runtime_lib", c.rtLib)
	c.LoadFunc("___runtime_lib?", c.rtIsLib)
	c.LoadFunc("___runtime_gc", c.rtGC)
	c.LoadFunc("___runtime_mem", c.rtMem)
	c.LoadFunc("___runtime_proc", c.rtProc)
	c.LoadFunc("___runtime_go_version", c.rtGoVersion)
	c.LoadFunc("___runtime_sys_info", c.rtSysInfo)

	// execution mode switching
	c.LoadFunc("bytecode", c.oakBytecodeCall)
	c.LoadFunc("interpreter", c.oakInterpreterCall)
}

func errObj(message string) ObjectValue {
	return ObjectValue{
		"type":  AtomValue("error"),
		"error": MakeString(message),
	}
}

func syscallErrObj(message string) ObjectValue {
	return ObjectValue{
		"type":  AtomValue("error"),
		"error": MakeString(message),
		"errno": IntValue(0),
		"r1":    IntValue(0),
		"r2":    IntValue(0),
	}
}

var (
	wsConnMu      sync.Mutex
	wsConnMap     = map[int64]*websocket.Conn{}
	wsNextConnID  int64
	socketConnMu  sync.Mutex
	socketConnMap = map[int64]*oakSocketConn{}
	socketNextID  int64
)

type oakSocketConn struct {
	conn      net.Conn
	reader    *bufio.Reader
	isTLS     bool
	serverTLS bool
}

func chanSentObj() ObjectValue {
	return ObjectValue{
		"type": AtomValue("sent"),
	}
}

func chanDataObj(data Value) ObjectValue {
	return ObjectValue{
		"type": AtomValue("data"),
		"data": data,
	}
}

func websocketObj(id int64) ObjectValue {
	return ObjectValue{
		"type": AtomValue("websocket"),
		"id":   IntValue(id),
	}
}

func socketObj(id int64) ObjectValue {
	return ObjectValue{
		"type": AtomValue("socket"),
		"id":   IntValue(id),
	}
}

func websocketEvent(messageType int, data []byte) ObjectValue {
	return ObjectValue{
		"type":   AtomValue("message"),
		"opcode": IntValue(messageType),
		"data":   MakeString(string(data)),
	}
}

func websocketClosedEvent(code int, reason string) ObjectValue {
	return ObjectValue{
		"type":   AtomValue("closed"),
		"code":   IntValue(code),
		"reason": MakeString(reason),
	}
}

func socketClosedEvent() ObjectValue {
	return ObjectValue{
		"type": AtomValue("closed"),
	}
}

func socketDataEvent(data []byte) ObjectValue {
	return ObjectValue{
		"type": AtomValue("data"),
		"data": MakeString(string(data)),
	}
}

func makeHeaderObject(headers http.Header) ObjectValue {
	out := ObjectValue{}
	for key, values := range headers {
		out[key] = MakeString(strings.Join(values, ","))
	}
	return out
}

func (c *Context) getWebsocket(arg Value, fnName string) (*websocket.Conn, int64, *runtimeError) {
	wsObj, ok := arg.(ObjectValue)
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a websocket, got %s", fnName, arg),
		}
	}

	typeVal, ok := wsObj["type"].(AtomValue)
	if !ok || typeVal != AtomValue("websocket") {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a websocket, got %s", fnName, arg),
		}
	}

	id, ok := wsObj["id"].(IntValue)
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("Websocket %s is malformed", arg),
		}
	}

	wsConnMu.Lock()
	conn, ok := wsConnMap[int64(id)]
	wsConnMu.Unlock()
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("Websocket %s is not available", arg),
		}
	}

	return conn, int64(id), nil
}

func storeWebsocket(conn *websocket.Conn) ObjectValue {
	wsConnMu.Lock()
	id := wsNextConnID
	wsNextConnID++
	wsConnMap[id] = conn
	wsConnMu.Unlock()

	return websocketObj(id)
}

func storeSocket(conn net.Conn, isTLS bool, serverTLS bool) ObjectValue {
	socketConnMu.Lock()
	id := socketNextID
	socketNextID++
	socketConnMap[id] = &oakSocketConn{
		conn:      conn,
		reader:    bufio.NewReader(conn),
		isTLS:     isTLS,
		serverTLS: serverTLS,
	}
	socketConnMu.Unlock()

	return socketObj(id)
}

func removeSocket(id int64) {
	socketConnMu.Lock()
	delete(socketConnMap, id)
	socketConnMu.Unlock()
}

func (c *Context) getSocket(arg Value, fnName string) (*oakSocketConn, int64, *runtimeError) {
	socketValue, ok := arg.(ObjectValue)
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a socket, got %s", fnName, arg),
		}
	}

	typeVal, ok := socketValue["type"].(AtomValue)
	if !ok || typeVal != AtomValue("socket") {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a socket, got %s", fnName, arg),
		}
	}

	id, ok := socketValue["id"].(IntValue)
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("Socket %s is malformed", arg),
		}
	}

	socketConnMu.Lock()
	conn, ok := socketConnMap[int64(id)]
	socketConnMu.Unlock()
	if !ok {
		return nil, 0, &runtimeError{
			reason: fmt.Sprintf("Socket %s is not available", arg),
		}
	}

	return conn, int64(id), nil
}

func removeWebsocket(id int64) {
	wsConnMu.Lock()
	delete(wsConnMap, id)
	wsConnMu.Unlock()
}

func (c *Context) getGoChan(arg Value, fnName string) (chan Value, *runtimeError) {
	chObj, ok := arg.(ObjectValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a channel, got %s", fnName, arg),
		}
	}

	typeVal, ok := chObj["type"].(AtomValue)
	if !ok || typeVal != AtomValue("channel") {
		return nil, &runtimeError{
			reason: fmt.Sprintf("First argument to %s must be a channel, got %s", fnName, arg),
		}
	}

	id, ok := chObj["id"].(IntValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Channel %s is malformed", arg),
		}
	}

	c.eng.chanLock.Lock()
	ch, ok := c.eng.chanMap[int64(id)]
	c.eng.chanLock.Unlock()
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Channel %s is not available", arg),
		}
	}

	return ch, nil
}

func rawMemoryRegion(addr uintptr, length int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(addr)), length)
}

func (c *Context) callbackify(syncFn builtinFn) builtinFn {
	return func(args []Value) (Value, *runtimeError) {
		if len(args) == 0 {
			return syncFn(args)
		}

		lastArg := args[len(args)-1]
		if !c.isCallableValue(lastArg) {
			return syncFn(args)
		}

		syncArgs := args[:len(args)-1]
		c.eng.Add(1)
		go func() {
			defer c.eng.Done()

			evt, err := syncFn(syncArgs)
			if err != nil {
				c.eng.reportErr(err)
				return
			}

			c.Lock()
			defer c.Unlock()
			_, err = c.EvalFnValueAsync(lastArg, false, evt)
			if err != nil {
				c.eng.reportErr(err)
				return
			}
		}()

		return null, nil
	}
}

// isCallableValue reports whether v can be invoked by evalFnCall.
func (c *Context) isCallableValue(v Value) bool {
	switch fn := v.(type) {
	case FnValue, BuiltinFnValue, ClassValue:
		return true
	case *closureVal:
		return fn.call != nil
	default:
		return false
	}
}

// isGoTargetValue reports whether v is a valid target for go().
// Builtins/classes are intentionally excluded to preserve existing behavior.
func (c *Context) isGoTargetValue(v Value) bool {
	switch fn := v.(type) {
	case FnValue:
		return true
	case *closureVal:
		return fn.call != nil
	default:
		return false
	}
}

func (c *Context) oakStdlibs(args []Value) (Value, *runtimeError) {
	if len(args) != 0 {
		return nil, &runtimeError{reason: "___stdlibs takes no arguments"}
	}

	obj := make(map[string]Value)
	for name, source := range stdlibs {
		obj[name] = MakeString(source)
	}
	return ObjectValue(obj), nil
}

func (c *Context) oakImport(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("import", args, 1); err != nil {
		return nil, err
	}

	pathBytes, ok := args[0].(*StringValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("path to import() must be a string, got %s", args[0]),
		}
	}
	pathStr := pathBytes.stringContent()

	// if a stdlib, just import the library from binary
	if isStdLib(pathStr) {
		return c.LoadLib(pathStr)
	}

	supportedExts := []string{".oak", ".ok", ".mag", ".mg"}
	candidatePaths := make([]string, 0, len(supportedExts))
	inputExt := filepath.Ext(pathStr)
	if inputExt != "" {
		for _, ext := range supportedExts {
			if strings.EqualFold(inputExt, ext) {
				candidatePaths = append(candidatePaths, pathStr)
				break
			}
		}
	}
	if len(candidatePaths) == 0 {
		for _, ext := range supportedExts {
			candidatePaths = append(candidatePaths, pathStr+ext)
		}
	}

	var (
		file     *os.File
		err      error
		filePath string
	)
	for _, candidate := range candidatePaths {
		resolved := candidate
		if !filepath.IsAbs(resolved) {
			resolved = filepath.Join(c.rootPath, resolved)
		}
		resolved = filepath.Clean(resolved)
		if !filepath.IsAbs(resolved) {
			if absPath, absErr := filepath.Abs(resolved); absErr == nil {
				resolved = absPath
			}
		}

		file, err = os.Open(resolved)
		if err == nil {
			filePath = resolved
			break
		}
	}
	if file == nil {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Could not open %s (%s)", pathStr, strings.Join(supportedExts, ", ")),
		}
	}
	defer file.Close()

	c.eng.importLock.RLock()
	imported, ok := c.eng.importMap[filePath]
	c.eng.importLock.RUnlock()
	if ok {
		return ObjectValue(imported.vars), nil
	}

	ctx := c.ChildContext(filepath.Dir(filePath))
	ctx.currentFile = filePath
	c.eng.importLock.Lock()
	c.eng.importMap[filePath] = ctx.scope
	c.eng.importLock.Unlock()
	ctx.LoadBuiltins()

	ctx.Unlock()
	_, err = ctx.Eval(file)
	ctx.Lock()
	if err != nil {
		if runtimeErr, ok := err.(*runtimeError); ok {
			return nil, runtimeErr
		} else {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Error importing %s: %s", pathStr, err.Error()),
			}
		}
	}

	return ObjectValue(ctx.scope.vars), nil
}

func (c *Context) oakInt(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("int", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case IntValue:
		return arg, nil
	case PointerValue:
		return IntValue(arg), nil
	case FloatValue:
		return IntValue(math.Floor(float64(arg))), nil
	case *StringValue:
		n, err := strconv.ParseInt(arg.stringContent(), 10, 64)
		if err != nil {
			return null, nil
		}
		return IntValue(n), nil
	default:
		return null, nil
	}
}

func (c *Context) oakFloat(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("float", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case IntValue:
		return FloatValue(arg), nil
	case FloatValue:
		return arg, nil
	case *StringValue:
		f, err := strconv.ParseFloat(arg.stringContent(), 64)
		if err != nil {
			return null, nil
		}
		return FloatValue(f), nil
	default:
		return null, nil
	}
}

func (c *Context) oakAtom(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("atom", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		return AtomValue(arg.stringContent()), nil
	case AtomValue:
		return arg, nil
	default:
		return AtomValue(arg.String()), nil
	}
}

func (c *Context) oakString(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("string", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		return arg, nil
	case AtomValue:
		return MakeString(string(arg)), nil
	default:
		return MakeString(arg.String()), nil
	}
}

func (c *Context) oakCodepoint(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("codepoint", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		if len(*arg) != 1 {
			return null, nil
		}
		return IntValue(int64((*arg)[0])), nil
	default:
		return null, nil
	}
}

func (c *Context) oakChar(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("char", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case IntValue:
		codepoint := int64(arg)
		if codepoint < 0 {
			codepoint = 0
		}
		if codepoint > 255 {
			codepoint = 255
		}
		return MakeSingleByteString(byte(codepoint)), nil
	default:
		return null, nil
	}
}

func (c *Context) oakType(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("type", args, 1); err != nil {
		return nil, err
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
	case AtomValue:
		return AtomValue("atom"), nil
	case *StringValue:
		return AtomValue("string"), nil
	case *ListValue:
		return AtomValue("list"), nil
	case ObjectValue:
		return AtomValue("object"), nil
	case FnValue, BuiltinFnValue, ClassValue, *closureVal:
		return AtomValue("function"), nil
	case PointerValue:
		return AtomValue("pointer"), nil
	}

	panic("Unreachable: unknown runtime value")
}

func (c *Context) oakName(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("name", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case AtomValue:
		return arg, nil
	case *StringValue:
		return AtomValue(arg.stringContent()), nil
	case ClassValue:
		return AtomValue(arg.defn.name), nil
	case FnValue:
		return AtomValue(arg.defn.name), nil
	case BuiltinFnValue:
		return AtomValue(arg.name), nil
	case PointerValue:
		c.eng.nameRefLock.Lock()
		nameVal, ok := c.eng.nameRefs[uintptr(arg)]
		c.eng.nameRefLock.Unlock()
		if ok {
			return AtomValue(nameVal.stringContent()), nil
		}
	}

	return AtomValue(args[0].String()), nil
}

func (c *Context) oakClassMatch(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("csof", args, 1); err != nil {
		return nil, err
	}

	if len(args) == 1 {
		_, ok := args[0].(ClassValue)
		return BoolValue(ok), nil
	}

	leftClass, leftIsClass := args[0].(ClassValue)
	rightClass, rightIsClass := args[1].(ClassValue)
	leftAtom, leftIsAtom := args[0].(AtomValue)
	rightAtom, rightIsAtom := args[1].(AtomValue)

	switch {
	case leftIsClass && rightIsClass:
		return BoolValue(leftClass.Eq(rightClass)), nil
	case leftIsClass && rightIsAtom:
		return BoolValue(leftClass.defn.name == string(rightAtom)), nil
	case leftIsAtom && rightIsClass:
		return BoolValue(string(leftAtom) == rightClass.defn.name), nil
	default:
		return BoolValue(false), nil
	}
}

func (c *Context) oakLen(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("string", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		return IntValue(len(*arg)), nil
	case *ListValue:
		return IntValue(len(*arg)), nil
	case ObjectValue:
		return IntValue(len(arg)), nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("%s does not support a len() call", arg),
		}
	}
}

func makeIntListUpTo(max int) Value {
	list := make(ListValue, max)
	for i := 0; i < max; i++ {
		list[i] = IntValue(i)
	}
	return &list
}

func (c *Context) oakKeys(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("print", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		return makeIntListUpTo(len(*arg)), nil
	case *ListValue:
		return makeIntListUpTo(len(*arg)), nil
	case ObjectValue:
		keys := make(ListValue, 0, len(arg))
		for key := range arg {
			keys = append(keys, MakeString(key))
		}
		return &keys, nil
	default:
		return MakeList(), nil
	}
}

func (c *Context) oakArgs(_ []Value) (Value, *runtimeError) {
	goArgs := os.Args
	args := make(ListValue, len(goArgs))
	for i, arg := range goArgs {
		args[i] = MakeString(arg)
	}
	return &args, nil
}

func (c *Context) oakEnv(_ []Value) (Value, *runtimeError) {
	envVars := ObjectValue{}
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		envVars[kv[0]] = MakeString(kv[1])
	}
	return envVars, nil
}

func (c *Context) oakTime(_ []Value) (Value, *runtimeError) {
	unixSeconds := float64(time.Now().UnixNano()) / 1e9
	return FloatValue(unixSeconds), nil
}

func (c *Context) oakNanotime(_ []Value) (Value, *runtimeError) {
	return IntValue(time.Now().UnixNano()), nil
}

func (c *Context) oakRand(_ []Value) (Value, *runtimeError) {
	return FloatValue(rand.Float64()), nil
}

func (c *Context) oakSrand(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("srand", args, 1); err != nil {
		return nil, err
	}

	bufLen, ok1 := args[0].(IntValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call srand(%s)", args[0]),
		}
	}

	buf := make([]byte, bufLen)
	_, err := crand.Read(buf)
	if err != nil {
		return null, nil
	}

	bytes := StringValue(buf)
	return &bytes, nil
}

func (c *Context) oakWait(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("wait", args, 1); err != nil {
		return nil, err
	}

	// in both Oak & Go, duration <= 0 results in immediate completion
	switch arg := args[0].(type) {
	case IntValue:
		time.Sleep(time.Duration(float64(arg) * float64(time.Second)))
	case FloatValue:
		time.Sleep(time.Duration(float64(arg) * float64(time.Second)))
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call wait(%s)", args[0]),
		}
	}

	return null, nil
}

func (c *Context) oakExit(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("exit", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case IntValue:
		os.Exit(int(arg))
		// unreachable
		return null, nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call exit(%s)", args[0]),
		}
	}
}

func (c *Context) oakExec(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("exec", args, 3); err != nil {
		return nil, err
	}

	path, ok1 := args[0].(*StringValue)
	cliArgs, ok2 := args[1].(*ListValue)
	stdin, ok3 := args[2].(*StringValue)
	if !ok1 || !ok2 || !ok3 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call exec(%s, %s, %s)", args[0], args[1], args[2]),
		}
	}

	argsList := make([]string, len(*cliArgs))
	for i, arg := range *cliArgs {
		if argStr, ok := arg.(*StringValue); ok {
			argsList[i] = argStr.stringContent()
		} else {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Mismatched types in call exec, arguments must be strings in %s", cliArgs),
			}
		}
	}

	cmd := exec.Command(path.stringContent(), argsList...)
	stdoutBuf := bytes.Buffer{}
	stderrBuf := bytes.Buffer{}
	cmd.Stdin = strings.NewReader(stdin.stringContent())
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Start()
	if err != nil {
		return errObj(fmt.Sprintf("Could not start command in exec(): %s", err.Error())), nil
	}

	err = cmd.Wait()
	exitCode := 0
	if err != nil {
		// if there is an err but err is just ExitErr, this means the process
		// ran successfully but exited with an error code. We consider this ok
		// and keep going.
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
		}
	}

	stdout, err := io.ReadAll(&stdoutBuf)
	if err != nil {
		return errObj(fmt.Sprintf("Could not read stdout from exec(): %s", err.Error())), nil
	}
	stdoutVal := StringValue(stdout)
	stderr, err := io.ReadAll(&stderrBuf)
	if err != nil {
		return errObj(fmt.Sprintf("Could not read stderr from exec(): %s", err.Error())), nil
	}
	stderrVal := StringValue(stderr)

	return ObjectValue{
		"type":   AtomValue("end"),
		"status": IntValue(exitCode),
		"stdout": &stdoutVal,
		"stderr": &stderrVal,
	}, nil
}

func (c *Context) oakSysProc(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("sysproc", args, 2); err != nil {
		return nil, err
	}

	library, ok1 := args[0].(*StringValue)
	name, ok2 := args[1].(*StringValue)
	if !ok1 || !ok2 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call sysproc(%s, %s)", args[0], args[1]),
		}
	}

	addr, err := lookupSysProc(library.stringContent(), name.stringContent())
	if err != nil {
		return errObj(fmt.Sprintf("Could not resolve procedure %s!%s: %s",
			library.stringContent(), name.stringContent(), err.Error())), nil
	}

	return ObjectValue{
		"type":    AtomValue("proc"),
		"library": MakeString(library.stringContent()),
		"name":    MakeString(name.stringContent()),
		"addr":    IntValue(addr),
	}, nil
}

func (c *Context) oakUTF16(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("utf16", args, 1); err != nil {
		return nil, err
	}

	arg, ok := args[0].(*StringValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call utf16(%s)", args[0]),
		}
	}

	encoded := utf16pkg.Encode([]rune(arg.stringContent()))
	buf := make([]byte, 0, (len(encoded)+1)*2)
	for _, word := range encoded {
		buf = append(buf, byte(word), byte(word>>8))
	}
	buf = append(buf, 0, 0)

	val := StringValue(buf)
	return &val, nil
}

func (c *Context) oakBits(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("bits", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *ListValue:
		buf := make([]byte, len(*arg))
		for i, val := range *arg {
			intVal, ok := val.(IntValue)
			if !ok || intVal < 0 || intVal > 255 {
				return nil, &runtimeError{
					reason: fmt.Sprintf("bits(list) expects byte values 0-255, got %s", val),
				}
			}
			buf[i] = byte(intVal)
		}
		bitsVal := StringValue(buf)
		return &bitsVal, nil
	case *StringValue:
		bytes := make(ListValue, len(*arg))
		for i, b := range *arg {
			bytes[i] = IntValue(b)
		}
		return &bytes, nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("bits() expects a list of bytes or a byte string, got %s", args[0]),
		}
	}
}

// oakPointer converts its argument to a pointer value.  It accepts an
// integer or a byte string; the latter behaves exactly like addr(), returning
// a pointer to the first element. Passing a pointer already returns it
// unchanged. The function exists purely to make pointer conversions explicit
// rather than relying on implicit casting from integers.
func (c *Context) oakPointer(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("pointer", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case IntValue:
		if arg < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("pointer() expects a non-negative integer, got %d", arg),
			}
		}
		return PointerValue(uintptr(arg)), nil
	case *StringValue:
		if len(*arg) == 0 {
			return PointerValue(0), nil
		}
		return PointerValue(uintptr(unsafe.Pointer(&(*arg)[0]))), nil
	case AtomValue:
		if arg == "" {
			return PointerValue(0), nil
		}

		atomName := string(arg)
		c.eng.nameRefLock.Lock()
		if ptr, ok := c.eng.namePtrs[atomName]; ok {
			c.eng.nameRefLock.Unlock()
			return PointerValue(ptr), nil
		}

		buf := StringValue([]byte(atomName))
		ptr := uintptr(unsafe.Pointer(&buf[0]))
		if c.eng.nameRefs == nil {
			c.eng.nameRefs = map[uintptr]*StringValue{}
		}
		if c.eng.namePtrs == nil {
			c.eng.namePtrs = map[string]uintptr{}
		}
		c.eng.nameRefs[ptr] = &buf
		c.eng.namePtrs[atomName] = ptr
		c.eng.nameRefLock.Unlock()

		return PointerValue(ptr), nil
	case PointerValue:
		return arg, nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("pointer() expects int, atom, or byte string, got %s", args[0]),
		}
	}
}

func (c *Context) memWriteBytes(val Value) (*StringValue, *runtimeError) {
	switch data := val.(type) {
	case *StringValue:
		return data, nil
	case AtomValue:
		s := StringValue([]byte(string(data)))
		return &s, nil
	case *ListValue:
		buf := make([]byte, len(*data))
		for i, el := range *data {
			byteVal, ok := el.(IntValue)
			if !ok || byteVal < 0 || byteVal > 255 {
				return nil, &runtimeError{reason: fmt.Sprintf("memwrite list payload expects byte values 0-255, got %s", el)}
			}
			buf[i] = byte(byteVal)
		}
		s := StringValue(buf)
		return &s, nil
	default:
		return nil, &runtimeError{reason: fmt.Sprintf("memwrite payload must be byte string, atom, or byte list, got %s", val)}
	}
}

func (c *Context) setNameRef(name AtomValue, data *StringValue) PointerValue {
	nameKey := string(name)
	if nameKey == "" {
		return PointerValue(0)
	}
	if len(*data) == 0 {
		c.eng.nameRefLock.Lock()
		if oldPtr, ok := c.eng.namePtrs[nameKey]; ok {
			delete(c.eng.nameRefs, oldPtr)
			delete(c.eng.namePtrs, nameKey)
		}
		c.eng.nameRefLock.Unlock()
		return PointerValue(0)
	}

	buf := StringValue(append([]byte(nil), (*data)...))
	ptr := uintptr(unsafe.Pointer(&buf[0]))

	c.eng.nameRefLock.Lock()
	if c.eng.nameRefs == nil {
		c.eng.nameRefs = map[uintptr]*StringValue{}
	}
	if c.eng.namePtrs == nil {
		c.eng.namePtrs = map[string]uintptr{}
	}
	if oldPtr, ok := c.eng.namePtrs[nameKey]; ok {
		delete(c.eng.nameRefs, oldPtr)
	}
	c.eng.nameRefs[ptr] = &buf
	c.eng.namePtrs[nameKey] = ptr
	c.eng.nameRefLock.Unlock()

	return PointerValue(ptr)
}

func (c *Context) oakAddr(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("addr", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		if len(*arg) == 0 {
			return PointerValue(0), nil
		}
		return PointerValue(uintptr(unsafe.Pointer(&(*arg)[0]))), nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("addr() expects a byte string, got %s", args[0]),
		}
	}
}

func (c *Context) oakMemRead(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("memread", args, 2); err != nil {
		return nil, err
	}

	// memread now accepts either an integer or pointer as its first argument.
	addrVal := args[0]
	length, ok2 := args[1].(IntValue)
	if !ok2 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call memread(%s, %s)", args[0], args[1]),
		}
	}
	var addr uintptr
	switch v := addrVal.(type) {
	case IntValue:
		if v < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("memread address must be non-negative, got %d", v),
			}
		}
		addr = uintptr(v)
	case PointerValue:
		addr = uintptr(v)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("memread address must be int or pointer, got %s", addrVal),
		}
	}
	if length < 0 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("memread length must be non-negative, got %d", length),
		}
	}
	if length == 0 {
		return MakeString(""), nil
	}
	if addr == 0 {
		return nil, &runtimeError{
			reason: "memread cannot read from null address",
		}
	}

	region := rawMemoryRegion(uintptr(addr), int(length))
	buf := append([]byte(nil), region...)
	value := StringValue(buf)
	return &value, nil
}

func (c *Context) oakMemWrite(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("memwrite", args, 2); err != nil {
		return nil, err
	}

	addrVal := args[0]
	data, dataErr := c.memWriteBytes(args[1])
	if dataErr != nil {
		return nil, dataErr
	}
	var addr uintptr
	switch v := addrVal.(type) {
	case IntValue:
		if v < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("memwrite address must be non-negative, got %d", v),
			}
		}
		addr = uintptr(v)
	case PointerValue:
		addr = uintptr(v)
	case AtomValue:
		ptr := c.setNameRef(v, data)
		if ptr == 0 {
			return IntValue(0), nil
		}
		return IntValue(len(*data)), nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("memwrite address must be int, pointer, or atom reference, got %s", addrVal),
		}
	}
	if len(*data) == 0 {
		return IntValue(0), nil
	}
	if addr == 0 {
		return nil, &runtimeError{
			reason: "memwrite cannot write to null address",
		}
	}

	region := rawMemoryRegion(addr, len(*data))
	return IntValue(copy(region, *data)), nil
}

func (c *Context) oakSyscall(args []Value) (Value, *runtimeError) {
	if len(args) < 1 {
		return syscallErrObj("syscall requires at least 1 argument (procedure or address)"), nil
	}

	var syscallTarget uintptr
	switch target := args[0].(type) {
	case IntValue:
		if target <= 0 {
			return syscallErrObj("invalid syscall target"), nil
		}
		syscallTarget = uintptr(target)
	case ObjectValue:
		if typeVal, ok := target["type"].(AtomValue); ok && typeVal == AtomValue("error") {
			return target, nil
		}

		typeVal, ok := target["type"].(AtomValue)
		if !ok || typeVal != AtomValue("proc") {
			return syscallErrObj(fmt.Sprintf("syscall target must be an integer or proc, got %s", args[0])), nil
		}

		addr, ok := target["addr"].(IntValue)
		if !ok || addr <= 0 {
			return syscallErrObj("invalid syscall target"), nil
		}
		syscallTarget = uintptr(addr)
	default:
		return syscallErrObj(fmt.Sprintf("syscall target must be an integer or proc, got %s", args[0])), nil
	}

	var syscallArgs []uintptr
	for i := 1; i < len(args); i++ {
		switch arg := args[i].(type) {
		case IntValue:
			syscallArgs = append(syscallArgs, uintptr(arg))
		case PointerValue:
			syscallArgs = append(syscallArgs, uintptr(arg))
		case BoolValue:
			if arg {
				syscallArgs = append(syscallArgs, 1)
			} else {
				syscallArgs = append(syscallArgs, 0)
			}
		case NullValue:
			syscallArgs = append(syscallArgs, 0)
		case *StringValue:
			if len(*arg) == 0 {
				syscallArgs = append(syscallArgs, 0)
				continue
			}
			syscallArgs = append(syscallArgs, uintptr(unsafe.Pointer(&(*arg)[0])))
		default:
			return syscallErrObj(fmt.Sprintf(
				"syscall argument %d must be int, pointer, bool, string, or ?, got %s", i, arg,
			)), nil
		}
	}

	r1, r2, err := oakSyscallN(syscallTarget, syscallArgs...)
	runtime.KeepAlive(args)

	if err != 0 {
		return ObjectValue{
			"type":  AtomValue("error"),
			"error": MakeString(err.Error()),
			"errno": IntValue(err),
			"r1":    IntValue(r1),
			"r2":    IntValue(r2),
		}, nil
	}

	return ObjectValue{
		"type":  AtomValue("success"),
		"errno": IntValue(0),
		"r1":    IntValue(r1),
		"r2":    IntValue(r2),
	}, nil
}

func (c *Context) oakWinMsgLoop(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("win_msg_loop", args, 1); err != nil {
		return nil, err
	}

	if runtime.GOOS != "windows" {
		return errObj("win_msg_loop is only supported on Windows"), nil
	}

	var hwnd uintptr
	switch v := args[0].(type) {
	case IntValue:
		if v < 0 {
			return nil, &runtimeError{reason: fmt.Sprintf("win_msg_loop expects a non-negative handle, got %d", v)}
		}
		hwnd = uintptr(v)
	case PointerValue:
		hwnd = uintptr(v)
	default:
		return nil, &runtimeError{reason: fmt.Sprintf("win_msg_loop expects int or pointer, got %s", args[0])}
	}

	getMessageAddr, err := lookupSysProc("user32.dll", "GetMessageW")
	if err != nil {
		return errObj("Could not resolve user32.dll!GetMessageW: " + err.Error()), nil
	}
	translateMessageAddr, err := lookupSysProc("user32.dll", "TranslateMessage")
	if err != nil {
		return errObj("Could not resolve user32.dll!TranslateMessage: " + err.Error()), nil
	}
	dispatchMessageAddr, err := lookupSysProc("user32.dll", "DispatchMessageW")
	if err != nil {
		return errObj("Could not resolve user32.dll!DispatchMessageW: " + err.Error()), nil
	}
	isWindowAddr, err := lookupSysProc("user32.dll", "IsWindow")
	if err != nil {
		return errObj("Could not resolve user32.dll!IsWindow: " + err.Error()), nil
	}

	msgBuf := make([]byte, 48)
	msgPtr := uintptr(unsafe.Pointer(&msgBuf[0]))

	for {
		alive, _, _ := oakSyscallN(isWindowAddr, hwnd)
		if alive == 0 {
			return IntValue(0), nil
		}

		r1, _, callErr := oakSyscallN(getMessageAddr, msgPtr, 0, 0, 0)
		if int32(r1) == -1 {
			return ObjectValue{
				"type":  AtomValue("error"),
				"error": MakeString(callErr.Error()),
				"errno": IntValue(callErr),
				"r1":    IntValue(r1),
				"r2":    IntValue(0),
			}, nil
		}
		if r1 == 0 {
			return IntValue(0), nil
		}

		oakSyscallN(translateMessageAddr, msgPtr)
		oakSyscallN(dispatchMessageAddr, msgPtr)
	}
}

func (c *Context) oakGo(args []Value) (Value, *runtimeError) {
	if len(args) < 1 {
		return nil, &runtimeError{
			reason: "go requires at least 1 argument (function to run)",
		}
	}

	fn := args[0]
	if !c.isGoTargetValue(fn) {
		return nil, &runtimeError{
			reason: fmt.Sprintf("go argument must be a function, got %s", fn),
		}
	}

	fnArgs := args[1:] // remaining arguments

	// Fork a lightweight context so this goroutine can evaluate code in
	// parallel with other goroutines.  The forked context shares the engine
	// (channels, imports, WaitGroup) but runs evaluations without competing
	// for asyncEvalLock.  Scope variable reads are safe (per-scope RWMutex);
	// local bindings created during evaluation stay isolated in new child
	// scopes.
	fork := c.forkContext()

	c.eng.Add(1)
	go func() {
		defer c.eng.Done()
		_, err := fork.EvalFnValueParallel(fn, false, fnArgs...)
		if err != nil {
			c.eng.reportErr(err)
		}
	}()

	return null, nil
}

func (c *Context) oakLockThread(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("lock_thread", args, 0); err != nil {
		return nil, err
	}
	runtime.LockOSThread()
	return null, nil
}

func (c *Context) oakUnlockThread(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("unlock_thread", args, 0); err != nil {
		return nil, err
	}
	runtime.UnlockOSThread()
	return null, nil
}

func (c *Context) oakMakeChan(args []Value) (Value, *runtimeError) {
	if len(args) > 1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("make_chan takes at most 1 argument, got %d", len(args)),
		}
	}

	bufSize := IntValue(0)
	if len(args) == 1 {
		sizeArg, ok := args[0].(IntValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Mismatched types in call make_chan(%s)", args[0]),
			}
		}
		if sizeArg < 0 {
			return nil, &runtimeError{
				reason: fmt.Sprintf("make_chan capacity must be non-negative, got %d", sizeArg),
			}
		}
		bufSize = sizeArg
	}

	c.eng.chanLock.Lock()
	id := c.eng.nextChanID
	c.eng.nextChanID++
	c.eng.chanMap[id] = make(chan Value, int(bufSize))
	c.eng.chanLock.Unlock()

	return ObjectValue{
		"type": AtomValue("channel"),
		"id":   IntValue(id),
		"cap":  bufSize,
	}, nil
}

func (c *Context) oakChanSend(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("chan_send", args, 2); err != nil {
		return nil, err
	}

	ch, err := c.getGoChan(args[0], "chan_send")
	if err != nil {
		return nil, err
	}

	value := args[1]
	if len(args) == 2 {
		c.Unlock()
		ch <- value
		c.Lock()
		return chanSentObj(), nil
	}

	if len(args) != 3 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("chan_send takes 2 arguments plus an optional callback, got %d", len(args)),
		}
	}

	callback := args[2]
	if !c.isCallableValue(callback) {
		return nil, &runtimeError{
			reason: fmt.Sprintf("chan_send callback must be a function, got %s", args[2]),
		}
	}

	c.eng.Add(1)
	go func() {
		defer c.eng.Done()
		ch <- value

		c.Lock()
		defer c.Unlock()
		_, err := c.EvalFnValueAsync(callback, false, chanSentObj())
		if err != nil {
			c.eng.reportErr(err)
		}
	}()

	return null, nil
}

func (c *Context) oakChanRecv(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("chan_recv", args, 1); err != nil {
		return nil, err
	}

	ch, err := c.getGoChan(args[0], "chan_recv")
	if err != nil {
		return nil, err
	}

	if len(args) == 1 {
		c.Unlock()
		value := <-ch
		c.Lock()
		return chanDataObj(value), nil
	}

	if len(args) != 2 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("chan_recv takes 1 argument plus an optional callback, got %d", len(args)),
		}
	}

	callback := args[1]
	if !c.isCallableValue(callback) {
		return nil, &runtimeError{
			reason: fmt.Sprintf("chan_recv callback must be a function, got %s", args[1]),
		}
	}

	c.eng.Add(1)
	go func() {
		defer c.eng.Done()
		value := <-ch

		c.Lock()
		defer c.Unlock()
		_, err := c.EvalFnValueAsync(callback, false, chanDataObj(value))
		if err != nil {
			c.eng.reportErr(err)
		}
	}()

	return null, nil
}

var inputReaderInit sync.Once
var inputReader *bufio.Reader

func initInputReader() {
	inputReader = bufio.NewReader(os.Stdin)
}

func (c *Context) oakInput(_ []Value) (Value, *runtimeError) {
	inputReaderInit.Do(initInputReader)
	str, err := inputReader.ReadString('\n')
	if err == io.EOF {
		return ObjectValue{
			"type":  AtomValue("error"),
			"error": MakeString("EOF"),
			// if any data was read before encountering EOF, ensure the caller
			// still gets that data.
			"data": MakeString(str),
		}, nil
	} else if err != nil {
		return errObj(fmt.Sprintf("Could not read input: %s", err.Error())), nil
	}

	inputStr := strings.TrimSuffix(str, "\n")

	return ObjectValue{
		"type": AtomValue("data"),
		"data": MakeString(inputStr),
	}, nil
}

func (c *Context) oakPrint(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("print", args, 1); err != nil {
		return nil, err
	}

	outputString, ok := args[0].(*StringValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Unexpected argument to print: %s", args[0]),
		}
	}

	n, _ := os.Stdout.Write(*outputString)
	return IntValue(n), nil
}

func (c *Context) oakLs(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("ls", args, 1); err != nil {
		return nil, err
	}

	dirPath, ok1 := args[0].(*StringValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call ls(%s)", args[0]),
		}
	}

	fileInfos, err := ioutil.ReadDir(dirPath.stringContent())
	if err != nil {
		return errObj(fmt.Sprintf("Could not list directory %s: %s", dirPath.stringContent(), err.Error())), nil
	}

	fileList := make(ListValue, len(fileInfos))
	for i, fi := range fileInfos {
		fileList[i] = ObjectValue{
			"name": MakeString(fi.Name()),
			"len":  IntValue(fi.Size()),
			"dir":  BoolValue(fi.IsDir()),
			"mod":  IntValue(fi.ModTime().Unix()),
		}
	}

	return ObjectValue{
		"type": AtomValue("data"),
		"data": &fileList,
	}, nil
}

func (c *Context) oakRm(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("rm", args, 1); err != nil {
		return nil, err
	}

	rmPath, ok1 := args[0].(*StringValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call rm(%s)", args[0]),
		}
	}

	err := os.RemoveAll(rmPath.stringContent())
	if err != nil {
		return errObj(fmt.Sprintf("Could not remove %s: %s", rmPath.stringContent(), err.Error())), nil
	}

	return ObjectValue{
		"type": AtomValue("end"),
	}, nil
}

func (c *Context) oakMkdir(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("mkdir", args, 1); err != nil {
		return nil, err
	}

	dirPath, ok1 := args[0].(*StringValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call mkdir(%s)", args[0]),
		}
	}

	err := os.MkdirAll(dirPath.stringContent(), 0755)
	if err != nil {
		return errObj(fmt.Sprintf("Could not make a new directory %s: %s", dirPath.stringContent(), err.Error())), nil
	}

	return ObjectValue{
		"type": AtomValue("end"),
	}, nil
}

func (c *Context) oakStat(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("stat", args, 1); err != nil {
		return nil, err
	}

	statPath, ok1 := args[0].(*StringValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call stat(%s)", args[0]),
		}
	}

	fileInfo, err := os.Stat(statPath.stringContent())
	if err != nil {
		if os.IsNotExist(err) {
			return ObjectValue{
				"type": AtomValue("data"),
				"data": null,
			}, nil
		}
		return errObj(fmt.Sprintf("Could not stat file %s: %s", statPath.stringContent(), err.Error())), nil
	}

	return ObjectValue{
		"type": AtomValue("data"),
		"data": ObjectValue{
			"name": MakeString(fileInfo.Name()),
			"len":  IntValue(fileInfo.Size()),
			"dir":  BoolValue(fileInfo.IsDir()),
			"mod":  IntValue(fileInfo.ModTime().Unix()),
		},
	}, nil
}

func (c *Context) oakOpen(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("open", args, 1); err != nil {
		return nil, err
	}

	// flags arg is optional
	if len(args) < 2 {
		args = append(args, AtomValue("readwrite"))
	}

	// perm arg is optional
	if len(args) < 3 {
		args = append(args, IntValue(0644))
	}

	pathString, ok1 := args[0].(*StringValue)
	flagsAtom, ok2 := args[1].(AtomValue)
	permInt, ok3 := args[2].(IntValue)
	if !ok1 || !ok2 || !ok3 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call open(%s, %s, %s)", args[0], args[1], args[2]),
		}
	}

	var flags int
	switch string(flagsAtom) {
	case "readonly":
		flags = os.O_RDONLY
	case "readwrite":
		flags = os.O_RDWR | os.O_CREATE
	case "append":
		flags = os.O_RDWR | os.O_CREATE | os.O_APPEND
	case "truncate":
		flags = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Invalid flag for open(): %s", flagsAtom),
		}
	}

	file, err := os.OpenFile(pathString.stringContent(), flags, os.FileMode(permInt))
	if err != nil {
		return errObj(fmt.Sprintf("Could not open file: %s", err.Error())), nil
	}

	fd := file.Fd()

	c.eng.fdLock.Lock()
	defer c.eng.fdLock.Unlock()
	c.eng.fileMap[fd] = file

	return ObjectValue{
		"type": AtomValue("file"),
		"fd":   IntValue(fd),
	}, nil
}

func (c *Context) oakClose(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("close", args, 1); err != nil {
		return nil, err
	}

	fdInt, ok1 := args[0].(IntValue)
	if !ok1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call close(%s)", args[0]),
		}
	}

	c.eng.fdLock.Lock()
	defer c.eng.fdLock.Unlock()
	file, ok := c.eng.fileMap[uintptr(fdInt)]

	if !ok {
		return errObj(fmt.Sprintf("Unknown fd %d", fdInt)), nil
	}

	err := file.Close()
	if err != nil {
		return errObj(fmt.Sprintf("Could not close file: %s", err.Error())), nil
	}

	delete(c.eng.fileMap, uintptr(fdInt))

	return ObjectValue{
		"type": AtomValue("end"),
	}, nil
}

func (c *Context) oakRead(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("read", args, 3); err != nil {
		return nil, err
	}

	fdInt, ok1 := args[0].(IntValue)
	offsetInt, ok2 := args[1].(IntValue)
	lengthInt, ok3 := args[2].(IntValue)
	if !ok1 || !ok2 || !ok3 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call read(%s, %s, %s)", args[0], args[1], args[2]),
		}
	}

	c.eng.fdLock.Lock()
	file, ok := c.eng.fileMap[uintptr(fdInt)]
	c.eng.fdLock.Unlock()

	if !ok {
		return errObj(fmt.Sprintf("Unknown fd %d", fdInt)), nil
	}

	offset := int64(offsetInt)
	length := int64(lengthInt)

	_, err := file.Seek(offset, 0)
	if err != nil {
		return errObj(fmt.Sprintf("Error reading file during seek: %s", err.Error())), nil
	}

	readBuf := make([]byte, length)
	count, err := file.Read(readBuf)
	if err != nil && err != io.EOF {
		return errObj(fmt.Sprintf("Error reading file: %s", err.Error())), nil
	}

	fileData := StringValue(readBuf[:count])
	return ObjectValue{
		"type": AtomValue("data"),
		"data": &fileData,
	}, nil
}

func (c *Context) oakWrite(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("write", args, 3); err != nil {
		return nil, err
	}

	fdInt, ok1 := args[0].(IntValue)
	offsetInt, ok2 := args[1].(IntValue)
	dataString, ok3 := args[2].(*StringValue)
	if !ok1 || !ok2 || !ok3 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call write(%s, %s, %s)", args[0], args[1], args[2]),
		}
	}

	c.eng.fdLock.Lock()
	file, ok := c.eng.fileMap[uintptr(fdInt)]
	c.eng.fdLock.Unlock()
	if !ok {
		return errObj(fmt.Sprintf("Unknown fd %d", fdInt)), nil
	}

	offset := int64(offsetInt)
	writeBuf := []byte(*dataString)

	var err error
	if offset == -1 {
		_, err = file.Seek(0, 2) // "2" is relative to end of file
	} else {
		_, err = file.Seek(offset, 0)
	}
	if err != nil {
		return errObj(fmt.Sprintf("Error writing file during seek: %s", err.Error())), nil
	}

	_, err = file.Write(writeBuf)
	if err != nil && err != io.EOF {
		return errObj(fmt.Sprintf("Error writing file: %s", err.Error())), nil
	}

	return ObjectValue{
		"type": AtomValue("end"),
	}, nil
}

type oakHTTPHandler struct {
	ctx         *Context
	oakCallback Value
}

func (h oakHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := h.ctx
	cb := h.oakCallback

	// unmarshal request
	method := r.Method
	url := r.URL.String()
	headers := ObjectValue{}
	for key, values := range r.Header {
		headers[key] = MakeString(strings.Join(values, ","))
	}
	var body *StringValue
	if r.ContentLength == 0 {
		body = MakeString("")
	} else {
		bodyBuf, err := io.ReadAll(r.Body)
		if err != nil {
			ctx.Lock()
			_, rtErr := ctx.EvalFnValueAsync(cb, false, errObj(
				fmt.Sprintf("Could not read request in listen(), %s", err.Error()),
			))
			ctx.Unlock()

			if rtErr != nil {
				ctx.eng.reportErr(rtErr)
			}
		}
		bodyStr := StringValue(bodyBuf)
		body = &bodyStr
	}

	// construct request object to pass to Oak, call handler
	responseEnded := false
	responses := make(chan Value, 1)
	endHandler := func(args []Value) (Value, *runtimeError) {
		if err := ctx.requireArgLen("listen/end", args, 1); err != nil {
			return nil, err
		}

		if responseEnded {
			ctx.eng.reportErr(&runtimeError{
				reason: "listen/end called more than once",
			})
		}

		responseEnded = true
		responses <- args[0]

		return null, nil
	}

	go func() {
		ctx.Lock()
		defer ctx.Unlock()

		_, err := ctx.EvalFnValueAsync(cb, false, ObjectValue{
			"type": AtomValue("req"),
			"req": ObjectValue{
				"method":  MakeString(method),
				"url":     MakeString(url),
				"headers": headers,
				"body":    body,
			},
			"end": BuiltinFnValue{
				name: "end",
				fn:   endHandler,
			},
		})
		if err != nil {
			ctx.eng.reportErr(err)
		}
	}()

	// validate responses
	resp := <-responses
	rsp, isObject := resp.(ObjectValue)
	if !isObject {
		ctx.eng.reportErr(&runtimeError{
			reason: fmt.Sprintf("listen/end should return a response, got %s", resp),
		})
		return
	}

	// unmarshal response from the return value
	// response = { status, headers, body }
	statusVal, okStatus := rsp["status"]
	headersVal, okHeaders := rsp["headers"]
	bodyVal, okBody := rsp["body"]

	resStatus, okStatus := statusVal.(IntValue)
	resHeaders, okHeaders := headersVal.(ObjectValue)
	resBody, okBody := bodyVal.(*StringValue)

	if !okStatus || !okHeaders || !okBody {
		ctx.eng.reportErr(&runtimeError{
			reason: fmt.Sprintf("listen/end returned malformed response, %s", rsp),
		})
		return
	}

	// write values to response
	// Content-Length is automatically set for us by Go
	for k, v := range resHeaders {
		if str, isStr := v.(*StringValue); isStr {
			w.Header().Set(k, str.stringContent())
		} else {
			ctx.eng.reportErr(&runtimeError{
				reason: fmt.Sprintf("Could not set response header, value %s was not a string", v),
			})
			return
		}
	}

	code := int(resStatus)
	// guard against invalid HTTP codes, which cause Go panics
	// https://golang.org/src/net/http/server.go
	if code < 100 || code > 599 {
		ctx.eng.reportErr(&runtimeError{
			reason: fmt.Sprintf("Could not set response status code, code %d is not valid", code),
		})
		return
	}

	// status code write must follow all other header writes, since it sends
	// the status
	w.WriteHeader(int(resStatus))
	_, err := w.Write(*resBody)
	if err != nil {
		ctx.Lock()
		defer ctx.Unlock()

		_, rtErr := ctx.EvalFnValueAsync(cb, false, errObj(
			fmt.Sprintf("Error writing request body in listen/end: %s", err.Error()),
		))
		if rtErr != nil {
			ctx.eng.reportErr(rtErr)
		}
	}
}

func (ctx *Context) oakListen(args []Value) (Value, *runtimeError) {
	if err := ctx.requireArgLen("listen", args, 2); err != nil {
		return nil, err
	}

	host, ok1 := args[0].(*StringValue)
	cb := args[1]
	ok2 := ctx.isCallableValue(cb)
	if !ok1 || !ok2 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call listen(%s, %s)", args[0], args[1]),
		}
	}

	sendErr := func(msg string) {
		ctx.Lock()
		defer ctx.Unlock()

		_, err2 := ctx.EvalFnValueAsync(cb, false, errObj(msg))
		if err2 != nil {
			ctx.eng.reportErr(err2)
		}
	}

	server := &http.Server{
		Addr: host.stringContent(),
		Handler: oakHTTPHandler{
			ctx:         ctx,
			oakCallback: cb,
		},
	}

	ctx.eng.Add(1)
	go func() {
		defer ctx.eng.Done()
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			sendErr(fmt.Sprintf("Error starting http server in listen(): %s", err.Error()))
		}
	}()

	closer := func(_ []Value) (Value, *runtimeError) {
		// attempt graceful shutdown, concurrently, without blocking Oak
		// evaluation thread
		ctx.eng.Add(1)
		go func() {
			defer ctx.eng.Done()

			err := server.Shutdown(context.Background())
			if err != nil {
				sendErr(fmt.Sprintf("Could not close server in listen/close: %s", err.Error()))
			}
		}()

		return null, nil
	}

	return BuiltinFnValue{
		name: "close",
		fn:   closer,
	}, nil
}

func (c *Context) oakReq(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("req", args, 1); err != nil {
		return nil, err
	}

	argErr := runtimeError{
		reason: fmt.Sprintf("Mismatched types in call req(%s)", args[0]),
	}

	data, ok1 := args[0].(ObjectValue)
	if !ok1 {
		return nil, &argErr
	}

	// unmarshal request data
	methodVal, ok1 := data["method"]
	urlVal, ok2 := data["url"]
	headersVal, ok3 := data["headers"]
	bodyVal, ok4 := data["body"]

	// default args
	if !ok1 {
		methodVal = MakeString("GET")
		ok1 = true
	}
	if !ok3 {
		headersVal = ObjectValue{}
		ok3 = true
	}
	if !ok4 {
		bodyVal = MakeString("")
		ok4 = true
	}

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, &argErr
	}

	method, ok1 := methodVal.(*StringValue)
	url, ok2 := urlVal.(*StringValue)
	headers, ok3 := headersVal.(ObjectValue)
	body, ok4 := bodyVal.(*StringValue)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, &argErr
	}

	client := &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			// do not follow redirects
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(
		method.stringContent(),
		url.stringContent(),
		strings.NewReader(body.stringContent()),
	)
	if err != nil {
		return errObj(fmt.Sprintf("Could not initialize request in req(): %s", err.Error())), nil
	}

	// construct headers
	// Content-Length is automatically set for us by Go
	req.Header.Set("User-Agent", "") // remove Go's default user agent header
	for k, v := range headers {
		if valStr, ok := v.(*StringValue); ok {
			req.Header.Set(k, valStr.stringContent())
		} else {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Could not set request header, value %s is not a string", v),
			}
		}
	}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return errObj(fmt.Sprintf("Could not send request: %s", err.Error())), nil
	}
	defer resp.Body.Close()

	respStatus := IntValue(resp.StatusCode)
	respHeaders := ObjectValue{}
	for key, values := range resp.Header {
		respHeaders[key] = MakeString(strings.Join(values, ","))
	}

	var respBody *StringValue
	if resp.ContentLength == 0 {
		respBody = MakeString("")
	} else {
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return errObj(fmt.Sprintf("Could not read response: %s", err.Error())), nil
		}
		strBuf := StringValue(buf)
		respBody = &strBuf
	}

	return ObjectValue{
		"type": AtomValue("resp"),
		"resp": ObjectValue{
			"status":  respStatus,
			"headers": respHeaders,
			"body":    respBody,
		},
	}, nil
}

type socketOptions struct {
	useTLS             bool
	serverTLS          bool
	serverName         string
	insecureSkipVerify bool
	certFile           string
	keyFile            string
}

func stringField(value Value, name string) (string, bool, *runtimeError) {
	if value == nil {
		return "", false, nil
	}
	str, ok := value.(*StringValue)
	if !ok {
		return "", false, &runtimeError{reason: fmt.Sprintf("%s must be a string, got %s", name, value)}
	}
	return str.stringContent(), true, nil
}

func boolField(value Value, name string) (bool, bool, *runtimeError) {
	if value == nil {
		return false, false, nil
	}
	b, ok := value.(BoolValue)
	if !ok {
		return false, false, &runtimeError{reason: fmt.Sprintf("%s must be a bool, got %s", name, value)}
	}
	return bool(b), true, nil
}

func parseSocketOptions(arg Value, name string) (socketOptions, *runtimeError) {
	if arg == nil {
		return socketOptions{}, nil
	}
	if _, ok := arg.(NullValue); ok {
		return socketOptions{}, nil
	}

	obj, ok := arg.(ObjectValue)
	if !ok {
		return socketOptions{}, &runtimeError{reason: fmt.Sprintf("%s must be an object, got %s", name, arg)}
	}

	options := socketOptions{}
	if value, ok := obj["tls"]; ok {
		parsed, _, err := boolField(value, name+".tls")
		if err != nil {
			return socketOptions{}, err
		}
		options.useTLS = parsed
	}
	if value, ok := obj["server"]; ok {
		parsed, _, err := boolField(value, name+".server")
		if err != nil {
			return socketOptions{}, err
		}
		options.serverTLS = parsed
	}
	if value, ok := obj["serverName"]; ok {
		parsed, _, err := stringField(value, name+".serverName")
		if err != nil {
			return socketOptions{}, err
		}
		options.serverName = parsed
	}
	if value, ok := obj["insecureSkipVerify"]; ok {
		parsed, _, err := boolField(value, name+".insecureSkipVerify")
		if err != nil {
			return socketOptions{}, err
		}
		options.insecureSkipVerify = parsed
	}
	if value, ok := obj["certFile"]; ok {
		parsed, _, err := stringField(value, name+".certFile")
		if err != nil {
			return socketOptions{}, err
		}
		options.certFile = parsed
	}
	if value, ok := obj["keyFile"]; ok {
		parsed, _, err := stringField(value, name+".keyFile")
		if err != nil {
			return socketOptions{}, err
		}
		options.keyFile = parsed
	}

	return options, nil
}

func tlsConfigFromOptions(options socketOptions) (*tls.Config, *runtimeError) {
	config := &tls.Config{
		InsecureSkipVerify: options.insecureSkipVerify,
		ServerName:         options.serverName,
	}

	if options.serverTLS {
		if options.certFile == "" || options.keyFile == "" {
			return nil, &runtimeError{reason: "TLS server mode requires certFile and keyFile"}
		}
		cert, err := tls.LoadX509KeyPair(options.certFile, options.keyFile)
		if err != nil {
			return nil, &runtimeError{reason: fmt.Sprintf("Could not load TLS certificate: %s", err.Error())}
		}
		config.Certificates = []tls.Certificate{cert}
	}

	return config, nil
}

func socketInfo(socket ObjectValue, conn net.Conn, isTLS bool) ObjectValue {
	return ObjectValue{
		"type":   AtomValue("ok"),
		"socket": socket,
		"remote": MakeString(conn.RemoteAddr().String()),
		"local":  MakeString(conn.LocalAddr().String()),
		"tls":    BoolValue(isTLS),
	}
}

func (c *Context) oakSocketConnect(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_connect", args, 1); err != nil {
		return nil, err
	}

	address, ok := args[0].(*StringValue)
	if !ok {
		return nil, &runtimeError{reason: fmt.Sprintf("Mismatched types in call socket_connect(%s)", args[0])}
	}

	options := socketOptions{}
	if len(args) >= 2 {
		parsed, err := parseSocketOptions(args[1], "socket_connect options")
		if err != nil {
			return nil, err
		}
		options = parsed
	}

	var conn net.Conn
	var dialErr error
	c.Unlock()
	if options.useTLS {
		config, err := tlsConfigFromOptions(options)
		c.Lock()
		if err != nil {
			return nil, err
		}
		c.Unlock()
		conn, dialErr = tls.Dial("tcp", address.stringContent(), config)
		c.Lock()
	} else {
		conn, dialErr = net.Dial("tcp", address.stringContent())
		c.Lock()
	}
	if dialErr != nil {
		return errObj(fmt.Sprintf("Could not connect socket: %s", dialErr.Error())), nil
	}

	socket := storeSocket(conn, options.useTLS, false)
	return socketInfo(socket, conn, options.useTLS), nil
}

func (c *Context) oakSocketSend(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_send", args, 2); err != nil {
		return nil, err
	}

	socketConn, _, err := c.getSocket(args[0], "socket_send")
	if err != nil {
		return nil, err
	}

	data, ok := args[1].(*StringValue)
	if !ok {
		return nil, &runtimeError{reason: fmt.Sprintf("Second argument to socket_send must be a string, got %s", args[1])}
	}

	c.Unlock()
	written, writeErr := socketConn.conn.Write([]byte(data.stringContent()))
	c.Lock()
	if writeErr != nil {
		return errObj(fmt.Sprintf("Could not write socket data: %s", writeErr.Error())), nil
	}

	return ObjectValue{
		"type":  AtomValue("sent"),
		"bytes": IntValue(written),
	}, nil
}

func (c *Context) oakSocketRecv(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_recv", args, 2); err != nil {
		return nil, err
	}

	socketConn, id, err := c.getSocket(args[0], "socket_recv")
	if err != nil {
		return nil, err
	}

	size, ok := args[1].(IntValue)
	if !ok || size < 0 {
		return nil, &runtimeError{reason: fmt.Sprintf("Second argument to socket_recv must be a non-negative int, got %s", args[1])}
	}
	if size == 0 {
		return socketDataEvent([]byte{}), nil
	}

	buf := make([]byte, int(size))
	c.Unlock()
	count, readErr := socketConn.reader.Read(buf)
	c.Lock()
	if readErr != nil {
		if readErr == io.EOF {
			if count > 0 {
				return socketDataEvent(buf[:count]), nil
			}
			removeSocket(id)
			return socketClosedEvent(), nil
		}
		return errObj(fmt.Sprintf("Could not read socket data: %s", readErr.Error())), nil
	}

	return socketDataEvent(buf[:count]), nil
}

func (c *Context) oakSocketRecvExact(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_recv_exact", args, 2); err != nil {
		return nil, err
	}

	socketConn, id, err := c.getSocket(args[0], "socket_recv_exact")
	if err != nil {
		return nil, err
	}

	size, ok := args[1].(IntValue)
	if !ok || size < 0 {
		return nil, &runtimeError{reason: fmt.Sprintf("Second argument to socket_recv_exact must be a non-negative int, got %s", args[1])}
	}
	if size == 0 {
		return socketDataEvent([]byte{}), nil
	}

	buf := make([]byte, int(size))
	c.Unlock()
	_, readErr := io.ReadFull(socketConn.reader, buf)
	c.Lock()
	if readErr != nil {
		if readErr == io.EOF || readErr == io.ErrUnexpectedEOF {
			removeSocket(id)
		}
		return errObj(fmt.Sprintf("Could not read exact socket data: %s", readErr.Error())), nil
	}

	return socketDataEvent(buf), nil
}

func (c *Context) oakSocketRecvLine(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_recv_line", args, 1); err != nil {
		return nil, err
	}

	socketConn, id, err := c.getSocket(args[0], "socket_recv_line")
	if err != nil {
		return nil, err
	}

	c.Unlock()
	line, readErr := socketConn.reader.ReadString('\n')
	c.Lock()
	if readErr != nil {
		if readErr == io.EOF {
			if line != "" {
				return socketDataEvent([]byte(strings.TrimRight(line, "\r\n"))), nil
			}
			removeSocket(id)
			return socketClosedEvent(), nil
		}
		return errObj(fmt.Sprintf("Could not read socket line: %s", readErr.Error())), nil
	}

	return socketDataEvent([]byte(strings.TrimRight(line, "\r\n"))), nil
}

func (c *Context) oakSocketStartTLS(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_starttls", args, 1); err != nil {
		return nil, err
	}

	socketConn, id, err := c.getSocket(args[0], "socket_starttls")
	if err != nil {
		return nil, err
	}

	options := socketOptions{}
	if len(args) >= 2 {
		parsed, parseErr := parseSocketOptions(args[1], "socket_starttls options")
		if parseErr != nil {
			return nil, parseErr
		}
		options = parsed
	}

	config, configErr := tlsConfigFromOptions(options)
	if configErr != nil {
		return nil, configErr
	}

	var tlsConn *tls.Conn
	c.Unlock()
	if options.serverTLS {
		tlsConn = tls.Server(socketConn.conn, config)
	} else {
		tlsConn = tls.Client(socketConn.conn, config)
	}
	handshakeErr := tlsConn.Handshake()
	c.Lock()
	if handshakeErr != nil {
		return errObj(fmt.Sprintf("Could not upgrade socket to TLS: %s", handshakeErr.Error())), nil
	}

	socketConn.conn = tlsConn
	socketConn.reader = bufio.NewReader(tlsConn)
	socketConn.isTLS = true
	socketConn.serverTLS = options.serverTLS

	return socketInfo(socketObj(id), tlsConn, true), nil
}

func (c *Context) oakSocketClose(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("socket_close", args, 1); err != nil {
		return nil, err
	}

	socketConn, id, err := c.getSocket(args[0], "socket_close")
	if err != nil {
		return nil, err
	}

	c.Unlock()
	closeErr := socketConn.conn.Close()
	c.Lock()
	removeSocket(id)
	if closeErr != nil {
		return errObj(fmt.Sprintf("Could not close socket: %s", closeErr.Error())), nil
	}

	return socketClosedEvent(), nil
}

func (ctx *Context) oakSocketListen(args []Value) (Value, *runtimeError) {
	if err := ctx.requireArgLen("socket_listen", args, 2); err != nil {
		return nil, err
	}

	address, okAddress := args[0].(*StringValue)
	cb := args[1]
	okCb := ctx.isCallableValue(cb)
	if !okAddress || !okCb {
		return nil, &runtimeError{reason: fmt.Sprintf("Mismatched types in call socket_listen(%s, %s)", args[0], args[1])}
	}

	options := socketOptions{}
	if len(args) >= 3 {
		parsed, err := parseSocketOptions(args[2], "socket_listen options")
		if err != nil {
			return nil, err
		}
		options = parsed
	}

	sendErr := func(msg string) {
		ctx.Lock()
		defer ctx.Unlock()

		_, err2 := ctx.EvalFnValueAsync(cb, false, errObj(msg))
		if err2 != nil {
			ctx.eng.reportErr(err2)
		}
	}

	var listener net.Listener
	var listenErr error
	if options.useTLS {
		options.serverTLS = true
		config, err := tlsConfigFromOptions(options)
		if err != nil {
			return nil, err
		}
		listener, listenErr = tls.Listen("tcp", address.stringContent(), config)
	} else {
		listener, listenErr = net.Listen("tcp", address.stringContent())
	}
	if listenErr != nil {
		return errObj(fmt.Sprintf("Could not listen on socket: %s", listenErr.Error())), nil
	}

	ctx.eng.Add(1)
	go func() {
		defer ctx.eng.Done()
		for {
			conn, acceptErr := listener.Accept()
			if acceptErr != nil {
				if ne, ok := acceptErr.(net.Error); ok && ne.Temporary() {
					sendErr(fmt.Sprintf("Temporary socket accept error: %s", acceptErr.Error()))
					continue
				}
				if strings.Contains(strings.ToLower(acceptErr.Error()), "closed") {
					return
				}
				sendErr(fmt.Sprintf("Error accepting socket connection: %s", acceptErr.Error()))
				return
			}

			socket := storeSocket(conn, options.useTLS, options.useTLS)

			ctx.Lock()
			_, cbErr := ctx.EvalFnValueAsync(cb, false, ObjectValue{
				"type":   AtomValue("connect"),
				"socket": socket,
				"remote": MakeString(conn.RemoteAddr().String()),
				"local":  MakeString(conn.LocalAddr().String()),
				"tls":    BoolValue(options.useTLS),
			})
			ctx.Unlock()
			if cbErr != nil {
				ctx.eng.reportErr(cbErr)
			}
		}
	}()

	closer := func(_ []Value) (Value, *runtimeError) {
		ctx.eng.Add(1)
		go func() {
			defer ctx.eng.Done()
			if err := listener.Close(); err != nil {
				sendErr(fmt.Sprintf("Could not close socket listener: %s", err.Error()))
			}
		}()
		return null, nil
	}

	return BuiltinFnValue{
		name: "close",
		fn:   closer,
	}, nil
}

func (c *Context) oakWSDial(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("ws_dial", args, 1); err != nil {
		return nil, err
	}

	urlVal, ok := args[0].(*StringValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call ws_dial(%s)", args[0]),
		}
	}

	requestHeaders := http.Header{}
	if len(args) >= 2 {
		headersObj, ok := args[1].(ObjectValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Second argument to ws_dial must be an object, got %s", args[1]),
			}
		}

		for key, value := range headersObj {
			valStr, ok := value.(*StringValue)
			if !ok {
				return nil, &runtimeError{
					reason: fmt.Sprintf("Header value %s for key %s must be a string", value, key),
				}
			}
			requestHeaders.Set(key, valStr.stringContent())
		}
	}

	dialer := websocket.Dialer{}
	conn, resp, err := dialer.Dial(urlVal.stringContent(), requestHeaders)
	if err != nil {
		status := IntValue(0)
		headers := ObjectValue{}
		if resp != nil {
			status = IntValue(resp.StatusCode)
			headers = makeHeaderObject(resp.Header)
			_ = resp.Body.Close()
		}

		return ObjectValue{
			"type":    AtomValue("error"),
			"error":   MakeString(err.Error()),
			"status":  status,
			"headers": headers,
		}, nil
	}

	status := IntValue(101)
	respHeaders := ObjectValue{}
	if resp != nil {
		status = IntValue(resp.StatusCode)
		respHeaders = makeHeaderObject(resp.Header)
		_ = resp.Body.Close()
	}

	return ObjectValue{
		"type":    AtomValue("ok"),
		"socket":  storeWebsocket(conn),
		"status":  status,
		"headers": respHeaders,
	}, nil
}

func (c *Context) oakWSSend(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("ws_send", args, 2); err != nil {
		return nil, err
	}

	conn, _, err := c.getWebsocket(args[0], "ws_send")
	if err != nil {
		return nil, err
	}

	data, ok := args[1].(*StringValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Second argument to ws_send must be a string, got %s", args[1]),
		}
	}

	messageType := websocket.TextMessage
	if len(args) >= 3 {
		opcode, ok := args[2].(IntValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Third argument to ws_send must be an int opcode, got %s", args[2]),
			}
		}
		messageType = int(opcode)
	}

	c.Unlock()
	writeErr := conn.WriteMessage(messageType, []byte(data.stringContent()))
	c.Lock()
	if writeErr != nil {
		return errObj(fmt.Sprintf("Could not write websocket message: %s", writeErr.Error())), nil
	}

	return ObjectValue{
		"type": AtomValue("sent"),
	}, nil
}

func (c *Context) oakWSRecv(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("ws_recv", args, 1); err != nil {
		return nil, err
	}

	conn, id, err := c.getWebsocket(args[0], "ws_recv")
	if err != nil {
		return nil, err
	}

	c.Unlock()
	messageType, payload, readErr := conn.ReadMessage()
	c.Lock()
	if readErr != nil {
		// Always remove the websocket on any read error: gorilla/websocket
		// marks the connection as permanently failed after any read error,
		// and a subsequent ReadMessage will panic with "repeated read on
		// failed websocket connection".
		removeWebsocket(id)

		if closeErr, ok := readErr.(*websocket.CloseError); ok {
			return websocketClosedEvent(closeErr.Code, closeErr.Text), nil
		}

		return errObj(fmt.Sprintf("Could not read websocket message: %s", readErr.Error())), nil
	}

	return websocketEvent(messageType, payload), nil
}

func (c *Context) oakWSClose(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("ws_close", args, 1); err != nil {
		return nil, err
	}

	conn, id, err := c.getWebsocket(args[0], "ws_close")
	if err != nil {
		return nil, err
	}

	closeCode := websocket.CloseNormalClosure
	closeText := ""
	if len(args) >= 2 {
		code, ok := args[1].(IntValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Second argument to ws_close must be an int close code, got %s", args[1]),
			}
		}
		closeCode = int(code)
	}
	if len(args) >= 3 {
		reason, ok := args[2].(*StringValue)
		if !ok {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Third argument to ws_close must be a string reason, got %s", args[2]),
			}
		}
		closeText = reason.stringContent()
	}

	// Remove from map first so concurrent ws_recv callers get "not available"
	// instead of racing with conn.Close() (which panics in gorilla/websocket).
	removeWebsocket(id)

	frame := websocket.FormatCloseMessage(closeCode, closeText)
	c.Unlock()
	_ = conn.WriteControl(websocket.CloseMessage, frame, time.Now().Add(2*time.Second))
	closeErr := conn.Close()
	c.Lock()

	if closeErr != nil {
		return errObj(fmt.Sprintf("Could not close websocket: %s", closeErr.Error())), nil
	}

	return websocketClosedEvent(closeCode, closeText), nil
}

func (ctx *Context) oakWSListen(args []Value) (Value, *runtimeError) {
	if err := ctx.requireArgLen("ws_listen", args, 3); err != nil {
		return nil, err
	}

	host, okHost := args[0].(*StringValue)
	pathVal, okPath := args[1].(*StringValue)
	cb := args[2]
	okCb := ctx.isCallableValue(cb)
	if !okHost || !okPath || !okCb {
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call ws_listen(%s, %s, %s)", args[0], args[1], args[2]),
		}
	}

	sendErr := func(msg string) {
		ctx.Lock()
		defer ctx.Unlock()

		_, err2 := ctx.EvalFnValueAsync(cb, false, errObj(msg))
		if err2 != nil {
			ctx.eng.reportErr(err2)
		}
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(pathVal.stringContent(), func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			sendErr(fmt.Sprintf("Could not upgrade websocket connection: %s", err.Error()))
			return
		}

		socket := storeWebsocket(conn)
		headers := makeHeaderObject(r.Header)

		ctx.Lock()
		defer ctx.Unlock()

		_, cbErr := ctx.EvalFnValueAsync(cb, false, ObjectValue{
			"type":   AtomValue("connect"),
			"socket": socket,
			"req": ObjectValue{
				"method":  MakeString(r.Method),
				"url":     MakeString(r.URL.String()),
				"headers": headers,
			},
		})
		if cbErr != nil {
			ctx.eng.reportErr(cbErr)
		}
	})

	server := &http.Server{
		Addr:    host.stringContent(),
		Handler: mux,
	}

	ctx.eng.Add(1)
	go func() {
		defer ctx.eng.Done()
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			sendErr(fmt.Sprintf("Error starting websocket server in ws_listen(): %s", err.Error()))
		}
	}()

	closer := func(_ []Value) (Value, *runtimeError) {
		ctx.eng.Add(1)
		go func() {
			defer ctx.eng.Done()

			err := server.Shutdown(context.Background())
			if err != nil {
				sendErr(fmt.Sprintf("Could not close websocket server in ws_listen/close: %s", err.Error()))
			}
		}()

		return null, nil
	}

	return BuiltinFnValue{
		name: "close",
		fn:   closer,
	}, nil
}

func (c *Context) oakSin(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("sin", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call sin(%s)", args[0]),
		}
	}

	return FloatValue(math.Sin(val)), nil
}

func (c *Context) oakCos(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("cos", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call cos(%s)", args[0]),
		}
	}

	return FloatValue(math.Cos(val)), nil
}

func (c *Context) oakTan(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("tan", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call tan(%s)", args[0]),
		}
	}

	return FloatValue(math.Tan(val)), nil
}

func (c *Context) oakAsin(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("asin", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call asin(%s)", args[0]),
		}
	}

	if val > 1 || val < -1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("asin() takes a number in range [-1, 1], got %f", val),
		}
	}

	return FloatValue(math.Asin(val)), nil
}

func (c *Context) oakAcos(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("acos", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call acos(%s)", args[0]),
		}
	}

	if val > 1 || val < -1 {
		return nil, &runtimeError{
			reason: fmt.Sprintf("acos() takes a number in range [-1, 1], got %f", val),
		}
	}

	return FloatValue(math.Acos(val)), nil
}

func (c *Context) oakAtan(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("atan", args, 1); err != nil {
		return nil, err
	}

	var val float64
	switch arg := args[0].(type) {
	case IntValue:
		val = float64(arg)
	case FloatValue:
		val = float64(arg)
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call atan(%s)", args[0]),
		}
	}

	return FloatValue(math.Atan(val)), nil
}

func (c *Context) oakPow(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("pow", args, 2); err != nil {
		return nil, err
	}

	var base float64
	var exp float64
	err := runtimeError{
		reason: fmt.Sprintf("Mismatched types in call pow(%s, %s)", args[0], args[1]),
	}

	switch arg := args[0].(type) {
	case IntValue:
		base = float64(arg)
	case FloatValue:
		base = float64(arg)
	default:
		return nil, &err
	}

	switch arg := args[1].(type) {
	case IntValue:
		exp = float64(arg)
	case FloatValue:
		exp = float64(arg)
	default:
		return nil, &err
	}

	if base == 0 && exp == 0 {
		return nil, &runtimeError{
			reason: "pow(0, 0) is not defined",
		}
	} else if base < 0 && float64(int64(exp)) != exp {
		return nil, &runtimeError{
			reason: "pow() of negative number to fractional exponent is not defined",
		}
	}

	return FloatValue(math.Pow(base, exp)), nil
}

func (c *Context) oakLog(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("log", args, 2); err != nil {
		return nil, err
	}

	var base float64
	var exp float64
	err := runtimeError{
		reason: fmt.Sprintf("Mismatched types in call log(%s, %s)", args[0], args[1]),
	}

	switch arg := args[0].(type) {
	case IntValue:
		base = float64(arg)
	case FloatValue:
		base = float64(arg)
	default:
		return nil, &err
	}

	switch arg := args[1].(type) {
	case IntValue:
		exp = float64(arg)
	case FloatValue:
		exp = float64(arg)
	default:
		return nil, &err
	}

	if base == 0 {
		return nil, &runtimeError{
			reason: "log(0, _) is not defined",
		}
	} else if exp == 0 {
		return nil, &runtimeError{
			reason: "log(_, 0) is not defined",
		}
	}

	// we use math.Log2 here because we want logs of base 2 to give exact
	// answers, where we care less about other bases
	return FloatValue(math.Log2(exp) / math.Log2(base)), nil
}

// ___runtime_lib returns the string content of the bundled standard library by
// the given name, or ? otherwise.
func (c *Context) rtLib(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("__runtime_lib", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		libName := arg.stringContent()
		if libSource, ok := stdlibs[libName]; ok {
			return MakeString(libSource), nil
		}
		return null, nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call ___runtime_lib(%s)", args[0]),
		}
	}
}

// ___runtime_lib? reports whether a bundled standard library by the given name exists
func (c *Context) rtIsLib(args []Value) (Value, *runtimeError) {
	if err := c.requireArgLen("__runtime_lib?", args, 1); err != nil {
		return nil, err
	}

	switch arg := args[0].(type) {
	case *StringValue:
		libName := arg.stringContent()
		_, ok := stdlibs[libName]
		return BoolValue(ok), nil
	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("Mismatched types in call ___runtime_lib?(%s)", args[0]),
		}
	}
}

// ___runtime_gc runs a garbage collection cycle for both Oak and the
// underlying Go runtime. It blocks until the GC cycle is complete.
func (c *Context) rtGC(_ []Value) (Value, *runtimeError) {
	runtime.GC()
	return null, nil
}

// ___runtime_mem reports a dictionary of memory usage statistics for diagnostics
func (c *Context) rtMem(_ []Value) (Value, *runtimeError) {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	return ObjectValue{
		// number of allocations
		"allocs": IntValue(memStats.Mallocs),
		"frees":  IntValue(memStats.Frees),
		"live":   IntValue(memStats.Mallocs - memStats.Frees),
		// number of bytes
		"heap": IntValue(memStats.HeapAlloc),
		"virt": IntValue(memStats.HeapSys),
		// total gc cycles count
		"gcs": IntValue(memStats.NumGC),
	}, nil
}

// ___runtime_proc returns metadata about the current process
func (c *Context) rtProc(_ []Value) (Value, *runtimeError) {
	var exeValue Value
	execPath, err := os.Executable()
	if err == nil {
		exeValue = MakeString(execPath)
	} else {
		exeValue = null
	}

	return ObjectValue{
		"pid": IntValue(os.Getpid()),
		"exe": exeValue,
	}, nil
}

// ___runtime_go_version returns the Go version
func (c *Context) rtGoVersion(_ []Value) (Value, *runtimeError) {
	return MakeString(runtime.Version()), nil
}

// ___runtime_sys_info returns system information
func (c *Context) rtSysInfo(_ []Value) (Value, *runtimeError) {
	return ObjectValue{
		"os":   MakeString(runtime.GOOS),
		"arch": MakeString(runtime.GOARCH),
		"cpus": IntValue(runtime.NumCPU()),
	}, nil
}

// oakBytecodeCall implements the bytecode() builtin.
// Usage: with bytecode([arg1, arg2]) fn(arg1, arg2) { body }
// Compiles the given function body to bytecode and executes it via the VM.
// Compiled chunks are cached per *fnNode so repeated calls skip recompilation.
// VM structs are pooled to reduce allocation overhead.
func (c *Context) oakBytecodeCall(args []Value) (Value, *runtimeError) {
	if len(args) < 2 {
		return nil, &runtimeError{
			reason: "bytecode() requires 2 arguments: argument list and function",
		}
	}

	argList, ok := args[0].(*ListValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("first argument to bytecode() must be a list, got %s", args[0]),
		}
	}
	fnArgs := []Value(*argList)

	fn := args[1]

	switch f := fn.(type) {
	case FnValue:
		// Look up or compile the chunk for this function definition
		chunk := c.getOrCompileChunk(f.defn)

		// Build initial locals: bind named params, then rest arg
		initVals := normalizeCallArgs(len(f.defn.args), fnArgs)
		if f.defn.restArg != "" {
			var restList ListValue
			if len(fnArgs) > len(f.defn.args) {
				restList = ListValue(fnArgs[len(f.defn.args):])
			} else {
				restList = ListValue{}
			}
			initVals = append(initVals[:len(f.defn.args)], &restList)
		}

		vm := acquireVM(chunk, c)
		vm.initLocals = initVals
		// Bridge the enclosing tree-walker scope so upvalues (map, filter, etc.) resolve
		vm.outerScope = bridgeToVmScope(&f.scope)
		result, err := vm.run()
		releaseVM(vm)
		return result, err

	case *closureVal:
		// If the closureVal has a preserved AST, compile via cache and run
		// in a pooled VM for consistent performance in both execution modes.
		if f.defn != nil {
			chunk := c.getOrCompileChunk(f.defn)
			initVals := normalizeCallArgs(len(f.defn.args), fnArgs)
			if f.defn.restArg != "" {
				var restList ListValue
				if len(fnArgs) > len(f.defn.args) {
					restList = ListValue(fnArgs[len(f.defn.args):])
				} else {
					restList = ListValue{}
				}
				initVals = append(initVals[:len(f.defn.args)], &restList)
			}
			vm := acquireVM(chunk, c)
			vm.initLocals = initVals
			// Use the closureVal's captured scope for upvalue resolution
			vm.outerScope = f.parentScope
			result, err := vm.run()
			releaseVM(vm)
			return result, err
		}
		if f.call != nil {
			return f.call(fnArgs)
		}
		return c.evalFnCall(fn, false, fnArgs)

	case BuiltinFnValue:
		return f.fn(fnArgs)

	default:
		return nil, &runtimeError{
			reason: fmt.Sprintf("second argument to bytecode() must be a function, got %s", fn),
		}
	}
}

// getOrCompileChunk returns a cached bytecodeChunk for the given fnNode,
// compiling it on the first call and caching for subsequent calls.
func (c *Context) getOrCompileChunk(defn *fnNode) *bytecodeChunk {
	eng := c.eng

	eng.bytecodeCacheLock.RLock()
	chunk, ok := eng.bytecodeCache[defn]
	eng.bytecodeCacheLock.RUnlock()
	if ok {
		return chunk
	}

	// Compile
	co := newCompiler()
	// Push a sentinel parent scope so the compiler knows this is a function body,
	// not top-level code. Without this, unknown identifiers (imports, builtins)
	// would be forward-declared as locals (loading null) instead of emitting
	// LOAD_UPVAL to search the enclosing scope chain at runtime.
	co.parentLocals = append(co.parentLocals, []string{"_outer_"})
	for _, paramName := range defn.args {
		co.declareLocal(paramName)
	}
	if defn.restArg != "" {
		co.declareLocal(defn.restArg)
	}
	co.compileNodeTail(defn.body, true)
	co.emitOp(opHalt)

	topNames := make([]string, len(co.locals))
	copy(topNames, co.locals)

	chunk = &bytecodeChunk{
		code:          co.code,
		constants:     co.constants,
		functions:     co.functions,
		topLevelNames: topNames,
		sourceMap:     co.sourceMap,
	}

	eng.bytecodeCacheLock.Lock()
	eng.bytecodeCache[defn] = chunk
	eng.bytecodeCacheLock.Unlock()

	return chunk
}

// oakInterpreterCall implements the interpreter() builtin.
// Usage: with interpreter([arg1, arg2]) fn(arg1, arg2) { body }
// Evaluates the given function using the tree-walking interpreter with TCO.
// When called in --bytecode mode, closureVals with preserved AST are converted
// to FnValues so the body is genuinely tree-walked instead of dispatched back
// to the bytecode VM.
func (c *Context) oakInterpreterCall(args []Value) (Value, *runtimeError) {
	if len(args) < 2 {
		return nil, &runtimeError{
			reason: "interpreter() requires 2 arguments: argument list and function",
		}
	}

	argList, ok := args[0].(*ListValue)
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("first argument to interpreter() must be a list, got %s", args[0]),
		}
	}
	fnArgs := []Value(*argList)

	fn := args[1]

	// If fn is a bytecode closureVal with preserved AST, convert to FnValue
	// so evalFnCall tree-walks the body instead of calling cv.call() (bytecode).
	if cv, ok := fn.(*closureVal); ok && cv.defn != nil {
		fn = FnValue{
			defn:  cv.defn,
			scope: c.bridgeVmScope(cv.parentScope),
		}
	}

	// Use thunkable=true to enable tail-call optimization
	v, err := c.evalFnCall(fn, true, fnArgs)
	if err != nil {
		return nil, err
	}
	// Unwrap thunks for TCO
	for {
		thunk, isThunk := v.(thunkValue)
		if !isThunk {
			return v, nil
		}
		v, err = c.evalExprWithOpt(thunk.defn.body, thunk.scope, true)
		if err != nil {
			return nil, err
		}
	}
}
