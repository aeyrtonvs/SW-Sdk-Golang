package issue_test

import (
	"log"
	"testing"

	"github.com/aeyrtonvs/cfdi-stampservice-go/services/issue"
	testhelp "github.com/aeyrtonvs/cfdi-stampservice-go/test"
)

var build = testhelp.Init()

func Test_IssueV1(t *testing.T) {
	testCases := []struct {
		Name      string
		Url       string
		Token     string
		ProxyPort int
		Proxy     string
		Xml       string
		Expected  string
	}{
		{
			Name:      "Success",
			Url:       "https://services.test.sw.com.mx",
			Token:     build.Token,
			ProxyPort: 0,
			Proxy:     "",
			Xml:       testhelp.GetXml(),
			Expected:  "success",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Log("Token:" + tc.Token)
			issueBase := issue.Issue(tc.Url, tc.Token, tc.ProxyPort, tc.Proxy)
			result := issue.StampV1(*issueBase, tc.Xml, false)
			if result == nil || result.Status != tc.Expected {
				t.Errorf("Test FAILED. Got: %s, Expected: %s", result.Status, tc.Expected)
				t.Fatalf(result.MessageDetail)
				t.Log(tc.Xml)

			} else {
				log.Print(result.Data)
			}
		})
	}
}

func Test_IssueV2(t *testing.T) {
	testCases := []struct {
		Name      string
		Url       string
		Token     string
		ProxyPort int
		Proxy     string
		Xml       string
		Expected  string
	}{
		{
			Name:      "Success",
			Url:       "https://services.test.sw.com.mx",
			Token:     build.Token,
			ProxyPort: 0,
			Proxy:     "",
			Xml:       testhelp.GetXml(),
			Expected:  "success",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			issueBase := issue.Issue(tc.Url, tc.Token, tc.ProxyPort, tc.Proxy)
			result := issue.StampV1(*issueBase, tc.Xml, false)
			if result == nil || result.Status != tc.Expected {
				t.Errorf("Test FAILED. Got: %s, Expected: %s", result.Status, tc.Expected)
				t.Fatalf(result.MessageDetail)
				t.Log(tc.Xml)
			} else {
				log.Print(result.Data)
			}
		})
	}
}
func Test_IssueV3(t *testing.T) {
	testCases := []struct {
		Name      string
		Url       string
		Token     string
		ProxyPort int
		Proxy     string
		Xml       string
		Expected  string
	}{
		{
			Name:      "Success",
			Url:       "https://services.test.sw.com.mx",
			Token:     build.Token,
			ProxyPort: 0,
			Proxy:     "",
			Xml:       testhelp.GetXml(),
			Expected:  "success",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			issueBase := issue.Issue(tc.Url, tc.Token, tc.ProxyPort, tc.Proxy)
			result := issue.StampV1(*issueBase, tc.Xml, false)
			if result == nil || result.Status != tc.Expected {
				t.Errorf("Test FAILED. Got: %s, Expected: %s", result.Status, tc.Expected)
				t.Fatalf(result.MessageDetail)
				t.Log(tc.Xml)
			} else {
				log.Print(result.Data)
			}
		})
	}
}
func Test_IssueV4(t *testing.T) {
	testCases := []struct {
		Name      string
		Url       string
		Token     string
		ProxyPort int
		Proxy     string
		Xml       string
		Expected  string
	}{
		{
			Name:      "Success",
			Url:       "https://services.test.sw.com.mx",
			Token:     build.Token,
			ProxyPort: 0,
			Proxy:     "",
			Xml:       testhelp.GetXml(),
			Expected:  "success",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			issueBase := issue.Issue(tc.Url, tc.Token, tc.ProxyPort, tc.Proxy)
			result := issue.StampV1(*issueBase, tc.Xml, false)
			if result == nil || result.Status != tc.Expected {
				t.Errorf("Test FAILED. Got: %s, Expected: %s", result.Status, tc.Expected)
				t.Fatalf(result.MessageDetail)
				t.Log(tc.Xml)
			} else {
				log.Print(result.Data)
			}
		})
	}
}
