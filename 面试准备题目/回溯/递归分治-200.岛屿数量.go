package main

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid{
		for column := range grid[line]{
			if grid[line][column] == '0'{
				continue
			}
			count ++
			_dfsdy(grid, line, column)
		}
	}
	return count
}

func _dfsdy(grid [][]byte, line int, column int) {
	// terminator
	if line < 0 || line >= len(grid) || column < 0 || column >= len(grid[line]){
		return
	}
	if grid[line][column] == '0'{
		return
	}
	grid[line][column] = '0'
	_dfsdy(grid, line+1, column)
	_dfsdy(grid, line-1, column)
	_dfsdy(grid, line, column+1)
	_dfsdy(grid, line, column-1)
}


