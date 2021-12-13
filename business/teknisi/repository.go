package teknisi

type Repository interface {
	InsertTeknisi(teknisiSpec Teknisi) error
	FindTeknisiByID(id int) (*Teknisi, error)
	FindAllTeknisi() ([]Teknisi, error)
	UpdateTeknisi(updateSpec TeknisiUpdateSpec) (bool, error)
	UpdateAntrian(id int, jumlahAntrian int, currentVersion int) (bool, error)
	DeleteTeknisi(id int) (bool, error)
}
