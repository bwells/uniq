package uniq

type Interface interface {
	Key() int
}

func Uniq(items []Interface) []Interface {

	set := make(map[int]Interface)
	for _, c := range items {
		set[c.Key()] = c
	}

	results := make([]Interface, len(set))
	i := 0
	for _, value := range set {
		results[i] = value
		i++
	}

	return results
}
