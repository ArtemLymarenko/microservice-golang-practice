package mapper

import (
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/interfaces/rest/dto"
)

func FromRegisterUserDTOToModel(registerDTO dto.RegisterUserRequest) model.User {
	return model.User{
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
		UserInfo: model.UserInfo{
			FirstName: registerDTO.FirstName,
			LastName:  registerDTO.LastName,
		},
	}
}

func FromLoginUserDTOToModel(loginDTO dto.LoginUserRequest) model.User {
	return model.User{
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}
}
