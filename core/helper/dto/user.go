package dto

import "mime/multipart"

type (
	UserAuthenticationRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserResponse struct {
		ID       string `json:"id"`
		Username string `json:"username,omitempty"`
		Role     string `json:"role,omitempty"`
		Picture  string `json:"picture,omitempty"`
	}

	UserUpdateRequest struct {
		ID       string `json:"id"`
		Username string `json:"username" form:"username"`
		Role     string `json:"role" form:"role"`
		Password string `json:"password" form:"password"`
	}

	UserChangePictureRequest struct {
		Picture *multipart.FileHeader `json:"picture" form:"picture"`
	}
)
