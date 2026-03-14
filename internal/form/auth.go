package form

type Register struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	PasswordConf string `json:"passwordConf" validate:"required,eqfield=Password"`
	Role         string `json:"role" validate:"required,oneof=user admin"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
