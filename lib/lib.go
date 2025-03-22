package lib

import (
	"bufio"
	"hash/fnv"
	"io"
	"os"
)

func FNV(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil && os.IsNotExist(err) {
		return nil, err
	}
	defer f.Close()

	br := bufio.NewReader(f)

	hash := fnv.New64a()

	for {
		buf := make([]byte, 4096)

		n, err := br.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n != 0 {
			hash.Write(buf[:n+1])
		} else {
			break
		}
	}
	return hash.Sum(nil), nil
}
