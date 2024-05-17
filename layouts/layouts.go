package layouts

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/layouts/app"
	"github.com/localrivet/templwind/layouts/base"
)

func BaseLayout(opts ...templwind.OptFunc[base.Opts]) templ.Component {
	return base.New(opts...)
}

func AppLayout(opts ...templwind.OptFunc[app.Opts]) templ.Component {
	return app.New(opts...)
}
