package main

import "github.com/ecdsa521/gowiki/lib"

func main() {

	wiki := gowiki.GoWiki{
		Config: map[string]interface{}{
			"port":     7777,
			"sitename": "Go Wiki",
		},
	}

	wiki.StartServer()

}
