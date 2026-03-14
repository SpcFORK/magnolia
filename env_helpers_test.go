package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestRequireArgLen(t *testing.T) {
	ctx := NewContext(".")

	if err := ctx.requireArgLen("f", []Value{IntValue(1)}, 1); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err := ctx.requireArgLen("f", []Value{}, 2)
	if err == nil {
		t.Fatalf("expected runtime error for missing arguments")
	}
	if !strings.Contains(err.reason, "f requires 2 arguments, got 0") {
		t.Fatalf("unexpected runtime error reason: %q", err.reason)
	}
}

func TestBuiltinFnValueStringAndEq(t *testing.T) {
	v := BuiltinFnValue{name: "demo"}
	if got := v.String(); got != "fn demo { <native fn> }" {
		t.Fatalf("unexpected string form: %q", got)
	}

	if !v.Eq(empty) {
		t.Fatalf("builtin function should Eq empty")
	}
	if !v.Eq(BuiltinFnValue{name: "demo"}) {
		t.Fatalf("builtin functions with same name should compare equal")
	}
	if v.Eq(BuiltinFnValue{name: "other"}) {
		t.Fatalf("builtin functions with different names should not compare equal")
	}
	if v.Eq(IntValue(1)) {
		t.Fatalf("builtin function should not equal unrelated types")
	}
}

func TestWebsocketObjectHelpers(t *testing.T) {
	obj := websocketObj(12)
	if typ, ok := obj["type"].(AtomValue); !ok || typ != AtomValue("websocket") {
		t.Fatalf("unexpected websocket object type field: %+v", obj)
	}
	if id, ok := obj["id"].(IntValue); !ok || id != IntValue(12) {
		t.Fatalf("unexpected websocket object id field: %+v", obj)
	}

	event := websocketEvent(2, []byte("hi"))
	if event["type"] != AtomValue("message") {
		t.Fatalf("unexpected websocket event type: %+v", event)
	}
	if event["opcode"] != IntValue(2) {
		t.Fatalf("unexpected websocket opcode: %+v", event)
	}
	if data, ok := event["data"].(*StringValue); !ok || data.stringContent() != "hi" {
		t.Fatalf("unexpected websocket event data: %+v", event)
	}

	closed := websocketClosedEvent(1000, "bye")
	if closed["type"] != AtomValue("closed") {
		t.Fatalf("unexpected websocket closed type: %+v", closed)
	}
	if closed["code"] != IntValue(1000) {
		t.Fatalf("unexpected websocket close code: %+v", closed)
	}
	if reason, ok := closed["reason"].(*StringValue); !ok || reason.stringContent() != "bye" {
		t.Fatalf("unexpected websocket close reason: %+v", closed)
	}
}

func TestMakeHeaderObject(t *testing.T) {
	h := http.Header{}
	h.Add("X-Test", "a")
	h.Add("X-Test", "b")
	h.Set("Accept", "application/json")

	out := makeHeaderObject(h)
	if s, ok := out["X-Test"].(*StringValue); !ok || s.stringContent() != "a,b" {
		t.Fatalf("unexpected joined multi header value: %+v", out)
	}
	if s, ok := out["Accept"].(*StringValue); !ok || s.stringContent() != "application/json" {
		t.Fatalf("unexpected single header value: %+v", out)
	}
}

func TestGetStoreRemoveWebsocket(t *testing.T) {
	ctx := NewContext(".")

	wsConnMu.Lock()
	oldMap := wsConnMap
	oldNext := wsNextConnID
	wsConnMap = map[int64]*websocket.Conn{}
	wsNextConnID = 0
	wsConnMu.Unlock()

	t.Cleanup(func() {
		wsConnMu.Lock()
		wsConnMap = oldMap
		wsNextConnID = oldNext
		wsConnMu.Unlock()
	})

	if _, _, err := ctx.getWebsocket(IntValue(1), "ws_recv"); err == nil || !strings.Contains(err.reason, "must be a websocket") {
		t.Fatalf("expected type-check error, got %v", err)
	}
	if _, _, err := ctx.getWebsocket(ObjectValue{"type": MakeString("websocket"), "id": IntValue(1)}, "ws_recv"); err == nil || !strings.Contains(err.reason, "must be a websocket") {
		t.Fatalf("expected non-atom type field check to fail, got %v", err)
	}
	if _, _, err := ctx.getWebsocket(ObjectValue{"type": AtomValue("not-websocket"), "id": IntValue(1)}, "ws_recv"); err == nil || !strings.Contains(err.reason, "must be a websocket") {
		t.Fatalf("expected wrong websocket type atom check to fail, got %v", err)
	}

	if _, _, err := ctx.getWebsocket(ObjectValue{"type": AtomValue("websocket")}, "ws_recv"); err == nil || !strings.Contains(err.reason, "malformed") {
		t.Fatalf("expected malformed websocket error, got %v", err)
	}

	if _, _, err := ctx.getWebsocket(ObjectValue{"type": AtomValue("websocket"), "id": IntValue(99)}, "ws_recv"); err == nil || !strings.Contains(err.reason, "not available") {
		t.Fatalf("expected unavailable websocket error, got %v", err)
	}

	obj := storeWebsocket(nil)
	conn, id, err := ctx.getWebsocket(obj, "ws_recv")
	if err != nil {
		t.Fatalf("expected stored websocket to be retrievable, got %v", err)
	}
	if conn != nil {
		t.Fatalf("expected nil connection placeholder, got %#v", conn)
	}
	if id != 0 {
		t.Fatalf("unexpected websocket id: %d", id)
	}

	removeWebsocket(id)
	if _, _, err := ctx.getWebsocket(obj, "ws_recv"); err == nil || !strings.Contains(err.reason, "not available") {
		t.Fatalf("expected websocket to be unavailable after removal, got %v", err)
	}
}

func TestMakeIntListUpTo(t *testing.T) {
	v := makeIntListUpTo(4)
	list, ok := v.(*ListValue)
	if !ok {
		t.Fatalf("expected ListValue, got %T", v)
	}
	if len(*list) != 4 {
		t.Fatalf("unexpected list length: %d", len(*list))
	}
	for i, el := range *list {
		iv, ok := el.(IntValue)
		if !ok || iv != IntValue(i) {
			t.Fatalf("unexpected list entry %d: %v", i, el)
		}
	}
}

func TestOakKeysBranches(t *testing.T) {
	ctx := NewContext(".")

	if _, err := ctx.oakKeys([]Value{}); err == nil || !strings.Contains(err.reason, "requires 1 arguments") {
		t.Fatalf("expected arity error, got %v", err)
	}

	str := MakeString("abc")
	v, err := ctx.oakKeys([]Value{str})
	if err != nil {
		t.Fatalf("unexpected error for string keys: %v", err)
	}
	list, ok := v.(*ListValue)
	if !ok || len(*list) != 3 {
		t.Fatalf("unexpected string keys list: %v", v)
	}

	inputList := MakeList(IntValue(9), IntValue(8))
	v, err = ctx.oakKeys([]Value{inputList})
	if err != nil {
		t.Fatalf("unexpected error for list keys: %v", err)
	}
	list, ok = v.(*ListValue)
	if !ok || len(*list) != 2 {
		t.Fatalf("unexpected list keys list: %v", v)
	}

	v, err = ctx.oakKeys([]Value{ObjectValue{"a": IntValue(1), "b": IntValue(2)}})
	if err != nil {
		t.Fatalf("unexpected error for object keys: %v", err)
	}
	list, ok = v.(*ListValue)
	if !ok || len(*list) != 2 {
		t.Fatalf("unexpected object keys list: %v", v)
	}

	v, err = ctx.oakKeys([]Value{IntValue(10)})
	if err != nil {
		t.Fatalf("unexpected error for default keys branch: %v", err)
	}
	list, ok = v.(*ListValue)
	if !ok || len(*list) != 0 {
		t.Fatalf("unexpected default keys list: %v", v)
	}
}

func TestOakConversionAndTypeHelpers(t *testing.T) {
	ctx := NewContext(".")

	if v, err := ctx.oakInt([]Value{FloatValue(4.9)}); err != nil || v != IntValue(4) {
		t.Fatalf("unexpected float->int conversion: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakInt([]Value{MakeString("17")}); err != nil || v != IntValue(17) {
		t.Fatalf("unexpected string->int conversion: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakInt([]Value{MakeString("bad")}); err != nil || v != null {
		t.Fatalf("unexpected invalid string->int behavior: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakInt([]Value{PointerValue(123)}); err != nil || v != IntValue(123) {
		t.Fatalf("unexpected pointer->int conversion: v=%v err=%v", v, err)
	}

	if v, err := ctx.oakFloat([]Value{IntValue(8)}); err != nil || v != FloatValue(8) {
		t.Fatalf("unexpected int->float conversion: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakFloat([]Value{MakeString("3.25")}); err != nil || v != FloatValue(3.25) {
		t.Fatalf("unexpected string->float conversion: v=%v err=%v", v, err)
	}

	if v, err := ctx.oakCodepoint([]Value{MakeString("A")}); err != nil || v != IntValue(65) {
		t.Fatalf("unexpected codepoint conversion: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakCodepoint([]Value{MakeString("AB")}); err != nil || v != null {
		t.Fatalf("expected null for multi-byte codepoint input: v=%v err=%v", v, err)
	}

	if v, err := ctx.oakChar([]Value{IntValue(-10)}); err != nil {
		t.Fatalf("unexpected char conversion error: %v", err)
	} else if s, ok := v.(*StringValue); !ok || len(*s) != 1 || (*s)[0] != 0 {
		t.Fatalf("unexpected negative-char conversion value: %v", v)
	}
	if v, err := ctx.oakChar([]Value{IntValue(300)}); err != nil {
		t.Fatalf("unexpected high char conversion error: %v", err)
	} else if s, ok := v.(*StringValue); !ok || len(*s) != 1 || (*s)[0] != 255 {
		t.Fatalf("unexpected clamped high-char conversion value: %v", v)
	}

	if v, err := ctx.oakLen([]Value{MakeString("abc")}); err != nil || v != IntValue(3) {
		t.Fatalf("unexpected string len result: v=%v err=%v", v, err)
	}
	if _, err := ctx.oakLen([]Value{IntValue(1)}); err == nil || !strings.Contains(err.reason, "does not support a len() call") {
		t.Fatalf("expected len type error, got %v", err)
	}

	if v, err := ctx.oakType([]Value{PointerValue(1)}); err != nil || v != AtomValue("pointer") {
		t.Fatalf("unexpected oakType(pointer) result: v=%v err=%v", v, err)
	}
	if v, err := ctx.oakType([]Value{BuiltinFnValue{name: "x"}}); err != nil || v != AtomValue("function") {
		t.Fatalf("unexpected oakType(function) result: v=%v err=%v", v, err)
	}

	if v, err := ctx.oakName([]Value{BuiltinFnValue{name: "demo"}}); err != nil || v != AtomValue("demo") {
		t.Fatalf("unexpected oakName(builtin) result: v=%v err=%v", v, err)
	}
	alphaNode := classNode{name: "Alpha"}
	if v, err := ctx.oakName([]Value{ClassValue{defn: &alphaNode}}); err != nil || v != AtomValue("Alpha") {
		t.Fatalf("unexpected oakName(class) result: v=%v err=%v", v, err)
	}

	ptrVal, err := ctx.oakPointer([]Value{AtomValue("Alpha")})
	if err != nil {
		t.Fatalf("unexpected pointer(:atom) error: %v", err)
	}
	ptr, ok := ptrVal.(PointerValue)
	if !ok || ptr == 0 {
		t.Fatalf("expected non-zero pointer from pointer(:atom), got %v", ptrVal)
	}
	if v, err := ctx.oakName([]Value{ptr}); err != nil || v != AtomValue("Alpha") {
		t.Fatalf("unexpected oakName(pointer(:atom)) result: v=%v err=%v", v, err)
	}
}

func TestOakClassMatch(t *testing.T) {
	ctx := NewContext(".")

	alphaNode := classNode{name: "Alpha"}
	betaNode := classNode{name: "Beta"}
	alpha := ClassValue{defn: &alphaNode}
	beta := ClassValue{defn: &betaNode}

	if v, err := ctx.oakClassMatch([]Value{alpha, alpha}); err != nil || v != BoolValue(true) {
		t.Fatalf("expected csof(class, same class) == true, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakClassMatch([]Value{alpha, beta}); err != nil || v != BoolValue(false) {
		t.Fatalf("expected csof(class, other class) == false, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakClassMatch([]Value{alpha, AtomValue("Alpha")}); err != nil || v != BoolValue(true) {
		t.Fatalf("expected csof(class, :name) == true, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakClassMatch([]Value{alpha, AtomValue("Beta")}); err != nil || v != BoolValue(false) {
		t.Fatalf("expected csof(class, different :name) == false, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakClassMatch([]Value{AtomValue("Alpha"), alpha}); err != nil || v != BoolValue(true) {
		t.Fatalf("expected csof(:name, class) == true, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakClassMatch([]Value{IntValue(1), AtomValue("Alpha")}); err != nil || v != BoolValue(false) {
		t.Fatalf("expected csof(non-class/atom pair) == false, got v=%v err=%v", v, err)
	}
}

func TestMemwriteAtomReferenceHelpers(t *testing.T) {
	ctx := NewContext(".")

	if v, err := ctx.oakMemWrite([]Value{AtomValue("slot"), AtomValue("hello")}); err != nil || v != IntValue(5) {
		t.Fatalf("expected memwrite(:slot, :hello) to write 5 bytes, got v=%v err=%v", v, err)
	}

	ptrVal, err := ctx.oakPointer([]Value{AtomValue("slot")})
	if err != nil {
		t.Fatalf("unexpected pointer(:slot) error: %v", err)
	}
	ptr, ok := ptrVal.(PointerValue)
	if !ok || ptr == 0 {
		t.Fatalf("expected non-zero pointer for :slot, got %v", ptrVal)
	}

	if v, err := ctx.oakName([]Value{ptr}); err != nil || v != AtomValue("hello") {
		t.Fatalf("expected name(pointer(:slot)) == :hello, got v=%v err=%v", v, err)
	}

	if v, err := ctx.oakMemWrite([]Value{AtomValue("slot"), MakeList(IntValue(65), IntValue(66), IntValue(67))}); err != nil || v != IntValue(3) {
		t.Fatalf("expected memwrite(:slot, [65,66,67]) to write 3 bytes, got v=%v err=%v", v, err)
	}
	if v, err := ctx.oakName([]Value{mustPointer(t, &ctx, AtomValue("slot"))}); err != nil || v != AtomValue("ABC") {
		t.Fatalf("expected name(pointer(:slot)) == :ABC after list write, got v=%v err=%v", v, err)
	}
}

func mustPointer(t *testing.T, ctx *Context, atom AtomValue) PointerValue {
	t.Helper()
	v, err := ctx.oakPointer([]Value{atom})
	if err != nil {
		t.Fatalf("unexpected pointer(%s) error: %v", atom, err)
	}
	ptr, ok := v.(PointerValue)
	if !ok {
		t.Fatalf("expected pointer result, got %T", v)
	}
	return ptr
}
