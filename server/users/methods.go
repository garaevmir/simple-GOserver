package users

import "fmt"

func NewUser(newname, newpassword string) User {
	return User{
		Name:     newname,
		password: newpassword,
	}
}

func (u User) String() string {
	return u.Name + "\n" + u.Info
}

func (u User) GetBalance() string {
	return fmt.Sprint(u.balance)
}

func (u *User) AddBalance(plus float64) {
	u.balance += plus
}

func (u *User) SubBalance(minus float64) {
	u.balance -= minus
}

func NewUsersStorage() UsersStorage {
	return UsersStorage{
		storage: make(map[string]*User),
	}
}

func (us UsersStorage) FindUser(name string) bool {
	_, ret := us.storage[name]
	return ret
}

func (us *UsersStorage) AddUser(user *User) bool {
	if us.FindUser(user.Name) {
		return false
	}
	us.storage[user.Name] = user
	return true
}

func (us UsersStorage) GetUser(name string) (*User, bool) {
	if us.FindUser(name) {
		return us.storage[name], true
	}
	return nil, false
}

func (us *UsersStorage) DeleteUser(name string) {
	if us.FindUser(name) {
		delete(us.storage, name)
	}
}

func (us UsersStorage) GetAll() string {
	var list string
	for i := range us.storage {
		list += i + "\n"
	}
	return list
}
