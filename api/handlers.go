package api

import (
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/chai2010/webp"
	"github.com/karmdip-mi/go-fitz"
)

const MaxBodySize = 10_000_000

func Convert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, MaxBodySize)

	format := r.URL.Query().Get("format")
	if format == "" {
		format = "webp"
		return
	}

	select {
	case <-r.Context().Done():
		http.Error(w, "context cancelled", http.StatusInternalServerError)
		return

	default:
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

		switch format {
		case "png":
			err = png.Encode(w, rgba)
		case "jpeg":
			err = jpeg.Encode(w, rgba, &jpeg.Options{
				Quality: 90,
			})
		case "webp":
			err = webp.Encode(w, rgba, &webp.Options{
				Quality: 85,
			})
		default:
			http.Error(w, "invalid format", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
