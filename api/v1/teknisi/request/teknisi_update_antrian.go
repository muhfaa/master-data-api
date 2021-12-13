package request

import (
	"master-data/business/teknisi"
)

type UpdateJumlahAntrianRequest struct {
	ID      int `json:"id" validate:"required"`
	Version int `json:"version" validate:"required"`
}

func (req *UpdateJumlahAntrianRequest) ToUpdateJumlahAntrianSpec() *teknisi.UpdateJumlahAntrian {
	var updateAntrian teknisi.UpdateJumlahAntrian
	updateAntrian.ID = req.ID
	updateAntrian.Version = req.Version

	return &updateAntrian
}
