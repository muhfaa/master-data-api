package request

import "master-data/business/teknisi"

type InsertTeknisiRequest struct {
	FullName   string `json:"full_name" validate:"required"`
	Specialist string `json:"specialist" validate:"required"`
	Platform   string `json:"platform" validate:"required"`
}

func (req *InsertTeknisiRequest) ToUpsertTeknisiSpec() *teknisi.Teknisi {
	var teknisiSpec teknisi.Teknisi

	teknisiSpec.FullName = req.FullName
	teknisiSpec.Specialist = req.Specialist
	teknisiSpec.Platform = req.Platform

	return &teknisiSpec

}
