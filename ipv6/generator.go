package ipv6

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
)

func GenerateRandomIPv6(prefix string, mask_bits int, verbose bool) string {
	fmt.Println("> Generate random ipv6:")
	seg_num := (128 - mask_bits) / 16
	var segs []string
	for i := 0; i < seg_num; i++ {
		// generate random ints in 0-65535
		var random_int uint16
		binary.Read(rand.Reader, binary.BigEndian, &random_int)
		// convert to hex string
		random_hex := fmt.Sprintf("%x", random_int)
		segs = append(segs, random_hex)
	}
	suffix := strings.Join(segs, ":")
	random_ipv6 := prefix + ":" + suffix
	if verbose {
		fmt.Println("  *", random_ipv6)
	}
	return random_ipv6
}
