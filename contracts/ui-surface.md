# UI Surface contract (source of truth pointer)

The normative Surface vocabulary lives in **`lib/ui-surface.oak`** (ops, props, enums, DrawList opcodes, fingerprint).

This directory documents the platform rule:

- **Append-only** codes — never renumber or reuse.
- **Drift guard** — `surface.fingerprint()` + `test/ui-surface.test.oak`.
- **Build bake** — `style-bake` / `transpile-ui` produce assets under VFS keys `ui/styles.bin` and `ui/font-atlas.json`.

See also: [docs/surface.md](../docs/surface.md).
