package authentication

import "github.com/aeyrtonvs/cfdi-stampservice-go/helpers/entities"

func ToAuthenticationResponse(ex error) *AuthenticationResponse {
	r := &entities.Response{
		Status:        "error",
		Message:       "Error inesperado.",
		MessageDetail: ex.Error(),
	}
	response := &AuthenticationResponse{
		Response: *r,
		Data:     AuthenticationDataResponse{},
	}
	return response
}
