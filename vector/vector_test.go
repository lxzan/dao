package vector

import "testing"

func TestUser_GetID(t *testing.T) {
	var docs Vector[User, string]
	docs = append(docs, User{ID: "a"})
	docs = append(docs, User{ID: "c"})
	docs = append(docs, User{ID: "c"})
	docs = append(docs, User{ID: "b"})
	docs.Uniq()
	docs.Sort()
	docs.Filter(func(i int, v User) bool {
		return v.ID == "b"
	})
	println(1)
}

type User struct {
	ID string
}

func (u User) GetID() string {
	return u.ID
}
