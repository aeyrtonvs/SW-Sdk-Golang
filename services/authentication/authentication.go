package authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Authentication(url string, user string, password string, proxyPort int, proxy string) *authenticationBase {
	base := new(authenticationBase)
	base.InitAuth(url, user, password, proxyPort, proxy)
	return base
}

func GetToken(authentication authenticationBase) *AuthenticationResponse {
	responseAuth := &AuthenticationResponse{}
	url := authentication.Url + "/security/authenticate"
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		e := ToAuthenticationResponse(err)
		return e
	}
	request.Header.Add("user", authentication.User)
	request.Header.Add("password", authentication.Password)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		e := ToAuthenticationResponse(err)
		return e
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		e := ToAuthenticationResponse(err)
		return e
	}
	err = json.Unmarshal(responseBytes, &responseAuth)
	if err != nil {
		e := ToAuthenticationResponse(err)
		return e
	}
	response.Body.Close()
	return responseAuth
}
