package usecase

import (
	"app/entity"
	"fmt"
)

func (uc *Usecase) RegisterUser(user entity.User) string {
	// Prepare the SQL statement for insertion
	if isEmailValid := isValidEmail(user.Email); !isEmailValid {
		return "Wrong email format, please check your email."
	}

	hashedPword, errHash := hashPassword(user.Password)
	if errHash != nil {
		return fmt.Sprintf("please select a different password %v \n", errHash)
	}
	err := uc.Provider.InsertUser(user.Name, user.Email, hashedPword)
	if err != "" {
		return err
	}

	return "Successfully Registered User, please login using your email and registered password."
}

func (uc *Usecase) LoginUser(param entity.User) (entity.User, string) {
	user, err := uc.Provider.GetUserByEmail(param.Email)
	if err != "" {
		return entity.User{}, err
	}

	if isPWMatch := verifyPassword(user.HashPassword, param.Password); !isPWMatch {
		return entity.User{}, "Incorrect Password"
	}

	return user, ""
}
