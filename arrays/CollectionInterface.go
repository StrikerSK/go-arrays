package arrays

type Collections interface {
	IsPresent(searchValue interface{}) (bool, error)
	FindIndex(searchedValue interface{}) (int, error)
	Get(searchedValue int) (interface{}, error)
}
