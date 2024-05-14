package appheader

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/link"
	"github.com/localrivet/templwind/utils"
)

type Opts struct {
	ID           string
	HideOnMobile bool
	LinkOpts     *link.Opts
	Title        string
}

// New creates a new component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewWithOpts creates a new component with the given opt
func NewWithOpts(opt *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opt)
}

// BuildOpts builds the options with the given opt
func BuildOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.BuildOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{
		ID:           utils.ToCamel(fmt.Sprintf("appHeader-%s", uuid.New().String())),
		HideOnMobile: false,
		Title:        "App Header",
	}
}

func WithID(id string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.ID = id
	}
}

func WithHideOnMobile(hide bool) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.HideOnMobile = hide
	}
}

func WithLinkOpts(linkOpts ...templwind.OptFunc[link.Opts]) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.LinkOpts = link.BuildOpts(linkOpts...)
	}
}

func WithTitle(title string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Title = title
	}
}
