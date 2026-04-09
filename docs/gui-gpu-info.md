# gui-gpu-info — GPU & Display Detection

`import('gui-gpu-info')` detects GPU adapters, VRAM, driver versions, and display capabilities using WMI and DXGI enumeration via PowerShell.

## Quick Start

```oak
gpu := import('gui-gpu-info')

// Quick structured GPU list
adapters := gpu.getGPUAdaptersParsed()
each(adapters, fn(a) {
    println(a.name + ' — ' + string(a.vramMB) + ' MB')
})

// Direct3D feature level
level := gpu.getD3DFeatureLevel()
println('D3D Feature Level: ' + level)

// Full dump
report := gpu.gpuCapabilityDump()
```

## API Reference

### `getGPUAdapters()`

Lists all GPU adapters with raw WMI output.

### `getGPUAdaptersParsed()`

Returns a structured list of GPU adapters, each with: `name`, `driver`, `vramMB`, `resolutionX`, `resolutionY`, `refreshRate`.

### `getDXGIAdapters()`

Enumerates DXGI adapters for detailed GPU information.

### `getD3DFeatureLevel()`

Detects the highest Direct3D feature level supported by the system.

### `getDisplayModes()`

Lists available display modes for the primary monitor.

### `getMonitorInfo()`

Gets details about connected monitors (manufacturer, serial, year, week).

### `gpuCapabilityDump()`

Generates a comprehensive GPU and display capability report combining all queries.

### `gpuCapabilityDumpParallel()`

Same as `gpuCapabilityDump()` but runs all queries concurrently for faster results.

## Notes

- All queries use PowerShell internally. Performance depends on WMI responsiveness.
- DXGI enumeration provides the most accurate VRAM data on modern systems.
