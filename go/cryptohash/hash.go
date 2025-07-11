package cryptohash

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Notes
For understanding colision resistance
colision resistance - computationally infeasible to find any 2 distinct inputs for xy such that hash(x) == hash(y)
*/

/*
Further research - prove the birthday paradox
looking at n bits, you’re solving a much smaller puzzle—one with an effective space of 2^n rather than 2^{256}. By the birthday paradox, you expect a collision after roughly \sqrt{2^n}=2^{n/2} trials, so even for n=32 you only need on the order of 2^{16}\approx65{,}536 hashes.
*/

type Birthday struct {
	day   int
	month string
}

func FindColission(nBits int) (a, b []byte, err error) {
	// determine realistic amount of bits to check for colission

	// seen := map[string]byte{}

	distinctInput := encodedCounter(nBits)

	// so testing 2^256 bits is a bad idea my guy. Instead, what if we had a mathematical solution to test a little... BIT of it instead?
	tinyWindowSnooper(distinctInput, 24)

	// for i := range nBits {
	// 	fmt.Println(i)

	// 	// digest := sha256.Sum256(distinctInput)

	// 	// prefix := extractingFirstNBits(digest, nBits)
	// 	// key := string(prefix)
	// 	// if other, ok := seen[key]; ok {
	// 	// 	return other, data, nil
	// 	// }
	// 	// seen[key] = data

	// }

	a = []byte{}
	b = []byte{}

	return a, b, err
}

func encodedCounter(myCoolAssInt int) []byte {
	// I take me cool ass int and then put that into 8 bytes
	// Nous faisons cela pour ne pas casser notre tout petit ordinateur
	if myCoolAssInt == 0 {
		return []byte{0}
	}

	/*
	   What it does:
	   Takes the binary representation of num
	   Shifts all bits 8 positions to the right
	   Discards the rightmost 8 bits
	   Assigns the result back to num

	   num := uint64(0x12345678)  // Binary: 00010010001101000101011001111000
	   num >>= 8                  // Binary: 00000000000100100011010001010110
	   // num is now 0x00123456
	*/

	var result []byte
	num := uint64(myCoolAssInt)
	for num > 0 {
		result = append([]byte{byte(num & 0xFF)}, result...)
		num >>= 8 // This shifts right by 8 bits to process the next byte
	}
	fmt.Printf("RESULT: %d \n", result)
	return result
}

func tinyWindowSnooper(digest []byte, nBits int) ([]byte, error) {
	fmt.Printf("The length of digest is %d\n", len(digest))
	if nBits < 0 || nBits > 256 {
		return nil, fmt.Errorf("nBits must be between 0 and 256, got %d", nBits)
	}
	// checkSum := sha256.Sum256(digest)
	// spoonin up des bowls-o-bytes

	fullBytesNeeded := nBits / 8
	remBits := nBits % 8

	spoon := make([]byte, fullBytesNeeded)
	copy(spoon, digest[0:fullBytesNeeded])

	if remBits > 0 {
		mask := byte(0xFF << (8 - remBits))
		next := digest[fullBytesNeeded] & mask
		spoon = append(spoon, next)
	}
	fmt.Println("A spoon full of sugar is enough to know si c'est merde ", spoon)
	return spoon, nil
}

func BirthdayParadoxProof(groupSize int, simulations int) (float64, error) {
	// In a group of 23 people there's a 50% chance that at least 2 people share a birthday
	// Number of comparisons = n(n-1)/2
	// 23(23-1)/2 = 253

	if groupSize <= 0 || simulations <= 0 {
		return 0, fmt.Errorf("groupSize and simulations must be positive integers")
	}

	// Days in each month
	daysInMonth := []int{
		31, // Jan
		28, // Feb
		31, // Mar
		30, // Apr
		31, // May
		30, // Jun
		31, // Jul
		31, // Aug
		30, // Sep
		31, // Oct
		30, // Nov
		31, // Dec
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	countWithSharedBirthday := 0

	for i := 0; i < simulations; i++ {
		birthdaysSeen := map[string]int{}
		for j := 0; j < groupSize; j++ {
			month := r.Intn(12) + 1                 // 1-12
			day := r.Intn(daysInMonth[month-1]) + 1 // 1 to max days in that month
			birthdayAsString := fmt.Sprintf("%02d-%02d", day, month)

			birthdaysSeen[birthdayAsString] += 1
			if birthdaysSeen[birthdayAsString] > 1 {
				countWithSharedBirthday++
				break
			}
		}
	}

	return float64(countWithSharedBirthday) / float64(simulations), nil
}
