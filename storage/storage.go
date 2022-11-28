package storage

import (
	"github.com/jmoiron/sqlx"
)

type StorageI interface {

}

type storagePg struct {

}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{ }
}