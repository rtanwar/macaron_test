package main

import ("gopkg.in/macaron.v1";
"fmt"
)

func main() {
	m := macaron.Classic()
        m.Use(macaron.Renderer())
//m.Use(macaron.Renderer(macaron.RenderOptions{
//			Directory:  "templates",              
//			Extensions: []string{".tmpl", ".html"}, 
//}))
	m.Get("/", func(ctx *macaron.Context) {
                fmt.Printf("%s\n", "/")
                ctx.Data["title"] = "Index Page"
		ctx.HTML(200, "pages/index")
	})
        m.Get("/login", func(ctx *macaron.Context) {
                fmt.Printf("%s\n", "/login") 
                ctx.Data["title"] = "Login Page"
                ctx.HTML(200, "pages/login")
        })
	m.Run()
}
