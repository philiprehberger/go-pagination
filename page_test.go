package pagination

import "testing"

func TestNewPage(t *testing.T) {
	items := []string{"a", "b", "c"}
	p := NewPage(items)

	if len(p.Items) != 3 {
		t.Errorf("expected 3 items, got %d", len(p.Items))
	}
	if p.Items[0] != "a" {
		t.Errorf("expected first item 'a', got %q", p.Items[0])
	}
}

func TestNewPage_WithOptions(t *testing.T) {
	items := []int{1, 2, 3}
	p := NewPage(items,
		WithTotal[int](100),
		WithHasNext[int](true),
		WithHasPrevious[int](false),
		WithStartCursor[int]("cursor_start"),
		WithEndCursor[int]("cursor_end"),
	)

	if p.PageInfo.Total != 100 {
		t.Errorf("expected Total=100, got %d", p.PageInfo.Total)
	}
	if !p.PageInfo.HasNextPage {
		t.Error("expected HasNextPage=true")
	}
	if p.PageInfo.HasPreviousPage {
		t.Error("expected HasPreviousPage=false")
	}
	if p.PageInfo.StartCursor != "cursor_start" {
		t.Errorf("expected StartCursor=cursor_start, got %q", p.PageInfo.StartCursor)
	}
	if p.PageInfo.EndCursor != "cursor_end" {
		t.Errorf("expected EndCursor=cursor_end, got %q", p.PageInfo.EndCursor)
	}
}

func TestNewPage_Empty(t *testing.T) {
	p := NewPage([]string{})

	if len(p.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(p.Items))
	}
	if p.PageInfo.Total != 0 {
		t.Errorf("expected Total=0, got %d", p.PageInfo.Total)
	}
	if p.PageInfo.HasNextPage {
		t.Error("expected HasNextPage=false")
	}
}
