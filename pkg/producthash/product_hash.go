package producthash

import (
	"fmt"

	"github.com/twmb/murmur3"
)

func HashProduct(productTitle string, produtPrice float32, stockQuantity uint16) uint32 {

	hasher := murmur3.New32()
	chunk1 := []byte(productTitle)
	chunk2 := []byte(fmt.Sprintf("%f", produtPrice))
	chunk3 := []byte(fmt.Sprintf("%d", stockQuantity))
	hasher.Write(chunk1)
	hasher.Write(chunk2)
	hasher.Write(chunk3)
	streamingHash := hasher.Sum32()

	return streamingHash

}
