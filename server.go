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
			fmt.Fprintf(w, "a scale of %v is too high. you can only use a scale of up to 100.", scale)
			return
		}

		log.Printf("grabbed cape for %v with scale %v", username, scale)

		capeBytes, err := getCapeBytes(username, scale)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		w.Write(capeBytes)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
