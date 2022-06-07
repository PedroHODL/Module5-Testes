package ordering

func Sort(lista []int) []int {
	var n = len(lista)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if lista[j] < lista[minIdx] {
				minIdx = j
			}
		}
		lista[i], lista[minIdx] = lista[minIdx], lista[i]
	}
	return lista
}
