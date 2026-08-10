package entities

type Profile struct {
	ID     *uint64
	UserID uint64
}

func NewProfile(uid uint64) *Profile {
	return &Profile{
		ID:     nil,
		UserID: uid,
	}
}
