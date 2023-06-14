package genshin

import (
	"fmt"

	"github.com/wsmbsbbz/predictor/util"
	predictor "github.com/wsmbsbbz/predictor/v1"
)

const (
	PredictCnt = 100000
)

func Main() {
	var g *GenshinPredictor
	var n int
	g = initGenshinPredictor(40, func(five, four []drawRecord) bool {
		util.Dprintf("five: %v, four: %v\n", five, four)
		for _, record := range five {
			if record.code == 0 {
				return true
			}
		}
		return false
	}, "40抽出海哥")
	util.Dprintf("g: %v\n", g)
	n = predictor.Predict(g, PredictCnt)
	fmt.Printf("%s: %d successful times of %d times(%f)\n", g.note, n, PredictCnt, float64(n)/PredictCnt)

	g = initGenshinPredictor(40, func(five, four []drawRecord) bool {
		util.Dprintf("five: %v, four: %v\n", five, four)
		for _, record := range four {
			if record.code == 2 {
				return true
			}
		}
		return false
	}, "40抽出香菱")
	util.Dprintf("g: %v\n", g)
	n = predictor.Predict(g, PredictCnt)
	fmt.Printf("%s: %d successful times of %d times(%f)\n", g.note, n, PredictCnt, float64(n)/PredictCnt)

	g = initGenshinPredictor(10, func(five, four []drawRecord) bool {
		util.Dprintf("five: %v, four: %v\n", five, four)
		return len(five) >= 3
	}, "10连三金")
	util.Dprintf("g: %v\n", g)
	n = predictor.Predict(g, PredictCnt)
	fmt.Printf("%s: %d successful times of %d times(%f)\n", g.note, n, PredictCnt, float64(n)/PredictCnt)
}

func getCurrent5StarsPool() []poolRecord {
	return []poolRecord{
		{"艾尔海森", 0, 5},
		// {"枫原万叶", 1, 5},
	}
}
func getCurrent4StarsPool() []poolRecord {
	return []poolRecord{
		{"香菱", 2, 4},
		{"瑶瑶", 3, 4},
		{"鹿野院平藏", 4, 4},
		{"随机非UP4星", 5, 4},
		{"随机非UP4星", 6, 4},
		{"随机非UP4星", 7, 4},
	}
}
func getCurrent3StarsPool() []poolRecord {
	return []poolRecord{
		{"3星垃圾", 5, 3},
	}
}

func initGenshinPredictor(n int, goalCheckFunc func(five, four []drawRecord) bool, note string) *GenshinPredictor {
	g := GenshinPredictor{}
	g.pinkballLeft = n
	g.fiveStarPool = getCurrent5StarsPool()
	g.fourStarPool = getCurrent4StarsPool()
	g.threeStarPool = getCurrent3StarsPool()
	g.goalCheckFunc = goalCheckFunc
	g.note = note
	return &g
}
