package output

import "testing"

type (
	testStruct struct {
		ID       int         `table:"id,default"`
		Name     string      `table:"name,default"`
		Children []testChild `table_children:"children"`
		Leaf     testLeaf    `table:"leaf,default"`
	}

	testLeaf struct {
		LeafID   int    `table:"leafid,default"`
		LeafName string `table:"leafname,default"`
	}

	testChild struct {
		ChildID       int              `table:"childid,default"`
		ChildName     string           `table:"childname,default"`
		GrandChildren []testGrandChild `table_children:"grandchildren"`
	}

	testGrandChild struct {
		GrandChildID   int    `table:"grandchildid,default"`
		GrandChildName string `table:"grandchildname,default"`
	}
)

var (
	testData = testStruct{
		ID:   1,
		Name: "test",
		Children: []testChild{
			{
				ChildID:   1,
				ChildName: "child1",
				GrandChildren: []testGrandChild{
					{
						GrandChildID:   1,
						GrandChildName: "grandchild11",
					},
					{
						GrandChildID:   2,
						GrandChildName: "grandchild12",
					},
				},
			},
			{
				ChildID:   2,
				ChildName: "child2",
				GrandChildren: []testGrandChild{
					{
						GrandChildID:   1,
						GrandChildName: "grandchild21",
					},
					{
						GrandChildID:   2,
						GrandChildName: "grandchild22",
					},
				},
			},
		},
		Leaf: testLeaf{
			LeafID:   1,
			LeafName: "leaf",
		},
	}
)

func TestFlatten(t *testing.T) {
	rows, err := Flatten(testData, DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 4 {
		t.Fatalf("expected 4 rows, got %d", len(rows))
	}
	if rows[0]["id"] != int64(1) {
		t.Errorf("expected id 1, got %d", rows[0]["id"])
	}
	if rows[0]["name"] != "test" {
		t.Errorf("expected name 'test', got '%s'", rows[0]["name"])
	}
	if rows[0]["children.childid"] != int64(1) {
		t.Errorf("expected children.childid 1, got %d", rows[0]["children.childid"])
	}
	if rows[0]["children.childname"] != "child1" {
		t.Errorf("expected children.childname 'child1', got '%s'", rows[0]["children.childname"])
	}
	if rows[0]["leaf.leafid"] != int64(1) {
		t.Errorf("expected leaf.leafid 1, got %d", rows[0]["leaf.leafid"])
	}
	if rows[0]["leaf.leafname"] != "leaf" {
		t.Errorf("expected leaf.leafname 'leaf', got '%s'", rows[0]["leaf.leafname"])
	}
}
