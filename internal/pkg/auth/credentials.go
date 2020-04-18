package auth

type Credentials interface {
	GetId() string
	GetUsername() string
	GetEmail() string
}
