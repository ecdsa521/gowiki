package gowiki

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//StartServer starts http server listener
func StartServer(port int) {
	fmt.Printf("Starting HTTP server on port %d\n", port)
	router := httprouter.New()

	router.GET("/", indexPage)
	router.GET("/view/:name", viewPage)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)

}
func init() {

}
