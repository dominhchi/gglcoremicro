package utils

func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
	for _, v := range slice1 {

		m[v]++
	}
	for _, v := range slice2 {

		times, _ := m[v]
		if times == 1 {

			n = append(n, v)
		}
	}
	return n
}

func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {

		m[v]++
	}
	for _, value := range slice1 {

		if m[value] == 0 {

			n = append(n, value)
		}
	}

	for _, v := range slice2 {

		if m[v] == 0 {

			n = append(n, v)
		}
	}
	return n
}
