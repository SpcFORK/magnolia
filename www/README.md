# oaklang.org

The [oaklang.org](https://oaklang.org) website hosts documentation, updates, releases, and an overview of the Oak programming language project. The website is a partly dynamic web application written entirely in Oak.

## Architecture

The site uses a **hybrid static + dynamic** design: a build-time generator produces HTML for library docs and blog posts, while a runtime HTTP server serves those pages alongside a dynamic syntax-highlighting proxy.

### Build step (`src/gen.oak`)

The static generator runs at deploy time and produces all pre-rendered HTML:

1. **Library pages** — scans `lib/*.oak`, syntax-highlights each file via `oak cat --html`, wraps the output in `tpl/lib.html`, and writes it to `static/lib/{name}.html`.
2. **Blog posts** — reads Markdown files from `content/`, parses YAML frontmatter for metadata, converts Markdown to HTML with `md.parse()`, highlights embedded Oak code blocks, renders through `tpl/post.html`, and writes to `static/posts/{slug}.html`.
3. **Blog index** — generates a reverse-chronological listing of all posts via `tpl/list.html`.

### Web server (`src/main.oak`)

An Oak HTTP server on port **9898** with the following route layout:

| Route | Behavior |
|---|---|
| `/` | Serves `static/index.html` (home / tutorial page) |
| `/lib/`, `/lib/:name` | Serves pre-generated library docs |
| `/posts/`, `/posts/:name` | Serves pre-generated blog posts |
| `/highlight/*url` | Dynamic proxy — fetches a remote `.oak` file, highlights it, and returns HTML (supports `?start=N&end=M` line ranges and `?embed` for bare output) |
| `/*path` | Static file fallback for CSS, JS, images |

### Client-side scripts (`src/app.js.oak`, `src/highlight.js.oak`)

Oak source files compiled to JavaScript via `oak build --web` and output to `static/js/`:

- **bundle.js** — auto-detects the visitor's OS from `navigator.userAgent` and expands the matching install instructions on the home page.
- **highlight.js** — powers the GitHub highlight form, converting `github.com/.../blob/...` URLs into `/highlight/` proxy URLs with optional line-range parameters.

### Templates (`tpl/`)

Five Handlebars-style (`{{ key }}`) HTML templates shared between the generator and the server:

| Template | Used by | Purpose |
|---|---|---|
| `lib.html` | gen.oak | Library source page wrapper |
| `post.html` | gen.oak | Blog post page wrapper |
| `list.html` | gen.oak | Blog index listing |
| `highlight.html` | main.oak | Full-page highlighted code view |
| `highlight-embed.html` | main.oak | Bare highlighted code (embed mode) |

### Styles (`static/css/`)

- **main.css** — responsive layout, typography (IBM Plex Sans / Mono), light/dark theme via `.dark` class.
- **lib.css** — Oak syntax-highlighting color scheme and table styles for library reference pages.

## Structure

- Index
	- Examples
		- HTTP echo server
		- File I/O
		- Fibonacci
	- The project + rationale + history
	- Download links proxied via dynamic API Backends to `oaklang.org/download`
- Docs + guides + tooling documentation, following `golang.org`
	- "How to write Oak programs" `https://golang.org/doc/code`
- Blog (updates + releases)
- Playground (powerful but simple online IDE)
- Source link -> github.com/thesephist/oak

