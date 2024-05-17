package components

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/appbar"
	"github.com/localrivet/templwind/components/apprail"
	"github.com/localrivet/templwind/components/appshell"
	"github.com/localrivet/templwind/components/card"
	"github.com/localrivet/templwind/components/dropdown"
	"github.com/localrivet/templwind/components/link"
	"github.com/localrivet/templwind/components/sidenav"
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

func Dropdown(opts ...templwind.OptFunc[dropdown.Opts]) templ.Component {
	return dropdown.New(opts...)
}

func Card(opts ...templwind.OptFunc[card.Opts]) templ.Component {
	return card.New(opts...)
}

func Link(opts ...templwind.OptFunc[link.Opts]) templ.Component {
	return link.New(opts...)
}

func Sidenav(opts ...templwind.OptFunc[sidenav.Opts]) templ.Component {
	return sidenav.New(opts...)
}
