package main

import(
	"net/http"
)

func main(){
	http.HandleFunc("/test",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hallo world"))
	})
	http.ListenAndServe(":8080",nil)
}