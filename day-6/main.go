package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type obstacle struct {
	x      int
	y      int
	marked bool
}

type Guard struct {
	direct string
	x      int
	y      int
}

func main() {
	// 读取文件
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var area [][]string
	var obstacles = make(map[string]obstacle)
	var guard Guard

	rows := 0

	// Read file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for j, str := range line {
			if str == "#" {
				obstacles[fmt.Sprintf("%d,%d", j, rows)] = obstacle{x: j, y: rows, marked: false}
			}
			if str == "^" {
				guard = Guard{direct: "top", x: j, y: rows}
			}
		}
		area = append(area, line)
		rows++
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 调用 part1 函数处理文件内容
	result1 := part1(guard, obstacles, area)

	// 输出结果
	fmt.Println("Result of part1:", result1)

}

func part1(guard Guard, obstacles map[string]obstacle, area [][]string) int {
	res := make(map[string]int)
	for {
		if move(&guard, obstacles, len(area), len(area[0])) {
			if _, ok := res[fmt.Sprintf("%d,%d", guard.x, guard.y)]; !ok {
				res[fmt.Sprintf("%d,%d", guard.x, guard.y)] = 1
			}

		} else {
			break
		}
	}
	return len(res)
}

func move(guard *Guard, obstacles map[string]obstacle, maxX, maxY int) bool {
	// 检查越界条件
	if guard.x < 0 || guard.y < 0 || guard.x >= maxX || guard.y >= maxY {
		return false
	}

	// 方向处理
	switch guard.direct {
	case "top":
		if guard.y == 0 || checkObstacle(obstacles, guard.x, guard.y-1) {
			// 如果上面是障碍，转到右边
			if guard.x+1 < maxX && !checkObstacle(obstacles, guard.x+1, guard.y) {
				guard.x++
				guard.direct = "right"
				return true
			}
			return false
		}
		guard.y-- // 向上移动
		return true
	case "right":
		if guard.x+1 == maxX || checkObstacle(obstacles, guard.x+1, guard.y) {
			// 如果右边是障碍，转到底部
			if guard.y+1 < maxY && !checkObstacle(obstacles, guard.x, guard.y+1) {
				guard.y++
				guard.direct = "bottom"
				return true
			}
			return false
		}
		guard.x++ // 向右移动
		return true
	case "bottom":
		if guard.y+1 == maxY || checkObstacle(obstacles, guard.x, guard.y+1) {
			// 如果下面是障碍，转到左边
			if guard.x > 0 && !checkObstacle(obstacles, guard.x-1, guard.y) {
				guard.x--
				guard.direct = "left"
				return true
			}
			return false
		}
		guard.y++ // 向下移动
		return true
	case "left":
		if guard.x == 0 || checkObstacle(obstacles, guard.x-1, guard.y) {
			// 如果左边是障碍，转到顶部
			if guard.y > 0 && !checkObstacle(obstacles, guard.x, guard.y-1) {
				guard.y--
				guard.direct = "top"
				return true
			}
			return false
		}
		guard.x-- // 向左移动
		return true
	}
	return false
}

// 辅助函数检查是否有障碍物
func checkObstacle(obstacles map[string]obstacle, x, y int) bool {
	_, has := obstacles[fmt.Sprintf("%d,%d", x, y)]
	return has
}
