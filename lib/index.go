package gowiki

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
}

func (wiki *GoWiki) indexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	genTemplate(w, r, "index.html", map[string]interface{}{
		"Title": wiki.Config["sitename"].(string),
	})

}
