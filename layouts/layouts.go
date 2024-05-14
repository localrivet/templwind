package layouts

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/layouts/base"
)

type BaseOpts struct {
	base.Opts
}

func BaseLayout(opts ...templwind.OptFunc[base.Opts]) templ.Component {
	return base.New(opts...)
}
