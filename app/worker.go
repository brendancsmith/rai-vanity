package app

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"runtime"

	"github.com/frankh/rai/address"
)

// StringPredicate -
type StringPredicate func(s string) bool

// GenerateSeeds -
func GenerateSeeds(ctx context.Context, addressPredicate StringPredicate) (<-chan string, error) {
	ch := make(chan string, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer func() {
				recover()
			}()

			select {
			case <-ctx.Done():
				return
			default:
				seedBytes := make([]byte, 32)
				rand.Read(seedBytes) // TODO: why this?
				seed := hex.EncodeToString(seedBytes)
				pub, _ := address.KeypairFromSeed(seed, 0)
				address := string(address.PubKeyToAddress(pub))

				if addressPredicate(address) {
					ch <- seed
				}

				// TODO: metrics
			}
		}()
	}

	return ch, nil
}
