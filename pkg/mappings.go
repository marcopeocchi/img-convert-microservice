package pkg

import (
	"errors"

	"github.com/h2non/bimg"
)

func MapStringToBimgType(format string) (bimg.ImageType, error) {
	switch format {
	case "avif":
		return bimg.AVIF, nil
	case "webp":
		return bimg.WEBP, nil
	case "jpeg":
		return bimg.JPEG, nil
	case "png":
		return bimg.PNG, nil
	case "gif":
		return bimg.GIF, nil
	case "tiff":
		return bimg.TIFF, nil
	default:
		return 0, errors.New("unsupported format")
	}
}

func MapImageTypeToContentType(t bimg.ImageType) string {
	switch t {
	case bimg.AVIF:
		return "image/avif"
	case bimg.WEBP:
		return "image/webp"
	case bimg.JPEG:
		return "image/jpeg"
	case bimg.PNG:
		return "image/png"
	case bimg.GIF:
		return "image/gif"
	case bimg.TIFF:
		return "image/tiff"
	default:
		return "application/octet-stream"
	}
}
