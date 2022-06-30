package bolthelper

import (
	"imageupload/pkg/utils/conversion"

	"github.com/boltdb/bolt"
)

func DBDeSerializeString(bucket *bolt.Bucket, key string) string {
	return string(bucket.Get([]byte(key)))
}

func DBDeSerializeBool(bucket *bolt.Bucket, key string) bool {
	return conversion.BoolInB(bucket.Get([]byte(key)))
}

func DBDeSerializeInt(bucket *bolt.Bucket, key string) int {
	return conversion.BtoI(bucket.Get([]byte(key)))
}

func DBDeSerializeUint16(bucket *bolt.Bucket, key string) uint16 {
	return conversion.B16toi(bucket.Get([]byte(key)))
}

func DBDeSerializeUint64(bucket *bolt.Bucket, key string) uint64 {
	return conversion.Btoi64(bucket.Get([]byte(key)))
}

func DBDeSerializeUint(bucket *bolt.Bucket, key string) uint {
	return uint(conversion.B16toi(bucket.Get([]byte(key))))
}

func DBDeSerializeBytes(bucket *bolt.Bucket, key string) []byte {
	return bucket.Get([]byte(key))
}
