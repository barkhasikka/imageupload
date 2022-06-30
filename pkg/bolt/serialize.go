package bolthelper

import (
	"imageupload/pkg/utils/conversion"

	"github.com/boltdb/bolt"
)

func DBSerializeString(bucket *bolt.Bucket, key, value string) {
	bucket.Put([]byte(key), []byte(value))
}

func DBSerializeBool(bucket *bolt.Bucket, key string, value bool) {
	valueByte := conversion.BinBool(value)
	bucket.Put([]byte(key), valueByte)
}

func DBSerializeUint16(bucket *bolt.Bucket, key string, value uint16) {
	valueByte := conversion.Uitob16(value)
	bucket.Put([]byte(key), valueByte)
}

func DBSerializeUint64(bucket *bolt.Bucket, key string, value uint64) {
	valueByte := conversion.Uitob64(value)
	bucket.Put([]byte(key), valueByte)
}

func DBSerializeInt(bucket *bolt.Bucket, key string, value int) {
	valueByte := conversion.Itob(int(value))
	bucket.Put([]byte(key), valueByte)
}
