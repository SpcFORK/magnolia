# Shader Code Generation (gui-shader-codegen)

## Overview

`gui-shader-codegen` generates GPU shader source code (GLSL and HLSL), handles
WebGL shader submission, and provides offline compilation wrappers for
`glslc`, `fxc`, and `dxc`. It also includes high-level assembly functions that
compose a complete shader from structured options.

## Import

```oak
cg := import('gui-shader-codegen')
```

All symbols are also re-exported by `gui-shader` and through the `GUI` facade.

## GLSL Helpers

| Function | Description |
|----------|-------------|
| `glslVersion(ver?)` | `#version` directive. Defaults to `300 es`. |
| `glslPrecision(prec?, type?)` | Precision declaration. Defaults to `mediump float`. |
| `glslStdUniforms()` | Standard uniforms block: `iTime`, `iResolution`, `iMouse`, `iFrame`. |
| `glslUniform(type, name)` | Single `uniform` declaration. |
| `glslUniforms(uniforms)` | Multiple uniforms from a list of `{ type, name }` dicts. |
| `glslIn(type, name)` | `in` declaration. |
| `glslOut(type, name)` | `out` declaration. |
| `glslQuadVertex()` | Full-screen triangle vertex shader (GLSL 300 es). |
| `glslQuadVertexCompat()` | Full-screen quad vertex shader (GLSL 100/attribute-based). |
| `glslFragmentWrap(body, version?)` | Wrap a `mainImage` body into a complete fragment shader. |
| `glslMathLib()` | Inline GLSL math library: hash, noise2D, fbm, SDF helpers, pingpong. |

## HLSL Helpers

| Function | Description |
|----------|-------------|
| `hlslStdCBuffer()` | Standard cbuffer with `iTime`, `iResolution`, `iMouse`, `iFrame`. |
| `hlslCBuffer(name, uniforms)` | Custom cbuffer from a list of `{ type, name }` dicts. |
| `hlslQuadVertex()` | Full-screen triangle vertex shader for D3D. |
| `hlslFragmentWrap(body)` | Wrap a pixel-shader body into a complete HLSL fragment. |
| `hlslMathLib()` | Inline HLSL math library (same functions as GLSL, using HLSL intrinsics). |

## WebGL Submission

| Function | Signature | Description |
|----------|-----------|-------------|
| `submitWebGL` | `(window, fragSource, vertSource?)` | Compile fragment + optional vertex shader and create a WebGL program. Returns `{ vertex, fragment, program }`. |
| `drawWebGL` | `(window, clearR?, clearG?, clearB?)` | Clear the viewport and draw a full-screen triangle. |
| `renderWebGL` | `(window, fragSource)` | Combined submit → draw → flush in one call. |

## Offline Compilation

| Function | Signature | Description |
|----------|-----------|-------------|
| `compileGLSL` | `(source, stage?, outputPath?)` | Invoke `glslc` to compile to SPIR-V. Defaults to fragment stage, `out.spv`. |
| `compileHLSL` | `(source, profile?, entry?, outputPath?)` | Invoke `fxc` to compile. Defaults to `ps_5_0`, entry `PS`, `out.fxo`. |
| `compileDXC` | `(source, profile?, entry?, outputPath?, spirv?)` | Invoke `dxc`. Defaults to `ps_6_0`. Pass `spirv? = true` for SPIR-V output. |

## Shader Source Assembly

High-level functions that assemble a complete shader from structured options.

### `assembleGLSL(opts)`

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `version` | string | `'300 es'` | GLSL version |
| `precision` | string | `'mediump'` | Precision qualifier |
| `stdUniforms` | bool | true | Include iTime/iResolution/iMouse/iFrame |
| `uniforms` | list | — | Extra `{ type, name }` uniform declarations |
| `mathLib` | bool | false | Include inline math library |
| `body` | string | — | GLSL source for the `mainImage` function body |

### `assembleHLSL(opts)`

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `stdUniforms` | bool | true | Include standard cbuffer fields |
| `uniforms` | list | — | Extra `{ type, name }` uniform declarations |
| `cbufferName` | string | `'ShaderUniforms'` | Name of the cbuffer |
| `mathLib` | bool | false | Include inline math library |
| `body` | string | — | HLSL source for the `mainImage` function body |
