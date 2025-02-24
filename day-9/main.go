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
	data := generateBlock(r)
	diskLayout := diskLayout(data)

	fmt.Println(resetOrder(diskLayout))
	return nil
}

// PartTwo solves the second problem of day 8 of Advent of Code 2024.
func PartTwo(r io.Reader, w io.Writer) error {

	return nil
}

func generateBlock(r io.Reader) []int {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil
	}

	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}

	diskMap := make([]int, len(data))
	for i, b := range data {
		if b < '0' || b > '9' {
			return nil
		}
		diskMap[i] = int(b - '0')
	}

	return diskMap
}

func diskLayout(diskMap []int) []int {
	diskSize := 0
	for _, n := range diskMap {
		diskSize += n
	}

	layout := make([]int, diskSize)

	position := 0
	fileID := 0
	isFile := true
	for _, v := range diskMap {
		if isFile {
			for range v {
				layout[position] = fileID
				position++
			}
			fileID++
			isFile = false
		} else {
			for range v {
				layout[position] = -1
				position++
			}
			isFile = true
		}
	}

	return layout
}

func resetOrder(l []int) int {
	res := 0

	left := 0
	right := len(l) - 1

	for left < right {
		if l[right] == -1 {
			right--
			continue
		}

		if l[left] == -1 {
			l[left] = l[right]
			l[right] = -1
			right--
			continue
		}

		left++
	}

	for i, v := range l {
		if v == -1 {
			continue
		}
		res += i * v
	}

	return res
}
