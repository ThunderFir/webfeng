package main

import (
	"feng/middlewares"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"feng"
)

type Engine struct {
}

func (eg *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "xxx")
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Hello [%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404")
	}
}

// main v3
func main() {
	r := feng.New()
	r.Use(middlewares.Logger())
	r.SetFuncMap(template.FuncMap{"FormatAsDate": FormatAsDate})
	r.LoadHTMLGlob("templates/*")
	r.Static("/sub", "./templates")
	r.Get("/", welcome)
	r.Get("/hello", helloQuery)
	r.Get("/hello/:name", helloParam)
	r.Get("/ping/*path", pingHandler)

	v1 := r.Group("/v1")
	v1.Get("/", welcome)
	v1.Get("/date", showTime)
	v2 := r.Group("/v2")
	v2.Use(middlewares.LoggerV2())
	v2.Post("/", headerHandler)
	v2.Get("/hello", helloParam)

	log.Fatal(r.Run(":9876"))
}

// main v2
//func main() {
//	engine := new(Engine)
//	log.Fatal(http.ListenAndServe(":9876", engine))
//}

// main v1
//func main()  {
//	http.HandleFunc("/ping", pingHandler)
//	http.HandleFunc("/header", headerHandler)
//	log.Fatal(http.ListenAndServe(":9876", nil))
//}

func showTime(c *feng.Context) {
	c.HTML(feng.StatusOK, "custom_func.tmpl", feng.H{
		"title": "NOW",
		"now":   time.Now(),
	})
}

func welcome(c *feng.Context) {
	c.HTML(feng.StatusOK, "welcome.tmpl", nil)
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func helloQuery(c *feng.Context) {
	c.String(feng.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func helloParam(c *feng.Context) {
	c.String(feng.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
}

func pingHandler(c *feng.Context) {
	c.JSON(feng.StatusOK, feng.H{"ping path": c.Param("path")})
}

func headerHandler(c *feng.Context) {
	for k, v := range c.Req.Header {
		_, _ = fmt.Fprintf(c.Writer, "Header[%q] = %q \n", k, v)
	}
}
