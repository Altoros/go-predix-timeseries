// Package measurement provides abstraction for measurements.
package measurement

import (
	"fmt"
)

type Int int
type Longint int64
type Float float32
type Double float64
type String string

type Measurement interface {
	Value() string
}

func (m Int) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m Longint) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m Float) Value() string {
	return fmt.Sprintf("%g", m)
}

func (m Double) Value() string {
	return fmt.Sprintf("%g", m)
}

func (m String) Value() string {
	return string(m)
}
