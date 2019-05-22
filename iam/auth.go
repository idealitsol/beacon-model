package iam

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"-"`

	Domain   string `json:"-"`
	Platform string `json:"-"`

	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type AuthResponse struct {
	Token         string      `json:"token"`
	UserID         string      `json:"userId"`
	Roles          interface{} `json:"roles,omitempty"`
	ForcePWDChange bool        `json:"forcePwdChange"`
	Mesg           string      `json:"message"`
	Err            error       `json:"err,omitempty"`
}