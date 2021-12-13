package kerusakan

type State string

const (
	StateAntrian    State = "antrian"
	StateDikerjakan State = "dikerjakan"
	StateSelesai    State = "selesai"
)

type Kerusakan struct {
	ID             int    `json:"_" db:"id"`
	JenisKerusakan string `json:"jenis_kerusakan" db:"jenis_kerusakan"`
	LamaPengerjaan string `json:"lama_pengerjaan" db:"lama_pengerjaan"`
	Harga          int    `json:"harga" db:"harga"`
	Version        int    `json:"version" db:"version"`
}

func NewKerusakan(
	jenis_kerusakan string,
	lama_pengerjaan string,
	harga int,
) Kerusakan {

	return Kerusakan{
		JenisKerusakan: jenis_kerusakan,
		LamaPengerjaan: lama_pengerjaan,
		Harga:          harga,
		Version:        1,
	}
}

// UpdateKerusakan
type UpdateKerusakanSpec struct {
	ID             int
	JenisKerusakan string
	LamaPengerjaan string
	Harga          int
	Version        int
}

func NewKerusakanUpdate(
	id int,
	jenis_kerusakan string,
	lama_pengerjaan string,
	harga int,
	version int,
) UpdateKerusakanSpec {

	return UpdateKerusakanSpec{
		ID:             id,
		JenisKerusakan: jenis_kerusakan,
		LamaPengerjaan: lama_pengerjaan,
		Harga:          harga,
		Version:        version,
	}
}

type InsertKerusakanSpec struct {
	JenisKerusakan string
	LamaPengerjaan string
	Harga          int
}
