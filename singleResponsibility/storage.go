package document

import (
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func NewDBStorage(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) Save(doc *Document) error {
	return nil
}

type File struct {
	path string
}

func NewFileStorage(path string) *File {
	return &File{
		path: path,
	}
}

func (file *File) Save(doc *Document) error {
	return nil
}