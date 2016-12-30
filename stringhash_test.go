package stringhash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hashed := New("https://images03.fasad.eu/228/400151/778787/large/5843669.jpg")
	fmt.Printf("%v\n", hashed)
	hashed = New("/dev/gocode/src/github.com/mysza/stringhash")
	fmt.Printf("%v\n", hashed)
	hashed = New("mysza")
	fmt.Printf("%v\n", hashed)
	hashed = New("mYsZa")
	fmt.Printf("%v\n", hashed)
	hashed = New("")
	fmt.Printf("%v\n", hashed)
	hashed = New("a a")
	fmt.Printf("%v\n", hashed)
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("https://example.website.com/long/path/to/file/filename.jpg")
	}
}
