package string

func AreValuesInSliceStrict(targetValues []string, providedValues []string, strict bool) bool {
	arePresent := false

	for index := range providedValues {
		valueInSlice := IsValueInSlice(targetValues, providedValues[index])

		if !strict && valueInSlice {
			arePresent = true
			break
		} else if !strict && !valueInSlice {
			continue
		} else if strict && valueInSlice {
			arePresent = true
			continue
		} else if strict && !valueInSlice {
			arePresent = false
			break
		}
	}

	return arePresent
}

func IsValueInSlice(stringSlice []string, searchedValue string) bool {
	for index := range stringSlice {
		if stringSlice[index] == searchedValue {
			return true
		}
	}

	return false
}
