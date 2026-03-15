package pagination

// PageInfo contains pagination metadata for a page of results.
type PageInfo struct {
	// HasNextPage indicates whether more items exist after this page.
	HasNextPage bool
	// HasPreviousPage indicates whether more items exist before this page.
	HasPreviousPage bool
	// StartCursor is the cursor of the first item in the page.
	StartCursor string
	// EndCursor is the cursor of the last item in the page.
	EndCursor string
	// Total is the total number of items across all pages.
	Total int
}

// Page represents a paginated set of items with associated metadata.
type Page[T any] struct {
	// Items contains the page's data.
	Items []T
	// PageInfo contains the pagination metadata.
	PageInfo PageInfo
}

// PageOption configures a Page.
type PageOption[T any] func(*Page[T])

// WithTotal sets the total item count on the page.
func WithTotal[T any](n int) PageOption[T] {
	return func(p *Page[T]) {
		p.PageInfo.Total = n
	}
}

// WithHasNext sets whether there is a next page.
func WithHasNext[T any](v bool) PageOption[T] {
	return func(p *Page[T]) {
		p.PageInfo.HasNextPage = v
	}
}

// WithHasPrevious sets whether there is a previous page.
func WithHasPrevious[T any](v bool) PageOption[T] {
	return func(p *Page[T]) {
		p.PageInfo.HasPreviousPage = v
	}
}

// WithStartCursor sets the start cursor on the page.
func WithStartCursor[T any](c string) PageOption[T] {
	return func(p *Page[T]) {
		p.PageInfo.StartCursor = c
	}
}

// WithEndCursor sets the end cursor on the page.
func WithEndCursor[T any](c string) PageOption[T] {
	return func(p *Page[T]) {
		p.PageInfo.EndCursor = c
	}
}

// NewPage creates a new Page with the given items and optional configuration.
func NewPage[T any](items []T, opts ...PageOption[T]) Page[T] {
	p := Page[T]{
		Items: items,
	}
	for _, opt := range opts {
		opt(&p)
	}
	return p
}
