// Generated by Fern. Do not edit.

package ir

type ObjectTypeDeclaration struct {
	// A list of other types to inherit from
	Extends    []*DeclaredTypeName `json:"extends,omitempty"`
	Properties []*ObjectProperty   `json:"properties,omitempty"`
}
