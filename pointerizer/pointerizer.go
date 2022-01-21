package pointerizer

import (
	"strings"
	"time"
)

func S(s string) *string { return &s }

func I(i int) *int { return &i }

func I32(i int32) *int32 { return &i }

func F64(f float64) *float64 { return &f }

func F32(f float32) *float32 { return &f }

func T(t time.Time) *time.Time { return &t }

func B(b bool) *bool { return &b }

func Ba(s []string) []byte {
	return []byte(strings.Join(s, ""))
}

// IntOrNil converts the pointer of int32 to int or nil
func IntOrNil(i *int32) *int {
	if i == nil {
		return nil
	}

	return I(int(*i))
}

// Int32OrNil converts the pointer of int to int32 or nil
func Int32OrNil(i *int) *int32 {
	if i == nil {
		return nil
	}

	return I32(int32(*i))
}
