// +build !bob, !jim

package swingers

func newSwinger() Swinger {
	return &swinger{"World"}
}
