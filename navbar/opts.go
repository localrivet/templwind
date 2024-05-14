package navbar

import (
	"github.com/a-h/templ"
)

type OptFunc func(*Opt)

type MenuItem struct {
	Icon  string
	Title string
	Link  string
}

type Opt struct {
	BrandName string
	BrandLogo string
	BrandLink string
	MenuItems []MenuItem
}

func defaultOpts() *Opt {
	return &Opt{}
}

func WithBrandName(name string) OptFunc {
	return func(o *Opt) {
		o.BrandName = name
	}
}

func WithBrandLogo(logo string) OptFunc {
	return func(o *Opt) {
		o.BrandLogo = logo
	}
}

func WithBrandURL(link string) OptFunc {
	return func(o *Opt) {
		o.BrandLink = link
	}
}

func WithMenu(items []MenuItem) OptFunc {
	return func(o *Opt) {
		o.MenuItems = items
	}
}

func New(opts ...OptFunc) templ.Component {
	opt := defaultOpts()
	for _, optFn := range opts {
		optFn(opt)
	}

	return navbar(opt)
}
