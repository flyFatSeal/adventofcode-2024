package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开 input.txt 文件
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 调用 PartOne 函数读取文件
	if err := PartOne(file, os.Stdout); err != nil {
		fmt.Println("Error in PartOne:", err)
		return
	}

	// 将文件指针重置为文件开头，准备处理 PartTwo
	_, err = file.Seek(0, io.SeekStart) // 重置文件指针到文件开头
	if err != nil {
		fmt.Println("Error resetting file pointer:", err)
		return
	}

	// 调用 PartTwo 函数读取文件
	if err := PartTwo(file, os.Stdout); err != nil {
		fmt.Println("Error in PartTwo:", err)
		return
	}
}

// PartOne solves the first problem of day 8 of Advent of Code 2024.
func PartOne(r io.Reader, w io.Writer) error {
	city, err := cityMapFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	count := countAntinodePositions(city, false)

	fmt.Println(count)

	return nil
}

// PartTwo solves the second problem of day 8 of Advent of Code 2024.
func PartTwo(r io.Reader, w io.Writer) error {
	city, err := cityMapFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	count := countAntinodePositions(city, true)

	_, err = fmt.Fprintf(w, "%d", count)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

type cityMap [][]byte

const empty = '.'

func countAntinodePositions(city cityMap, resonantHarmonics bool) int {
	nodeIndex := make(map[byte][]vector)
	for row := range city {
		for col, node := range city[row] {
			if node == empty {
				continue
			}
			position := vector{row, col}
			nodeIndex[node] = append(nodeIndex[node], position)
		}
	}

	positionHasAntinode := make(map[vector]struct{})
	for _, positions := range nodeIndex {
		for _, positionA := range positions {
			for _, positionB := range positions {
				if positionA == positionB {
					continue
				}

				delta := positionA.minus(positionB)

				if !resonantHarmonics {
					antinodePosition := positionA.plus(delta)
					if withinCityBounds(city, antinodePosition) {
						positionHasAntinode[antinodePosition] = struct{}{}
					}
					continue
				}

				antinodePosition := positionA
				for withinCityBounds(city, antinodePosition) {
					positionHasAntinode[antinodePosition] = struct{}{}
					antinodePosition = antinodePosition.plus(delta)
				}
			}
		}
	}

	return len(positionHasAntinode)
}

func withinCityBounds(city cityMap, position vector) bool {
	return position.row >= 0 && position.row < len(city) &&
		position.col >= 0 && position.col < len(city[position.row])
}

type vector struct {
	row, col int
}

func (v vector) plus(w vector) vector {
	return vector{
		row: v.row + w.row,
		col: v.col + w.col,
	}
}

func (v vector) minus(w vector) vector {
	return vector{
		row: v.row - w.row,
		col: v.col - w.col,
	}
}

func cityMapFromReader(r io.Reader) (cityMap, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)

	return bytes.Split(data, []byte("\n")), nil
}
