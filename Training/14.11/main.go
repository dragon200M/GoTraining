package main

import "fmt"

func main() {

	buckets := make([][]int, 0)

	for i := 0; i < 10; i++ {

		buckets = append(buckets, []int{})

	}

	tab := []int{3, 2, 1, 1, 1, 4, 4, 4, 4, 4, 5, 5, 5, 8, 5, 5, 5}

	for j, i := range tab {
		buckets[i] = append(buckets[i], j)

	}

	zm := 0
	max := 0
	for i := 0; i < len(buckets); i++ {

		if len(buckets[i]) > max {
			zm = i
			max = len(buckets[i])
		}

	}

	fmt.Println("Liczba:", zm, ",indeks:", buckets[zm])
}

/*
p22:
. Mając tablicę intów trzeba było napisać metodę, która zwróci jakąkolwiek pozycję,
w której znajduje się wartość dominująca, czyli taka, która zajmuje więcej niż 50% miejsc w tablicy
np. table[] = {3, 2, 3, 3, 1, 4, 3} zwróci wartość 0, 2, 3 lub 6

*/
