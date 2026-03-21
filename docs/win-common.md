# win-common

Shared Win32 low-level helpers used across multiple `gui-native-win-*`
submodules.

Centralises pointer-reading, COM v-table dispatch, and HRESULT helpers so
each Windows-specific module imports them from one place.

## Functions

### `_ptrRead(address)`

Reads a native-width pointer (4 or 8 bytes) at `address` depending on the
current `windows.ptrSize()`.

### `_comCall(comObj, methodIndex, args...)`

Calls a COM method through the object's v-table.  `methodIndex` is the
zero-based slot in the virtual function table.

### `_readOutPointer(outBuf)`

Reads a native-width pointer from a buffer previously passed as an
out-parameter to a Win32 / COM call.

### `_releaseOk?(res)`

Returns `true` when `res` indicates a successful Win32 call **or** the
`ERROR_GEN_FAILURE` (errno 31) that some COM Release calls return on
final release.
