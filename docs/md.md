# Markdown Library (md)

## Overview

`libmd` implements a Markdown parser and renderer for Oak, supporting standard Markdown syntax for converting text to HTML or other formats.

## Import

```oak
md := import('md')
// Typically used for parsing and rendering Markdown documents
```

## Features

- **Inline formatting**: Bold, italic, code, links, images
- **Block elements**: Headings, paragraphs, lists, code blocks, blockquotes
- **Lists**: Ordered and unordered lists
- **Links and images**: Standard Markdown syntax
- **Code blocks**: Fenced and indented code blocks
- **URL handling**: Automatic percent-decoding for links

## Key Components

### Reader Generic Iterator

The library uses a generic `Reader` interface that works with both strings and lists, providing methods for:
- `peek()` - Look at current character
- `next()` - Advance to next character
- `readUntil(char)` - Read until delimiter
- `readUntilPrefix(prefix)` - Read until matching prefix
- `readUntilMatchingDelim(delim)` - Handle nested brackets/parens

### Text Tokenization

The `tokenizeText()` function handles inline Markdown elements:
- Bold: `**text**` or `__text__`
- Italic: `*text*` or `_text_`
- Code: `` `code` ``
- Links: `[text](url)`
- Images: `![alt](url)`

## Usage Example

```oak
md := import('md')

markdown := '
# Hello World

This is **bold** and this is *italic*.

## Features

- Easy to use
- Fast parsing
- Standard compliant
'

// Parse and render would be called here
// (specific API methods depend on implementation)
```

## Implementation Notes

- Uses a character-by-character parser
- Handles nested delimiters for parsing balanced brackets
- UTF-8 aware word character detection
- Percent-decodes URLs in links
- Maintains position tracking for error reporting

## Limitations

- Specific rendering API may vary
- May not support all Markdown extensions (tables, task lists, etc.)
- Parser is character-based (not line-based) which affects some constructs

## See Also

- `http.percentDecode()` - For URL decoding in links
- `str` library - For string manipulation
- `fmt` library - For formatted output
