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

		var (
			format     = r.URL.Query().Get("f")
			qualityStr = r.URL.Query().Get("q")
			widthStr   = r.URL.Query().Get("w")
			heigthStr  = r.URL.Query().Get("h")
		)

		quality, err := strconv.Atoi(qualityStr)
		if err != nil {
			quality = DEFAULT_QUALITY
		}

		width, err := strconv.Atoi(widthStr)
		if err != nil {
			width = 0
		}

		heigth, err := strconv.Atoi(heigthStr)
		if err != nil {
			heigth = 0
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
				Width:   width,
				Height:  heigth,
			})
			if err != nil {
				return nil
			}

			return res
		})

		work <- req

		res := <-resultChan
		close(resultChan)

		if res == nil {
			http.Error(w, "file conversion failed", http.StatusInternalServerError)
			return
		}

		w.Write(*res)
		w.Header().Add("Content-Type", pkg.MapImageTypeToContentType(imgType))
		w.WriteHeader(http.StatusOK)
	}
}
