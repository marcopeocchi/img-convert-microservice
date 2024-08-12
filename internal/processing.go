package internal

import (
	"time"

	"github.com/h2non/bimg"
)

type ProcessingOptions struct {
	Image   *bimg.Image
	ImgType bimg.ImageType
	Quality int
	Width   int
	Height  int
}

func Process(opt *ProcessingOptions) (*[]byte, error) {
	start := time.Now()

	defer func() {
		TimePerOpGauge.Set(float64(time.Since(start)))
		OpsCounter.Inc()
	}()

	options := bimg.Options{
		Type:    opt.ImgType,
		Quality: opt.Quality,
	}

	if opt.Width > 0 {
		options.Width = opt.Width
	}
	if opt.Height > 0 {
		options.Height = opt.Height
	}

	buf, err := opt.Image.Process(options)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
