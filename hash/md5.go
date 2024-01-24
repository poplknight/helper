package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5[T, V ~[]byte | ~string](v T) V {
	hash := md5.New()
	hash.Write([]byte(v))
	return V(fmt.Sprintf("%x", hash.Sum(nil)))
}
