//go:build ignore
// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	var err error

	println("generating assets ...")

	err = vfsgen.Generate(http.FileSystem(http.Dir("./web/dist/")), vfsgen.Options{
		Filename:     "./internal/assets/appcode.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatal(err)
	}
}
