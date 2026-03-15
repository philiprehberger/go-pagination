package pagination

import (
	"net/url"
	"strconv"
)

const (
	defaultOffsetPage = 1
	defaultOffsetSize = 20
	maxOffsetSize     = 100
)

// OffsetParams holds the parsed parameters for offset-based pagination.
type OffsetParams struct {
	// Page is the current page number (1-based).
	Page int
	// Size is the number of items per page.
	Size int
	// Offset is the computed offset: (Page - 1) * Size.
	Offset int
}

// OffsetOption configures offset pagination parsing.
type OffsetOption func(*offsetConfig)

type offsetConfig struct {
	defaultSize int
	maxSize     int
}

// WithDefaultSize sets the default page size for offset pagination.
func WithDefaultSize(n int) OffsetOption {
	return func(c *offsetConfig) {
		if n > 0 {
			c.defaultSize = n
		}
	}
}

// WithMaxSize sets the maximum page size for offset pagination.
func WithMaxSize(n int) OffsetOption {
	return func(c *offsetConfig) {
		if n > 0 {
			c.maxSize = n
		}
	}
}

// ParseOffset parses offset pagination parameters from URL query values.
// It reads "page" and "size" (or "per_page") query parameters.
// Default page is 1, default size is 20, maximum size is 100.
func ParseOffset(query url.Values) OffsetParams {
	return ParseOffsetWithOptions(query)
}

// ParseOffsetWithOptions parses offset pagination parameters with configurable options.
func ParseOffsetWithOptions(query url.Values, opts ...OffsetOption) OffsetParams {
	cfg := &offsetConfig{
		defaultSize: defaultOffsetSize,
		maxSize:     maxOffsetSize,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	params := OffsetParams{
		Page: defaultOffsetPage,
		Size: cfg.defaultSize,
	}

	if v := query.Get("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			params.Page = n
		}
	}

	if v := query.Get("size"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			params.Size = n
		}
	} else if v := query.Get("per_page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			params.Size = n
		}
	}

	if params.Size > cfg.maxSize {
		params.Size = cfg.maxSize
	}

	params.Offset = (params.Page - 1) * params.Size

	return params
}

// LimitOffset returns the SQL-friendly limit and offset values from the given OffsetParams.
func LimitOffset(params OffsetParams) (limit, offset int) {
	return params.Size, params.Offset
}
