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

func Convert(work chan<- internal.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		resultChan := make(chan *[]byte)

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

		req := internal.NewRequest(resultChan, func() *[]byte {
			buffer, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil
			}

			image := bimg.NewImage(buffer)

			res, err := internal.Process(&internal.ProcessingOptions{
				Image:   image,
				ImgType: imgType,
				Quality: quality,
			})
			if err != nil {
				return nil
			}

			return res
		})

		work <- req

		res := <-resultChan
		if res == nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Write(*res)
		w.Header().Add("Content-Type", pkg.MapImageTypeToContentType(imgType))
		w.WriteHeader(http.StatusOK)
	}
}
