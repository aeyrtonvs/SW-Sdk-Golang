package authentication

import (
	"github.com/aeyrtonvs/cfdi-stampservice-go/helpers/entities"
)

type AuthenticationResponse struct {
	entities.Response
	Data AuthenticationDataResponse `json:"data"`
}
type AuthenticationDataResponse struct {
	Token      string `json:"token"`
	Expires_In int    `json:"expires_in"`
}
