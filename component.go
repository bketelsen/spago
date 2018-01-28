package spago

import "github.com/gobuffalo/plush"

type Componenter interface {
	Render() string
	Template() string
}

// use gopherpen style asset bundling
// determine if server or client
// wrap bootstrap as proof of concept

// Component is embedded into your thing
type Component struct {
}

// Render outputs the tag
func (c Component) Render(x Component) string {
	ctx := plush.NewContext()
	ctx.Set("this", x)
	s, err := plush.Render(x.Template(), ctx)
	if err != nil {
		return err.Error()
	}
	return s
}
