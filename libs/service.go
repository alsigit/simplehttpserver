package libs

/*
Name       : userService
Type       : struct
Properties :
	- db   : array of User
*/
type userService struct {
	db []*User
}

/*
Name       : User
Type       : struct
Properties :
	- name : string
*/
type User struct {
	Nama   string
	Alamat string
}

/*
Name	: Register
Type	: implementation function of interface
Return	: string
Parent	: userService
*/
func (u *userService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " registered"
}

/*
Name	: GetUser
Type	: implementation function of interface
Return	: string
Parent	: userService
*/
func (u *userService) GetUser() []*User {
	return u.db
}

/*
Name       : UserInterface
Type       : interface
Method	   :
	- Register : function
	- GetUser  : function
*/
type UserInterface interface {
	Register(u *User) string
	GetUser() []*User
}

func MyUsers(db []*User) UserInterface {
	return &userService{
		db: db,
	}
}
