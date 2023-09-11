package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"path"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (

	//go:embed temp/*
	localFS embed.FS
)

type data struct {
	Name string
}

func main() {
	// single file
	// fmt.Println(x)

	// using a file system
	// file, _ := y.ReadFile("temp/empty.html")
	// fmt.Println(string(file))

	// you would go to localhost:3000/static/temp/...
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(y))))
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	// g.StaticFS("/templates", http.FS(y))

	g.HTMLRender = loadTemplates(localFS)
	// Step 1
	// rf := readFromFS(localFS)
	// name, data, err := rf("temp/layout/base.html")
	// if err != nil {
	// 	fmt.Println("error here: ", err)
	// }
	// fmt.Println(name, string(data))

	// Step 2
	// rf := readFromFS(localFS)
	// customParse(rf)

	g.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index", nil)
	})
	g.Run(":3333")
	// emptyTemp, err := y.ReadFile("temp/empty.html")
	// if err != nil {
	// 	panic(err)
	// }
	// empty := template.New("empty.html")
	// emptyTemp, err := empty.ParseFS(y, "temp/empty.html")
	// if err != nil {
	// 	fmt.Println("error while parsing the template file")
	// }

	// index := template.New("index.html")
	// indexTemp, err := index.ParseFS(y, "temp/layout/base.html", "temp/content/index.html")
	// if err != nil {
	// 	fmt.Println("error while parsing the template file")
	// }

	// emptyContent := template.New("empty_content.html")
	// emptyContentTemp, err := emptyContent.ParseFS(y, "temp/layout/base.html", "temp/empty_content.html")
	// if err != nil {
	// 	fmt.Println("error while parsing the template file")
	// }

	// err = emptyContent.ExecuteTemplate(os.Stdout, "base.html", data{
	// 	Name: "Yoooo",
	// })
	// if err != nil {
	// 	fmt.Println("error while executing template", err)
	// }

	// fmt.Println(index.DefinedTemplates())

	// r.Add("index", indexTemp)
	// r.Add("empty", emptyTemp)
	// r.Add("empty_content.html", emptyContentTemp)
	// g.HTMLRender = r

	// g.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index", nil)
	// })

	// g.GET("/empty", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "empty", data{
	// 		Name: "Aldrick",
	// 	})
	// })

	// g.GET("/emptyc", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "content", data{
	// 		Name: "Aldrick",
	// 	})
	// })

}

func loadTemplates(fileSystem embed.FS) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.Add(parsing("index", fileSystem, "temp/layout/base.html", "temp/content/index.html"))
	return r
}

func parsing(name string, fileSystem embed.FS, files ...string) (string, *template.Template) {
	tmpl := template.Must(customParse(readFromFS(fileSystem), files...))
	return name, tmpl
}

func customParse(readFile func(string) (string, []byte, error), filenames ...string) (*template.Template, error) {
	var t *template.Template

	for _, filename := range filenames {
		name, b, err := readFile(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		fmt.Println(name, s)

		var tmpl *template.Template
		t = template.New(name)

		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		_, err = t.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func readFromFS(fileSystem fs.FS) func(string) (string, []byte, error) {
	return func(file string) (string, []byte, error) {
		name := path.Base(file)
		b, err := fs.ReadFile(fileSystem, file)
		return name, b, err
	}
}
