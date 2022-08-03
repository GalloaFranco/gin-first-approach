package services

type ILoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() ILoginService {
	return &loginService{ // Here we should call a DB or make a request to extern API, not leave the values hardcoded
		authorizedPassword: "123",
		authorizedUsername: "admin",
	}
}

func (loginService loginService) Login(username string, password string) bool {
	return loginService.authorizedPassword == password && loginService.authorizedUsername == username
}
