package response

type TeknisiUpdateResponseData struct {
	ID int `json:"id"`
}

func NewTeknisiUpdateResponse(id int) TeknisiUpdateResponseData {
	var responseUpdate TeknisiUpdateResponseData
	responseUpdate.ID = id

	return responseUpdate
}
