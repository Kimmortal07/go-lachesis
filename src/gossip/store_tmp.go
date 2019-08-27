package gossip

import (
	"fmt"
	"sync"

	"github.com/Fantom-foundation/go-lachesis/src/common/bigendian"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb"
)

type (
	tmpDb struct {
		Db     kvdb.KeyValueStore
		Tables interface{}
	}

	tmpDbs struct {
		min map[string]uint64
		seq map[string]map[uint64]tmpDb

		sync.Mutex
	}
)

func (s *Store) initTmpDbs() {
	s.tmpDbs.min = make(map[string]uint64)
	s.tmpDbs.seq = make(map[string]map[uint64]tmpDb)

	// load mins
	it := s.table.TmpDbs.NewIterator()
	for it.Next() {
		min := bigendian.BytesToInt64(it.Value())
		s.tmpDbs.min[string(it.Key())] = min
	}
	if it.Error() != nil {
		s.Fatal(it.Error())
	}
	it.Release()
}

func (s *Store) getTmpDb(name string, ver uint64, makeTables func(kvdb.KeyValueStore) interface{}) interface{} {
	s.tmpDbs.Lock()
	defer s.tmpDbs.Unlock()

	if min, ok := s.tmpDbs.min[name]; !ok {
		s.tmpDbs.min[name] = ver
		s.tmpDbs.seq[name] = make(map[uint64]tmpDb)
		err := s.table.TmpDbs.Put([]byte(name), bigendian.Int64ToBytes(ver))
		if err != nil {
			s.Fatal(err)
		}
	} else if ver < min {
		return nil
	}

	if tmp, ok := s.tmpDbs.seq[name][ver]; ok {
		return tmp.Tables
	}

	db := s.makeDb(tmpDbName(name, ver))
	tables := makeTables(db)
	s.tmpDbs.seq[name][ver] = tmpDb{
		Db:     db,
		Tables: tables,
	}

	return tables
}

func (s *Store) delTmpDb(name string, ver uint64) {
	s.tmpDbs.Lock()
	defer s.tmpDbs.Unlock()

	min, ok := s.tmpDbs.min[name]
	if !ok {
		return
	}

	for i := min; i <= ver; i++ {
		tmp := s.tmpDbs.seq[name][i]
		if tmp.Db != nil {
			tmp.Db.Close()
			tmp.Db.Drop()
		}
		delete(s.tmpDbs.seq[name], i)
	}

	ver += 1
	s.set(s.table.TmpDbs, []byte(name), &ver)
}

func tmpDbName(scope string, ver uint64) string {
	return fmt.Sprintf("tmp_%s_%d", scope, ver)
}