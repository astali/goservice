package lib

import (
	"crypto/md5"
	"fmt"
)

func MD5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
