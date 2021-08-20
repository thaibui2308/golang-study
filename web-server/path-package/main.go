package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()

	// Add path to the Path Resolution
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)

	http.ListenAndServe("3000", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	checkPath := req.Method + req.URL.Path
	for pattern, handleFunc := range p.handlers {
		if match, err := path.Match(pattern, checkPath); match && err == nil {
			handleFunc(res, req)
		} else if err != nil {
			fmt.Fprint(res, req)
		}
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Default Name"
	}
	fmt.Fprint(res, "Hello ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")

	name := parts[2]

	if name == "" {
		name = "Default User"
	}
	fmt.Fprint(res, "Hello ", name)
}
