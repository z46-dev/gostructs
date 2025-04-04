package misc

import "github.com/z46-dev/go-logger"

// BitSet represents a bit set for the Bloom filter
// It is a slice of uint64 where each bit represents an element in the filter
// The size of the bit set is determined by the number of bits needed to represent the maximum number of elements in the filter

type BloomFilter[T any] struct {
	BitSet []uint64
	Hashes []func(T) int
}

func NewBloomFilter[T any](size int, hashes []func(T) int) *BloomFilter[T] {
	return &BloomFilter[T]{
		BitSet: make([]uint64, (size+63)/64),
		Hashes: hashes,
	}
}

func (bf *BloomFilter[T]) Add(value T) {
	for _, hash := range bf.Hashes {
		index := hash(value) % (len(bf.BitSet) * 64)
		bf.BitSet[index/64] |= 1 << (index % 64)
	}
}

func (bf *BloomFilter[T]) Contains(value T) bool {
	for _, hash := range bf.Hashes {
		index := hash(value) % (len(bf.BitSet) * 64)
		if bf.BitSet[index/64]&(1<<(index%64)) == 0 {
			return false
		}
	}

	return true
}

func (bf *BloomFilter[T]) Clear() {
	for i := range bf.BitSet {
		bf.BitSet[i] = 0
	}
}

func BloomFilterTests() {
	var hashes []func(string) int = []func(string) int{
		func(value string) int {
			return len(value) % 100
		},
		func(value string) (hash int) {
			for _, c := range value {
				hash += int(c)
			}

			return hash
		},
		func(value string) (hash int) {
			for _, c := range value {
				hash ^= int(c)
			}

			return hash
		},
	}

	var filter *BloomFilter[string] = NewBloomFilter(100, hashes)

	var wordsToAdd []string = []string{"hello", "world", "foo", "bar", "baz"}
	var wordsToNotAdd []string = []string{"qux", "quux", "corge", "grault", "garply"}

	var log *logger.Logger = logger.NewLogger().SetPrefix("[BloomFilter]", logger.BoldBlue)

	log.Basicf("Created a bloom filter with size %d (%d bits), and %d hash functions\n", len(filter.BitSet), len(filter.BitSet)*64, len(filter.Hashes))
	log.Basicf("Adding %d words to the filter\n", len(wordsToAdd))

	for _, word := range wordsToAdd {
		filter.Add(word)
		log.Statusf("Added %s to the filter\n", word)
	}

	log.Basicf("Checking if the filter contains %d words\n", len(wordsToAdd))
	for _, word := range wordsToAdd {
		if filter.Contains(word) {
			log.Successf("Filter contains %s\n", word)
		} else {
			log.Errorf("Filter does not contain %s\n", word)
		}
	}

	for _, word := range wordsToNotAdd {
		if filter.Contains(word) {
			log.Warningf("Filter contains %s, but it should no\nt", word)
		} else {
			log.Successf("Filter does not contain %s\n", word)
		}
	}
}
