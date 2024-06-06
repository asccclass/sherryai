// router.go
package main

import(
   "net/http"
   "github.com/asccclass/sherryserver"
)

func NewRouter(srv *SherryServer.Server, documentRoot string)(*http.ServeMux) {
   router := http.NewServeMux()

   // Static File server
   staticfileserver := SherryServer.StaticFileServer{documentRoot, "index.html"}
   staticfileserver.AddRouter(router)
/*
   // App router
   router.Handle("/homepage", oauth.Protect(http.HandlerFunc(Home)))
   router.Handle("/upload", oauth.Protect(http.HandlerFunc(Upload)))
   router.Handle("/download/{fileName}", oauth.Protect(http.HandlerFunc(Download)))
   router.Handle("/logout", http.HandlerFunc(oauth.Logout))
*/	
   return router
}