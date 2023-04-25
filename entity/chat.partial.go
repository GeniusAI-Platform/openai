package entity

type Role uint8

const (
	SYSTEM = iota
	USER
	ASSISTANT
)

func (r Role) String() string {
	switch r {
	case SYSTEM:
		return "system"
	case USER:
		return "user"
	case ASSISTANT:
		return "assistant"
	default:
		return "assistant"
	}
}
