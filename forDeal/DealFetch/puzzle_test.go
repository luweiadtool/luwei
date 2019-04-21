package DealFetch

import (
	"fmt"
	"goDemo/Project/Luwei/config"
	"myTool/file"
	"testing"
)

func TestFetch(t *testing.T) {
	value := file.PathExist(".../source/arial.ttf")
	fmt.Println(value)
}

func TestPuzzle(t *testing.T) {
	dir := "/Users/qianjianeng/Downloads/forDeal/70000294"
	con := config.LoadConfig()
	Puzzle(dir, con.ForDeal,dir+"/70000294_puzzle", true, true)
}