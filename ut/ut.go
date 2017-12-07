package ut

func flatten(chromosome [][]int) []int {
	flattench := make([]int)
	for _, route := range chromosome {
		for _, node := range route {
			flattench = append(flattench, node)
		}
	}
	return flattench
}
