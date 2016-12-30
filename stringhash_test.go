package stringhash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hashed := Hash("https://images03.fasad.eu/228/400151/778787/large/5843669.jpg")
	fmt.Printf("%v\n", hashed)
	hashed = Hash("/dev/gocode/src/github.com/mysza/stringhash")
	fmt.Printf("%v\n", hashed)
	hashed = Hash("mysza")
	fmt.Printf("%v\n", hashed)
	hashed = Hash("mYsZa")
	fmt.Printf("%v\n", hashed)
	hashed = Hash("")
	fmt.Printf("%v\n", hashed)
	hashed = Hash("a a")
	fmt.Printf("%v\n", hashed)
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Hash("https://example.website.com/long/path/to/file/filename.jpg")
	}
}
