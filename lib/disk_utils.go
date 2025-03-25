package lib

import (
	"log/slog"
	"strconv"
	"strings"
)

type DiskBlock struct {
	ID    int64
	Value string
	Size  int64
}

func NewBlockMap(s string) (string, []DiskBlock) {
	id := 0 // block ID
	dm := []string{}
	blx := []DiskBlock{}
	for i, s0 := range strings.Split(s, "") {
		v, _ := strconv.Atoi(s0)
		var buff []string
		if i%2 == 0 { // represents a disk block
			for j := 0; j < v; j++ {
				buff = append(buff, strconv.Itoa(id))
				blx = append(blx, DiskBlock{
					ID:    int64(id),
					Value: s0,
					Size:  int64(v),
				})
			}
			id++ // increment our block id
		} else { // represents free space
			for j := 0; j < v; j++ {
				buff = append(buff, ".")
				blx = append(blx, DiskBlock{
					ID:    int64(0),
					Value: ".",
					Size:  int64(v),
				})
			}
		}
		dm = append(dm, strings.Join(buff, ""))
	}
	return strings.Join(dm, ""), blx
}

func Defragment(blx []DiskBlock) []DiskBlock {
	slog.Debug("Defragment", "input", blx)

	lst := make([]DiskBlock, len(blx))
	copy(lst, blx)

	i := 0

	for j := len(lst) - 1; j >= 0; j-- {
		if lst[j].Value == "." {
			slog.Debug("found a dot, continuing", "j", j)
			continue // keep going until we find a non-dot character
		}
		// at this point, j points to the right-most non-dot character

		for i = 0; i < j && lst[i].Value != "."; i++ {
			slog.Debug("Defragment", "lst", lst, "j", j, "i", i)
		}
		// at this point, i points to the left-most dot character

		if i == j { // we met in the middle
			slog.Debug("i and j have converged", "i", i, "j", j)
			break
		}

		if lst[i].Value == "." && lst[j].Value != "." {
			buf := lst[i]
			lst[i] = lst[j]
			lst[j] = buf
		}
	}

	return lst
}

func BlockChecksum(lst []DiskBlock) (sum int64) {
	for i, s := range lst {
		sum += int64(i) * s.ID
	}
	return
}
