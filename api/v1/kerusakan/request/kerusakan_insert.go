package request

import (
	"master-data/business/kerusakan"
)

type InsertKerusakanRequest struct {
	JenisKerusakan string `json:"jenis_kerusakan" validate:"required"`
	LamaPengerjaan string `json:"lama_pengerjaan" validate:"required"`
	Harga          int    `json:"harga" validate:"required"`
}

func (req *InsertKerusakanRequest) ToUpsertKerusakanSpec() kerusakan.Kerusakan {
	var insertSpec kerusakan.Kerusakan
	insertSpec.JenisKerusakan = req.JenisKerusakan
	insertSpec.LamaPengerjaan = req.LamaPengerjaan
	insertSpec.Harga = req.Harga

	return insertSpec
}
