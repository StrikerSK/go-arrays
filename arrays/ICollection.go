package arrays

// ICollections represents interface for slices and arrays
type ICollections interface {
	IsPresent(searchValue interface{}) (bool, error)
	FindIndex(searchedValue interface{}) (int, error)
	GetByIndex(searchedValue int) (interface{}, error)
	Add(newValue interface{}) error
	RemoveByIndex(index int) error
}
