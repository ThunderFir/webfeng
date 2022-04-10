package main

import (
	"fmt"
	"log"
	"net/http"

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
	r.Get("/", pingHandler)
	r.Get("/hello", headerHandler)
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

func pingHandler(c *feng.Context) {
	c.String(feng.StatusOK, "ping %q success\n", c.Path)
}

func headerHandler(c *feng.Context) {
	for k, v := range c.Req.Header {
		_, _ = fmt.Fprintf(c.Writer, "Header[%q] = %q \n", k, v)
	}
}
