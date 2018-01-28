package quote

import (
	"github.com/gobuffalo/plush"
)

type Quote struct {
	///spago.Component
	Text   string
	Author string
}

// Render outputs the tag
func (q Quote) Render() string {
	ctx := plush.NewContext()
	ctx.Set("this", q)
	s, err := plush.Render(q.Template(), ctx)
	if err != nil {
		return err.Error()
	}
	return s
}
func (q Quote) Template() string {
	return `<div><%=this.Text%> by <%= this.Author %></div>`
}
