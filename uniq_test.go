package uniq

import "testing"

type Int int

func (i Int) Key() int {
	return int(i)
}

type IntSlice []Int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Get(i int) Keyer {
	return Keyer(s[i])
}

func (s IntSlice) Set(i int, item Keyer) {
	s[i] = item.(Int)
}

func (s IntSlice) Slice(left, right int) Interface {
	return s[left:right]
}

func (s IntSlice) New(size int) Interface {
	return make(IntSlice, size)
}

func TestUniq(t *testing.T) {

	items := IntSlice([]Int{1, 1, 2, 2, 3, 3, 3, 2, 1})

	items = Uniq(items).(IntSlice)

	if len(items) != 3 {
		t.Error("Expected len==3, got ", len(items))
	}
}
