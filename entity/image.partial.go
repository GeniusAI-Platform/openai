package entity

type (
	ImageSize           uint8
	ImageResponseFormat uint8
)

const (
	ImageSize256x256 ImageSize = iota
	ImageSize512x512
	ImageSize1024x1024
)

func (i ImageSize) String() string {
	switch i {
	case ImageSize256x256:
		return "256x256"
	case ImageSize512x512:
		return "512x512"
	case ImageSize1024x1024:
		return "1024x1024"
	default:
		return "512x512"
	}
}

const (
	ImageResponseFormatURL ImageResponseFormat = iota
	ImageResponseFormatB64JSON
)

func (i ImageResponseFormat) String() string {
	switch i {
	case ImageResponseFormatURL:
		return "url"
	case ImageResponseFormatB64JSON:
		return "b64_json"
	default:
		return "url"
	}
}
