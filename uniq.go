package uniq

type Keyer interface {
	Key() int
}

type Interface interface {
	Len() int
	Get(i int) Keyer
	Set(i int, item Keyer)
	Slice(left, right int) Interface
	New(size int) Interface
}

// func Uniq(items Interface) []Interface {

// 	set := make(map[int]Interface)
// 	for _, c := range items {
// 		set[c.Key()] = c
// 	}

// 	results := make([]Interface, len(set))
// 	i := 0
// 	for _, value := range set {
// 		results[i] = value
// 		i++
// 	}

// 	return results
// }

func Uniq(items Interface) Interface {
	set := make(map[int]Keyer)
	var item Keyer
	for i := 0; i < items.Len(); i++ {
		item = items.Get(i)
		set[item.Key()] = item
	}

	uniqs := items.New(len(set))

	i := 0
	for _, value := range set {
		uniqs.Set(i, value)
		i++
	}

	return uniqs
}
