func findCircleNum(isConnected [][]int) int{
	n := len(isConnected)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(city int){
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	provinces := 0
	for city := 0; city < n; city++ {
		if !visited[city] {
			dfs(city)
			provinces++
		}
	}

	return provinces
}


// approach bfs

func findCircleNum(isConnected [][]int) int{
	n := len(isConnected)
	visited := make([]bool, n)
	provinces := 0

	for city := 0; city < n; city++ {
		if !visited[city] {
			queue := []int{city}
			visited[city] = true

			for len(queue) > 0 {
				currentCity := queue[0]
				queue = queue[1:]
				for neighbor := 0; neighbor < n; neighbor++ {
					if isConnected[currentCity][neighbor] == 1 && !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
			provinces++
		}

	}

	return provinces
}

// approach dsu

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootY] = rootX
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	count :=0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}
	return count
}
