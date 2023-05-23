package generatorexec

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type RegistryType uint8

const (
	RegistryTypeNpm RegistryType = iota + 1
	RegistryTypeMaven
	RegistryTypePypi
)

func (x RegistryType) String() string {
	switch x {
	default:
		return strconv.Itoa(int(x))
	case RegistryTypeNpm:
		return "NPM"
	case RegistryTypeMaven:
		return "MAVEN"
	case RegistryTypePypi:
		return "PYPI"
	}
}

func (x RegistryType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", x.String())), nil
}

func (x *RegistryType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "NPM":
		value := RegistryTypeNpm
		*x = value
	case "MAVEN":
		value := RegistryTypeMaven
		*x = value
	case "PYPI":
		value := RegistryTypePypi
		*x = value
	}
	return nil
}
