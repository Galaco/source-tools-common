package entity

import "github.com/galaco/vmf"

// Build an entity list
// Constructs from the root node of Vmf entity data
func FromVmfNodeTree(entityNodes vmf.Node) List {
	numEntities := len(*entityNodes.GetAllValues())

	entityList := List{
		entities: make([]Entity, numEntities),
	}

	for i := 0; i < numEntities; i++ {
		mapEnt := entityList.Get(i)
		var e *EPair
		eNode := (*entityNodes.GetAllValues())[i].(vmf.Node)
		for _,kv := range *eNode.GetAllValues() {
			n := kv.(vmf.Node)
			e = parseEPair(&n)
			e.Next = mapEnt.EPairs
			mapEnt.EPairs = e
		}
	}

	return entityList
}


func parseEPair(node *vmf.Node) *EPair {
	if len(*node.GetAllValues()) > 1 {
		return nil
	}

	switch (*node.GetAllValues())[0].(type) {
	case string:
		return &EPair{
			Next: nil,
			Key: *(*node).GetKey(),
			Value: (*node.GetAllValues())[0].(string),
		}
	default:
		return nil
	}
}
