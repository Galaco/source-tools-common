package texdatastringtable

import (
	"reflect"
	"testing"
)

var texDataString = "foo\x00bar\x00baz\x00bat"
var texDataData = []int32{
	0,
	4,
	8,
	12,
}

func TestNewTable(t *testing.T) {
	expected := TexDataStringTable{}
	actual := *NewTable(texDataString, texDataData)
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Unexpected type. Expected: %s, actual: %s", reflect.TypeOf(expected), reflect.TypeOf(actual))
	}
}

func TestTexDataStringTable_AddOrFindString(t *testing.T) {
	table := NewTable(texDataString, texDataData)
	if r, _ := table.AddOrFindString("baz"); r != 2 {
		t.Errorf("Failed to find string. Expected: 2, actual: %d", r)
	}

	if r, _ := table.AddOrFindString("newstring"); r != 4 {
		t.Errorf("Failed to add new string. Expected: 4, actual: %d", r)
	}
}

func TestTexDataStringTable_GetString(t *testing.T) {
	table := NewTable(texDataString, texDataData)
	if r, _ := table.GetString(0); r != "foo" {
		t.Errorf("Failed to get string. Expected: foo, actual: %s", r)
	}

	if r, _ := table.GetString(3); r != "bat" {
		t.Errorf("Failed to get string. Expected: bat, actual: %s", r)
	}

	if r, _ := table.GetString(999); r != "" {
		t.Errorf("Failed to get string. Expected: '', actual: %s", r)
	}
}
