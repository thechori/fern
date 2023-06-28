// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type SingleUnionTypeProperties struct {
	PropertiesType         string
	SamePropertiesAsObject *DeclaredTypeName
	SingleProperty         *SingleUnionTypeProperty
	NoProperties           any
}

func NewSingleUnionTypePropertiesFromSamePropertiesAsObject(value *DeclaredTypeName) *SingleUnionTypeProperties {
	return &SingleUnionTypeProperties{PropertiesType: "samePropertiesAsObject", SamePropertiesAsObject: value}
}

func NewSingleUnionTypePropertiesFromSingleProperty(value *SingleUnionTypeProperty) *SingleUnionTypeProperties {
	return &SingleUnionTypeProperties{PropertiesType: "singleProperty", SingleProperty: value}
}

func NewSingleUnionTypePropertiesFromNoProperties(value any) *SingleUnionTypeProperties {
	return &SingleUnionTypeProperties{PropertiesType: "noProperties", NoProperties: value}
}

func (s *SingleUnionTypeProperties) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		PropertiesType string `json:"_type,omitempty"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.PropertiesType = unmarshaler.PropertiesType
	switch unmarshaler.PropertiesType {
	case "samePropertiesAsObject":
		value := new(DeclaredTypeName)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SamePropertiesAsObject = value
	case "singleProperty":
		value := new(SingleUnionTypeProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SingleProperty = value
	case "noProperties":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.NoProperties = value
	}
	return nil
}

func (s SingleUnionTypeProperties) MarshalJSON() ([]byte, error) {
	switch s.PropertiesType {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", s.PropertiesType, s)
	case "samePropertiesAsObject":
		var marshaler = struct {
			PropertiesType string `json:"_type,omitempty"`
			*DeclaredTypeName
		}{
			PropertiesType:   s.PropertiesType,
			DeclaredTypeName: s.SamePropertiesAsObject,
		}
		return json.Marshal(marshaler)
	case "singleProperty":
		var marshaler = struct {
			PropertiesType string `json:"_type,omitempty"`
			*SingleUnionTypeProperty
		}{
			PropertiesType:          s.PropertiesType,
			SingleUnionTypeProperty: s.SingleProperty,
		}
		return json.Marshal(marshaler)
	case "noProperties":
		var marshaler = struct {
			PropertiesType string `json:"_type,omitempty"`
			NoProperties   any    `json:"noProperties,omitempty"`
		}{
			PropertiesType: s.PropertiesType,
			NoProperties:   s.NoProperties,
		}
		return json.Marshal(marshaler)
	}
}

type SingleUnionTypePropertiesVisitor interface {
	VisitSamePropertiesAsObject(*DeclaredTypeName) error
	VisitSingleProperty(*SingleUnionTypeProperty) error
	VisitNoProperties(any) error
}

func (s *SingleUnionTypeProperties) Accept(v SingleUnionTypePropertiesVisitor) error {
	switch s.PropertiesType {
	default:
		return fmt.Errorf("invalid type %s in %T", s.PropertiesType, s)
	case "samePropertiesAsObject":
		return v.VisitSamePropertiesAsObject(s.SamePropertiesAsObject)
	case "singleProperty":
		return v.VisitSingleProperty(s.SingleProperty)
	case "noProperties":
		return v.VisitNoProperties(s.NoProperties)
	}
}
