package gowiki

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Page struct. Actual content is in Text above.
type Page struct {
	title string
	body  string
	id    int
}

func viewPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "<h1>Page:</h1><hr>\n%v\n<hr>\n", ps)
}
func init() {

}
func (p *Page) save() {

}
