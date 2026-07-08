package users_transport_http

type UsersHTTPHandler struct {
	userService UsersService
}

type UsersService interface {

}

func NewUsersHTTPHandler (
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: usersService,
	}
}