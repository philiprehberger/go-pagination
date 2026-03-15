package pagination

import (
	"net/url"
	"testing"
)

func TestParseCursor_Defaults(t *testing.T) {
	q := url.Values{}
	p := ParseCursor(q)

	if p.First != 20 {
		t.Errorf("expected First=20, got %d", p.First)
	}
	if p.Last != 0 {
		t.Errorf("expected Last=0, got %d", p.Last)
	}
	if p.After != "" {
		t.Errorf("expected After empty, got %q", p.After)
	}
	if p.Before != "" {
		t.Errorf("expected Before empty, got %q", p.Before)
	}
}

func TestParseCursor_WithParams(t *testing.T) {
	q := url.Values{}
	q.Set("after", "abc123")
	q.Set("first", "10")

	p := ParseCursor(q)

	if p.After != "abc123" {
		t.Errorf("expected After=abc123, got %q", p.After)
	}
	if p.First != 10 {
		t.Errorf("expected First=10, got %d", p.First)
	}
}

func TestParseCursor_MaxPageSize(t *testing.T) {
	q := url.Values{}
	q.Set("first", "500")

	p := ParseCursor(q)

	if p.First != 100 {
		t.Errorf("expected First capped at 100, got %d", p.First)
	}
}

func TestParseCursor_MaxPageSizeCustom(t *testing.T) {
	q := url.Values{}
	q.Set("first", "30")

	p := ParseCursorWithOptions(q, WithMaxPageSize(25))

	if p.First != 25 {
		t.Errorf("expected First capped at 25, got %d", p.First)
	}
}

func TestEncodeCursor(t *testing.T) {
	encoded := EncodeCursor("42")
	if encoded == "" {
		t.Fatal("expected non-empty encoded cursor")
	}
	if encoded == "42" {
		t.Fatal("expected encoded cursor to differ from input")
	}
}

func TestDecodeCursor(t *testing.T) {
	encoded := EncodeCursor("42")
	decoded, err := DecodeCursor(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != "42" {
		t.Errorf("expected 42, got %q", decoded)
	}
}

func TestDecodeCursor_Invalid(t *testing.T) {
	_, err := DecodeCursor("not-valid-base64!!!")
	if err == nil {
		t.Fatal("expected error for invalid base64")
	}
}
