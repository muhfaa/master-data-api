package response

import (
	"master-data/business/kerusakan"
)

type KerusakanResponseData struct {
	ID             int    `json:"id"`
	JenisKerusakan string `json:"jenis_kerusakan"`
	LamaPengerjaan string `json:"lama_pengerjaan"`
	Harga          int    `json:"harga"`
	Version        int    `json:"version"`
}

func NewKerusakanDataResponse(datas []kerusakan.Kerusakan) []KerusakanResponseData {
	var (
		responseData  KerusakanResponseData
		responseDatas []KerusakanResponseData
	)

	if datas == nil {
		return responseDatas
	}

	for _, kerusakan := range datas {
		responseData.ID = kerusakan.ID
		responseData.JenisKerusakan = kerusakan.JenisKerusakan
		responseData.LamaPengerjaan = kerusakan.LamaPengerjaan
		responseData.Harga = kerusakan.Harga
		responseData.Version = kerusakan.Version

		responseDatas = append(responseDatas, responseData)
	}

	return responseDatas
}

func NewKerusakanResponse(data *kerusakan.Kerusakan) KerusakanResponseData {
	var response KerusakanResponseData
	response.ID = data.ID
	response.JenisKerusakan = data.JenisKerusakan
	response.LamaPengerjaan = data.LamaPengerjaan
	response.Harga = data.Harga
	response.Version = data.Version

	return response
}
