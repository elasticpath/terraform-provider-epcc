package field

import "fmt"

type Type string

const (
	String       Type = "string"
	Integer      Type = "integer"
	Boolean      Type = "boolean"
	Float        Type = "float"
	Date         Type = "date"
	Relationship Type = "relationship"
)

func (t Type) AsString() string {
	return string(t)
}

func (t Type) Validate() error {
	switch t {
	case String:
		fallthrough
	case Integer:
		fallthrough
	case Boolean:
		fallthrough
	case Float:
		fallthrough
	case Date:
		fallthrough
	case Relationship:
		return nil
	}
	return fmt.Errorf("field type must be one of %v", []string{
		String.AsString(),
		Integer.AsString(),
		Boolean.AsString(),
		Float.AsString(),
		Date.AsString(),
		Relationship.AsString(),
	})
}
