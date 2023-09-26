package internal

import (
	"context"
	"errors"
	"io"

	"github.com/h2non/bimg"
)

type ProcessingOptions struct {
	Image   *bimg.Image
	ImgType bimg.ImageType
	Quality int
}

func Process(ctx context.Context, w io.Writer, opt *ProcessingOptions) error {
	select {
	case <-ctx.Done():
		return errors.New("context cancelled")

	default:
		return process(w, opt)
	}
}

func process(w io.Writer, opt *ProcessingOptions) error {
	buf, err := opt.Image.Process(bimg.Options{
		Type:    opt.ImgType,
		Quality: opt.Quality,
	})
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
