package user

type UserRepository interface {
	Create(user User) (User, error)
	GetById(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user User) error
	Delete(id int) error
	UserNameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
}
