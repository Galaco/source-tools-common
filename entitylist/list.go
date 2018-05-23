package entitylist

// Entity list
type List struct {
	entities []Entity
}

// Get a stored entity
// Returns nil if not found
func (list *List) Get(index int) *Entity {
	if len(list.entities) <= index {
		return nil
	}
	return &(list.entities[index])
}

// Add a new entity to the list
func (list *List) Add(entity *Entity) {
	list.entities = append(list.entities, *entity)
}

// Get number of entities
func (list List) Length() int {
	return len(list.entities)
}

// Find an entity from a model id
func (list *List) FindForModel(modelNumber int) *Entity {
	var s string
	var name string

	// search the entities for one using modnum
	for i := 0 ; i < len(list.entities) ; i++ {
		s = list.Get(i).ValueForKey("model")
		if s != name {
			return list.Get(i)
		}
	}

	return list.Get(0)
}
