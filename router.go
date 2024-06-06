// router.go
package main

import(
	"fmt"
   "net/http"
   "github.com/asccclass/sherryserver"
	"github.com/asccclass/sherryai/botmanager"
)

func NewRouter(srv *SherryServer.Server, documentRoot string)(*http.ServeMux) {
   router := http.NewServeMux()

   // Static File server
   staticfileserver := SherryServer.StaticFileServer{documentRoot, "index.html"}
   staticfileserver.AddRouter(router)

	// add ai
	ai, err := botManager.NewBotManager("Taide")
	if err != nil {
      fmt.Println("Initial boot manager error:" + err.Error())
		return nil
	}
	fmt.Println(ai.Url)

/*
   // App router
   router.Handle("/homepage", oauth.Protect(http.HandlerFunc(Home)))
   router.Handle("/upload", oauth.Protect(http.HandlerFunc(Upload)))
   router.Handle("/download/{fileName}", oauth.Protect(http.HandlerFunc(Download)))
   router.Handle("/logout", http.HandlerFunc(oauth.Logout))
*/	
   return router
}