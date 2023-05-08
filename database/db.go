package database

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type DB struct {
	logger   *log.Logger
	BadgerDB *badger.DB
}

func NewDatabase(log *log.Logger) *DB {
	badgerOptions := badger.DefaultOptions("tmp/badger")
	badgerOptions.Logger = nil //TODO implement a custom logger
	db, err := badger.Open(badgerOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{BadgerDB: db, logger: log}
}

func (db *DB) WriteRecord(duration int) {
	db.logger.Println("Recording score:", duration)
	err := db.BadgerDB.Update(func(txn *badger.Txn) error {
		error := txn.Set([]byte(fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))), []byte(fmt.Sprintf("%d", duration)))
		return error
	})

	if err != nil {
		db.logger.Println(err)
	}
}

func (db *DB) ShowScores() {
	err := db.BadgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("%s : %s ms\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		db.logger.Println(err)
	}
}
