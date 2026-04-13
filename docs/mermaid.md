# Mermaid Library (mermaid)

## Overview

`libmermaid` provides functions for programmatically creating [Mermaid](https://mermaid.js.org/) diagrams — flowcharts, sequence diagrams, class diagrams, pie charts, and more — and exporting them as Mermaid text, standalone HTML, or image files.

## Import

```oak
mermaid := import('mermaid')
```

## Quick Start

```oak
mermaid := import('mermaid')

// Create a flowchart
g := mermaid.graph(:td)
g |> mermaid.node('A', 'Start', mermaid.ShapeStadium)
g |> mermaid.node('B', 'Process')
g |> mermaid.node('C', 'End', mermaid.ShapeCircle)
g |> mermaid.edge('A', 'B', 'begins')
g |> mermaid.edge('B', 'C', 'finishes')

// Render as text
text := mermaid.render(g)

// Save as HTML preview
mermaid.saveHTML(g, 'diagram.html')

// Save as PNG (requires mermaid-cli)
mermaid.saveImage(g, 'diagram.png')
```

## Graph Types

| Constant | Value | Description |
|----------|-------|-------------|
| `FlowchartLR` | `:lr` | Left-to-right flowchart |
| `FlowchartTD` | `:td` | Top-down flowchart (default) |
| `FlowchartRL` | `:rl` | Right-to-left flowchart |
| `FlowchartBT` | `:bt` | Bottom-to-top flowchart |
| `SequenceDiagram` | `:sequence` | Sequence diagram |
| `ClassDiagram` | `:classDiagram` | Class diagram |
| `StateDiagram` | `:stateDiagram` | State diagram |
| `PieChart` | `:pie` | Pie chart |
| `Gantt` | `:gantt` | Gantt chart |
| `ERDiagram` | `:erDiagram` | Entity-relationship diagram |
| `GitGraph` | `:gitGraph` | Git graph |

## Node Shapes

| Constant | Mermaid Syntax | Appearance |
|----------|---------------|------------|
| `ShapeRect` | `A["label"]` | Rectangle |
| `ShapeRound` | `A("label")` | Rounded rectangle |
| `ShapeStadium` | `A(["label"])` | Stadium / pill |
| `ShapeCircle` | `A(("label"))` | Circle |
| `ShapeRhombus` | `A{"label"}` | Diamond |
| `ShapeHexagon` | `A{{"label"}}` | Hexagon |
| `ShapeTrapezoid` | `A[/"label"\]` | Trapezoid |
| `ShapeDefault` | `A` | Default (rectangle) |

## Edge Styles

| Constant | Mermaid Syntax | Appearance |
|----------|---------------|------------|
| `EdgeArrow` | `-->` | Solid arrow (default) |
| `EdgeDotted` | `-.->` | Dotted arrow |
| `EdgeThick` | `==>` | Thick arrow |
| `EdgeOpen` | `---` | Open line (no arrow) |

## Functions

### `graph(direction)`

Creates a new graph object.

- `direction` — One of the graph type constants (default: `:td`)

### `node(g, id, label, shape)`

Adds a node to the graph. Returns the graph for chaining.

- `g` — Graph object
- `id` — Unique node identifier
- `label` — Display text (default: same as `id`)
- `shape` — Shape constant (default: `ShapeDefault`)

### `edge(g, from, to, label, style)`

Adds an edge between two nodes. Returns the graph for chaining.

- `g` — Graph object
- `from` — Source node ID
- `to` — Target node ID
- `label` — Edge label (default: `''`)
- `style` — Edge style constant (default: `EdgeArrow`)

### `subgraph(g, id, label, builderFn)`

Adds a subgraph container. `builderFn` receives a fresh graph object to populate.

### `raw(g, line)`

Appends a raw Mermaid syntax line to the diagram.

### `setTitle(g, title)`

Sets the diagram title.

### `setTheme(g, theme)`

Sets the Mermaid theme (`'default'`, `'forest'`, `'dark'`, `'neutral'`).

### `render(g)`

Returns the complete Mermaid diagram as text.

### `renderHTML(g)`

Returns a standalone HTML page that renders the diagram using the Mermaid.js CDN.

### `save(g, filepath)`

Writes the Mermaid text to a `.mmd` file.

### `saveHTML(g, filepath)`

Writes a standalone HTML preview file.

### `saveImage(g, filepath, opts)`

Exports the diagram to PNG, SVG, or PDF via [mermaid-cli](https://github.com/mermaid-js/mermaid-cli) (`mmdc`).

**Options:**
- `format` — `'png'`, `'svg'`, or `'pdf'` (default: inferred from filepath extension)
- `width` — Image width in pixels (default: `800`)
- `height` — Image height in pixels (default: `600`)
- `theme` — Mermaid theme override
- `bg` — Background color (default: `'white'`)
- `mmdc` — Path to the mmdc binary (default: `'mmdc'`)

> Requires `@mermaid-js/mermaid-cli` to be installed: `npm install -g @mermaid-js/mermaid-cli`

### Sequence Diagram Helpers

- `seqMessage(g, from, to, label, lineType)` — Adds a message between actors
- `seqNote(g, position, actor, text)` — Adds a note
- `seqActivate(g, actor)` / `seqDeactivate(g, actor)` — Activation bars
- `seqLoop(g, label, builderFn)` — Loop block
- `seqAlt(g, label, builderFn, elseLabel, elseBuilderFn)` — Alt block

### Pie Chart Helpers

- `pieSlice(g, label, value)` — Adds a slice to a pie chart

### Module Dependency Graph

- `depGraph(moduleNodes, entryPath)` — Creates a Mermaid dependency graph from a module map (as produced by the build system's `ModuleNodes`)

## Build Target: `--graph`

The `--graph` build flag generates a Mermaid module dependency graph for your project:

```bash
magnolia build --entry main.oak --output deps.mmd --graph
magnolia build --entry main.oak --output deps.html --graph
```

- `.mmd` output: Raw Mermaid text
- `.html` output: Standalone HTML preview with embedded Mermaid.js renderer

The graph shows all modules and their import relationships, with the entry point highlighted using a stadium shape.
