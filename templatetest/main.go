package main

import (
	"bytes"
	"fmt"

	"github.com/bketelsen/spago/markup"
	"github.com/kr/pretty"
)

const template = `<hello Greeting='greeting' Name='name'>inside the tags</hello>`

func main() {
	hmarkup := "<Hello greeting='Welcome'></Hello>"
	v2 := bytes.NewBufferString(hmarkup)
	dec := markup.NewTagDecoder(v2)
	var t markup.Tag
	dec.Decode(&t)
	pretty.Println(t)
	fmt.Println(t.Text)
	fmt.Println(t.IsComponent())
}

type Hello struct {
	Greeting    string
	Name        string
	Placeholder string
	TextBye     bool
	ChildErr    bool
}

func (h *Hello) Render() string {
	return `
<div>
	<h1>{{html .Greeting}}</h1>
	<input type="text" placeholder="{{.Placeholder}}" onchange="Name">
	<p>
		{{if .Name}}
			<world name="{{html .Name}}">
		{{else}}
			<span>World</span>
		{{end}}
	</p>


	{{if .TextBye}}
		Goodbye person.
	{{else}}
		<span>Goodbye</span>
		<p>world</p>
	{{end}}
</div>
	`
}

type World struct {
	Name string
}

func (w *World) Render() string {
	return `
<div>
	{{html .Name}}!
</div>
	`
}
