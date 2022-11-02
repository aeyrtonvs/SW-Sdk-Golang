package stamp

import "github.com/aeyrtonvs/cfdi-stampservice-go/helpers/entities"

func ToStampResponseV1(ex error) *StampResponseV1 {
	r := entities.Response{
		Status:        "error",
		Message:       "Error inesperado.",
		MessageDetail: ex.Error(),
	}
	response := &StampResponseV1{
		Response: r,
		Data:     StampResponseDataV1{},
	}
	return response
}
func ToStampResponseV2(ex error) *StampResponseV2 {
	r := entities.Response{
		Status:        "error",
		Message:       "Error inesperado.",
		MessageDetail: ex.Error(),
	}
	response := &StampResponseV2{
		Response: r,
		Data:     StampResponseDataV2{},
	}
	return response
}
func ToStampResponseV3(ex error) *StampResponseV3 {
	r := entities.Response{
		Status:        "error",
		Message:       "Error inesperado.",
		MessageDetail: ex.Error(),
	}
	response := &StampResponseV3{
		Response: r,
		Data:     StampResponseDataV3{},
	}
	return response
}
func ToStampResponseV4(ex error) *StampResponseV4 {
	r := entities.Response{
		Status:        "error",
		Message:       "Error inesperado.",
		MessageDetail: ex.Error(),
	}
	response := &StampResponseV4{
		Response: r,
		Data:     StampResponseDataV4{},
	}
	return response
}
