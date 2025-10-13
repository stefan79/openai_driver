package output

import "testing"

type parentTest struct {
	ID       int         `table:"id,default"`
	Name     string      `table:"name,default"`
	Children []childTest `table_children:"mychildren"`
}
type childTest struct {
	ID   int    `table:"id,default"`
	Name string `table:"name,default"`
}
type badParentTest struct {
	ID       int         `table:"id,default"`
	Children []childTest `table_children:"kids"`
	More     []childTest `table_children:"morekids"`
}

func TestDefaultColumnsFromType_NegativeCase(t *testing.T) {
	if _, err := DefaultColumnsFromType(badParentTest{}); err == nil {
		t.Fatalf("expected error for multiple table_children tags")
	}
}
func TestDefaultColumnsFromType_PositiveCase(t *testing.T) {
	cols, err := DefaultColumnsFromType(parentTest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cols) != 4 {
		t.Errorf("expected 4 columns, got %d", len(cols))
	}
	if cols[0].Header != "id" {
		t.Errorf("expected first column to be 'id', got %s", cols[0].Header)
	}
	if cols[1].Header != "name" {
		t.Errorf("expected second column to be 'name', got %s", cols[1].Header)
	}
	if cols[2].Header != "mychildren.id" {
		t.Errorf("expected third column to be 'mychildren.id', got %s", cols[2].Header)
	}
	if cols[3].Header != "mychildren.name" {
		t.Errorf("expected fourth column to be 'mychildren.name', got %s", cols[3].Header)
	}
}

func TestDefaultColumnsFromType_MultipleChildren(t *testing.T) {
	if _, err := DefaultColumnsFromType(badParentTest{}); err == nil {
		t.Fatalf("expected error for multiple table_children tags")
	}
}
