package teknisi

import (
	"database/sql"
	"errors"
	"master-data/business"
	teknisiBusiness "master-data/business/teknisi"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type MySQL struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQL {
	return &MySQL{
		db,
	}
}

// InsertTeknisi
func (repo *MySQL) InsertTeknisi(teknisiSpec teknisiBusiness.Teknisi) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}
	insertQuery := `INSERT INTO teknisi (
		full_name,
		specialist,
		platform,
		jumlah_antrian,
		version
	)
	VALUES
	(?,?,?,?,?)`

	_, err = tx.Exec(insertQuery, teknisiSpec.FullName, teknisiSpec.Specialist, teknisiSpec.Platform, teknisiSpec.JumlahAntrian, 1)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return err
		}

		err = errors.New("resource error")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	return nil
}

// FindTeknisiByID
func (repo *MySQL) FindTeknisiByID(id int) (*teknisiBusiness.Teknisi, error) {
	var teknisi teknisiBusiness.Teknisi

	selectQuery := `SELECT * FROM teknisi WHERE id = ?`

	err := repo.db.QueryRowx(selectQuery, id).StructScan(&teknisi)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, business.ErrNotFound
		}
		err = errors.New("resource error")
		return nil, err
	}

	return &teknisi, nil
}

// FindAllTeknisi
func (repo *MySQL) FindAllTeknisi() ([]teknisiBusiness.Teknisi, error) {
	var teknisi teknisiBusiness.Teknisi
	var allTeknisi []teknisiBusiness.Teknisi

	selectQuery := `SELECT * FROM teknisi`

	row, err := repo.db.Queryx(selectQuery)
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return nil, err
	}

	for row.Next() {
		err = row.StructScan(&teknisi)
		if err != nil {
			log.Error(err)
			err = errors.New("resource error")
			return nil, err
		}

		allTeknisi = append(allTeknisi, teknisi)
	}
	return allTeknisi, nil
}

// UpdateTeknisi
func (repo *MySQL) UpdateTeknisi(updateSpec teknisiBusiness.TeknisiUpdateSpec) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE teknisi 
		SET
		full_name = ?,
		specialist = ?,
		platform = ?,
		version = ?
		WHERE
		id = ?`

	_, err = tx.Exec(insertQuery, updateSpec.FullName, updateSpec.Specialist, updateSpec.Platform, updateSpec.Version+1, updateSpec.ID)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return false, err
		}

		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}

// UpdateAntrian
func (repo *MySQL) UpdateAntrian(id int, jumlahAntrian int, currentVersion int) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE teknisi 
		SET
		jumlah_antrian = ?,
		version = ?
		WHERE
		id = ?`

	_, err = tx.Exec(insertQuery, jumlahAntrian, currentVersion+1, id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return false, err
		}

		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}

// DeleteTeknisi
func (repo *MySQL) DeleteTeknisi(id int) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	deleteQuery := `DELETE FROM teknisi WHERE id = ?`

	_, err = tx.Exec(deleteQuery, id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}
