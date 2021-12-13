package teknisi

import (
	"master-data/business"
)

type Service interface {
	InsertTeknisi(teknisiSpec Teknisi) error
	FindTeknisiByID(id int) (*Teknisi, error)
	FindAllTeknisi() ([]Teknisi, error)
	UpdateTeknisi(updateSpec TeknisiUpdateSpec) (bool, error)
	AddAntrian(id int, currentVersion int) (bool, error)
	EraseAntrian(id int, currentVersion int) (bool, error)
	DeleteTeknisi(id int) (bool, error)
}

type service struct {
	teknisiRepository Repository
}

func NewService(teknisiRepository Repository) Service {

	return &service{
		teknisiRepository,
	}
}

func (s *service) InsertTeknisi(teknisiSpec Teknisi) error {

	teknisi := NewTeknisi(
		teknisiSpec.ID,
		teknisiSpec.FullName,
		teknisiSpec.Specialist,
		teknisiSpec.Platform,
		teknisiSpec.JumlahAntrian,
		teknisiSpec.Version)

	err := s.teknisiRepository.InsertTeknisi(teknisi)
	if err != nil && err != business.ErrNotFound {
		return err
	}

	return nil
}

func (s *service) FindTeknisiByID(id int) (*Teknisi, error) {

	return s.teknisiRepository.FindTeknisiByID(id)
}

func (s *service) FindAllTeknisi() ([]Teknisi, error) {

	return s.teknisiRepository.FindAllTeknisi()
}

func (s *service) UpdateTeknisi(updateSpec TeknisiUpdateSpec) (bool, error) {
	existingTeknisi, err := s.FindTeknisiByID(updateSpec.ID)
	if err != nil {
		return false, err
	} else if existingTeknisi == nil {
		return false, business.ErrNotFound
	} else if existingTeknisi.Version != updateSpec.Version {
		return false, business.ErrHasBeenModified
	}

	updateTeknisi := NewTeknisiUpdate(
		updateSpec.ID,
		updateSpec.FullName,
		updateSpec.Specialist,
		updateSpec.Platform,
		updateSpec.Version)

	result, err := s.teknisiRepository.UpdateTeknisi(updateTeknisi)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *service) AddAntrian(id int, currentVersion int) (bool, error) {
	existingTeknisi, err := s.FindTeknisiByID(id)

	if err != nil {
		return false, err
	} else if existingTeknisi == nil {
		return false, business.ErrNotFound
	} else if existingTeknisi.Version != currentVersion {
		return false, business.ErrHasBeenModified
	}

	jumlahAntrian := existingTeknisi.JumlahAntrian + 1

	result, err := s.teknisiRepository.UpdateAntrian(id, jumlahAntrian, currentVersion)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *service) EraseAntrian(id int, currentVersion int) (bool, error) {
	existingTeknisi, err := s.FindTeknisiByID(id)

	if err != nil {
		return false, err
	} else if existingTeknisi == nil {
		return false, business.ErrNotFound
	} else if existingTeknisi.Version != currentVersion {
		return false, business.ErrHasBeenModified
	}

	jumlahAntrian := existingTeknisi.JumlahAntrian - 1

	result, err := s.teknisiRepository.UpdateAntrian(id, jumlahAntrian, currentVersion)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *service) DeleteTeknisi(id int) (bool, error) {
	existingTeknisi, err := s.FindTeknisiByID(id)
	if err != nil {
		return false, err
	} else if existingTeknisi == nil {
		return false, business.ErrNotFound
	}

	isDeleted, err := s.teknisiRepository.DeleteTeknisi(id)
	if err != nil || !isDeleted {
		return false, err
	}

	return isDeleted, nil
}
