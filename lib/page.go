package gowiki

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/julienschmidt/httprouter"
)

//Page struct. Actual content is in Text above.
type Page struct {
	Title   string `yaml:"title"`
	url     string
	Content string `yaml:"content"`
	Size    int
	Version int `yaml:"version"`
	exists  bool
	Date    time.Time `yaml:"date"`
	id      int
}

func eq(d1 interface{}, d2 interface{}, s string, e string) string {
	if d1 == d2 {
		return s
	}
	return e

}
func neq(d1 interface{}, d2 interface{}, s string, e string) string {
	if d1 != d2 {
		return s
	}
	return e

}
func (wiki *GoWiki) navbar(active string, page *Page) []map[string]interface{} {

	data := []map[string]interface{}{
		{"Name": "View", "Href": fmt.Sprintf("/view/%s", page.url), "Class": eq("view", active, "active", "")},
		{"Name": "Edit", "Href": fmt.Sprintf("/edit/%s", page.url), "Class": eq("edit", active, "active", "")},
		{"Name": "History", "Href": fmt.Sprintf("/history/%s", page.url), "Class": eq("history", active, "active", "")},
	}

	return data
}
func (wiki *GoWiki) editPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page := Page{
		url: ps.ByName("name"),
	}
	page.load()
	fmt.Printf("f: %v\n", r.Method)
	if r.Method == "GET" {
		genTemplate(w, r, "edit.html", map[string]interface{}{
			"Navbar":  wiki.navbar("edit", &page),
			"Title":   eq(page.Title, "", page.url, page.Title),
			"Content": page.Content,
			"URL":     page.url,
		})
	} else if r.Method == "POST" {
		r.ParseForm()
		page.Title = r.FormValue("title")
		page.Content = r.FormValue("content")
		page.save()
		http.Redirect(w, r, fmt.Sprintf("/view/%s", page.url), 301)
	}
}
func (wiki *GoWiki) viewPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page := Page{
		url: ps.ByName("name"),
	}
	if page.url == "" {
		page.url = "index"
	}
	if page.load() {
		genTemplate(w, r, "view.html", map[string]interface{}{
			"Navbar":  wiki.navbar("view", &page),
			"Title":   eq(page.Title, "", page.url, page.Title), //if page.Title == "" return page.url, else page.Title
			"Content": page.Content,
		})
	} else {
		genTemplate(w, r, "view.html", map[string]interface{}{
			"Navbar":  wiki.navbar("view", &page),
			"Title":   page.url,
			"Content": "There is no text on this page.",
		})
	}

}
func init() {

}

//normalize file name
func name(i string) string {
	r := regexp.MustCompile("[\\W]+")
	i = r.ReplaceAllString(i, "-")
	return i
}
func (p *Page) save() bool {
	savePage("content/"+name(p.url)+".txt", p)

	return true
}

func (p *Page) load() bool {
	if _, err := os.Stat("content/" + name(p.url) + ".txt"); os.IsNotExist(err) {
		return false
	}
	p.url = name(p.url)

	loadPage("content/"+name(p.url)+".txt", p)

	return true
}
