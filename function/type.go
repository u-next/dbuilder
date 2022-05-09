package function

import "fmt"

// Type determines if a node belongs to particular type.
type Type struct {
	// targetType to filter for
	targetType string

	// invert controls whether a NOT operator should be appended to the front to invert the results
	invert bool
}

// NewType return Type
func NewType(targetType string, invert bool) Type {
	return Type{targetType: targetType, invert: invert}
}

func (op Type) String() string {
	cls := fmt.Sprintf("type(%s)", op.targetType)

	if op.invert {
		return fmt.Sprintf("NOT %s", cls)
	}

	return fmt.Sprintf("%s", cls)
}
