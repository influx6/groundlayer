package pages

import (
	"github.com/influx6/groundlayer/pkg/styled"
	"github.com/influx6/groundlayer/pkg/styled/base"
)

var Theme = &base.Theme{
	Colors:    nil,
	Screens:   styled.TailwindUIScreens,
	Resolvers: base.ThemeResolvers{},
}
