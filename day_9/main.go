package main

import (
	. "AdventOfCode2024/lib"
	"bufio"
	"flag"
	"log/slog"
	"os"
)

func puzzle1(l string) int64 {
	_, dbm := NewBlockMap(l)
	//defragmented := DefragmentRecursive(dbm)
	slog.Info("block map: %v", dbm)
	defragmented := Defragment(dbm)
	return BlockChecksum(defragmented)
}

func puzzle2() int64 {

	return int64(2)
}

func main() {
	flag.BoolFunc("debug", "enable debug logging", func(s string) (err error) {
		slog.SetLogLoggerLevel(slog.Level(slog.LevelDebug))
		return
	})

	flag.Parse()

	stdioScanner := bufio.NewScanner(os.Stdin)

	var line string

	for stdioScanner.Scan() {
		line = stdioScanner.Text()
	}
	slog.Debug("generated diskMap", "diskMap", line)

	sum := puzzle1(line)
	slog.Info("found puzzle1 sum", "sum", sum)

	sum = puzzle2()
	slog.Info("found puzzle2 sum", "sum", sum)
}
