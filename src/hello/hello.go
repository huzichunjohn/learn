package main

import "github.com/codegangsta/martini"

func main() {

    server := martini.Classic()

    server.Get("/", func() string {
	return "<h1>hello world!</h1>"
    })

    server.Get("/:who", func(args martini.Params) string {
	return "<h1>hello " + args["who"] + "</h1>"
    })

    server.Run()
}
