package genshin

import (
	"fmt"
	"math/rand"

	"github.com/wsmbsbbz/predictor/util"
)

const ()

type poolRecord struct {
	name    string
	code    int
	quality int
}

// example: drawRecord{"Xiangling", 8, 4, 26, 9}
type drawRecord struct {
	name    string
	code    int
	quality int

	fiveStarAccumulate int
	fourStarAccumulate int
}

type GenshinPredictor struct {
	pinkballLeft int
	// key: code, val: name, such as (8:Xiangling)
	fiveStarPool       []poolRecord
	fourStarPool       []poolRecord
	threeStarPool      []poolRecord
	fiveStarAccumulate int // probability is 0.6%
	fourStarAccumulate int // probability is 5.1%
	guarantee          bool

	treasure5Star []drawRecord
	treasure4Star []drawRecord
	treasure3Star []drawRecord
	score         int
	goalCheckFunc func(five, four []drawRecord) bool
	// Note for construct and print the predict result.
	note string
}

func (g *GenshinPredictor) Draw() {
	for g.pinkballLeft > 0 {
		g.SingleDraw()
		g.pinkballLeft--
	}
}

func (g *GenshinPredictor) SingleDraw() {
	util.Dprintf("g: %v\n", g)
	// 5 star
	var fiveProb float64
	if g.fiveStarAccumulate <= 73 {
		fiveProb = 0.6
	} else if g.fiveStarAccumulate == 90 {
		fiveProb = 100
	} else {
		fiveProb = float64(g.fiveStarAccumulate-73)*6 + 0.6
	}

	if rand.Float64()*100 < fiveProb {
		c := g.Draw5Star()
		g.treasure5Star = append(g.treasure5Star, poolToDraw(c, g.fiveStarAccumulate, g.fourStarAccumulate))
		return
	}

	// 4 star
	var fourProb float64
	if g.fourStarAccumulate < 9 {
		fourProb = 5.1
	} else {
		fourProb = 100
	}
	if rand.Float64()*100 < fourProb {
		c := g.Draw4Star()
		g.treasure4Star = append(g.treasure4Star, poolToDraw(c, g.fiveStarAccumulate, g.fourStarAccumulate))
		return
	}

	// 3 star
	c := g.Draw3Star()
	g.treasure3Star = append(g.treasure3Star, poolToDraw(c, g.fiveStarAccumulate, g.fourStarAccumulate))
}

func poolToDraw(p poolRecord, five, four int) drawRecord {
	return drawRecord{
		name:               p.name,
		code:               p.code,
		quality:            p.quality,
		fiveStarAccumulate: five,
		fourStarAccumulate: four,
	}
}

func (g *GenshinPredictor) Draw5Star() poolRecord {
	// Acquiring a 5-star character does not increase the cumulative progress of a 4-star character
	g.fiveStarAccumulate = 0
	c := randPoolRecord(g.fiveStarPool)
	util.Dprintf("Draw5Star: get character %v\n", c)
	return c
}

func (g *GenshinPredictor) Draw4Star() poolRecord {
	g.fiveStarAccumulate++
	g.fourStarAccumulate = 0
	c := randPoolRecord(g.fourStarPool)
	util.Dprintf("Draw4Star: get character %v\n", c)
	return c
}

func (g *GenshinPredictor) Draw3Star() poolRecord {
	g.fiveStarAccumulate++
	g.fourStarAccumulate++
	c := randPoolRecord(g.threeStarPool)
	util.Dprintf("Draw3Star: get character %v\n", c)
	return c
}

func randPoolRecord(ps []poolRecord) poolRecord {
	n := len(ps)
	return ps[rand.Intn(n)]
}

func (g *GenshinPredictor) CheckAchieve() bool {
	// fmt.Println(g)
	return g.goalCheckFunc(g.treasure5Star, g.treasure4Star)
}

func (g *GenshinPredictor) String() string {
	return fmt.Sprintf(`GenshinPredictor:
note: %s,
pinkballLeft: %d,
fiveStarPool: %v,
fourStarPool: %v,
threeStarPool: %v,
fiveStarAccumulate: %d,
fourStarAccumulate: %d,
guarantee: %v,
treasure5Star: %v,
treasure4Star: %v,
treasure3Star: %v,
`, g.note, g.pinkballLeft, g.fiveStarPool, g.fourStarPool, g.threeStarPool,
		g.fiveStarAccumulate, g.fourStarAccumulate, g.guarantee,
		g.treasure5Star, g.treasure4Star, g.treasure3Star)
}
