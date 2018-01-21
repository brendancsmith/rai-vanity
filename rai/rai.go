package rai

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/a-h/round"
	"github.com/frankh/rai"
	"github.com/frankh/rai/address"
)

// GenerateVanityAddress -
func GenerateVanityAddress(substring string, index int, estimatedIter int) (string, rai.Account, error) {
	c := make(chan string, 100)
	progress := make(chan int, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(c chan string, progress chan int) {
			defer func() {
				recover()
			}()
			num := 0
			for {
				num++
				if num%(500+i) == 0 {
					progress <- num
					num = 0
				}
				seedBytes := make([]byte, 32)
				rand.Read(seedBytes)
				seed := hex.EncodeToString(seedBytes)
				pub, _ := address.KeypairFromSeed(seed, 0)
				address := string(address.PubKeyToAddress(pub))

				if substring == address[index:index+len(substring)] {
					c <- seed
					break
				}
			}
		}(c, progress)
	}

	go func(progress chan int) {
		total := 0
		fmt.Println()
		for {
			num, ok := <-progress
			if !ok {
				break
			}
			total += num
			percent := float64(total) / float64(estimatedIter) * 100
			fmt.Printf("\033[1A\033[KTried %d (~%.2f%%)\n", total, percent)
		}
	}(progress)

	seed := <-c
	pub, _ := address.KeypairFromSeed(seed, 0)
	account := address.PubKeyToAddress(pub)
	if !address.ValidateAddress(account) {
		return "", "", fmt.Errorf("Address generated had an invalid checksum!\nPlease create an issue on github: https://github.com/frankh/rai-vanity")
	}

	close(c)
	close(progress)

	return seed, account, nil
}

// EstimateIterations -
func EstimateIterations(substring string, index int) int {
	// single iteration needed for empty match
	searchSpace := 1

	digitsRemaining := len(substring)
	if digitsRemaining == 0 {
		return searchSpace
	}

	// no iterations needed for "xrb_"
	if index < 4 {
		digitsRemaining = int(math.Max(float64(digitsRemaining-(4-index)), 0))
		if digitsRemaining == 0 {
			return searchSpace
		}
	}

	// 2 possibilities for the first address digit (1 or 3)
	if index <= 4 {
		searchSpace = searchSpace * 2
		digitsRemaining = digitsRemaining - 1
		if digitsRemaining == 0 {
			return searchSpace
		}
	}

	// charset of 32 for the rest of the digits
	searchSpace = searchSpace * int(round.AwayFromZero(math.Pow(32, float64(digitsRemaining)), 0))

	// discovery time is assumed to be the mean
	return searchSpace / 2
}

func validateSegment(segment string, index int) error {
	length := len(segment)

	if length == 0 {
		return errors.New("segment length must be nonzero")
	} else if index+length >= 63 {
		return fmt.Errorf("segment exceeds address bounds (%d + %d >= 63)", index, length)
	}

	if index < 4 {
		lenPrefixOverlap := int(math.Min(float64(4-index), float64(length)))
		prefixOverlap := segment[:lenPrefixOverlap]
		prefixOverlapReference := "xrb_"[index : index+lenPrefixOverlap]
		if prefixOverlap != prefixOverlapReference {
			return fmt.Errorf(
				"Segment %s at index %d should start with %s",
				segment,
				index,
				prefixOverlapReference)
		}
	}

	if index < 5 && index+length >= 5 {
		firstDigit := segment[4-index]
		if firstDigit != '1' && firstDigit != '3' {
			return fmt.Errorf("First address digit must be '1' or '3'. %c is invalid", firstDigit)
		}
	}

	if index+length >= 6 {
		for _, c := range segment[5-index:] {
			if !strings.Contains(address.EncodeXrb, string(c)) {
				return fmt.Errorf("Character %c is invalid for a RaiBlocks address", c)
			}
		}
	}

	return nil
}
