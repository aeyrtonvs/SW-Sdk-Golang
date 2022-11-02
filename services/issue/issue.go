package issue

import (
	"encoding/json"

	"github.com/aeyrtonvs/cfdi-stampservice-go/helpers"
	"github.com/aeyrtonvs/cfdi-stampservice-go/services/stamp"
)

func Issue(url string, token string, proxyPort int, proxy string) *IssueBase {
	base := new(IssueBase)
	base.Init(url, token, proxyPort, proxy)
	return base
}

func StampV1(issue IssueBase, xml string, isB64 bool) *stamp.StampResponseV1 {
	responseV1 := &stamp.StampResponseV1{}
	path := "/cfdi33/issue/v1"
	response, err := helpers.PostForm(issue.Url, path, issue.Token, xml)
	if err != nil {
		e := stamp.ToStampResponseV1(err)
		return e
	}
	err = json.Unmarshal(response, &responseV1)
	if err != nil {
		e := stamp.ToStampResponseV1(err)
		return e
	}
	return responseV1
}
func StampV2(issue IssueBase, xml string, isB64 bool) *stamp.StampResponseV2 {
	responseV2 := &stamp.StampResponseV2{}
	path := "/cfdi33/issue/v1"
	response, err := helpers.PostForm(issue.Url, path, issue.Token, xml)
	if err != nil {
		e := stamp.ToStampResponseV2(err)
		return e
	}
	err = json.Unmarshal(response, &responseV2)
	if err != nil {
		e := stamp.ToStampResponseV2(err)
		return e
	}
	return responseV2
}
func StampV3(issue IssueBase, xml string, isB64 bool) *stamp.StampResponseV3 {
	responseV3 := &stamp.StampResponseV3{}
	path := "/cfdi33/issue/v3"
	response, err := helpers.PostForm(issue.Url, path, issue.Token, xml)
	if err != nil {
		e := stamp.ToStampResponseV3(err)
		return e
	}
	err = json.Unmarshal(response, &responseV3)
	if err != nil {
		e := stamp.ToStampResponseV3(err)
		return e
	}
	return responseV3
}
func StampV4(issue IssueBase, xml string, isB64 bool) *stamp.StampResponseV4 {
	responseV4 := &stamp.StampResponseV4{}
	path := "/cfdi33/issue/v4"
	response, err := helpers.PostForm(issue.Url, path, issue.Token, xml)
	if err != nil {
		e := stamp.ToStampResponseV4(err)
		return e
	}
	err = json.Unmarshal(response, &responseV4)
	if err != nil {
		e := stamp.ToStampResponseV4(err)
		return e
	}
	return responseV4
}
