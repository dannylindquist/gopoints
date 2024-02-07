package store

type Store struct {
	container map[string]map[int]interface{}
}

func NewStore() *Store {
	s := &Store{
		container: make(map[string]map[int]interface{}),
	}
	return s
}

func (s *Store) GetItems(dataType string) map[int]interface{} {
	if s.container[dataType] == nil {
		s.container[dataType] = make(map[int]interface{})
	}
	return s.container[dataType]
}

func (s *Store) SetItems(dataType string, id int, val interface{}) {
	s.GetItems(dataType)[id] = val
}
