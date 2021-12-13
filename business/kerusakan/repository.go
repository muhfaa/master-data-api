package kerusakan

type Repository interface {
	InsertKerusakan(kerusakanSpec Kerusakan) error
	FindKerusakanByID(id int) (*Kerusakan, error)
	FindAllKerusakan() ([]Kerusakan, error)
	UpdateKerusakan(updateSpec UpdateKerusakanSpec) (bool, error)
	DeleteKerusakan(id int) (bool, error)
}
