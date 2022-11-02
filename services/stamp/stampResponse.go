package stamp

import (
	"github.com/aeyrtonvs/cfdi-stampservice-go/helpers/entities"
)

type StampResponseV1 struct {
	entities.Response
	Data StampResponseDataV1 `json:"data"`
}
type StampResponseV2 struct {
	entities.Response
	Data StampResponseDataV2 `json:"data"`
}
type StampResponseV3 struct {
	entities.Response
	Data StampResponseDataV3 `json:"data"`
}
type StampResponseV4 struct {
	entities.Response
	Data StampResponseDataV4 `json:"data"`
}
type StampResponseDataV1 struct {
	Tfd string `json:"tfd"`
}
type StampResponseDataV2 struct {
	Tfd  string `json:"tfd"`
	Cfdi string `json:"cfdi"`
}
type StampResponseDataV3 struct {
	Cfdi string `json:"cfdi"`
}
type StampResponseDataV4 struct {
	CadenaOriginalSat string `json:"cadenaOriginalSAT"`
	NoCertificadoSat  string `json:"noCertificadoSAT"`
	NoCertificadoCfdi string `json:"noCertificadoCFDI"`
	Uuid              string `json:"uuid"`
	SelloSat          string `json:"selloSAT"`
	SelloCfdi         string `json:"selloCFDI"`
	FechaTimbrado     string `json:"fechaTimbrado"`
	QrCode            string `json:"qrCode"`
	Cfdi              string `json:"cfdi"`
}
