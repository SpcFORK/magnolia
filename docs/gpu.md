# GPU Library (gpu)

## Overview

`libgpu` provides low-level helpers for writing GPU interop code in Oak, enabling direct access to CUDA and OpenCL drivers through Oak's syscall bridge.

## Import

```oak
gpu := import('gpu')
{ scan: scan, available: available, cuda: cuda, opencl: opencl } := import('gpu')
```

## Functions

### `scan()`

Scans for available GPU backends (CUDA and OpenCL) on the system.

**Returns:** List of backend info objects

```oak
{ scan: scan } := import('gpu')

backends := scan()
each(backends, fn(backend) {
    println('Backend: ' + backend.name)
    println('Available: ' + string(backend.available))
    if backend.available -> {
        println('Library: ' + backend.library)
        println('Probe: ' + backend.probe)
    }
})

// Example output:
// {
//   name: 'cuda'
//   available: true
//   library: 'libcuda.so.1'
//   probe: 'cuInit'
//   address: 0x7f...
// }
```

### `available()`

Returns list of only the available GPU backends.

```oak
{ available: available } := import('gpu')

gpus := available()
if len(gpus) = 0 -> {
    println('No GPU backends available')
} else {
    println('Found ' + string(len(gpus)) + ' GPU backend(s)')
    each(gpus, fn(gpu) {
        println('- ' + gpu.name + ' (' + gpu.library + ')')
    })
}
```

### `cuda(symbol)`

Resolves a CUDA driver API symbol from the CUDA library.

**Returns:** Procedure object or error

```oak
{ cuda: cuda, call: call } := import('gpu')

cuInit := cuda('cuInit')
if cuInit.type {
    :proc -> {
        // Initialize CUDA
        result := call(cuInit, 0)
        println('CUDA initialized: ' + string(result))
    }
    :error -> println('Failed to load CUDA: ' + cuInit.error)
}
```

### `opencl(symbol)`

Resolves an OpenCL API symbol from the OpenCL library.

```oak
{ opencl: opencl, call: call } := import('gpu')

clGetPlatformIDs := opencl('clGetPlatformIDs')
if clGetPlatformIDs.type = :proc -> {
    // Use OpenCL...
}
```

### `resolve(library, symbol)`

Resolves a procedure from a specific library.

```oak
{ resolve: resolve } := import('gpu')

proc := resolve('libcuda.so.1', 'cuDeviceGetCount')
if proc.type = :proc -> {
    // Use procedure...
}
```

### `call(procOrAddress, args...)`

Calls a resolved GPU procedure with arguments.

```oak
{ cuda: cuda, call: call } := import('gpu')

cuDeviceGetCount := cuda('cuDeviceGetCount')
deviceCount := 0

// Call with pointer to deviceCount
result := call(cuDeviceGetCount, addr(deviceCount))
println('GPU Devices: ' + string(deviceCount))
```

## Platform-Specific Libraries

### CUDA Libraries (by OS)

```oak
// Windows
['nvcuda.dll', 'cudart64_12.dll', 'cudart64_11.dll']

// Linux
['libcuda.so.1', 'libcudart.so']

// macOS
[] // CUDA not available on macOS
```

### OpenCL Libraries (by OS)

```oak
// Windows
['OpenCL.dll']

// Linux
['libOpenCL.so.1', 'libOpenCL.so']

// macOS
['/System/Library/Frameworks/OpenCL.framework/OpenCL']
```

## Examples

### Check GPU Availability

```oak
{ available: available } := import('gpu')

gpus := available()

if len(gpus) = 0 -> {
    println('No GPU detected. Running on CPU.')
    useCPU()
} else {
    println('GPU detected: ' + gpus.0.name)
    useGPU(gpus.0)
}
```

### Initialize CUDA

```oak
{ cuda: cuda, call: call } := import('gpu')

fn initCUDA {
    cuInit := cuda('cuInit')
    if cuInit.type != :proc -> {
        println('CUDA not available')
        ? // Return null
    } else {
        result := call(cuInit, 0)
        if result = 0 {
            true -> println('CUDA initialized successfully')
            _ - println('CUDA initialization failed: ' + string(result))
        }
        result
    }
}

if initCUDA() = 0 -> {
    // Proceed with CUDA operations...
}
```

### Get Device Count

```oak
{ cuda: cuda, call: call } := import('gpu')

fn getCUDADeviceCount {
    cuDeviceGetCount := cuda('cuDeviceGetCount')
    if cuDeviceGetCount.type != :proc -> 0
    
    count := 0
    result := call(cuDeviceGetCount, addr(count))
    
    if result = 0 {
        true -> count
        _ -> 0
    }
}

devices := getCUDADeviceCount()
println('CUDA Devices: ' + string(devices))
```

### OpenCL Platform Detection

```oak
{ opencl: opencl, call: call } := import('gpu')

fn getOCLPlatformCount {
    clGetPlatformIDs := opencl('clGetPlatformIDs')
    if clGetPlatformIDs.type != :proc -> 0
    
    count := 0
    // Call with NULL platforms to get count
    call(clGetPlatformIDs, 0, 0, addr(count))
    count
}

platforms := getOCLPlatformCount()
println('OpenCL Platforms: ' + string(platforms))
```

## Backends

### CUDA Backend

**Probe function:** `cuInit`

Used for NVIDIA GPU computing via CUDA driver API.

### OpenCL Backend

**Probe function:** `clGetPlatformIDs`

Used for cross-platform GPU/accelerator computing.

## Workflow

1. **Scan** for available backends
2. **Resolve** needed API functions
3. **Call** GPU procedures with arguments
4. **Handle** return codes and errors

## Implementation Notes

- Uses Oak's `sysproc()` to load library symbols
- Uses Oak's `syscall()` to invoke resolved procedures
- Automatically tries multiple library paths per platform
- Returns error if procedure cannot be resolved
- Does not manage GPU memory automatically
- Does not provide high-level GPU abstractions

## Limitations

- Low-level API requiring manual FFI work
- No automatic memory management
- No type safety for procedure calls
- Platform-specific library paths must be known
- No GPU memory allocation helpers
- No kernel compilation support
- Must understand underlying GPU API (CUDA/OpenCL)
- No error code interpretation

## Use Cases

- Direct CUDA/OpenCL interop
- Custom GPU kernel execution
- Performance-critical GPU computing
- Integrating with existing GPU codebases
- Low-level hardware access

## Safety Considerations

⚠️ **This is a low-level API**:
- Invalid syscalls can crash the program
- No bounds checking on memory access
- Must match C calling conventions exactly
- Incorrect arguments can corrupt memory
- Requires understanding of GPU driver APIs

## See Also

- Oak built-in `sysproc()` - Load library procedure
- Oak built-in `syscall()` - Call foreign procedure
- Oak built-in `___runtime_sys_info()` - System information
- CUDA Driver API documentation
- OpenCL API documentation
