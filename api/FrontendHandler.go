package api

import "net/http"

func FrontendHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./dist")).ServeHTTP(w, r)

}
