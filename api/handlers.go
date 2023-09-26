package api

import (
	"fuku/internal"
	"fuku/pkg"
	"io"
	"net/http"
	"strconv"

	"github.com/h2non/bimg"
)

const (
	MAX_BODY_SIZE   = 10_000_000
	DEFAULT_QUALITY = 85
)

func Convert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, MAX_BODY_SIZE)

	format := r.URL.Query().Get("format")

	quality, err := strconv.Atoi(r.URL.Query().Get("quality"))
	if err != nil {
		quality = DEFAULT_QUALITY
	}

	imgType, err := pkg.MapStringToBimgType(format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	image := bimg.NewImage(buffer)

	err = internal.Process(r.Context(), w, &internal.ProcessingOptions{
		Image:   image,
		ImgType: imgType,
		Quality: quality,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", pkg.MapImageTypeToContentType(imgType))
	w.WriteHeader(http.StatusOK)
}
