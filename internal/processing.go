package internal

import (
	"time"

	"github.com/h2non/bimg"
)

type ProcessingOptions struct {
	Image   *bimg.Image
	ImgType bimg.ImageType
	Quality int
}

func Process(opt *ProcessingOptions) ([]byte, error) {
	start := time.Now()

	defer func() {
		TimePerOpGauge.Set(float64(time.Since(start)))
		OpsCounter.Inc()
	}()

	buf, err := opt.Image.Process(bimg.Options{
		Type:    opt.ImgType,
		Quality: opt.Quality,
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}
