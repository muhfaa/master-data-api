package kerusakan

import (
	"master-data/business"
)

type Service interface {
	InsertKerusakan(kerusakanSpec Kerusakan) error
	FindKerusakanByID(id int) (*Kerusakan, error)
	FindAllKerusakan() ([]Kerusakan, error)
	UpdateKerusakan(updateSpec UpdateKerusakanSpec) (bool, error)
	DeleteKerusakan(id int) (bool, error)
}

type service struct {
	kerusakanRepository Repository
}

func NewService(kerusakanRepository Repository) Service {

	return &service{
		kerusakanRepository,
	}
}

func (s *service) InsertKerusakan(kerusakanSpec Kerusakan) error {

	kerusakan := NewKerusakan(
		kerusakanSpec.JenisKerusakan,
		kerusakanSpec.LamaPengerjaan,
		kerusakanSpec.Harga,
	)

	err := s.kerusakanRepository.InsertKerusakan(kerusakan)
	if err != nil && err != business.ErrNotFound {
		return err
	}

	return nil
}

func (s *service) FindKerusakanByID(id int) (*Kerusakan, error) {
	return s.kerusakanRepository.FindKerusakanByID(id)
}

func (s *service) FindAllKerusakan() ([]Kerusakan, error) {
	return s.kerusakanRepository.FindAllKerusakan()
}

func (s *service) UpdateKerusakan(updateSpec UpdateKerusakanSpec) (bool, error) {
	existingKerusakan, err := s.FindKerusakanByID(updateSpec.ID)
	if err != nil {
		return false, err
	} else if existingKerusakan == nil {
		return false, business.ErrNotFound
	} else if existingKerusakan.Version != updateSpec.Version {
		return false, business.ErrHasBeenModified
	}

	kerusakanUpdate := NewKerusakanUpdate(
		updateSpec.ID,
		updateSpec.JenisKerusakan,
		updateSpec.LamaPengerjaan,
		updateSpec.Harga,
		updateSpec.Version,
	)

	result, err := s.kerusakanRepository.UpdateKerusakan(kerusakanUpdate)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *service) DeleteKerusakan(id int) (bool, error) {
	existingKerusakan, err := s.FindKerusakanByID(id)

	if err != nil {
		return false, err
	} else if existingKerusakan == nil {
		return false, business.ErrNotFound
	}

	result, err := s.kerusakanRepository.DeleteKerusakan(id)
	if err != nil {
		return false, err
	}

	return result, nil
}
