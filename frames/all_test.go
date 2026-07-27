package frames

import "testing"

func TestList(t *testing.T) {
	list := List()
	if len(list) != len(All) {
		t.Errorf("List() returned %d items, expected %d", len(list), len(All))
	}
	// Verify defensive copy
	if len(list) > 0 {
		orig := list[0]
		list[0] = nil
		if All[0] == nil {
			t.Errorf("List() did not return a defensive copy")
		}
		list[0] = orig // restore
	}
}

func TestByName(t *testing.T) {
	if len(ByName) != len(All) {
		t.Errorf("ByName has %d items, expected %d", len(ByName), len(All))
	}
	for _, d := range All {
		if ByName[d.Name] != d {
			t.Errorf("ByName[%s] = %v, expected %v", d.Name, ByName[d.Name], d)
		}
	}
}
