package maybe

type User struct {
	name string
}

func getUser(id int) *User {
	if id >= 0 {
		return &User{name: "Mickey"}
	}

	return nil
}
