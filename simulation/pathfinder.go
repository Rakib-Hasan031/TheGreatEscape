package main

import	"fmt"

func getPath(m *[][]tile, from *tile, to *tile) ([]*tile, bool){
	// TODO: should 'to' be a list of tiles? (all doors)
	// or should there be no 'to', it just searches for any door?	
	// map to keep track of the final path

	var parentOf map[*tile]*tile
	parentOf = make(map[*tile]*tile)
	//initialise 'costqueue', start-0, other-infinite
	costQueue := queue{}
	
	for i,list := range *m {
		for j, _ := range list {		
			costQueue.Add(&(*m)[i][j], 1000)   // 1000~infinite		
		}
	}
	
	costQueue.Update(from, 0)

//	checkedQueue := costQueue   TODO: implement this later for a more efficient algorithm

	current := tileCost{}
	//essential loop
	for len(costQueue) != 0 && current.tile != to{	
	
		current = (&costQueue).Pop()	
		neighbors := getNeighbors(current.tile)
		for _, neighbor := range neighbors {		
			cost := current.cost + stepCost(*neighbor)
			// TODO: 1 default cost improve!? depending on heat, smoke etc		
			if cost < costQueue.costOf(neighbor) {			
				parentOf[neighbor] = current.tile
				costQueue.Update(neighbor, cost)			
			}
		}
	//	checkedQueue.AddTC(current)
	//	costQueue.Remove(current.tile)

	}
	return compactPath(parentOf, from, to)
}
 
func stepCost(t tile) float32{
	cost := float32(1)
	cost += float32(t.heat)/5   //TODO how much cost for fire etc??
	cost += float32(t.fireLevel)*1000
	return cost
}

func getNeighbors(current *tile) []*tile{
	neighbors := []*tile{}

	if validTile(current.neighborNorth) {neighbors = append(neighbors, current.neighborNorth)}
	if validTile(current.neighborEast) {neighbors = append(neighbors, current.neighborEast)}
	if validTile(current.neighborWest) {neighbors = append(neighbors, current.neighborWest)}
	if validTile(current.neighborSouth) {neighbors = append(neighbors, current.neighborSouth)}

	return neighbors
}

func validTile(t *tile) bool {
	if t == nil {
		return false
	}
	return !t.wall 
}

func compactPath(parentOf map[*tile]*tile, from *tile, to *tile) ([]*tile, bool) {
	path := []*tile{to}

	current := to

	for current.xCoord != from.xCoord || current.yCoord != from.yCoord {		
		path = append([]*tile{parentOf[current]}, path...)
	
		ok := true
		current, ok = parentOf[current]
		if  !ok{
			return nil, false	
		}
	}
	return path, true	
}

func printPath(path []*tile) {
	if path == nil {
		fmt.Println("No valid path exists")
	}
	for i, t := range path {
		fmt.Println(i , ":", t.xCoord ,"," ,t.yCoord)
	}
}

func mainPath() {

	workingPath()
	fmt.Println("--------------")
	blockedPath()
	fmt.Println("--------------")
	firePath()
	fmt.Println("--------------")
	doorsPath()
}

func workingPath() {
	matrix := [][]int {
		{0,1,0,0},
		{0,0,1,0},
		{0,0,0,0}, 
		{0,0,1,0}}
	testmap := TileConvert(matrix)
	printTileMap(testmap)
	
	path, _ := getPath(&testmap, &testmap[0][0], &testmap[0][2])

	fmt.Println("\nWorking path:")
	printPath(path)
}


func blockedPath(){
	matrix := [][]int {
		{0,1,0,0},
		{0,0,1,0},
		{0,0,1,0}, 
		{0,0,1,0}}
	testmap := TileConvert(matrix)
	printTileMap(testmap)
	
	path, _ := getPath(&testmap, &testmap[0][0], &testmap[3][3])

	fmt.Println("\nBlocked path:")
	printPath(path)

}

func firePath() {
	matrix := [][]int {
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,0,0,0}, 
		{0,0,0,0,0,0,0}} 

	testmap := TileConvert(matrix)
	SetFire(&(testmap[3][2]))
	for i:= 0; i < 10; i++ {
		FireSpread(testmap)
	}

	printTileMap(testmap)

	path, _ := getPath(&testmap, &testmap[0][3], &testmap[6][2])
	fmt.Println("\nFire path:")
	printPath(path)
}

func doorsPath() {
	matrix := [][]int {
		{0,0,0,1,0,0,0},
		{0,0,0,2,0,0,0},
		{1,1,1,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,2,0,0,0}, 
		{0,0,0,1,0,0,0}}

	testmap := TileConvert(matrix)
	printTileMap(testmap)
	path, _ := getPath(&testmap, &testmap[0][0], &testmap[6][0])
	fmt.Println("\nDoors path:")
	printPath(path)
}
