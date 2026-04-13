# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\mermaid.oak`

- `fs` · `import(...)`
- `path` · `import(...)`
- `FlowchartLR` · `:lr`
- `FlowchartTD` · `:td`
- `FlowchartRL` · `:rl`
- `FlowchartBT` · `:bt`
- `SequenceDiagram` · `:sequence`
- `ClassDiagram` · `:classDiagram`
- `StateDiagram` · `:stateDiagram`
- `PieChart` · `:pie`
- `Gantt` · `:gantt`
- `ERDiagram` · `:erDiagram`
- `GitGraph` · `:gitGraph`
- `ShapeRect` · `:rect`
- `ShapeRound` · `:round`
- `ShapeStadium` · `:stadium`
- `ShapeCircle` · `:circle`
- `ShapeRhombus` · `:rhombus`
- `ShapeHexagon` · `:hexagon`
- `ShapeTrapezoid` · `:trapezoid`
- `ShapeDefault` · `:default`
- `EdgeArrow` · `:arrow`
- `EdgeDotted` · `:dotted`
- `EdgeThick` · `:thick`
- `EdgeOpen` · `:open`
### `graph(direction)`

> returns `:object`

### `node(g, id, label, shape)`

### `edge(g, from, to, label, style)`

### `subgraph(g, id, label, builderFn)`

### `raw(g, line)`

### `setTitle(g, t)`

### `setTheme(g, t)`

### `escapeLabel(s)`

### `renderShape(id, label, shape)`

### `renderEdge(e)`

### `renderSubgraph(sg, indent)`

### `render(g)`

### `renderFlowchart(g)`

### `seqMessage(g, from, to, label, lineType)`

### `seqNote(g, position, actor, text)`

### `seqActivate(g, actor)`

### `seqDeactivate(g, actor)`

### `seqLoop(g, label, builderFn)`

### `seqAlt(g, label, builderFn, elseLabel, elseBuilderFn)`

### `pieSlice(g, label, value)`

### `renderHTML(g)`

> returns `:string`

### `save(g, filepath)`

### `saveHTML(g, filepath)`

### `saveImage(g, filepath, opts)`

### `inferFormat(filepath)`

> returns `:string`

### `depGraph(moduleNodes, entryPath)`

### `extractImportEdges(g, fromId, node, sanitize)`

