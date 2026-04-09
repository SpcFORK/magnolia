# data-xml — XML Parser & Serializer

`import('data-xml')` provides XML parsing and serialization with support for elements, attributes, text, CDATA sections, comments, and self-closing tags.

## Quick Start

```oak
xml := import('data-xml')

// Parse XML
nodes := xml.parse('<root attr="val"><item>Hello</item><item>World</item></root>')
// nodes.0.tag => 'root'
// nodes.0.attrs.attr => 'val'

// Build and serialize XML
doc := xml.element('root', {version: '1.0'}, [
    xml.element('item', {}, [xml.text('Hello')])
    xml.element('item', {}, [xml.text('World')])
    xml.comment('end of list')
])
text := xml.serialize(doc)

// Query elements
first := xml.querySelector(nodes, 'item')
all := xml.querySelectorAll(nodes, 'item')
content := xml.textContent(nodes.0)
```

## Node Constructors

### `element(tag, attrs, children)`

Creates an element node: `{type: :element, tag, attrs, children}`.

### `text(content)`

Creates a text node: `{type: :text, content}`.

### `cdata(content)`

Creates a CDATA section node: `{type: :cdata, content}`.

### `comment(content)`

Creates a comment node: `{type: :comment, content}`.

## Serialization & Parsing

### `serialize(node)`

Converts an XML AST node (or list of nodes) to an XML string. Handles attribute escaping and entity encoding (`&amp;`, `&lt;`, `&gt;`, `&quot;`, `&apos;`).

### `parse(text)`

Parses XML text into a list of AST nodes. Supports:
- Elements with attributes
- Self-closing tags (`<tag/>`)
- CDATA sections (`<![CDATA[...]]>`)
- Comments (`<!--...-->`)
- Processing instructions (`<?...?>`)
- Text nodes with entity unescaping

## Query Functions

### `querySelector(nodes, sel)`

Finds the first element matching a tag name (depth-first search).

### `querySelectorAll(nodes, sel)`

Finds all elements matching a tag name.

### `textContent(node)`

Extracts all text content from a node and its children recursively.

## Notes

- Entity references `&amp;`, `&lt;`, `&gt;`, `&quot;`, `&apos;` are handled automatically.
- Attribute values support both `"` and `'` delimiters.
- No namespace or DTD support.
