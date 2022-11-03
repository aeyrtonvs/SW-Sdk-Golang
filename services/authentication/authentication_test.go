package authentication_test

import (
	"testing"

	"github.com/aeyrtonvs/cfdi-stampservice-go/services/authentication"
	testhelp "github.com/aeyrtonvs/cfdi-stampservice-go/test"
)

var build = testhelp.Init()

func Test_Authentication(t *testing.T) {

	testCases := []struct {
		Name      string
		Url       string
		User      string
		Password  string
		ProxyPort int
		Proxy     string
		Expected  string
	}{
		{
			Name:      "Success",
			Url:       "https://services.test.sw.com.mx",
			User:      build.User,
			Password:  build.Password,
			ProxyPort: 0,
			Proxy:     "",
			Expected:  "success",
		},
		{
			Name:      "Error_InvalidCredentials",
			Url:       "https://services.test.sw.com.mx",
			User:      build.User,
			Password:  "invalidPassword",
			ProxyPort: 0,
			Proxy:     "",
			Expected:  "error",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			auth := authentication.Authentication(tc.Url, tc.User, tc.Password, tc.ProxyPort, tc.Proxy)
			result := authentication.GetToken(*auth)
			if result == nil || result.Status != tc.Expected {
				t.Errorf("Test FAILED. Got: %s, Expected: %s", result.Status, tc.Expected)
			}
		})
	}
}
