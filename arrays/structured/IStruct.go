package structured

// IStruct - Custom interface used in following StructArray implementation
type IStruct interface {
	// CompareValues - Validation value API that needs to have implemented which fields should be parameter value be checked against
	CompareValues(input interface{}) (bool, error)

	// CompareTypes - Validation type API needs to have implemented checker for types of parameter
	CompareTypes(input interface{}) error
}
