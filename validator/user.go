package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var errMsg map[string]string

func init() {
	errMsg = make(map[string]string)
	errMsg["Username"] = "非法的用户名"
	errMsg["Password"] = "无效的密码"
	errMsg["Age"] = "非法的年龄"
}

type User struct {
	Username string `validate:"required"` //首字母必须大写，否则验证不了
	Password string `validate:"required"`
	Age      uint8  `validate:"gte=0,lte=130"`
}

//如何返回error 类型错误

func ValidateUser() {
	u := &User{
		Username: "",
		Password: "",
		Age:      180,
	}
	validate := validator.New()
	err := validate.Struct(u)
	// fmt.Println(err)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, flErr := range validationErrors {
			fmt.Println(errMsg[flErr.Field()])
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

type User2 struct {
	Email string `validate:"required,email"`
}

func ValidteUser2() {
	validate := validator.New()
	userInfo := &User2{
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
