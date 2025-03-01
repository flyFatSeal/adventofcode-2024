package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type diskLayout []int

const empty = -1

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
	diskMap, err := diskMapFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	layout := diskLayoutFromDiskMap(diskMap)
	layout.compact()
	checksum := layout.checksum()

	_, err = fmt.Fprintf(w, "%d", checksum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}
	return nil
}

// PartTwo solves the second problem of day 8 of Advent of Code 2024.
func PartTwo(r io.Reader, w io.Writer) error {

	diskMap, err := diskMapFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	layout := diskLayoutFromDiskMap(diskMap)
	layout.compact()
	checksum := layout.checksum()

	_, err = fmt.Fprintf(w, "%d", checksum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil

}

func (l diskLayout) String() string {
	var b bytes.Buffer
	for _, v := range l {
		if v == empty {
			b.WriteByte('.')
		} else {
			b.WriteByte(byte(v) + '0')
		}
	}
	return b.String()
}

func (l diskLayout) compact() {
	// Start with the highest file ID and try to move files left
	right := len(l) - 1
	var maxFileID int
	for ; right >= 0; right-- {
		if l[right] != empty {
			maxFileID = l[right]
			break
		}
	}

	// Try to move each file in decreasing order of ID
	for fileID := maxFileID; fileID >= 0; fileID-- {
		// Find where the file starts and its size
		fileStart, size, found := l.findFile(fileID, right+1)
		if !found {
			panic(fmt.Sprintf("missing file %d", fileID))
		}
		// Find space for the file
		emptyStart, found := l.findEmpty(size)
		if !found {
			continue
		}

		// If there is enough space, move the file
		if fileStart <= emptyStart {
			continue
		}

		l.moveFile(fileStart, emptyStart, size)
	}
}

// Move file from `from` position to `to` position
func (l diskLayout) moveFile(from, to, size int) {
	for i := 0; i < size; i++ {
		l[to+i] = l[from+i]
		l[from+i] = empty
	}
}

// Find the start and size of a file with `id` before position `before`
func (l diskLayout) findFile(id, before int) (int, int, bool) {
	last := -1
	for position := before - 1; ; position-- {
		if position < 0 {
			return 0, 0, false
		}
		if l[position] == id {
			last = position
			break
		}
	}
	if last == -1 {
		return 0, 0, false
	}
	for first := last - 1; first >= 0; first-- {
		if l[first] != id {
			return first + 1, last - first, true
		}
	}
	return 0, last, true
}

// Find empty space of given `size`
func (l diskLayout) findEmpty(size int) (int, bool) {
	sizeSoFar := 0
	for position := 0; position < len(l)-size; position++ {
		if l[position] == empty {
			sizeSoFar++
			if sizeSoFar == size {
				return position - size + 1, true
			}
		} else {
			sizeSoFar = 0
		}
	}
	return 0, false
}

// Checksum calculates the checksum for the disk layout
func (l diskLayout) checksum() int {
	checksum := 0
	for position, blockID := range l {
		if blockID == empty {
			continue
		}
		checksum += position * blockID
	}
	return checksum
}

// Convert disk map from input data into disk layout
func diskLayoutFromDiskMap(diskMap []int) diskLayout {
	diskSize := 0
	for _, n := range diskMap {
		diskSize += n
	}

	layout := make(diskLayout, diskSize)

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
				layout[position] = empty
				position++
			}
			isFile = true
		}
	}

	return layout
}

func diskMapFromReader(r io.Reader) ([]int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil, fmt.Errorf("no input")
	}

	diskMap := make([]int, len(data))
	for i, b := range data {
		if b < '0' || b > '9' {
			return nil, fmt.Errorf("expected a digit, got %q", b)
		}
		diskMap[i] = int(b - '0')
	}

	return diskMap, nil
}
