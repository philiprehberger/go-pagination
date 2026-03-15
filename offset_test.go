package pagination

import (
	"net/url"
	"testing"
)

func TestParseOffset_Defaults(t *testing.T) {
	q := url.Values{}
	p := ParseOffset(q)

	if p.Page != 1 {
		t.Errorf("expected Page=1, got %d", p.Page)
	}
	if p.Size != 20 {
		t.Errorf("expected Size=20, got %d", p.Size)
	}
	if p.Offset != 0 {
		t.Errorf("expected Offset=0, got %d", p.Offset)
	}
}

func TestParseOffset_WithParams(t *testing.T) {
	q := url.Values{}
	q.Set("page", "3")
	q.Set("size", "15")

	p := ParseOffset(q)

	if p.Page != 3 {
		t.Errorf("expected Page=3, got %d", p.Page)
	}
	if p.Size != 15 {
		t.Errorf("expected Size=15, got %d", p.Size)
	}
}

func TestParseOffset_PerPage(t *testing.T) {
	q := url.Values{}
	q.Set("per_page", "25")

	p := ParseOffset(q)

	if p.Size != 25 {
		t.Errorf("expected Size=25, got %d", p.Size)
	}
}

func TestParseOffset_MaxSize(t *testing.T) {
	q := url.Values{}
	q.Set("size", "500")

	p := ParseOffset(q)

	if p.Size != 100 {
		t.Errorf("expected Size capped at 100, got %d", p.Size)
	}
}

func TestParseOffset_ComputesOffset(t *testing.T) {
	q := url.Values{}
	q.Set("page", "3")
	q.Set("size", "10")

	p := ParseOffset(q)

	if p.Offset != 20 {
		t.Errorf("expected Offset=20, got %d", p.Offset)
	}
}

func TestLimitOffset(t *testing.T) {
	q := url.Values{}
	q.Set("page", "5")
	q.Set("size", "10")

	p := ParseOffset(q)
	limit, offset := LimitOffset(p)

	if limit != 10 {
		t.Errorf("expected limit=10, got %d", limit)
	}
	if offset != 40 {
		t.Errorf("expected offset=40, got %d", offset)
	}
}
