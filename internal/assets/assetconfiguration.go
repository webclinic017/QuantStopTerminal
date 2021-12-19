//go:build dev
// +build dev

package assets

import (
	"github.com/shurcooL/httpfs/filter"
	"net/http"
	"os"
	pathpkg "path"
)

/*
 * Skip files we don't need to include
 * https://medium.com/swlh/single-binary-web-apps-in-go-and-vue-part-3-73f65e9cccf3
 */
var skipFunc = func(path string, fi os.FileInfo) bool {
	fname := fi.Name()
	extension := pathpkg.Ext(fname)

	return extension == ".go" ||
		extension == ".DS_Store" ||
		extension == ".md" ||
		extension == ".svg" ||
		fname == "LICENSE"
}

var Assets = filter.Skip(http.Dir("./web/dist/"), skipFunc)
