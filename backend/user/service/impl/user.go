package impl

type UserServiceImpl struct{}

func (u *UserServiceImpl) Login(username string, password string) (string, error) {
	return "Login", nil
}

func (u *UserServiceImpl) Register(username string, password string) (string, error) {
	return "Register", nil
}

func (u *UserServiceImpl) Info(username string) (string, error) {
	return "Info", nil
}
