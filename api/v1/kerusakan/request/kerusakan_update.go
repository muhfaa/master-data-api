package request

import (
	"master-data/business/kerusakan"
)

type UpdateKerusakanRequest struct {
	ID             int    `json:"id" validate:"required"`
	JenisKerusakan string `json:"jenis_kerusakan"`
	LamaPengerjaan string `json:"lama_pengerjaan"`
	Harga          int    `json:"harga"`
	Version        int    `json:"version" validate:"required"`
}

func (req *UpdateKerusakanRequest) ToUpdateKerusakanSpec() *kerusakan.UpdateKerusakanSpec {
	var updateSpec kerusakan.UpdateKerusakanSpec
	updateSpec.ID = req.ID
	updateSpec.JenisKerusakan = req.JenisKerusakan
	updateSpec.LamaPengerjaan = req.LamaPengerjaan
	updateSpec.Harga = req.Harga
	updateSpec.Version = req.Version

	return &updateSpec
}
