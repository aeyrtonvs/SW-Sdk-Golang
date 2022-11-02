package stamp

import (
	"encoding/json"

	"github.com/aeyrtonvs/cfdi-stampservice-go/helpers"
)

func Stamp(url string, token string, proxyPort int, proxy string) *StampBase {
	base := new(StampBase)
	base.Init(url, token, proxyPort, proxy)
	return base
}

// Servicio de Timbrado de un CFDI sellado.
// Regresa StampResponseV1
func StampV1(stamp StampBase, xml string, isB64 bool) *StampResponseV1 {
	responseV1 := &StampResponseV1{}
	path := "/cfdi33/stamp/v1"
	response, err := helpers.PostForm(stamp.Url, path, stamp.Token, xml)
	if err != nil {
		return ToStampResponseV1(err)
	}
	err = json.Unmarshal(response, &responseV1)
	if err != nil {
		return ToStampResponseV1(err)
	}
	return responseV1
}

// Servicio de Timbrado de un CFDI sellado.
// Regresa StampResponseV2
func StampV2(stamp StampBase, xml string, isB64 bool) *StampResponseV2 {
	responseV2 := &StampResponseV2{}
	path := "/cfdi33/stamp/v1"
	response, err := helpers.PostForm(stamp.Url, path, stamp.Token, xml)
	if err != nil {
		return ToStampResponseV2(err)
	}
	err = json.Unmarshal(response, &responseV2)
	if err != nil {
		return ToStampResponseV2(err)
	}
	return responseV2
}

// Servicio de Timbrado de un CFDI sellado.
// Regresa StampResponseV3
func StampV3(stamp StampBase, xml string, isB64 bool) *StampResponseV3 {
	responseV3 := &StampResponseV3{}
	path := "/cfdi33/stamp/v3"
	response, err := helpers.PostForm(stamp.Url, path, stamp.Token, xml)
	if err != nil {
		return ToStampResponseV3(err)
	}
	err = json.Unmarshal(response, &responseV3)
	if err != nil {
		return ToStampResponseV3(err)
	}
	return responseV3
}

// Servicio de Timbrado de un CFDI sellado.
// Regresa StampResponseV4
func StampV4(stamp StampBase, xml string, isB64 bool) *StampResponseV4 {
	responseV4 := &StampResponseV4{}
	path := "/cfdi33/stamp/v4"
	response, err := helpers.PostForm(stamp.Url, path, stamp.Token, xml)
	if err != nil {
		return ToStampResponseV4(err)
	}
	err = json.Unmarshal(response, &responseV4)
	if err != nil {
		return ToStampResponseV4(err)
	}
	return responseV4
}
