package controllers

import (
	"errors"

	"main.go/models"
)

func userCheck(user models.User) (bool, models.User) {

	for _, vale := range UserDemo {
		if vale.Email == user.Email {

			return true, vale
		}
	}
	return false, models.User{}
}

func passCheck(user models.User) error {
	for _, vale := range UserDemo {
		if vale.Password == user.Password {

			return nil
		}
	}
	return errors.New("pass word not match")
}
func useCheck(Token string) (error, models.User) {

	for _, vale := range UserDemo {
		if vale.Token == Token {

			return nil, vale
		}
	}
	return errors.New("token Cant Find"), models.User{}
}

func TokenSet(id, token string) (error, models.User) {
	for i, vale := range UserDemo {
		if vale.ID == id {
			UserDemo[i].Token = token
			return nil, vale
		}
	}
	return errors.New("token Cant Find"), models.User{}
}

// bcrypt
// http only cookie
// *http.Cookie
// set-cookie header
//jwt works accecc refersh token

// func setSeqrity(id, token string) error {
// 	for i, vale := range UserDemo {
// 		if vale.ID == id {
// 			UserDemo[i].Token = token
// 			return nil
// 		}
// 	}
// 	return errors.New("token cant set")
// }

// func ValidateToken(token string) (bool, error) {
// 	for _, vale := range UserDemo {
// 		if vale.ID == token {

// 			return true, nil
// 		}
// 	}
// 	return false, errors.New("somthin went Wrong")

// }
