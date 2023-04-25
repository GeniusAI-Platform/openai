package models

type Moderation uint8

const (
	// TEXT_MODERATION_STABLE Almost as capable as the latest model, but slightly older
	TEXT_MODERATION_STABLE Moderation = iota + 1
	// TEXT_MODERATION_LATEST Most capable moderation model. Accuracy will be slighlty higher than the stable model
	TEXT_MODERATION_LATEST
)

func (m Moderation) String() string {
	switch m {
	case TEXT_MODERATION_STABLE:
		return "text-moderation-stable"
	case TEXT_MODERATION_LATEST:
		return "text-moderation-latest"
	default:
		return ""
	}
}
