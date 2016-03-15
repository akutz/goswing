package swingers

// Swinger is the interface for types that can be shared.
type Swinger interface {

	// Name returns the name of the Swinger.
	Name() string
}

type swinger struct {
	name string
}

func (s *swinger) Name() string {
	return s.name
}

// NewSwinger returns a new swinger.
func NewSwinger() Swinger {
	return newSwinger()
}
