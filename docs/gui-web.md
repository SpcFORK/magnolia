# gui-web

Canvas/WebGL middleware helpers for web and JS runtimes. Implements a command queue model that records drawing and WebGL calls into `window.messages`.

Key exports

- `createCanvas(window, id?, options?)` — configure canvas metadata
- `initWebGL(window, contextName?, attrs?)` — initialize WebGL middleware
- `webglCreateShader`, `webglCreateProgram`, `webglUseProgram`, `webglDrawArrays`, `webglFlush` — command-recording shims

Notes

- The web backend does not call native host APIs directly; it records operations for a host bridge to execute.
