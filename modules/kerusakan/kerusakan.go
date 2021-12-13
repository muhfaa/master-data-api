package kerusakan

import (
	"database/sql"
	"errors"
	"master-data/business"
	kerusakanBusiness "master-data/business/kerusakan"
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

// InsertKerusakan
func (repo *MySQL) InsertKerusakan(kerusakanSpec kerusakanBusiness.Kerusakan) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}
	insertQuery := `INSERT INTO kerusakan (
		jenis_kerusakan,
		lama_pengerjaan,
		harga,
		version
	)
	VALUES
	(?,?,?,?)`

	_, err = tx.Exec(insertQuery, kerusakanSpec.JenisKerusakan, kerusakanSpec.LamaPengerjaan, kerusakanSpec.Harga, 1)
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

// FindKerusakanByID
func (repo *MySQL) FindKerusakanByID(id int) (*kerusakanBusiness.Kerusakan, error) {
	var kerusakan kerusakanBusiness.Kerusakan

	selectQuery := `SELECT * FROM kerusakan WHERE id = ?`

	err := repo.db.QueryRowx(selectQuery, id).StructScan(&kerusakan)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, business.ErrNotFound
		}
		err = errors.New("resource error")
		return nil, err
	}

	return &kerusakan, nil
}

// FindAllKerusakan
func (repo *MySQL) FindAllKerusakan() ([]kerusakanBusiness.Kerusakan, error) {
	var kerusakan kerusakanBusiness.Kerusakan
	var allKerusakan []kerusakanBusiness.Kerusakan

	selectQuery := `SELECT * FROM kerusakan`

	row, err := repo.db.Queryx(selectQuery)
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return nil, err
	}

	for row.Next() {
		err = row.StructScan(&kerusakan)
		if err != nil {
			log.Error(err)
			err = errors.New("resource error")
			return nil, err
		}

		allKerusakan = append(allKerusakan, kerusakan)
	}
	return allKerusakan, nil
}

// UpdateKerusakan
func (repo *MySQL) UpdateKerusakan(updateSpec kerusakanBusiness.UpdateKerusakanSpec) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE kerusakan 
		SET
		jenis_kerusakan = ?,
		lama_pengerjaan = ?,
		harga = ?,
		version = ?
		WHERE
		id = ?`

	_, err = tx.Exec(insertQuery, updateSpec.JenisKerusakan, updateSpec.LamaPengerjaan, updateSpec.Harga, updateSpec.Version+1, updateSpec.ID)
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

// DeleteKerusakan
func (repo *MySQL) DeleteKerusakan(id int) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	deleteQuery := `DELETE FROM kerusakan WHERE id = ?`

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
