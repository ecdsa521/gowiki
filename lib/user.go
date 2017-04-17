package gowiki

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//User struct
type User struct {
	id       int
	name     string
	password string
}

func init() {

}
func (wiki *GoWiki) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	genTemplate(w, r, "user.html", map[string]interface{}{
		"Hello": "world",
		"Title": "new user",
		"Loop": []map[string]interface{}{
			{"Name": "Test 1", "Href": "https://google.com"},
			{"Name": "Test 2", "Href": "https://pl.wikipedia.org"}},
	})
}

func (wiki *GoWiki) loginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (p *User) load() {

}
func (p *User) save() {

}
