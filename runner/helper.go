package runner

import (
	"bufio"
	"io"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([HeritageDataCount]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result [HeritageDataCount]int
	var i = 0
	for scanner.Scan() || i < HeritageDataCount {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result[i] = x
		i++
	}
	return result, scanner.Err()
}

func Abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}
