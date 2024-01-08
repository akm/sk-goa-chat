package testuser

const DefaultPassword = "Passw0rd!"

func Foo() *User {
	return &User{Email: "foo@example.com", Name: "Foo", Password: DefaultPassword}
}

func Bar() *User {
	return &User{Email: "bar@example.com", Name: "Bar", Password: DefaultPassword}
}
