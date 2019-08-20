package migong

import "fmt"

func Myset(mymap *[8][7]int, i, j int) bool {
	if mymap[6][5] == 2 {
		// 找到出路
		return true
	} else {
		if mymap[i][j] == 0 {
			// 先假设该点是个通路, 但还需继续探索
			mymap[i][j] = 2
			// 上下左右
			/*
				1 1 1 1 1 1 1
				1 2 2 2 2 2 1
				1 2 2 2 2 2 1
				1 1 1 2 2 2 1
				1 3 3 2 2 2 1
				1 3 3 2 2 2 1
				1 3 3 2 2 2 1
				1 1 1 1 1 1 1
			*/
			// 右下左上
			/*
				1 1 1 1 1 1 1
				1 3 3 3 3 3 1
				1 0 0 0 0 0 1
				1 1 1 0 0 0 1
				1 0 0 0 0 0 1
				1 0 0 0 0 0 1
				1 0 0 0 0 0 1
				1 1 1 1 1 1 1
			*/
			// 下右上左
			/*
				1 1 1 1 1 1 1
				1 2 0 0 0 0 1
				1 2 2 2 0 0 1
				1 1 1 2 0 0 1
				1 0 0 2 0 0 1
				1 0 0 2 0 0 1
				1 0 0 2 2 2 1
				1 1 1 1 1 1 1
			*/
			if Myset(mymap, i+1, j) {
				return true
			} else if Myset(mymap, i, j+1) {
				return true
			} else if Myset(mymap, i-1, j) {
				return true
			} else if Myset(mymap, i, j-1) {
				return true
			} else {
				mymap[i][j] = 3
				return false
			}
		} else {
			// 改点为墙(1), 不能探索
			return false
		}
	}
}
func main() {
	var mymap [8][7]int
	/*
		0: 还未探索的路
		1: 墙
		2: 通路
		3: 探索过, 死路
	*/
	mymap[3][1] = 1
	mymap[3][2] = 1
	for i := 0; i < 8; i++ {
		mymap[i][0] = 1
		mymap[i][6] = 1
	}

	for i := 0; i < 7; i++ {
		mymap[0][i] = 1
		mymap[7][i] = 1
	}
	Myset(&mymap, 1, 1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mymap[i][j], " ")
		}
		fmt.Println("")
	}
}
