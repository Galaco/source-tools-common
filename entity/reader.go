package entity

import "github.com/galaco/vmf"

// FromVmfNode parses a single Vmf node that represents an
// entity
func FromVmfNode(entityNode *vmf.Node) Entity {
	var e *EPair
	mapEnt := Entity{}
	for _, kv := range *entityNode.GetAllValues() {
		n := kv.(vmf.Node)
		e = parseEPair(&n)
		// ignore array children
		if e == nil {
			continue
		}
		e.Next = mapEnt.EPairs
		mapEnt.EPairs = e
	}

	return mapEnt
}

// FromVmfNodeTree
// Build an entity list
// Constructs from the root node of Vmf entity data
func FromVmfNodeTree(entityNodes vmf.Node) List {
	numEntities := len(*entityNodes.GetAllValues())

	entities := make([]Entity, numEntities)

	for i := 0; i < numEntities; i++ {
		eNode := (*entityNodes.GetAllValues())[i].(vmf.Node)
		entities[i] = FromVmfNode(&eNode)
	}
	entityList := List{
		entities: entities,
	}

	return entityList
}

// parseEPair scoped Function for reading entity keyvalue pair
func parseEPair(node *vmf.Node) *EPair {
	if len(*node.GetAllValues()) > 1 {
		return nil
	}

	switch (*node.GetAllValues())[0].(type) {
	case string:
		return &EPair{
			Next:  nil,
			Key:   *(*node).GetKey(),
			Value: (*node.GetAllValues())[0].(string),
		}
	default:
		return nil
	}
}
