package validator

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

var errMsg map[string]string

func init() {
	errMsg = map[string]string{
		"Username": "非法的用户名",
		"Password": "无效的密码",
	}
}

type User struct {
	Username string `validate:"required"` //首字母必须大写，否则验证不了
	Password string `validate:"required"`
	Age      uint8  `validate:"gte=0,lte=130"`
}

func ValidateUser() {
	u := &User{
		Username: "",
		Password: "",
		Age:      180,
	}
	validate := validator.New()
	err := validate.Struct(u)
	fmt.Println(err)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, flErr := range validationErrors {
			if flErr.Field() == "UserName" {

			}
			/*
				fmt.Println(flErr)
				fmt.Println(flErr.Namespace())
				fmt.Println(flErr.Field())
				fmt.Println(flErr.StructNamespace()) // can differ when a custom TagNameFunc is registered or
				fmt.Println(flErr.StructField())     // by passing alt name to ReportError like below
				fmt.Println(flErr.Tag())
				fmt.Println(flErr.ActualTag())
				fmt.Println(flErr.Kind())
				fmt.Println(flErr.Type())
				fmt.Println(flErr.Value())
				fmt.Println(flErr.Param())
			*/
		}
	}

}
