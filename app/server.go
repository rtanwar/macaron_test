package main

import ("gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
m.Use(macaron.Renderer(macaron.RenderOptions{
			Directory:  "templates",              
			Extensions: []string{".tmpl", ".html"}, 
}))
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "pages/index")
	})
	m.Run()
}
