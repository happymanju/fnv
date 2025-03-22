package cli

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/happymanju/fnv/lib"
)

func Run(args []string) int {
	hash, err := lib.FNV(args[0])
	if err != nil {
		log.Printf("%v/n", err)
		return 1
	}

	asHex := hex.EncodeToString(hash)
	fmt.Println(asHex)
	return 0
}
