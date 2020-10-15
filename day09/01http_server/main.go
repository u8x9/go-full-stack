package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, 世界"
	w.Write([]byte(str))
}
func main() {
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe("127.0.0.1:19090", nil); err != nil {
		panic(err)
	}
}
