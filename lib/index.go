package gowiki

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
}

func indexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "<h1>Index:</h1><hr>\n%v\n<hr>\n", r.URL.Path)
}
