package response

import (
	"master-data/business/teknisi"
)

type TeknisiResponseData struct {
	ID            int    `json:"id"`
	FullName      string `json:"full_name"`
	Specialist    string `json:"specialist"`
	Platform      string `json:"platform"`
	JumlahAntrian int    `json:"jumlah_antrian"`
	Version       int    `json:"version"`
}

func NewTeknisiDataResponse(datas []teknisi.Teknisi) []TeknisiResponseData {
	var (
		responseData  TeknisiResponseData
		responseDatas []TeknisiResponseData
	)

	if datas == nil {
		return responseDatas
	}

	for _, teknisiData := range datas {

		responseData.ID = teknisiData.ID
		responseData.FullName = teknisiData.FullName
		responseData.Specialist = teknisiData.Specialist
		responseData.Platform = teknisiData.Platform
		responseData.JumlahAntrian = teknisiData.JumlahAntrian
		responseData.Version = teknisiData.Version

		responseDatas = append(responseDatas, responseData)
	}

	return responseDatas
}
