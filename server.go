package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

const (
	port = 8080
)

func main() {

	router := httprouter.New()
	router.GET("/croppedcape/:username", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username := ps.ByName("username")
		var scale int

		keys, ok := r.URL.Query()["scale"]

		if !ok || len(keys[0]) < 1 {
			// scale isn't present
			scale = 1
		} else {
			// scale is present
			scale, _ = strconv.Atoi(keys[0])
		}

		if scale > 100 {
			http.Error(w, fmt.Sprintf("a scale of %v is too high. you can only use a scale of up to 100.", scale), 400)
			return
		}

		log.Printf("grabbed cape for %v with scale %v", username, scale)

		capeBytes, err := getCapeBytes(username, scale)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(capeBytes)
	})

	router.ServeFiles("/info/*filepath", http.Dir("./static"))

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.Redirect(w, r, "/info", http.StatusPermanentRedirect)
	})

	fmt.Printf("Listening on port %v", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
