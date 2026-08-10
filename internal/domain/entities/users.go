package entities

type UserEntity struct {
	ID        *uint64
	Firstname string
	Lastname  string
	Username  string
	DoB       string
	CreatedAt string
	Email     string
}

func NewUser(firstname, lastname, username, dob, email string) *UserEntity {
	return &UserEntity{
		ID:        nil,
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		DoB:       dob,
		Email:     email,
	}
}

func (u *UserEntity) GetUsername() string {
	return u.Username
}
func (u *UserEntity) GetFullname() string {
	return u.Firstname + " " + u.Lastname
}
func (u *UserEntity) GetEmail() string {
	return u.Email
}

func (u *UserEntity) GetDoB() string {
	return u.DoB
}
