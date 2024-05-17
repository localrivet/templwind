package dropdown

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/utils"
)

type Link struct {
	Icon  string
	Title string
	Link  string
	Click string
}

type Opts struct {
	ID    string
	Links []Link
}

// New creates a new component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewWithOpt creates a new component with the given opt
func NewWithOpts(opts *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opts)
}

// WithOpts builds the options with the given opt
func WithOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.WithOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{
		ID: utils.ToCamel(fmt.Sprintf("dropdown-%s", uuid.New().String())),
	}
}

func Links(items []Link) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Links = items
	}
}
