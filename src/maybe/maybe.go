package maybe

// Maybe an interface that exposes the possible behavior when
// a User is present or not
type Maybe interface {
	IfPresent(f func(c *User))
	WhenAbsent(f func())
}

func MaybeUser(user *User) Maybe {
	if user == nil {
		return &Absent{}
	}

	return &Present{u: user}
}

// Present is the implementation of the Maybe when the user is present
type Present struct {
	u *User
}

// IfPresent interface implementation
func (m *Present) IfPresent(f func(c *User)) {
	f(m.u)
}

// WhenAbsent interface implementation
func (m *Present) WhenAbsent(f func()) {
}

// Absent is the implementation of the Maybe when the user is absent
type Absent struct{}

// Absent interface implementation
func (n *Absent) IfPresent(f func(c *User)) {
}

// WhenAbsent interface implementation
func (n *Absent) WhenAbsent(f func()) {
	f()
}
