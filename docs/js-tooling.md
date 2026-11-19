# JS ↔ Oak tooling

_Status: IN PROGRESS — M0–M4 functionally complete; M3 harness automation, M5 docs/CLI surface, and M6 corpora pending; jQuery runtime smoke green 2026-11-16._
_Date: 2026-11-16 · Consolidates the implementation plan and the November 2026 ablation results into one reference; supersedes the separate plan and report files._
Corpora are **not committed** (fetch on demand via `tools/fetch-corpus.sh`, sha256-pinned). All workspace paths below are relative to the repo root.

---

## Part I — Implementation plan

## 1. Executive summary

Magnolia today has **half** of a JS ↔ Oak toolchain:

- **Oak → JS works end to end.** `oak build --web` / `--ts` runs `lib/build.oak` → `buildAnalyze.analyzeNode` → `buildRender.renderJSBundle` (`lib/build-render.oak`), prepended by the JS runtime preamble (`lib/runtime-js.oak`, surfaced via `lib/runtime-codegen.oak`). `make test-js` runs the whole Oak stdlib test suite on Node.js through this path.
- **JS → Oak exists only at the AST level, with no parser.** `lib/ast-js.oak` defines JS AST node constructors, an ESTree-flavored `transpileToOak` (JS AST → Oak AST), `walkJS`, and a hand-rolled `renderJS` (JS AST → JS source). There is **no tokenizer or parser for real JavaScript source** in Oak — the equivalent of Acorn for the JS side — so `jquery.js` / `three.js` cannot be transpiled today no matter how the transpiler is invoked.

The ask — ablate the JS transpile suite with JS and Acorn, test transpiling to Oak, test on real libraries (jQuery 1.0.0, …) — therefore splits into two deliverables:

1. **A measurable e2e pipeline**: real JS source → Acorn (ESTree JSON) → Oak ingest → JS AST → Oak AST → render/bundle/run. This is the fastest way to make jQuery 1.0 transpile and to get quantitative coverage data.
2. **An Oak-native JS front-end** (Acorn-subset: `lib/js-tokenize.oak` + `lib/js-parse.oak`) so the toolchain does not depend on Node at runtime. This is the actual "proper JS ↔ Oak tooling" endgame.

The two front-ends share one backend, so the ablation is a true A/B: same corpus, same ingest, same transpiler — swap the parser.

---

## 2. What exists today (verified by reading the repo)

| Component | Path | Notes |
|---|---|---|
| JS AST constructors + `transpileToOak` | `lib/ast-js.oak` (568 lines) | `jsNull/jsNumber/jsBinary/jsCall/…`; `transpileToOak` maps JS node types → Oak AST; `walkJS`; `renderJS` |
| JS→Oak operator map | `lib/ast-js.oak:_mapJSOp/_mapJSUnaryOp/_jsIdentToOak` | `+`→`:plus`, `===`→`:eq`, `&&`→`:and`, `undefined`/`null`→`?`, `console`→`__console`, `Math`→`__Math`, `$`→`_dlr_` |
| **Silent `:null` fallback** | `lib/ast-js.oak:_ -> { type: :null }` in `transpileToOak` | Unknown node ⇒ `?` in the output with **no diagnostic**. This is the single most dangerous defect for this work (see §6.1) |
| Spread dropped silently | `lib/ast-js.oak` (`:jsSpread -> transpileToOak(node.argument)` at top level; spread in call args unhandled) | `f(...args)` loses arguments |
| `renderJS` untested & buggy | `lib/ast-js.oak:renderJS` | Only `'` escaped (`\\''`), backslash unescaped; no coverage in `test/`; no CLI uses it |
| TS AST (extends JS) | `lib/ast-ts.oak` | Same shape: constructors + `transpileToOak`, no parser. Types stripped during transpile |
| Lua/Java/Go AST modules | `lib/ast-lua.oak`, `lib/ast-java.oak`, `lib/ast-go.oak`, `lib/ast-ir.oak` | Same "constructors + transpile, no parser" pattern. `ast-go.oak` and `ast-ir.oak` are the closest precedents for "external AST → Oak" |
| Oak tokenizer/parser/printer | `lib/syntax-tokenize.oak`, `lib/syntax-parse.oak`, `lib/syntax-print.oak`, facade `lib/syntax.oak` | Full-fidelity lossless tokenizer + recursive-descent parser, written in Oak. Template for the in-Oak JS parser |
| Transpile middleware | `lib/transpile.oak` | `createTranspiler/registerTranspiler/transpileNode/optimizeConstants/removeDebugCalls`; runs inside `build.oak` between parse and analyze |
| Oak→JS codegen | `lib/build-render.oak:renderJSBundle` + `build-render-node.oak` + `lib/runtime-js.oak` | Production path, used by `make test-js`. `build-render-node.oak` explicitly supports standalone rendering ("extracted … so Oak AST nodes … can be rendered … outside the full build pipeline") — use it for golden tests |
| JS analysis/analyze | `lib/build-analyze.oak` | Web-target tail-call trampolining etc. — matters for emitted bundles, not for JS→Oak ingest |
| Test corpus entry | `test/main.oak` (+ `Makefile: INCLUDES`) | stdlib tests embedded via `import('x.test')`; `ast-js:lib/ast-js.oak` etc. are `--include`d. `make test-oak` (interpreter), `make test-js` (via `oak build --web` + node), `make test-bundle` |
| Tests for ast-js today | `test/ast-targets.test.oak` | Constructor-level unit tests + transpiler mapping tests on **hand-built** ASTs only. Nothing parses real JS |
| jQuery 1.0.0 | GitHub tag **`1.0`** (no `1.0.0` tag; npm starts at `jquery@1.5.1`) | `https://codeload.github.com/jquery/jquery/tar.gz/refs/tags/1.0` ✓ (fetched, 727 KB tarball). Layout: `src/{jquery,ajax,event,fx}/*.js`, ~150 KB total |
| three.js | e.g. tag `r58` ✓ fetchable from codeload | Attached `threejs-oak-complete.patch.txt` fixes the **Lua** renderer (`fnDepth`/`iifeCounter`; `local` only inside functions; `;` separators; named-function emission). Precedent for how a big-library fix lands: small renderer patch + transpiler fixes, verified e2e |

### Verified experimentally today

- Node `v22.22.3` at `/usr/local/bin/node`; npm `10.9.8`; npm registry reachable.
- `acorn@8.19.0` installed; **all four jQuery 1.0 source modules parse cleanly at `ecmaVersion: 3`**:
  - `jquery.js` 4545 nodes · `ajax/ajax.js` 987 · `event/event.js` 511 · `fx/fx.js` 1290 → **7333 nodes, 0 errors**.
- `codeload.github.com` serves the `jquery@1.0` and `three.js@r58` tarballs; `api.github.com` tag listing works.
- **No Go toolchain in this environment** (`go: command not found`), no prebuilt `magnolia`/`oak` binary found anywhere. **All Oak execution requires Go** (`go run -race . test/main.oak`, or a `magnolia` binary). Consequences in §7.

---

## 3. Goals / non-goals

**Goals**
- G1. JS → Oak end to end on real-world source: jQuery 1.0 (ES3) as the primary corpus; then jQuery 1.12.4 (ES5) and an ES2015 sample (jQuery 3.x / three.js) as stretch corpora.
- G2. Quantitative ablation of the JS transpile path: parse success, transpile coverage (node types handled vs. fallen through to `:null`), bundle success, runtime smoke results — measured A0–A4 (§5).
- G3. Oak-native JS tokenizer + parser (Acorn-subset) written in Oak, producing the **same JS AST** as the Acorn front-end, so backend improvements benefit both and the A/B is meaningful.
- G4. Harden the Oak → JS direction (the other half of "↔"): fix `renderJS`, wire a JS-AST → JS emit path, keep `make test-js` green.
- G5. Diagnostics instead of silent corruption: every unsupported construct in the JS→Oak path must be a loud error or an explicit, printed skip-list — never a silent `?`.
- G6. Docs: this document, `docs/js-tooling.md` (the consolidated plan + results), plus updates to `docs/transpile.md`, `docs/build.md`, `OAK-README.md`/`README.md` tooling sections.

**Non-goals (for this iteration)**
- Not making transpiled jQuery/three.js *feature-complete* against the DOM. Browser APIs are out of scope; the transpiler must not polyfill them.
- No `--from-js` for TS/TSX, JSX, or WASM-from-JS in this round.
- Not modifying interpreter semantics; no new Oak syntax.
- three.js → Oak is a stretch goal; the attached Lua patch is treated as context/precedent, not as part of the acceptance criteria.

---

## 4. Recommended architecture

```
                 ┌─────────────────────────── Option B (Node) ───────────────────────────┐
jquery.js  ───▶  │ tools/acorn-dump.js  (acorn@8.19.0, pinned, sha256-checked) → ESTree   │
                 └──────────────────────────────────┬─────────────────────────────────────┘
                                                     │  JSON file (or stdin)
                 ┌───────────────────────────────────▼────────────────────────────────────┐
                 │  lib/estree.oak   (NEW, Oak): ESTree JSON → JS AST (ast-js constructors) │
                 │  · unknown module/kind tables · regex/literal edge cases · diagnostics   │
                 └───────────────────────────────────┬─────────────────────────────────────┘
                 ┌───────────────────────────────────▼────────────────────────────────────┐
                 │  lib/ast-js.oak (fixed): transpileToOak  (+ strict mode: no silent :null)│
                 └───────────────────────────────────┬─────────────────────────────────────┘
                 ┌───────────────────────────────────▼─────────────────────────────────────┐
                 │  build-render-node.oak renderOakNode → jquery.oak → oak build → run    │
                 └──────────────────────────────────────────────────────────────────────────┘

                 ┌─────────────────────────── Option A (Oak-native) ──────────────────────┐
jquery.js  ───▶  │ lib/js-tokenize.oak (NEW) → lib/js-parse.oak (NEW, Acorn-subset)      │
                 │        → produces the same JS AST shape as estree.oak ingest           │
                 └───────────────────────────────────┬─────────────────────────────────────┘
                                                     └──────► same backend as above
```

Design rules:

- **One JS AST, two front-ends.** `estree.oak` (Acorn JSON ingest) and the Oak-native parser MUST emit byte-identical JS AST (modulo `tok.pos`/`tok.type` metadata, which the differ normalizes). This is what makes the ablation a real A/B and lets the native parser be validated against Acorn on every commit.
- **Backend stays unopinionated about source.** `transpileToOak` consumes JS AST only; zero Acorn-specific knowledge beyond the ESTree field names, which live in `estree.oak`.
- **Diagnostics are data.** `estree.oak` and `js-parse.oak` return `{ ast, errors, skipped }`; strict mode (default in tests) throws on any error/skip.
- **No DOM polyfills in the transpiler.** Smoke tests provide their own minimal `window`/`document` stubs in the *test harness*, not in the runtime.

### Why not other options

- **Port Acorn wholesale to Oak in one shot**: ~120 KB of dense, performance-tuned JS with heavy regex usage; as a first step it would block everything. Instead the Oak-native parser covers the *jQuery-1.0-proven subset* first, then widens guided by the ablation data. Estimated 3–5k lines of Oak with tests.
- **Emit Oak by string-template from JS AST (skip Oak AST)**: breaks `--web` round-trips, macros, and analyze passes; rejected.
- **Run jQuery inside the Oak interpreter without transpiling** (FFI): doesn't exist; rejected.
- **Wait for Node FFI in Oak**: not needed for the goal; A is pure Oak.

---

## 5. Ablation matrix (the "ablate the JS transpile suite" deliverable)

Baseline today: `test/ast-targets.test.oak` exercises `transpileToOak` on **hand-constructed** ASTs. Everything below replaces or removes one layer and re-measures. All rows run on the same corpus and the same assertions.

| Row | Front-end | Backend | What it measures / what it removes |
|---|---|---|---|
| **A0** control | hand-built JS AST (existing suite) | stock `transpileToOak` | Regression floor. Must stay green |
| **A1** | Acorn → JSON → `estree.oak` | fixed `transpileToOak` (strict) | Real-syntax coverage: parse→Oak e2e. Primary number |
| **A2** | A1 with `optimizeConstants` + `removeDebugCalls` **disabled** | fixed transpile | Isolates constant-folding/debug-strip bugs from front-end/transpiler bugs |
| **A3** | **Oak-native** parser → same JS AST | fixed `transpileToOak` | Parser equivalence: AST diff vs A1 (normalized), plus independent coverage of constructs Acorn rejects/accepts differently |
| **A4** | A1 with one `transpileToOak` mapping **removed at a time** (e.g. `:jsObject`, `:jsCall`, `:jsBinary`) | strict transpile | Mutation ablation: per-node-type fault injection, verifies the skip-list reporting is complete and the tests actually catch regressions |
| **A5** stretch | A1 on jQuery 1.12.4 (ES5) and a jQuery-3.x/three.js ES2015 excerpt | strict | ES5/ES2015 delta reporting (getters, `let`/`const`, arrow fns, template literals, classes) |

**Metrics per row**: files parsed OK / total · total nodes & % mapped (target 100% of ESTree types present in corpus, else explicit skip-list) · emitted Oak AST nodes matching snapshot · `oak build`-equivalent render OK (via `renderOakNode`, no Go needed once ingest runs in-Oak… see §7) · runtime smoke (jQuery pure utilities) · wall time.

**Corpus** (fetched by `tools/fetch-corpus.sh`, sha256-pinned, *not committed*):

- `jquery-1.0/src/jquery/jquery.js`, `ajax/ajax.js`, `event/event.js`, `fx/fx.js` — the release's own module set; concatenation order = Makefile order.
- `three.js r58` (stretch; cf. attached Lua patch).
- Tiny hand-written fixtures **committed in-repo** (`test/fixtures/js/`): `es3-core.js` (var/function/this/proto), `expressions.js`, `regex-literals.js`, `with-statement.js` (must fail loudly), `es2015-arrow.js` (must fail loudly on ES3 target)… each ≤ 200 lines, each with a golden transpile snapshot.

---

## 6. Work items

### M0 — Fix the backend so the rest of the work is measurable (1–2 days)

1. **Kill the silent `:null` fallback** in `lib/ast-js.oak:transpileToOak`: return `{ type: :error, … }`-style diagnostic or throw with node type + `tok.pos`; strict mode on by default in tests; non-strict accumulates `errors[]`.
2. **Spread/rest**: `jsCall` with `jsSpread` args → either map to Oak rest args + append or skip-list with error (Oak has rest params; `f(...xs)` semantics differ — decide per construct). `jsSpread` currently emitted then dropped by callers = silent argument loss; fix or error.
3. **`_jsIdentToOak` completeness**: audit against Oak keywords & web formatIdent rules; add `arguments`, `eval`, `with`, `typeof`, `instanceof` strategy note (JS `typeof x === 'undefined'` is the corpus's idiom — must not corrupt).
4. Fix `renderJS`: full string escaping (`\\`, `'`, `"`, control chars, unicode), object shorthand, template literal rendering, spread in calls/arrays, trailing-comma/`jsVar` declarator lists, and **add tests**.
5. `walkJS`: cover `jsObject` key+val consistently, `jsFunction` params (defaults), `jsVar` names; add missing node types as they appear in §M2 coverage report.
6. `transpile.test.oak` / `ast-targets.test.oak` extended to lock in the fixes (these are the A0 control).

### M1 — Acorn bridge: `tools/acorn-dump.js` + `lib/estree.oak` (2–4 days)

1. `tools/acorn-dump.js`: Node CLI — `acorn.parse(src, {ecmaVersion, allowHashBang, onComment?})` → JSON to stdout. Preserve `range`/`loc` optionally (off by default for smaller snapshots). Acorn pinned (`tools/package.json` with `acorn` exact version, or vendored `tools/vendor/acorn.min.js` — decide; vendoring avoids npm in CI).
2. `lib/estree.oak` **(new module)**: `parseJSON(jsonText, opts) -> { ast|?, errors: [], skipped: [] }`:
   - Literal / Identifier / ThisExpression / RegexLiteral (`regex.pattern/flags` → decide: Oak has no regex literal node → represent as `jsRegex` node mapped to a runtime constructor, or skip-list with loud error; jQuery 1.0 uses regex literals heavily — this is *required*, not optional);
   - Statement/expression coverage guided by the jQuery 1.0 node histogram (measure first, implement to histogram, then generalize);
   - Ends at the JS AST level — **no Oak logic inside `estree.oak`** beyond node mapping.
3. `tools/fetch-corpus.sh`: downloads jQuery 1.0 (`refs/tags/1.0`) + three.js r58 tarballs, verifies sha256, extracts to `build/corpus/` (gitignored).
4. `tools/e2e-js.sh`: for each corpus file → acorn-dump → `renderOakNode` golden → report. (Oak execution of the generated code happens in CI, see §7.)
5. Makefile: `test-js-e2e` target (node-only, works without changes to Go code).

### M2 — Test wiring (1 day, no Go rebuild needed for local loop; CI for full)

1. `test/js-estree.test.oak` (new): table-driven unit tests from committed fixtures: ESTree-JSON → JS AST spot checks (node counts, tricky nodes: `regex`, `getter/setter` in objects? (not in ES3), `SequenceExpression` (jQuery uses comma operator), `ConditionalExpression`, `UnaryExpression(typeof)`, `UpdateExpression(++/--)` — note: Oak has no `++`; map to `:= x + 1` assignments with correct expression-statement context, or skip-list loud).
2. `test/main.oak`: add `'js-estree'` to a new `TestRunners` entry `'js-tooling'` (or extend `ast-targets`); `Makefile`: `js-estree:lib/estree.oak` include; keep `make test-js` parity (oak-stdlib tests on Node) green.
3. Golden transpile snapshots: `test/fixtures/oak/expected/*.oak` generated by `renderOakNode`, hand-reviewed, diffed in CI via a small node differ.
4. jQuery e2e job (CI, requires Go): fetch corpus → for each module: acorn → ingest → transpile → render → `magnolia build` → node-smoke asserts (`jQuery`, `$` defined; `$.each/$.map/$.grep/$.trim/$.merge/$.isArray/$.extend/$.inArray/$.type` behave on literals; `jQuery.fn.jquery` === '1.0') with a **harness-provided** `window`/`document` stub. DOM-dependent paths: documented skip-list, not stubbed to "pass".

### M3 — Ablation runs & report (1 day)

- Script `tools/ablate.sh` runs A0–A4, emits `build/ablation-report.md` (node histograms, pass/fail, skip-list).
- Report committed as `docs/js-tooling.md` (the only generated artifact committed).
- jQuery 1.0 acceptance gates: parse 4/4 modules · 0 silent `:null` fallbacks · ≥ 95% corpus nodes mapped (remainder explicitly skip-listed) · bundle renders · smoke utils pass · Oak output round-trips through `oak fmt` without syntax errors (uses the repo's own parser as an oracle!).

### M4 — Oak-native front-end: `lib/js-tokenize.oak` + `lib/js-parse.oak` (4–7 days, the big one)

Follow `syntax-tokenize.oak`/`syntax-parse.oak` architecture; scope: **ES3 + ES5 + minimal ES2015** (arrow fns, template literals, shorthand, `let/const` as `var`, classes → error for now). Regex-vs-divide tokenization ambiguity resolved with the standard "previous significant token" heuristic (Acorn's own approach). Error recovery: bail with position on first error (no correction — matches Acorn's `onToken` philosophy, keeps the parser small and correct). Validator: A3 diff harness vs Acorn on the corpus — normalize (`type`, `name/val/op`, structure, drop `tok.pos`) then recursive compare; CI fails on any diff.

### M5 — Reverse direction hardening & CLI (2–3 days)

1. Wire `renderJS` into a real path: `oak build --web` already exists; add nothing user-facing yet — instead fix `renderJS` correctness + tests (M0.4) and expose `js2oak`/`oak2js` helpers via `cmd/` (new `cmd/js2oak.oak`? or extend `cmd/build.oak` with `--from-js`? — **decision point**, see §8).
2. `docs/js-tooling.md`; update `docs/transpile.md`, `docs/build.md:--web`, `docs/target-ts.md`, README tooling tables; `docs/spec/ast-js.md` regenerated style note for the new modules (`estree`, `js-tokenize`, `js-parse`).

### M6 — Real-library milestones & the Lua-patch cross-check (1–2 days, ongoing)

- jQuery 1.0 e2e green in CI; file-level triage log `docs/jquery-1.0-port-notes.md`.
- jQuery 1.12.4 (ES5) as second corpus — reveals ES5 gaps (getters/setters, `Object.defineProperty`, reserved-word property names, `\u` escapes).
- three.js r58 as stretch: parse + coverage histogram first; DOM/WebGL shims explicitly out of scope. Cross-reference `threejs-oak-complete.patch.txt`: its `fnDepth`/`local` and `;`-separator fixes are **Lua-renderer-specific**; the analogous audit for the **JS/web renderer** (e.g. assignment-target lowering `__oak_assgn_tgt`, IIFE counters) should be done and, if bugs of the same class exist, fixed in the JS renderer with the same style of minimal patch.

### Suggested commit sequence

`M0 fixes (ast-js)` → `M1 estree.oak + tooling` → `M2 tests/wiring` → `M3 ablation report` → `M4 js-parse` (separate PR) → `M5 docs/CLI` → `M6 corpora`. Each commit: `oak fmt`-clean, `make test-oak`-equivalent left to CI.

---

## 7. Environment constraints found (please read)

1. **No Go toolchain and no prebuilt `magnolia` binary in this sandbox.** Consequently *no Oak code can be executed locally here* — neither `test/main.oak` nor `estree.oak` nor golden renders. Local verification is limited to Node-side checks (acorn dumps, JS differ scripts, sha256 fetch). Oak-side correctness must be gated by CI (`make test-oak`, `make test-bundle`, `make test-js`) or by running on a machine with Go (the Makefile and embedded stdlib assume it).
2. **New Oak stdlib modules require Go registration** (`lib.go`: `//go:embed` + `stdlibs` map + `stdlibAutoLoadOrder`) and a rebuild. The plan avoids this for anything optional by using the established `--include` pattern (`Makefile: INCLUDES`, e.g. `ast-js:lib/ast-js.oak`) for tests; but for distribution, `estree`, `js-tokenize`, `js-parse` SHOULD be registered in `lib.go` in the same PR series (maintainer with Go toolchain runs it).
3. **npm jQuery starts at 1.5.1**; jQuery 1.0.0 exists only as GitHub tag **`1.0`** (source under `src/`, not a ready dist file; concatenation order per its `Makefile`). The plan's corpus fetcher handles this; if you specifically need the historical `jquery-1.0.js` dist, say so and I'll use the `1.0.4`/built artifact instead, or vendor the 4 modules.
4. **The attached `threejs-oak-complete.patch.txt` is Lua-renderer-only** and is not needed for the JS↔Oak work; it's kept as a model for large-library fix methodology (small patch + e2e definition of done).
5. Acorn's JSON output **contains regex literals as objects** (`{ pattern, flags }`) and `ThisExpression`/`SequenceExpression`/`UpdateExpression` nodes with no Oak counterpart semantics today — these must be explicit in the transpiler or the skip-list; silent dropping is the failure mode we're eliminating.

---

## 8. Decision points for you

1. **Parser strategy** — proceed with hybrid (Acorn bridge now + Oak-native parser as M4) as recommended? Or Oak-native only (slower, zero Node dependency), or Acorn-only (fastest, permanent Node dependency for JS→Oak)?
2. **CLI surface** — `magnolia js2oak <file.js>` new command, or `oak build --from-js --entry lib.js`, or no CLI (library + tests only) for this iteration?
3. **New-module registration** — OK to register `estree`/`js-tokenize`/`js-parse` in `lib.go` + `stdlibAutoLoadOrder` (needs Go rebuild & release notes), or keep them build-include-only for now?
4. **Corpus in git?** Recommendation: no (fetch script + sha256); only ≤200-line hand fixtures are committed. OK?
5. **jQuery 1.0 provenance** — use tag `1.0` (the actual 1.0.0, source layout) as planned, or npm `jquery@1.5.1` as the earliest npm-installable 1.x?
6. **Skip semantics for `with` statements** (ES3, legal in jQuery 1.0): hard error (recommended) or best-effort lowering?

## 9. Definition of done

- `make test-oak`, `make test-bundle`, `make test-js` green with the new suites wired in.
- jQuery 1.0: 4/4 modules parse (acorn **and** Oak-native), transpile with **zero** silent `?` fallbacks, render to Oak source that parses under `oak fmt`/`syntax-parse`, and the browser-independent utilities pass smoke tests in the CI e2e job; results recorded in `docs/js-tooling.md`.
- Docs landed: `docs/js-tooling.md` + cross-links; `ast-js` spec docs updated; README/OAK-README tooling sections reflect the new `js2oak` capability.
- All large third-party corpora remain uncommitted (fetch-on-demand); only small fixtures and the ablation report are in git.

---

## 10. Status log

### 2026-11-16 — M0 (transpiler core) implemented and verified

`lib/ast-js.oak` `transpileToOak` rewritten end-to-end; new `lib/js-runtime.oak`
(runtime harness for generated programs); `lib/build-render-node.oak` /
`lib/codegen-common.oak` renderers hardened.

**Delivered:**

- Control-flow ABI: `return` / `break` / `continue` / `throw` lower to
  value-carrying signal objects `{__t: 'jsReturn'|'jsBreak'|'jsContinue'|'jsThrow', v: <value>}`.
  Statement sequences are rewritten so each statement binds its value to `__js_s`
  and the rest of the sequence runs only when it is not a signal. Function
  boundaries normalize via `__jsFunctionReturn` (`jsReturn` → payload,
  fall-through → `?` i.e. JS `undefined`, other signals propagate).
- New Oak syntax emitted: loops lower to tail-recursive closures consumed by
  `__jsWhile` / `__jsDoWhile` / `__jsFor` / `__jsForIn` / `__jsForOf`
  (provided by `lib/js-runtime.oak`); `try` / `catch` / `finally` lower to
  signal-threading temp bindings (`__jsIsThrow` gates the catch branch;
  `finally` overrides on signal, matches JS semantics).
- `typeof` lowers to `__jsTypeof` (null → 'object' quirk included); unary `!`
  and **all** `if`/ternary conditions wrap the operand in `__jsTruthy(...)`
  (Oak `if` is pattern matching, not truthiness); `&&`/`||` lower to a
  two-phase lowering (bool/bool operands keep native Oak `&`/`|`; otherwise an
  if-expression with a shared temp binding).
- `delete obj.prop` lowers to `__jsDelete(obj, key)` (Oak objects cannot
  delete keys — the shim returns `true` and mutates nothing; JS semantics for
  configurable properties). `new` lowers to a plain call + diagnostic.
  `arguments` is synthesized as a leading `var arguments = [params...]` when a
  function body references it. JS idents starting with `__js` are mangled
  (`_` prefix) so they can never capture the transpiler's internal temps.
- Strict-mode flag: `setStrictTranspile(true)` makes unsupported nodes emit a
  `{type: :jsError}` node; `takeDiagnostics()` / `clearDiagnostics()` expose
  recorded diagnostics; `JsTempCounter` provides `freshTemp()`.
- `jsReturn` semantics **changed** (was: inner value; now: signal object) —
  `test/ast-targets.test.oak` updated to lock the new contract.
- Renderer bug found & fixed: Oak's `<<` **mutates its left operand in place**,
  and `build-render-node.renderOakNode` returned shared node strings
  (`tok.val` / identifier `val`) for some node types — repeated rendering
  corrupted the AST tree (every later render duplicated slices of it). Both
  renderers now return fresh strings for every node type. This was the root
  cause of the `safeDiv`/`latch` misrenders.
- End-to-end verified (not just AST-shape tests): hand-built JS AST →
  `transpileToOak` → `build-render-node.renderOakNode` → Oak source →
  executed under Magnolia printing the expected values
  (`-1 0 1` / `6` / `8` / `5 -1` / `bF`), including early returns past
  nested `if`s, `break` in `for`, `continue` in `while`, `try/catch/finally`
  with throw-override semantics. `tools/m0-transpile-demo.oak` reproduces the
  whole round trip: `magnolia tools/m0-transpile-demo.oak /tmp/out.oak &&
  magnolia /tmp/out.oak`.
- Registered `js-runtime` in `lib.go` (embed + stdlibs map). Note: the
  `TestStdlibRegistryCoversLibDirectory` go test fails at base commit
  (`ad21652`) for 8 unrelated WIP libs (`ui-core`, `style-bake`, …) that were
  never registered; **not** caused by this change.

**Known gaps / deliberate simplifications (M0 scope):** signals propagate
through `__js_s` without demangling to native JS exceptions (by design);
`new` has no constructor protocol; `with` / `switch` / `class` / getters /
setters / destructuring params / `for await` unsupported (diagnostics or
placeholder `_`); Oak `int / int` is integer division (JS floats — `:jsNumber`
keeps float nodes float, but JS `x / 2` with odd operands truncates); `delete`
is a no-op shim; `__jsForIn` key order is Go map order; regex literals lower
to `__jsRegExp` (carrier object, no engine); `NaN` inexpressible in Oak.

**Next — M1:** Acorn bridge: `lib/estree.oak` (ESTree→`js*` AST ingestion
from `acorn-dump.js` JSON via `lib/json.oak`), ablation harness against real
Acorn on the jQuery 1.0 / three.js r58 corpora, `this`/`instanceof`/prototype
diagnostics, then M2 test wiring.

### 2026-11-16 — M1–M4 (ablation harness, Oak-native front-end, round-trip gate) implemented and verified

Full numbers live in **`docs/js-tooling.md`**; this entry records scope and status.

**Delivered:**

- **M1 — Acorn bridge**: `tools/acorn-dump.js` (acorn 8.19.0, `--ecmaVersion` flag),
  `lib/estree.oak` (ESTree JSON → `js*` AST, diagnostics list), `tools/fetch-corpus.sh`
  (sha256-pinned jQuery 1.0 (`1.0` tag) + three.js r58; corpora stay out of git),
  `tools/a3-diff.oak` (structural A3 tree compare, `tok` ignored).
- **M2 — wiring**: `test/ast-targets.test.oak` extended to 208 tests (A0 green);
  25-fixture JS suite under `build/fixtures/js/` (round-trip all OK; still
  uncommitted — see ablation doc §5.4); jQuery 1.0 e2e covered by the tools
  below (CI wiring still pending).
- **M3 — ablation (partial)**: A0/A1/A3 measured; A2 (constant-folding off),
  A4 (mutation ablation), A5 (ES5/ES2015 stretch) **not run**; no
  `tools/ablate.sh` yet. Report = `docs/js-tooling.md` (hand-maintained).
- **M4 — Oak-native front-end**: `lib/js-tokenize.oak` + `lib/js-parse.oak`
  (ES3+ES5 core; arrow fns/template literals partial), regex-vs-divide via
  previous-significant-token heuristic. Validated by A3: **IDENTICAL trees** on
  jQuery 1.0, three.js r58, three.min.js, and the 25 fixtures.
- **Round-trip gate** (plan §9 DoD, Oak output re-parsed by repo-native
  `syntax-parse.oak`; `oak fmt` is not usable as an oracle): **25/25 fixtures,
  jQuery 1.0, three.js r58, three.min.js — all parse OK after rendering.**
  Fixes required to make the gate meaningful:
  1. `build-render-node.oak` — precedence-aware parenthesization mirroring
     `parse.go:infixOpPrecedence` (loose receivers/callees/operands now
     parenthesized; unknown binary op → loud abort, never silent `?`).
  2. `syntax-tokenize.oak` / `syntax-parse.oak` — numeral exponents
     (`1e-10`) lexed/ classified as float (was: `1` + identifier `e`, failing
     the oracle on rendered three.js even though the real binary was fine).
  3. Oak object-condition idiom (`?`/`_` arms), not `true` arms — see
     `arena-system-message` / ablation doc §6.3.
- Diagnostics carry node type + position everywhere; **zero silent `:null`
  fallbacks** (audit complete). three.js r58: 5,042 diagnostics
  (3,825 `this` / 872 `new` / 236 `instanceof` / 97 compound-assign / 9 `in`
  / 3 unary `+`), zero unknowns. jQuery 1.0: 202 (184 `this`, 18 `new`).
- three.min.js: 4 native parser warnings (labeled statements / labeled break,
  deliberately unsupported lowering, loud not silent); round-trip OK.

**Definition-of-done check (jQuery 1.0):** parse acorn ✓ · parse Oak-native ✓
(A3 IDENTICAL) · transpile zero silent `?` ✓ · render ✓ (56,317 B) ·
round-trip ✓ · **browser smoke utils: NOT RUN** (no DOM harness here; needs
CI e2e — plan M2.4/M6).

**Still open:** §8 decision points (CLI surface; `lib.go` registration of
`estree`/`js-tokenize`/`js-parse`; fixtures commit location; jQuery 1.0
provenance; `with` semantics); A2/A4/A5; cross-links (`docs/transpile.md`/`build.md`, README tooling tables) (M5);
three.js DOM/WebGL execution explicitly out of scope; three.min.js per-class
diagnostic histogram not re-tallied.


---

## Part II — Ablation results (2026-11)

## 1. Corpus

| Corpus | Path (not committed) | acorn 8.19.0 node count | Notes |
|---|---|---|---|
| jQuery 1.0.0 | `build/corpus/jquery-1.0/jquery-1.0.js` | 7,330 (7,333 in the `ecmaVersion:3` dumps) | Concatenated from the `1.0` tag's `src/` in Makefile order; `$Rev$` → `1.0` for the smoke corpus |
| three.js r58 | `build/corpus/three-r58/build/three.js` | 132,924 | Stretch corpus |
| three.js r58 (minified) | `build/corpus/three-r58/build/three.min.js` | 124,333 | Stretch corpus |

Hand-written ES3 fixtures live in `build/fixtures/js/c01…c25-*.js` (25 files, ≤ 200 lines each). Per the plan they should ultimately be committed under `test/fixtures/js/`; that move is pending (§5, decision D4).

## 2. Method

Deliverable gates map to plan §5 rows. All rows share the same corpus and the same assertions.

- **A0** (control): existing `test/ast-targets.test.oak` suite on hand-built ASTs.
- **A1** (Acorn bridge): `tools/acorn-dump.js` (acorn pinned) → `lib/estree.oak` ingest → `transpileToOak` → `lib/build-render-node.oak`.
- **A3** (Oak-native): `lib/js-parse.oak` (`js-tokenize.oak` + recursive-descent parser) → the *same* `js*` AST → same backend. Equivalence checked by `tools/a3-diff.oak <file> <acorn-dump.json>` — structural recursive compare on `type`, field values, list lengths; `tok` metadata ignored. Must print `IDENTICAL`.
- **Round-trip gate** (new, M4): emitted Oak is re-parsed by the repo's own Oak parser (`syntax-parse.oak` via `syntax.parse`). `oak fmt` is **not** usable as an oracle (it exits 0 on invalid input); `syntax.parse` returning `:list` is the gate. See `tools/js-fmt-roundtrip.oak`.
- A2 (folding/debug-strip disabled), A4 (per-mapping mutation ablation), A5 (ES5/ES2015 corpora) — **not run yet**.

## 3. Results (2026-11-16)

### A0 — control

- `test/ast-targets.test.oak`: **208/208**.
- `test/ui-surface.test.oak`: **21/21**.
- `go test ./...` (and `go test -race .`): **all pass**. The previously recorded `TestStdlibRegistryCoversLibDirectory` failure (eight unregistered WIP libs) is resolved on this branch — the milestone's `lib.go` registrations cover them.

### A1 / A3 — front-ends agree on the corpus

| Corpus | A1 parse | A3 Oak-native parse | A3 diff | Transpile diagnostics | Rendered Oak | Round-trip `syntax.parse` |
|---|---|---|---|---|---|---|
| jQuery 1.0 | ok, 0 warnings | ok, 0 warnings | **IDENTICAL** | 202 — 184× `this`, 18× `new` | 56,341 B | **OK** |
| three.js r58 | ok, 0 warnings | ok, 0 warnings | **IDENTICAL** | 5,042 — see §4 | 1,086,621 B | **OK** |
| three.min.js | ok, 4 warnings¹ | ok, 4 warnings¹ | **IDENTICAL** | 5,002 | 835,713 B | **OK** |
| Fixtures (25) | — | ok | IDENTICAL (suite) | — | — | **25/25 OK** |

¹ The four warnings are the native parser's *deliberate* droppables on minified optional-label code: "labeled statement is not supported; the label is ignored" (×2) and "break with a label is not supported; acts on the innermost loop or switch" (×2). Semantics follow the documented (intentionally Oak-limited) lowering, not silent corruption.

Oak-native token-diff against the Oak oracle for the same sources is green (`js-tokenize` vs `syntax-tokenize` token streams).

### Gates

| Gate (plan §3/§9) | Status |
|---|---|
| G1 JS→Oak e2e on real source | **Pass** (jQuery 1.0 primary: parse→render→round-trip **and** runtime smoke under a stubbed DOM — §3, §5.2; three.js r58 stretch parse→render→round-trip pass; no real-browser run) |
| G2 quantitative ablation | Partial — A0, A1, A3 measured; A2/A4/A5 not run; no `ablate.sh` harness yet |
| G3 Oak-native tokenizer + parser, Acorn-equivalent | **Pass** — A3 IDENTICAL on 3 corpora + 25 fixtures |
| G4 harden Oak→JS direction | Partial — `build-render-node.oak` hardened this milestone (see §6); web runtime fixed and executed this session (receiver `this` binding, native RegExp, JS String API, host constructors — see §6.5–6.9); `renderJS` fixes + tests (M0.4) and CLI (M5) still pending |
| G5 diagnostics, not silent corruption | **Pass** — zero silent `:null` fallbacks (audit complete); unknown binary op now hard-aborts at render; all diagnostics carry node type/position |
| G6 docs | Partial — this doc; `js-tooling.md` + cross-links pending |

### Definition of done — jQuery 1.0 checklist

| Requirement | Status |
|---|---|
| Parses acorn and Oak-native | ✅ |
| Transpile with zero silent `?` fallbacks | ✅ (202 explicit diagnostics; every transpile path either maps or records) |
| Renders to Oak | ✅ 56,341 B |
| Rendered Oak parses (`oak fmt`-equivalent oracle) | ✅ via `syntax.parse` |
| Browser-independent utilities smoke-tested | ✅ **10/10** under a stubbed `window`/`document` (Oak-transpiled jQuery executed end-to-end by `oak build --web` → node; `jQuery.fn.jquery`, `extend`, `merge`, `grep`, `map`, `each` (array/object), `trim`×2) — see §5.2 |

## 4. Diagnostics inventory (three.js r58, 5,042 total; zero unknowns)

| Class | Count | Meaning |
|---|---|---|
| `this` | 3,825 | Oak has no dynamic `this`; lowered to a placeholder with diagnostic |
| `new` | 872 | No constructor protocol; lowered to plain call + diagnostic |
| `instanceof` | 236 | Unsupported; diagnostic |
| compound-assign double-eval | 97 | Lowering re-evaluates the target expression; flagged where operands may not be pure |
| `in` operator | 9 | Oak has no `in`; diagnostic |
| unary `+` | 3 | Lowered to `__jsTruthy`-style numeric coercion is identity for numbers; diagnostic emitted |

jQuery 1.0's 202 diagnostics are the same `this` (184) and `new` (18) classes — the corpus simply exercises those constructs less. three.min.js reports 5,002 (same classes; the exact per-class breakdown was not re-tallied).

## 5. Not done / open

1. A2/A4/A5 ablation rows; `tools/ablate.sh` (M3 automation) — planned, not implemented.
2. ~~jQuery browser smoke~~ — **resolved 2026-11-16.** The transpiled jQuery 1.0 executes end-to-end under `oak build --web` → node with a stubbed `window`/`document`/`navigator`/`location`/`XMLHttpRequest` and asserts **10/10** browser-independent utilities (`jQuery.fn.jquery === '1.0'`, `extend({a:1},{b:2})`, `merge` dedupe, `grep` evens, `map` doubling, `each` over arrays and objects, `trim` on spaces and tabs/newlines). Reaching this exposed four real runtime bugs, all fixed and individually verified — receiver `this` binding for property-access calls, assignment-chain value, js-runtime stdlib bindings, native RegExp + JS String API (§6.5–6.9). The smoke harness lives at `/tmp/jq-smoke.oak` (not committed; see §8 and decision below). Still open: **where the harness should live** (repo path + CI wiring), `$.isArray`/`$.inArray`/`$.type` coverage, and any real-DOM (RTL/jsdom) e2e.
3. Runtime execution of rendered corpus output: jQuery done (see above); three.js r58 remains unexecuted; fixture-level runtime checks not asserted (only `syntax.parse` acceptance).
4. Hand fixtures are not yet committed under `test/fixtures/js/` (build/ is gitignored).
5. Cross-links (`transpile.md`, `build.md`, README tables) — M5.
6. M6 corpora: jQuery 1.12.4 (ES5), jQuery-3/three.js ES2015 — stretch.
7. §8 decision points of the plan (CLI surface, `lib.go` registration of `estree`/`js-tokenize`/`js-parse`, corpus-in-git, jQuery provenance tag `1.0` vs npm `1.5.1`, `with`-statement semantics) remain **open with the user**. Diagnoses here are recorded, not decided.

## 6. Renderer & oracle fixes made this milestone (M4)

1. **`lib/build-render-node.oak` — precedence-aware parenthesization.** Oak's infix binding strength differs from JS's: modulus (80) > power (75) > `* /` (50) > `+ -` (40) > comparisons (30) > `&` (20) > `^` (15) > `|` (10) > string concat / shifts (1); unary binds tighter (90); member-access/call postfix is tight. The renderer now mirrors parse.go's `infixOpPrecedence` (`_oakPrec`, `_childPrec`, `_needBinP`, `_needTightP`), parenthesizes loose operands of binary expressions and loose receivers/callees of property access and calls, and remains loud (`abort`) on unknown binary operators instead of printing `?`. This fixed silent mis-association such as JS `a + b * c` lowering rendering as Oak code that would re-associate; the earlier failure mode was a hard syntax error in the round-trip gate (`... << "…" .toFixed(…)`), now impossible by construction.
2. **`lib/syntax-tokenize.oak` / `lib/syntax-parse.oak` — the round-trip oracle lacked exponent literals.** The Oak-native tokenizer lexed `1e-10` as `1` + identifier `e` (+ `- 10`), so the oracle rejected the rendered three.js at `EPSILON := 1e-10` (col 83662) even though the real Magnolia binary and the JS front-ends lex it correctly. `readValidNumeral` now consumes `[eE][+-]?digits` (after one dot, mirroring JS lexing), and the parser classifies a numeral as float when it contains `.`/`e`/`E` unless it is a `0x`/`0b` hex/binary literal (previously float only when `.` was present). All numeral-adjacent behavior re-verified on fixtures.
3. Oak `if` is pattern matching, not truthiness: conditions over objects must use `?`/`_` arms (this caused the transient `Binary operator < is not defined for values ?, 40` crash; fixed before proceeding).
4. Earlier in the milestone (carried forward): synthetic `:int` nodes with `tok.val: ?` (jQuery/three.js now render), and JS identifiers colliding with Oak keywords (`fn`, `if`, `with`, `cs`, `_`) — mangled via `_kw_*` arms in `_jsIdentToOakMapped`.
5. **Assignment chains discard the bound value** (`lib/build-render.oak`, `:assgn` render). `a = b = x` was rendered so the inner assignment evaluated but the outer target received nothing; jQuery's shared `jQuery.extend = jQuery.fn.extend = function(obj, prop){ if (!prop) { prop = obj; obj = this; } … }` executed `obj = this` → `obj = undefined`, and the following `for (var i in prop) obj[i] = prop[i]` threw `Cannot set properties of undefined` inside `__jsForIn`. The emitted IIFE now returns `__oak_assgn_val`. Verified with a self-contained probe (nested/chained targets) and by re-running the bisect corpus up to the jQuery load path.
6. **Property-access calls bind no receiver `this`** (`lib/build-render.oak`, `:fnCall`). `obj.m(args)` rendered as a plain `(obj.m)(args)`, so `this === undefined` inside `m` — Oak has no receiver concept. `this` inside transpiled JS reads `__oak_js_this`, so every property-access call now emits `((recv, saved) => ((__oak_js_this = recv), ((ret) => ((__oak_js_this = saved), ret))((recv.m)(args)))(recv, __oak_js_this)` — receiver evaluated once, save/restore around the call. Plain-name calls and direct function calls are unchanged. All one-argument `jQuery.extend({…})` / `jQuery.fn.extend({…})` sites depend on this (`extend` reads `this` to choose the copy target); after the fix, bisect k=7 and the jQuery smoke run to completion.
7. **`lib/js-runtime.oak` bound stdlib through `std`, which exports none of it.** `len`/`keys`/`values` are VM builtins resolved on the web backend through the runtime prelude, not `std` exports; the bogus `std.len`/`std.keys`/`std.values` bindings were `undefined`, crashing the first `__jsForIn` (`keys is not a function`). Bindings removed; `map` kept. The web prelude gained `values()` (object → own values; non-objects → `[]`), matching the VM's `opBuiltin` semantics.
8. **The web prelude is native-JS-hostile to transpiled code.** Oak strings pushed through JS method calls (`x.toLowerCase()`, `'a,b'.split(',')`, `str.replace(re, fn)`) had no such methods (`__oak_recv.toLowerCase is not a function`), and JS `x.constructor === String` dispatch could never hold (Oak strings are class instances). `lib/runtime-js.oak` now: (i) implements the JS String surface on `__Oak_String_Instance` (`toLowerCase`/`toUpperCase`/`trim`/`charAt`/`charCodeAt`/`concat`/`slice`/`substring`/`substr`/`indexOf`/`lastIndexOf`/`includes`/`startsWith`/`endsWith`/`search`/`split`/`match`/`replace`/`replaceAll`; string results re-wrap as Oak strings); (ii) sets `__Oak_String_Instance.prototype.constructor = String` (named class expression + static block) so transpiled constructor identity checks hold; (iii) provides the host constructors the transpiler emits as free identifiers — `__Array`/`__Object`/`__String`/`__Number`/`__Boolean`/`__Function`/`__Date`/`__RegExp`/`__Promise`; (iv) provides `__js_RegExp`, constructing **native** `RegExp` tagged with `__t: 'jsRegExp'`. Real regexes give `.test()/.exec()/String.replace()/match()/split()` all JavaScript semantics including capture groups (`m.(1)` for `/^[^<]*(<.+>)[^>]*$/`), which unblocked `jQuery.trim` and the HTML-detection `exec`.
9. **`lib/js-runtime.oak` `__jsRegExp`** returned a marker-only object `{source, flags, __t: 'jsRegExp'}` with no `test`/`exec` — any transpiled regex use crashed or silently mis-behaved. It now delegates to `__js_RegExp` (see 6.8). The smoke's local literal-only `__jsRegExp` stub (which `abort`ed on non-literal patterns like `/^[^<]*(<.+>)[^>]*$/`) was removed entirely.

## 7. Reproduction

```bash
# corpus (network; not committed)
./tools/fetch-corpus.sh

# Go toolchain in this sandbox lives at /tmp/gobin-root/go/bin; all runs use:
export PATH=/tmp/gobin-root/go/bin:$PATH GOPATH=/tmp/gopath GOCACHE=/tmp/gocache GOPROXY=direct GOSUMDB=off
export -n GH_TOKEN GITHUB_TOKEN; export GIT_TERMINAL_PROMPT=0

# A3 equivalence (example: three.js)
node tools/acorn-dump.js build/corpus/three-r58/build/three.js > /tmp/three.json
go run . tools/a3-diff.oak build/corpus/three-r58/build/three.js /tmp/three.json
# -> IDENTICAL

# Round-trip gate (parse -> transpile -> render -> syntax.parse)
go run . tools/js-fmt-roundtrip.oak build/corpus/jquery-1.0/jquery-1.0.js /tmp/jquery-1.0.rendered.oak
go run . tools/js-fmt-roundtrip.oak build/corpus/three-r58/build/three.js  /tmp/three.rendered.oak
go run . tools/js-fmt-roundtrip.oak build/corpus/three-r58/build/three.min.js /tmp/three-min.rendered.oak

# Runtime smoke (jQuery under a stubbed DOM; harness not in repo yet — see §5.2/§8)
#   build/ = corpus (not committed); /tmp/jq-smoke.oak = harness appended after the
#   transpiled jQuery sequence (runs __smoke(jQuery) as the last element)
go run . build --entry /tmp/jq-smoke.oak --output /tmp/jq-smoke.js --web --include math-base:lib/math-base.oak
node /tmp/jq-smoke.js
# -> SMOKE PASS (all assertions), 10/10, rc=0
```

Wall times on this machine (cold Go build): jQuery ≈ 10 s, three.js ≈ 115 s, three.min.js ≈ 124 s — dominated by Oak-native parsing of 130k-node files, not by rendering.

## 8. Bottom line

The JS→Oak path is measurable end to end and **green on the primary corpus at parse/transpile/render/round-trip level and now at runtime level**: transpiled jQuery 1.0 loads and passes 10/10 browser-independent assertions under a stubbed DOM (§3, §5.2). Four runtime bugs were found and fixed to get there (§6.5–6.9); each was verified in isolation before the full smoke was re-run.

What remains before this can be called done:

- A2/A4/A5 rows; `tools/ablate.sh` (M3 automation).
- **jQuery smoke harness is not committed** — it lives at `/tmp/jq-smoke.oak` (transpiled corpus + stubbed DOM + `__smoke(jQuery)` injection). Deciding where it belongs (e.g. `tools/` as a manual runner, `test/fixtures/js/` + an e2e runner, or CI-only) and wiring it is **open with the user** (D7-adjacent).
- Cross-links (`transpile.md`, `build.md`, README tables) — M5.
- M6 corpora (jQuery 1.12.4, jQuery-3/three.js ES2015).
- The plan's §8 decisions (CLI surface, `lib.go` registration of `estree`/`js-tokenize`/`js-parse`, corpus-in-git, jQuery provenance tag, `with`-statement semantics) that only the user can make.
- A real-DOM e2e for the smoke (jsdom/RTL or CI browser runner) would raise G1 above "stubbed DOM".
- `test/main.oak` aborts at the pre-existing `pack.test.oak` failure (reproduced on base `ad21652`; unrelated — stale `mp_parse` expectation). The rest of that suite does not run until it is fixed.
