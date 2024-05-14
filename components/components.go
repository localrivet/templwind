package components

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/appbar"
	"github.com/localrivet/templwind/components/apprail"
	"github.com/localrivet/templwind/components/appshell"
	"github.com/localrivet/templwind/components/card"
	"github.com/localrivet/templwind/components/dropdown"
)

func AppBar(opts ...templwind.OptFunc[appbar.Opts]) templ.Component {
	return appbar.New(opts...)
}

func AppRail(opts ...templwind.OptFunc[apprail.Opts]) templ.Component {
	return apprail.New(opts...)
}

func AppShell(opts ...templwind.OptFunc[appshell.Opts]) templ.Component {
	return appshell.New(opts...)
}

func Card(opts ...templwind.OptFunc[card.Opts]) templ.Component {
	return card.New(opts...)
}

func Dropdown(opts ...templwind.OptFunc[dropdown.Opts]) templ.Component {
	return dropdown.New(opts...)
}
