package gowiki

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

//GoWiki is an object for wiki interface
type GoWiki struct {
	Config map[string]interface{}
	router *httprouter.Router
}

//StartServer starts http server listener
func (wiki *GoWiki) StartServer() {
	port := wiki.Config["port"]
	fmt.Printf("Starting HTTP server on port %d\n", port)
	wiki.router = httprouter.New()

	wiki.router.GET("/", wiki.viewPage)
	wiki.router.GET("/view/:name", wiki.viewPage)
	wiki.router.GET("/edit/:name", wiki.editPage)
	wiki.router.POST("/edit/:name", wiki.editPage)
	wiki.router.GET("/user/create", wiki.createUser)
	wiki.router.GET("/user/login", wiki.loginUser)

	wiki.router.ServeFiles("/static/*filepath", http.Dir("static"))
	http.ListenAndServe(fmt.Sprintf(":%d", port), wiki.router)

}
func genTemplate(w http.ResponseWriter, r *http.Request, file string, data map[string]interface{}) {
	var t *template.Template
	t = template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"markDown": markDowner}).ParseFiles("template/layout.html", "template/navbar.html", "template/"+file))

	t.ExecuteTemplate(w, "layout", data)

}

func loadPage(name string, data *Page) error {
	text, _ := ioutil.ReadFile(name)
	return yaml.Unmarshal([]byte(text), &data)

}
func savePage(name string, data *Page) error {
	data.Version++

	data.Date = time.Now()
	text, err := yaml.Marshal(&data)
	if err == nil {

		return ioutil.WriteFile(name, text, 0666)
	}

	return err

}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func init() {

}
