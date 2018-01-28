package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/bketelsen/spago/markup"
	"github.com/gobuffalo/packr"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	bui := markup.NewCompBuilder()
	bui.Register(&Hello{})
	bui.Register(&World{})

	env := markup.NewEnv(bui)

	hello := &Hello{
		Greeting:    "Welcome",
		Name:        "Brian",
		Placeholder: "something here please",
		TextBye:     true,
	}

	root, _ := env.Mount(hello)
	var v bytes.Buffer
	enc := markup.NewTagEncoder(&v, env)
	enc.Encode(root)

	/*	hmarkup := "<Hello greeting='Welcome'></Hello>"
		v2 := bytes.NewBufferString(hmarkup)
		dec := markup.NewTagDecoder(v2)
		var t markup.Tag
		dec.Decode(&t)
		pretty.Println(t)
		fmt.Println(t.IsComponent())
	*/
	box := packr.NewBox("./assets")

	js.Global.Get("document").Call("write", v.String())

	js.Global.Get("document").Call("write", box.String("something.txt"))
	// http(v.String())

}

func httplisten(response string) {

	//fmt.Println(v.String())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, response)
	})
	http.ListenAndServe(":8080", nil)
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
