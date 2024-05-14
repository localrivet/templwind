package templwind

import "github.com/a-h/templ"

// Component interface with generic methods
type Component[T any] interface {
	New(opts ...OptFunc[T]) templ.Component
	NewWithOpt(opt *T) templ.Component
	BuildOpts(opts ...OptFunc[T]) *T
}

// OptFunc is a generic function type for options
type OptFunc[T any] func(*T)

// New creates a new templ.Component with the given options
func New[T any](defaultOpts func() *T, tpl func(*T) templ.Component, opts ...OptFunc[T]) templ.Component {
	opt := defaultOpts()
	for _, optFn := range opts {
		optFn(opt)
	}
	return tpl(opt)
}

// NewWithOpt creates a new templ.Component with the given opt
func NewWithOpts[T any](tpl func(*T) templ.Component, opts *T) templ.Component {
	return tpl(opts)
}

// BuildOpts constructs the options with the given option functions
func BuildOpts[T any](defaultOpts func() *T, opts ...OptFunc[T]) *T {
	opt := defaultOpts()
	for _, optFn := range opts {
		optFn(opt)
	}
	return opt
}
