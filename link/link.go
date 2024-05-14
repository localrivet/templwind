package link

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

type Opts struct {
	ID        string
	Class     string
	HXGet     string
	HXPost    string
	HXPut     string
	HXPatch   string
	HXDelete  string
	Target    string
	HXSwap    Swap
	HXTarget  string
	HXTrigger []string
	HXPushURL bool
	XOn       string
}

// NewDropdown creates a new Dropdown component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewDropdownWithOpt creates a new Dropdown component with the given opts
func NewWithOpts(opts *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opts)
}

// BuildDropdownOpts builds the options for Dropdown component
func BuildOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.BuildOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{
		HXSwap:   InnerHTML,
		HXTarget: "#content",
	}
}

func WithID(id string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ID = id
	}
}

func WithClass(class string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Class = class
	}
}

func WithHXGet(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXGet = href
	}
}

func WithHXPost(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPost = href
	}
}

func WithHXPut(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPut = href
	}
}

func WithHXPatch(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPatch = href
	}
}

func WithHXDelete(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXDelete = href
	}
}

func WithTarget(target string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Target = target
	}
}

type Swap string

func (s Swap) String() string {
	return string(s)
}

const (
	// the default, puts the content inside the target element
	InnerHTML Swap = "innerHTML"
	// replaces the entire target element with the returned content
	OuterHTML Swap = "outerHTML"
	// prepends the content before the first child inside the target
	AfterBegin Swap = "afterbegin"
	// prepends the content before the target in the target’s parent element
	BeforeBegin Swap = "beforebegin"
	// appends the content after the last child inside the target
	BeforeEnd Swap = "beforeend"
	// appends the content after the target in the target’s parent element
	AfterEnd Swap = "afterend"
	// deletes the target element regardless of the response
	Delete Swap = "delete"
	// does not append content from response (Out of Band Swaps and Response Headers will still be processed)
	None Swap = "none"
)

func WithHXSwap(hxSwap Swap) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXSwap = hxSwap
	}
}

func WithHXTarget(hxTarget string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXTarget = hxTarget
	}
}

func WithHXTrigger(hxTrigger []string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXTrigger = hxTrigger
	}
}

func WithHXPushURL(hxPushURL bool) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPushURL = hxPushURL
	}
}

func WithXOn(xOn string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.XOn = xOn
	}
}
