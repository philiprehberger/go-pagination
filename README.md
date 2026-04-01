# go-pagination

[![CI](https://github.com/philiprehberger/go-pagination/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-pagination/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-pagination.svg)](https://pkg.go.dev/github.com/philiprehberger/go-pagination)
[![Last updated](https://img.shields.io/github/last-commit/philiprehberger/go-pagination)](https://github.com/philiprehberger/go-pagination/commits/main)

Cursor and offset pagination helpers for Go. Generic, zero dependencies

## Installation

```bash
go get github.com/philiprehberger/go-pagination
```

## Usage

### Cursor Pagination

```go
import "github.com/philiprehberger/go-pagination"

// Parse cursor params from an HTTP request
params := pagination.ParseCursor(r.URL.Query())

// Encode/decode opaque cursors
cursor := pagination.EncodeCursor("42")
id, err := pagination.DecodeCursor(cursor)

// Build a page response with cursor info
page := pagination.NewPage(items,
    pagination.WithTotal[Item](total),
    pagination.WithHasNext[Item](true),
    pagination.WithStartCursor[Item](pagination.EncodeCursor(items[0].ID)),
    pagination.WithEndCursor[Item](pagination.EncodeCursor(items[len(items)-1].ID)),
)
```

### Offset Pagination

```go
import "github.com/philiprehberger/go-pagination"

// Parse offset params from an HTTP request
params := pagination.ParseOffset(r.URL.Query())

// Use limit/offset for SQL queries
limit, offset := pagination.LimitOffset(params)
// SELECT * FROM items LIMIT $1 OFFSET $2
```

### Generic Page Response

```go
import "github.com/philiprehberger/go-pagination"

items := []User{...}
page := pagination.NewPage(items,
    pagination.WithTotal[User](250),
    pagination.WithHasNext[User](true),
    pagination.WithHasPrevious[User](true),
)
// page.Items, page.PageInfo.Total, page.PageInfo.HasNextPage, etc.
```

### Edge Cases

```go
import pagination "github.com/philiprehberger/go-pagination"

// Empty results
page := pagination.NewPage([]string{}, pagination.WithTotal[string](0))
// Page{Items: [], Total: 0, HasNext: false}

// Single item page
page = pagination.NewPage([]string{"only"}, pagination.WithTotal[string](1))
```

### Combining Cursor and Offset

```go
import (
    "net/url"
    pagination "github.com/philiprehberger/go-pagination"
)

query, _ := url.ParseQuery("cursor=abc123&limit=10")
params := pagination.ParseCursorWithOptions(query,
    pagination.WithDefaultPageSize(20),
    pagination.WithMaxPageSize(100),
)
// params.After = "abc123", params.PageSize = 10
```

## API

| Type / Function | Description |
|---|---|
| `CursorParams` | Parsed cursor pagination parameters |
| `ParseCursor(query)` | Parse cursor params from URL query |
| `ParseCursorWithOptions(query, ...CursorOption)` | Parse with custom defaults/limits |
| `WithDefaultPageSize(n)` | Set default cursor page size |
| `WithMaxPageSize(n)` | Set max cursor page size |
| `EncodeCursor(id)` | Encode an ID to an opaque cursor |
| `DecodeCursor(cursor)` | Decode an opaque cursor to an ID |
| `OffsetParams` | Parsed offset pagination parameters |
| `ParseOffset(query)` | Parse offset params from URL query |
| `ParseOffsetWithOptions(query, ...OffsetOption)` | Parse with custom defaults/limits |
| `WithDefaultSize(n)` | Set default offset page size |
| `WithMaxSize(n)` | Set max offset page size |
| `LimitOffset(params)` | Get SQL-friendly limit and offset |
| `Page[T]` | Generic paginated result set |
| `PageInfo` | Pagination metadata |
| `NewPage[T](items, ...PageOption[T])` | Create a page with options |
| `WithTotal[T](n)` | Set total count |
| `WithHasNext[T](v)` | Set has-next-page flag |
| `WithHasPrevious[T](v)` | Set has-previous-page flag |
| `WithStartCursor[T](c)` | Set start cursor |
| `WithEndCursor[T](c)` | Set end cursor |

## Development

```bash
go test ./...
go vet ./...
```

## Support

If you find this project useful:

⭐ [Star the repo](https://github.com/philiprehberger/go-pagination)

🐛 [Report issues](https://github.com/philiprehberger/go-pagination/issues?q=is%3Aissue+is%3Aopen+label%3Abug)

💡 [Suggest features](https://github.com/philiprehberger/go-pagination/issues?q=is%3Aissue+is%3Aopen+label%3Aenhancement)

❤️ [Sponsor development](https://github.com/sponsors/philiprehberger)

🌐 [All Open Source Projects](https://philiprehberger.com/open-source-packages)

💻 [GitHub Profile](https://github.com/philiprehberger)

🔗 [LinkedIn Profile](https://www.linkedin.com/in/philiprehberger)

## License

[MIT](LICENSE)
