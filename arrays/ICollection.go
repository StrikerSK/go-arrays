package arrays

// ICollections represents interface for slices and arrays
type ICollections interface {
	IsPresent(searchValue interface{}) (bool, error)
	FindIndex(searchedValue interface{}) (int, error)
	Get(searchedValue int) (interface{}, error)
	Add(newValue interface{}) error
}
