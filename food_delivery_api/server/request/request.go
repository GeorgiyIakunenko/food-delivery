package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Age       int    `json:"age" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
	Address   string `json:"address" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

/*type RefreshRequest struct {
	Token string `json:"token"`
}*/

type ChangePasswordRequest struct {
	//Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type PasswordResetEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordResetRequest struct {
	Email       string `json:"email" validate:"required,email"`
	ResetCode   string `json:"reset_code" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type OrderProductRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required"`
}

type OrderRequest struct {
	TotalPrice    float64               `json:"total_price" validate:"required"`
	PaymentMethod string                `json:"payment_method" validate:"required"`
	Address       string                `json:"address" validate:"required"`
	Products      []OrderProductRequest `json:"product" validate:"required"`
}

func ValidateRequest(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func ParseBody(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
