package main

import (
	"image/jpeg"
	"net/http"

	"github.com/karmdip-mi/go-fitz"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	doc, err := fitz.NewFromReader(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rgba, err := doc.Image(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = jpeg.Encode(w, rgba, &jpeg.Options{
		Quality: 90,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
