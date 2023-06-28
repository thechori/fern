// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ExampleTypeReferenceShape struct {
	Type      string
	Primitive *ExamplePrimitive
	Container *ExampleContainer
	Unknown   any
	Named     *ExampleNamedType
}

func NewExampleTypeReferenceShapeFromPrimitive(value *ExamplePrimitive) *ExampleTypeReferenceShape {
	return &ExampleTypeReferenceShape{Type: "primitive", Primitive: value}
}

func NewExampleTypeReferenceShapeFromContainer(value *ExampleContainer) *ExampleTypeReferenceShape {
	return &ExampleTypeReferenceShape{Type: "container", Container: value}
}

func NewExampleTypeReferenceShapeFromUnknown(value any) *ExampleTypeReferenceShape {
	return &ExampleTypeReferenceShape{Type: "unknown", Unknown: value}
}

func NewExampleTypeReferenceShapeFromNamed(value *ExampleNamedType) *ExampleTypeReferenceShape {
	return &ExampleTypeReferenceShape{Type: "named", Named: value}
}

func (e *ExampleTypeReferenceShape) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "primitive":
		var valueUnmarshaler struct {
			Primitive *ExamplePrimitive `json:"primitive,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Primitive = valueUnmarshaler.Primitive
	case "container":
		var valueUnmarshaler struct {
			Container *ExampleContainer `json:"container,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Container = valueUnmarshaler.Container
	case "unknown":
		var valueUnmarshaler struct {
			Unknown any `json:"unknown,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Unknown = valueUnmarshaler.Unknown
	case "named":
		value := new(ExampleNamedType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Named = value
	}
	return nil
}

func (e ExampleTypeReferenceShape) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "primitive":
		var marshaler = struct {
			Type      string            `json:"type,omitempty"`
			Primitive *ExamplePrimitive `json:"primitive,omitempty"`
		}{
			Type:      e.Type,
			Primitive: e.Primitive,
		}
		return json.Marshal(marshaler)
	case "container":
		var marshaler = struct {
			Type      string            `json:"type,omitempty"`
			Container *ExampleContainer `json:"container,omitempty"`
		}{
			Type:      e.Type,
			Container: e.Container,
		}
		return json.Marshal(marshaler)
	case "unknown":
		var marshaler = struct {
			Type    string `json:"type,omitempty"`
			Unknown any    `json:"unknown,omitempty"`
		}{
			Type:    e.Type,
			Unknown: e.Unknown,
		}
		return json.Marshal(marshaler)
	case "named":
		var marshaler = struct {
			Type string `json:"type,omitempty"`
			*ExampleNamedType
		}{
			Type:             e.Type,
			ExampleNamedType: e.Named,
		}
		return json.Marshal(marshaler)
	}
}

type ExampleTypeReferenceShapeVisitor interface {
	VisitPrimitive(*ExamplePrimitive) error
	VisitContainer(*ExampleContainer) error
	VisitUnknown(any) error
	VisitNamed(*ExampleNamedType) error
}

func (e *ExampleTypeReferenceShape) Accept(v ExampleTypeReferenceShapeVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "primitive":
		return v.VisitPrimitive(e.Primitive)
	case "container":
		return v.VisitContainer(e.Container)
	case "unknown":
		return v.VisitUnknown(e.Unknown)
	case "named":
		return v.VisitNamed(e.Named)
	}
}
