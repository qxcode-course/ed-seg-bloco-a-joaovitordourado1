package main

import "fmt"

type bomb struct {
	x    int
	y    int
	raio int
}

func main() {
	n := 0
    co:=0
	fmt.Scan(&n,&co)
	bomba := make([]bomb, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&bomba[i].x, &bomba[i].y, &bomba[i].raio)
	}
	graph := buildgraph(bomba)
	max := 0
	for i := 0; i < n; i++ {
		ans := bfs(i, graph, n)
		if ans > max {
			max = ans
		}
	}

	fmt.Println(max)
}
func buildgraph(bombas []bomb) [][]int {
	k := len(bombas)
	graph := make([][]int, k)
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			if i == j {
				continue
			}
			l := bombas[i].y - bombas[j].y
			o := bombas[i].x - bombas[j].x
			dq := (l * l) + (o * o)
			rq := bombas[i].raio * bombas[i].raio

			if dq <= rq {
				graph[i] = append(graph[i], j)
			}
		}
	}
	return graph
}
func bfs(start int, g [][]int, n int) int {
	vis := make([]bool, n)
	q := []int{start}
	vis[start] = true
	c := 1
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if !vis[v] {
				vis[v] = true
				q = append(q, v)
				c++
			}
		}
	}
	return c
}
