package entity

// List is an entity list
type List struct {
	entities []Entity
}

// Get gets a stored entity
// Returns nil if not found
func (list *List) Get(index int) *Entity {
	if len(list.entities) <= index {
		return nil
	}
	return &(list.entities[index])
}

// Add adds a new entity to the list
// Returns the index of the newly added entity
func (list *List) Add(entity *Entity) int {
	list.entities = append(list.entities, *entity)
	return len(list.entities) - 1
}

// Length gets number of entities
func (list List) Length() int {
	return len(list.entities)
}

// FindByKeyValue finds an entity by a key/value pair
// Note: Returns the first found entity, so non-unique pairing
// may be problematic
func (list *List) FindByKeyValue(key string, value string) *Entity {
	var v string

	// search the entities for one using modnum
	for i := 0; i < len(list.entities); i++ {
		v = list.Get(i).ValueForKey(key)
		if v == value {
			return list.Get(i)
		}
	}

	return nil
}

// FindForModel finds an entity from a model id
func (list *List) FindForModel(modelNumber int) *Entity {
	var s string
	var name string

	// search the entities for one using modnum
	for i := 0; i < len(list.entities); i++ {
		s = list.Get(i).ValueForKey("model")
		if s != name {
			return list.Get(i)
		}
	}

	return list.Get(0)
}

// NewEntityList
// Create a new entity list from an existing slice of entities
func NewEntityList(entities []Entity) List {
	return List{
		entities: entities,
	}
}
