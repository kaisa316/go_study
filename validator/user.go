package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email string `validate:"required,email"`
}

func ValidteUser() {
	validate := validator.New()
	userInfo := &User{
		Email: "333",
	}
	err := validate.Struct(userInfo)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(err.Error())
}
