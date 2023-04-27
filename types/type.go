package types

type (
	ID string
)

func (i ID) IsEmpty() bool {
	if len(i) == 0 {
		return true
	}
	return false
}
