package main

type DBCleaner struct {}

func NewDBCleaner() *DBCleaner {
	return &DBCleaner{}
}

func (c *DBCleaner) Clean() error {
	// handle DB Clean
	return nil
}