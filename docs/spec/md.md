# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\md.oak`

### `cs Reader(s)`

> returns `:object`

### `uword?(c)`

> returns `:bool`

### `tokenizeText(line)`

### `unifyTextNodes(nodes, joiner)`

### `parseText(tokens)`

### `uListItemLine?(line)`

> returns `:bool`

### `oListItemLine?(line)`

> returns `:bool`

### `listItemLine?(line)`

> returns `:bool`

### `tableLine?(line)`

> returns `:bool`

### `tableSepLine?(line)`

### `trimUListGetLevel(reader)`

### `trimOListGetLevel(reader)`

### `lineNodeType(line)`

> returns `?`

### `parse(text)`

### `parseDoc(lineReader)`

### `parseHeader(nodeType, lineReader)`

> returns `:object`

### `parseBlockQuote(lineReader)`

> returns `:object`

### `parseCodeBlock(lineReader)`

> returns `:object`

### `parseRawHTML(lineReader)`

> returns `:object`

### `parseList(lineReader, listType)`

> returns `:object`

### `parseTableRow(line)`

### `parseTableAlign(sepLine)`

### `parseTable(lineReader)`

> returns `:object`

### `parseParagraph(lineReader)`

> returns `:object`

### `compile(nodes)`

### `wrap(tag, node)`

> returns `:string`

### `sanitizeAttr(attr)`

### `sanitizeURL(url)`

> returns `:string`

### `compileNode(node)`

### `transform(text)`

