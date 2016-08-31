package runs

import (
	"sort"
)

type byMapped struct {
	arr    []interface{}
	mapper func(interface{}) int64
}

func (a byMapped) Len() int           { return len(a.arr) }
func (a byMapped) Swap(i, j int)      { a.arr[i], a.arr[j] = a.arr[j], a.arr[i] }
func (a byMapped) Less(i, j int) bool { return a.mapper(a.arr[i]) < a.mapper(a.arr[j]) }

// Detect will go over a list of `things` and group them according to a distance function called `belongs`.
// You will need to provide a mapper function, that maps a thing to an ordinal (int64) value.
//
// For example:
//	 grouped := runs.Detect(things, func(thing interface{}) int64 {
//	  	return thing.(os.FileInfo).Size()
//	 }, func(a, b int64) bool {
//	  	return math.Abs(float64(a-b)) < 1000
//	 })
func Detect(things []interface{}, mapper func(interface{}) int64, belongs func(a, b int64) bool) map[int64][]interface{} {
	sort.Sort(byMapped{arr: things, mapper: mapper})
	sentinel := mapper(things[0])
	last := sentinel
	clusters := map[int64][]interface{}{}
	for _, rec := range things {
		ordinal := mapper(rec)
		if belongs(last, ordinal) == false {
			sentinel = ordinal
			clusters[sentinel] = []interface{}{}
		}
		clusters[sentinel] = append(clusters[sentinel], rec)
		last = ordinal
	}

	return clusters
}
