package request

import "master-data/business/teknisi"

type UpdateTeknisiRequest struct {
	ID         int    `json:"id" validate:"required"`
	FullName   string `json:"full_name"`
	Specialist string `json:"specialist"`
	Platform   string `json:"platform"`
	Version    int    `json:"version" validate:"required"`
}

func (req *UpdateTeknisiRequest) ToUpdateTeknisiSpec() *teknisi.TeknisiUpdateSpec {
	var updateTeknisiRequest teknisi.TeknisiUpdateSpec
	updateTeknisiRequest.ID = req.ID
	updateTeknisiRequest.FullName = req.FullName
	updateTeknisiRequest.Specialist = req.Specialist
	updateTeknisiRequest.Platform = req.Platform
	updateTeknisiRequest.Version = req.Version

	return &updateTeknisiRequest
}
