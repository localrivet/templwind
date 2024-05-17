package link

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

type Opts struct {
	ID         string
	Title      string
	Badge      templ.Component
	Icon       string
	Class      string
	HXGet      string
	HXPost     string
	HXPut      string
	HXPatch    string
	HXDelete   string
	Target     string
	HXSwap     Swap
	HXTarget   string
	HXTrigger  []string
	HXPushURL  bool
	XOnTrigger string
	Submenu    []*Opts
}

// New creates a new component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewWithOpts creates a new component with the given opts
func NewWithOpts(opts *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opts)
}

// WithOpts builds the options for the component
func WithOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.WithOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{
		HXSwap:   InnerHTML,
		HXTarget: "#content",
	}
}

func ID(id string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ID = id
	}
}

func Title(title string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Title = title
	}
}

func Badge(badge templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Badge = badge
	}
}

func Icon(icon string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Icon = icon
	}
}

func Class(class string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Class = class
	}
}

func HXGet(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXGet = href
	}
}

func HXPost(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPost = href
	}
}

func HXPut(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPut = href
	}
}

func HXPatch(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPatch = href
	}
}

func HXDelete(href string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXDelete = href
	}
}

func Target(target string) templwind.OptFunc[Opts] {
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

func HXSwap(hxSwap Swap) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXSwap = hxSwap
	}
}

func HXTarget(hxTarget string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXTarget = hxTarget
	}
}

func HXTrigger(hxTrigger []string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXTrigger = hxTrigger
	}
}

func HXPushURL(hxPushURL bool) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HXPushURL = hxPushURL
	}
}

func XOnTrigger(xOnTrigger string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.XOnTrigger = xOnTrigger
	}
}

func Submenu(linkOpts ...*Opts) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Submenu = append(opts.Submenu, linkOpts...)
	}
}
