package domain

import (
	"errors"
	"testing"
)

func FuzzNewLabel(f *testing.F) {
	for label := range allowedLabelsMap {
		f.Add(string(label))
	}

	f.Fuzz(func(t *testing.T, label string) {
		got, err := NewLabel(label)

		_, allowed := allowedLabelsMap[Label(label)]

		if allowed {
			if err != nil {
				t.Fatalf("NewLabel(%q) returned unexpected error: %v", label, err)
			}

			if got != Label(label) {
				t.Fatalf("NewLabel(%q) = %q, want %q", label, got, label)
			}
		} else {
			if err == nil {
				t.Fatalf("NewLabel(%q) returned no error for invalid label", label)
			}

			if !errors.Is(err, ErrInvalidLabel) {
				t.Fatalf("NewLabel(%q) error = %v, want ErrInvalidLabel", label, err)
			}
		}

	})
}
