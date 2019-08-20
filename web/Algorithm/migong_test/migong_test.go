package migong_test

import (
	"eilieili/web/Algorithm/migong"
	"testing"
)

func TestMigong(t *testing.T) {
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
	if migong.Myset(&mymap, 1, 1) {
		t.Error(`migong.Myset(mymap, 1, 1) is failed`)
	}
}
