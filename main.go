package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", abouthandler)
	// mux.HandleFunc("/product", productHandler)
	mux.HandleFunc("/buku", bukuHandler)

	log.Println("starting server :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello world!"))
}

func abouthandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ini halaman about"))
}

func bukuHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "ini adalah produk ke : %d", idNumb)

}