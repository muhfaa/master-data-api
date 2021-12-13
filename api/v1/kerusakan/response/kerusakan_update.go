package response

import "master-data/api/v1/kerusakan/request"

type KerusakanUpdateResponseData struct {
	ID int `json:"id"`
}

func NewKerusakanUpdateResponse(updateKerusakan request.UpdateKerusakanRequest) KerusakanUpdateResponseData {
	var response KerusakanUpdateResponseData
	response.ID = updateKerusakan.ID

	return response
}
