package hashmap_test

import (
	"testing"

	"github.com/murtaza-u/lab/go/data-structures/hashmap"
)

func TestHash(t *testing.T) {
	h := hashmap.DefaultHasher(10)
	hash := h.Hash("Mark")
	if hash != 5 {
		t.Fatalf("want: 5 got: %d", hash)
	}
}
