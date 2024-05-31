package hashmap

type Hasher interface {
	Hash(string) int
}

func DefaultHasher(arraySize int) Hasher {
	return hasher{arraySize: arraySize}
}

type hasher struct {
	arraySize int
}

func (h hasher) Hash(data string) int {
	var sum int
	for _, c := range data {
		sum += int(c)
	}
	return sum % h.arraySize
}
