package leveldb_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	keyPrefix = []byte("key_")
	valPrefix = []byte("val_")
)

var db *leveldb.DB

func init() {
	d, err := leveldb.OpenFile("./level_test.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	db = d
}

func generateKV() ([]byte, []byte) {
	rnd := strconv.FormatUint(rand.Uint64(), 10)
	return append(keyPrefix, rnd...), append(valPrefix, rnd...)
}

func TestLevelDB(t *testing.T) {

	rnd := strconv.FormatUint(rand.Uint64(), 10)
	//val2 := make([]byte, 12)
	//binary.BigEndian.PutUint64(rnd, rand.Uint64())
	//base64.URLEncoding.Encode(val2, rnd)

	key := append(keyPrefix, rnd...)
	newVal := append(valPrefix, rnd...)

	val, err := db.Get(key, nil)
	assert.NotEqual(t, nil, err)

	err = db.Put(key, newVal, nil)
	if err != nil {
		t.Fatal(err)
	}

	val, err = db.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("key = %s, new  = %s\n", key, val)
}

func TestLevelI(t *testing.T) {
	it := db.NewIterator(nil, nil)
	for it.Next() {
		log.Printf("%s\n", it.Key())
	}
}

func BenchmarkLevelDB_Put(b *testing.B) {
	key, val := generateKV()
	for i := 0; i < b.N; i++ {
		err := db.Put(key, val, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLevelDB_Get(b *testing.B) {
	key, _ := generateKV()
	for i := 0; i < b.N; i++ {
		_, err := db.Get(key, nil)
		if err != nil {
			//b.Fatal(err)
		}
	}
}
