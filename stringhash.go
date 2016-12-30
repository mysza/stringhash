package stringhash

import "github.com/spaolacci/murmur3"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_"
	bits     = 6
	mask     = 0x3f
	max      = 11
)

func alphabetise(hash uint64) string {
	ret := make([]byte, max)
	for i := range ret {
		ret[i] = alphabet[int(hash&mask)]
		hash >>= bits
	}
	return string(ret)
}

// Hash creates an 11 letter hash of the given string.
func Hash(s string) string {
	return alphabetise(murmur3.Sum64([]byte(s)))
}
