// Package pagination provides cursor and offset pagination helpers.
package pagination

import (
	"encoding/base64"
	"net/url"
	"strconv"
)

const (
	defaultCursorPageSize = 20
	maxCursorPageSize     = 100
)

// CursorParams holds the parsed parameters for cursor-based pagination.
type CursorParams struct {
	// After is the cursor pointing to the item after which results should start.
	After string
	// Before is the cursor pointing to the item before which results should end.
	Before string
	// First is the number of items to return from the start.
	First int
	// Last is the number of items to return from the end.
	Last int
}

// CursorOption configures cursor pagination parsing.
type CursorOption func(*cursorConfig)

type cursorConfig struct {
	defaultPageSize int
	maxPageSize     int
}

// WithDefaultPageSize sets the default page size for cursor pagination.
func WithDefaultPageSize(n int) CursorOption {
	return func(c *cursorConfig) {
		if n > 0 {
			c.defaultPageSize = n
		}
	}
}

// WithMaxPageSize sets the maximum page size for cursor pagination.
func WithMaxPageSize(n int) CursorOption {
	return func(c *cursorConfig) {
		if n > 0 {
			c.maxPageSize = n
		}
	}
}

// ParseCursor parses cursor pagination parameters from URL query values.
// It reads "after", "before", "first", and "last" query parameters.
// Default first is 20, maximum is 100.
func ParseCursor(query url.Values) CursorParams {
	return ParseCursorWithOptions(query)
}

// ParseCursorWithOptions parses cursor pagination parameters with configurable options.
func ParseCursorWithOptions(query url.Values, opts ...CursorOption) CursorParams {
	cfg := &cursorConfig{
		defaultPageSize: defaultCursorPageSize,
		maxPageSize:     maxCursorPageSize,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	params := CursorParams{
		After:  query.Get("after"),
		Before: query.Get("before"),
		First:  cfg.defaultPageSize,
	}

	if v := query.Get("first"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			params.First = n
		}
	}
	if params.First > cfg.maxPageSize {
		params.First = cfg.maxPageSize
	}

	if v := query.Get("last"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			params.Last = n
		}
	}
	if params.Last > cfg.maxPageSize {
		params.Last = cfg.maxPageSize
	}

	return params
}

// EncodeCursor encodes an ID into an opaque base64 cursor string.
func EncodeCursor(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

// DecodeCursor decodes a base64 cursor string back into the original ID.
// Returns an error if the cursor is not valid base64.
func DecodeCursor(cursor string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
