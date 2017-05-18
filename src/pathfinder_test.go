package main

import (
	"math"
	"testing"
	"fmt"
//	"sync"
)
/*
func TestWorkingPath(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0}}
	testmap := TileConvert(matrix)

	path, ok := getPath(&testmap, &testmap[0][0])

	if !ok {t.Errorf("Expected a valid path")}
	if len(path) != 8 {t.Errorf("Expected pathlength: 8, but got pathlength: %d", len(path))}
	if *path[0] != testmap[0][0] {t.Errorf("Expected starttile: %d, but got starttile: %d", testmap[0][0], path[0])}
	if *path[7] != testmap[0][2] {t.Errorf("Expected lasttile: %d, but got lasttile: %d", testmap[0][2], path[7])}
}

func TestBlockedPath(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0}}
	testmap := TileConvert(matrix)

	path, ok := getPath(&testmap, &testmap[0][0])
	if len(path) > 0 {t.Errorf("Expected a empty path, but got a path of length: %d", len(path))}
	if ok {t.Errorf("Expected an invalid path")}
}

func TestFirePath(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0}}
	testmap := TileConvert(matrix)
	SetFire(&(testmap[3][2]))
	for i := 0; i < 10; i++ {
		FireSpread(testmap)
	}

	path, ok := getPath(&testmap, &testmap[0][3])
	if !ok {t.Errorf("Expected a valid path")}
	if len(path) != 7 {t.Errorf("Expected pathlength: 7, but got pathlength: %d", len(path))}
	if *path[0] != testmap[0][3] {t.Errorf("Expected starttile: %d, but got starttile: %d", testmap[0][3], path[0])}
	if *path[6] != testmap[6][3] {t.Errorf("Expected lasttile: %d, but got lasttile: %d", testmap[6][3], path[6])}
}

func TestDoorsPath(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 1, 0, 0, 0}}

	testmap := TileConvert(matrix)

	path, ok := getPath(&testmap, &testmap[0][0])
	if !ok {t.Errorf("Expected a valid path")}
	if len(path) != 13 {t.Errorf("Expected pathlength: 13, but got pathlength: %d", len(path))}
	if *path[0] != testmap[0][0] {t.Errorf("Expected starttile: %d, but got starttile: %d", testmap[0][0], path[0])}
	if *path[12] != testmap[6][0] {t.Errorf("Expected lasttile: %d, but got lasttile: %d", testmap[6][0], path[12])}	
}

*/

/*
func TestGetPath(t *testing.T) {
     tests above... 
}*/



func TestStepCost(t *testing.T) {

	ti := makeNewTile(0, 0, 0)

	if stepCost(ti) != 1 {
		t.Errorf("Expected stepcost: 1, but got stepcost: %d", stepCost(ti))
	}

	for i := float32(0); i < 10; i++ { //TODO: om vi ändrar cost för heat så redigera testet!
		if stepCost(ti) != float32(i/5+1) {
			t.Errorf("Expected stepcost: %g, but got stepcost: %g", float32(i/5+1), stepCost(ti))
		}
		ti.heat += 1
	}
	SetFire(&ti)
	if !math.IsInf(float64(stepCost(ti)), 1) {
		t.Errorf("Expected stepcost: %g, but got stepcost: %g", float32(math.Inf(1)), stepCost(ti))
	}

	// empty tile = 1
	// heatlvl tile = 1 + heatlvl/5
	// fire tile = infinity
}

func mapToQueue(m [][]tile) queue{
	q := queue{}
	for i, list := range m {
		for j, _ := range list {
			v := float32(0)
			q = append(q, tileCost{&m[i][j], &v})
		}
	}
	return q
}

func TestGetNeighbors(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0, 1, 0, 1, 0}, 
		{1, 1, 1, 1, 1, 1, 1}, 
		{0, 0, 0, 0, 0, 0, 0},
		{0, 3, 3, 3, 3, 3, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0}}
	testmap := TileConvert(matrix)
	tileQ := mapToQueue(testmap)
//	fmt.Println(len(tileQ))
//	for _, tl := range tileQ {fmt.Println(tl.tile.xCoord, tl.tile.yCoord)}
	
	for i, list := range testmap {
		for j, ti := range list {
			neighbors := getNeighbors(&ti, tileQ)
			if i == 0 && validTile(&ti) {
				if len(neighbors) != 0 {					
					t.Errorf("Expected 0 neigbors, but got %d neighbors", len(neighbors))
				}
			} else if i == 2 {
				if len(neighbors) != 2 {
					if len(neighbors) > 0 {
						fmt.Println(tileQ.inQueue(&ti))					
						fmt.Println(ti)
						fmt.Println(neighbors[0])}
					t.Errorf("Expected 2 neigbors, but got %d neighbors", len(neighbors))
				}
			} else if (i == 5 || i == 6) && j > 0 && j < 6 {
				if len(neighbors) != 8 {
					t.Errorf("Expected 8 neigbors, but got %d neighbors", len(neighbors))
				}
			}
		}
	}
}

func TestValidTile(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{3, 3, 3, 3}}
	testmap := TileConvert(matrix)

	for i, list := range testmap {
		for _, ti := range list {
			if i < 2 {
				if !validTile(&ti) {
					t.Errorf("Expected validtile, but got invalidtile")
				}
			} else {
				if validTile(&ti) {
					t.Errorf("Expected invalidvalidtile, but got validtile")
				}
			}
		}
	}
	if validTile(nil) {
		t.Errorf("Expected invalidvalidtile, but got validtile")
	}
}

func TestCompactPath(t *testing.T) {
	matrix := [][]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,2}}
	testmap := TileConvert(matrix)
	
	parentOf := make(map[*tile]*tile)
	
	previous := &testmap[0][0]
	
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			parentOf[previous] = &testmap[i][j]
			previous = &testmap[i][j]
		}
	}
	
	path, ok := compactPath(parentOf, &testmap[5][4], &testmap[0][0])
	
	if ok {
		if len(path) != 30 {t.Errorf("Expected pathlength: %d, but got pathlangth: %d", 30, len(path))}
		ind := 29
		for i := 0; i <= 5; i++ {
			for j := 0; j <= 4; j++ {
				if *path[ind] != testmap[i][j] {t.Errorf("Expected pathtile: %d, but got pathtile: %d", testmap[i][j], path[ind])}
				ind--
			}
		}		
	}
}

/*
func TestTime(t *testing.T) {
	q := queue{}

	size := 10000

	matrix := [][]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,2}}
	testmap := TileConvert(matrix)

	for i := 0; i < size; i++ {
		q.Add()
	}

}*/

func makeTestMap(xSize, ySize int) [][]tile{
	testMatrix := [][]int{}

	for x := 0; x < xSize; x++ {
		row := []int{}
		for y := 0; y < ySize; y++ {
			row = append(row, y)
		}
		testMatrix = append(testMatrix, row)
	}
		
	return TileConvert(testMatrix)
}
/*
func TestTwo(t *testing.T) {
	matrix := [][]int{}
	xS := 10
	yS := 10

	for x := 0; x < xS; x++ {
		row := []int{}
		for y := 0; y < yS; y++ {
			row = append(row, 0)
		}		
		matrix = append(matrix, row)
	}
	matrix[xS - 1][yS - 1] = 2
	testmap := TileConvert(matrix)

	ok1 := false
	ok2 := false
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, ok1 = getPath(&testmap, &testmap[0][0])}()
	go func() {
		defer wg.Done()
		_, ok2 = getPath(&testmap, &testmap[1][1])}()

	if !ok1 {t.Errorf("Expected a valid path, but got a invalid one")}
	if !ok2 {t.Errorf("Expected a valid path, but got a invalid one")}
<<<<<<< HEAD:simulation/pathfinder_test.go

//	_, ok := getPath(&testmap, &testmap[0][0])

//	if !ok {t.Errorf("Expected a valid path, buut got a invalid one")}	
} */

/*
--- PASS: TestLargeMap (45.73s)
=== RUN   TestLargeMap2
--- PASS: TestLargeMap2 (26.45s)
*/

	
	/*	ppl := &testmap[0][0], &testmap[1][1]
	for pers := range ppl {
		go func(p *Tile) {
			getPath(&testmap, p)	
		}(pers)
	}*/
//	_, ok := getPath(&testmap, &testmap[0][0])

//	if !ok {t.Errorf("Expected a valid path, buut got a invalid one")}	
//}

func TestLargeMap(t *testing.T) {
	matrix := [][]int{}
	xS := 10
	yS := 10

	for x := 0; x < xS; x++ {
		row := []int{}
		for y := 0; y < yS; y++ {
			row = append(row, 0)
		}		
		matrix = append(matrix, row)
	}
	matrix[xS - 1][yS - 1] = 2
	testmap := TileConvert(matrix)
	_, ok := getPath(&testmap, &testmap[0][0])

	if !ok {t.Errorf("Expected a valid path, but got a invalid one")}
} 

func inNbrs(nbrs []*tile, t *tile) bool{
	for _, n := range nbrs {
		if *n == *t {return true}
	}
	return false
}


// 100*100 took 1.48 s
func TestLargeMap2(t *testing.T) {
	matrix := [][]int{}
	xS := 200
	yS := 200

	for x := 0; x < xS; x++ {
		row := []int{}
		for y := 0; y < yS; y++ {
			row = append(row, 0)
		}		
		matrix = append(matrix, row)
	}
	matrix[xS - 1][yS - 1] = 2
	testmap := TileConvert(matrix)
	_, ok := getPath2(&testmap, &testmap[0][0])

	if !ok {t.Errorf("Expected a valid path, but got a invalid one")}
}


// 200*200 ended after 10min: took to long!
// 100*100 took 71.66 s
// 50*50 took 2.22s
/*
func TestManyPeople(t *testing.T) {
	matrix := [][]int{}
	xS := 100
	yS := 100

	for x := 0; x < xS; x++ {
		row := []int{}
		for y := 0; y < yS; y++ {
			row = append(row, 0)
		}		
		matrix = append(matrix, row)
	}
	matrix[xS - 1][yS - 1] = 2
	testmap := TileConvert(matrix)
	var wg sync.WaitGroup
	wg.Add(xS)
	for x := 0; x < xS; x++ {
		go func(i int) {
			defer wg.Done()
			_, ok := getPath2(&testmap, &testmap[i][0])
			if !ok {t.Errorf("Expected a valid path, but got a invalid one")} 
		}(x)
	}
	wg.Wait()
}*/


/*

func TestGetNeighborsPruned(t *testing.T) {
	matrix := [][]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,2}}
	testmap := TileConvert(matrix)


	nbrs := []*tile{}
	
	nbrs  = getNeighborsPruned(&testmap[1][1], Direction{0,1})
//	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[0][1]) {t.Errorf("whoopsie")}

	nbrs = getNeighborsPruned(&testmap[1][1], Direction{0,-1})
//	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[2][1]) {t.Errorf("whoopsie")}

	
	nbrs = getNeighborsPruned(&testmap[1][1], Direction{1,0})
//	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[2][1]) {t.Errorf("whoopsie")}

	nbrs = getNeighborsPruned(&testmap[1][1], Direction{-1,0})
	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[2][1]) {t.Errorf("whoopsie")}
	
	nbrs = getNeighborsPruned(&testmap[1][1], Direction{0,-1})
	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[2][1]) {t.Errorf("whoopsie")}
	
	nbrs = getNeighborsPruned(&testmap[1][1], Direction{0,-1})
	fmt.Println(nbrs[0])
	if !inNbrs(nbrs, &testmap[2][1]) {t.Errorf("whoopsie")}
}
*/


func TestGetJumpPoint(t *testing.T) {
	matrix := [][]int {
		{0,0,1,0,0,0},
		{0,0,1,0,0,0},
		{0,0,0,0,0,0},
		{0,0,0,0,0,0},
		{0,0,0,1,0,0},
		{0,0,0,1,0,2}}
	testmap := TileConvert(matrix)


	// höger
	jp1 := getJumpPoint(&testmap[3][1], Direction{0,1})
	if jp1.jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp1.jp == testmap[3][4]) {t.Errorf("Expected jp: 3 3, but got jp: %d %d", jp1.jp.xCoord, jp1.jp.yCoord)}

	// vänster
	jp2 := getJumpPoint(&testmap[2][4], Direction{0,-1})
	if jp2.jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp2.jp == testmap[2][1]) {t.Errorf("Expected jp: 2 1, but got jp: %d %d", jp2.jp.xCoord, jp2.jp.yCoord)}
	
	
	matrix2 := [][]int {
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0}}
	testmap = TileConvert(matrix2)
	
	
	// höger
	jp1 = getJumpPoint(&testmap[2][1], Direction{0,1})
	if jp1.jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp1.jp == testmap[2][4]) {t.Errorf("Expected jp: 2 4, but got jp: %d %d", jp1.jp.xCoord, jp1.jp.yCoord)}
//	fmt.Println(jp1.fn[0])
//	fmt.Println(jp1.fn[1])
	
	// vänster
	jp2 = getJumpPoint(&testmap[2][5], Direction{0,-1})
	if jp2.jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp2.jp == testmap[2][2]) {t.Errorf("Expected jp: 2 2, but got jp: %d %d", jp2.jp.xCoord, jp2.jp.yCoord)}

	
	/*
	jp := getJumpPoint(&testmap[1][1], Direction{1,1})
	if jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp== testmap[3][3]) {t.Errorf("Expected jp: 2 2, but got jp: %d %d", jp.xCoord, jp.yCoord)}*/
}

func TestSneJP(t *testing.T) {
	matrix := [][]int {
		{0,0,0,0,0,0},
		{0,0,0,0,0,0},
		{0,0,0,0,0,0},
		{0,0,0,0,0,0},
		{0,0,1,0,0,0},
		{0,0,1,0,0,0}}
	testmap := TileConvert(matrix)
	
	// nw
	jp1 := sneJP(&testmap[3][3], Direction{-1,-1})
	if jp1.jp == nil {t.Errorf("Expected a valid jp, but got an invalid one")}
	if !(*jp1.jp == testmap[3][3]) {t.Errorf("Expected jp: 2 2, but got jp: %d %d", jp1.jp.xCoord, jp1.jp.yCoord)}
	fmt.Println(jp1.fn[0])


}
/*
<<<<<<< HEAD:simulation/pathfinder_test.go
func TestGetPath2(t *testing.T) {
	matrix := [][]int {
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,1,0,0},
		{0,0,1,0,2}}
	testmap := TileConvert(matrix)

	path, ok := getPath2(&testmap, &testmap[0][0])
	if !ok {t.Errorf("Expected a valid path, but got an invalid one")}
	if !(*path[0] == testmap[2][2]) {t.Errorf("Expected jp: 2 2, but got jp: %d %d", path[0].xCoord, path[0].yCoord)}
=======
func TestThis(t *testing.T) {
	matrix := [][]int{}
	xS := 100
	yS := 100

	for x := 0; x < xS; x++ {
		row := []int{}
		for y := 0; y < yS; y++ {
			row = append(row, 0)
		}		
		matrix = append(matrix, row)
	}
	matrix[xS - 1][yS - 1] = 2
	testmap := TileConvert(matrix)

	list := [][]int{matrix[0]}
	nrOfPpl := yS
	for x := 0; x < nrOfPpl; x++ {
		list = append(list, )
	}


	ppl := PeopleInit(testmap, list)
	MovePeople(&testmap, ppl)
>>>>>>> master:src/pathfinder_test.go
}*/






// tests moved from pathfinder!




func mainPath() {

	workingPath()
	fmt.Println("--------------")
/*	blockedPath()
	fmt.Println("--------------")
	firePath()*/
	fmt.Println("--------------")
	doorsPath()
}

func workingPath() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0}}
	testmap := TileConvert(matrix)

	path, _ := getPath(&testmap, &testmap[0][0])

	fmt.Println("\nWorking path:")
	printPath(path)
}

func blockedPath() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0}}
	testmap := TileConvert(matrix)

	path, _ := getPath(&testmap, &testmap[0][0])

	fmt.Println("\nBlocked path:")
	printPath(path)

}

func firePath() {
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0}}
	testmap := TileConvert(matrix)
	SetFire(&(testmap[3][2]))
	for i := 0; i < 10; i++ {
		FireSpread(testmap)
	}

	path, _ := getPath(&testmap, &testmap[0][3])
	fmt.Println("\nFire path:")
	printPath(path)
}

func doorsPath() {
	matrix := [][]int{
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 1, 0, 0, 0}}

	testmap := TileConvert(matrix)

	path, _ := getPath(&testmap, &testmap[0][0])
	fmt.Println("\nDoors path:")
	printPath(path)
}

func Whut() {
	matrix := [][]int {
		{0,0,0,0,0,0,0},
		{0,0,1,0,0,0,0},
		{1,1,1,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,0,0,0},
		{2,0,0,1,1,0,0}}

	testmap := TileConvert(matrix)	

	path, _ := getPath2(&testmap, &testmap[0][6])
	printPath(path)
}

