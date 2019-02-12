package entity

import (
	"github.com/go-gl/mathgl/mgl32"
	"reflect"
	"testing"
)

var entities = []Entity{
	{
		Origin: mgl32.Vec3{1, 0, 0},
	},
	{
		Origin: mgl32.Vec3{0, 1, 0},
		EPairs: &EPair{
			Next:  nil,
			Key:   "model",
			Value: "8",
		},
	},
	{
		Origin: mgl32.Vec3{0, 0, 1},
	},
}

func TestNewEntityList(t *testing.T) {
	expected := List{}
	actual := NewEntityList(entities)

	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Unexpected type. Expected: %s, actual: %s", reflect.TypeOf(expected), reflect.TypeOf(actual))
	}
}

func TestList_Get(t *testing.T) {
	list := NewEntityList(entities)

	ent := list.Get(1)
	if ent.Origin.Y() != 1 {
		t.Errorf("Entity fetched does not match expected")
	}
}

func TestList_Add(t *testing.T) {
	list := NewEntityList(entities)

	newEnt := Entity{
		Origin: mgl32.Vec3{2, 2, 2},
	}

	index := list.Add(&newEnt)

	ent := list.Get(index)
	if ent.Origin.Y() != 2 {
		t.Errorf("Entity fetched does not match the added entity")
	}
}

func TestList_Length(t *testing.T) {
	expected := 3
	list := NewEntityList(entities)

	if list.Length() != expected {
		t.Errorf("List has unexpected length, expected: %d, actual: %d", expected, list.Length())
	}

}

func TestList_FindForModel(t *testing.T) {
	list := NewEntityList(entities)

	ent := list.FindForModel(8)

	if ent == nil {
		t.Errorf("Failed to find entity with model number: 8")
	}

}

func TestList_FindByKeyValue(t *testing.T) {
	list := NewEntityList(entities)

	ent := list.FindByKeyValue("model", "8")

	if ent == nil {
		t.Errorf("Failed to find entity with keyvalue: model: 8")
	}

}
