package iterator

type User struct {
	Name string
	Age  int
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	GetNext() *User
}

// Collection interface
type Collection interface {
	CreateIterator() Iterator
}

// Implementation of the Iterator by User
type UserIterator struct {
	index int
	users []*User
}

func (iterator *UserIterator) HasNext() bool {
	return iterator.index < len(iterator.users)
}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// Implementation of the Collection by User
type UserCollection struct {
	users []*User
}

func (c *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: c.users,
	}
}

// Factory Method for create UserCollection
func CreateUserCollection(users []*User) *UserCollection {
	return &UserCollection{users: users}
}
