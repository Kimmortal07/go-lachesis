package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Fantom-foundation/go-lachesis/src/gossip"
	"github.com/Fantom-foundation/go-lachesis/src/internal"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb/leveldb"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb/memorydb"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb/table"
	"github.com/Fantom-foundation/go-lachesis/src/metrics"
	"github.com/Fantom-foundation/go-lachesis/src/poset"
)

func makeStorages(makeDb internal.DbProducer) (*gossip.Store, *poset.Store) {
	db := makeDb("lachesis")

	g := table.New(db, []byte("g_"))
	p := table.New(db, []byte("p_"))

	return gossip.NewStore(g, makeDb),
		poset.NewStore(p, makeDb)
}

func dbProducer(dbdir string) internal.DbProducer {
	if dbdir == "inmemory" {
		return func(name string) kvdb.KeyValueStore {
			return memorydb.New()
		}
	}

	return func(name string) kvdb.KeyValueStore {
		db, err := openDb(dbdir, name)
		if err != nil {
			panic(err)
		}

		return db
	}
}

func openDb(dir, name string) (
	db kvdb.KeyValueStore,
	err error,
) {
	err = os.MkdirAll(dir, 0600)
	if err != nil {
		return
	}

	f := filepath.Join(dir, name+"-ldb")

	stopWatcher := metrics.StartFileWatcher(name+"_db_file_size", f)

	onClose := func() error {
		stopWatcher()
		return nil
	}
	onDrop := func() error {
		return os.Remove(f)
	}

	db, err = leveldb.New(f, 16, 0, "", onClose, onDrop)
	if err != nil {
		panic(fmt.Sprintf("can't create temporary database: %v", err))
	}

	return
}