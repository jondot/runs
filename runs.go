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

// GroupVisitor marks a group visitor that is put into a walker. Use this
// as a guide to implement your own visitors.
// For example:
//		type GroupPrinter struct {
//		}
//
//		func (g *GroupPrinter) VisitHeader(group interface{}) {
//			groupName := fmt.Sprintf("group_%s", group.(os.FileInfo).Name())
//			fmt.Printf("%s\n", groupName)
//		}
//
//		func (g *GroupPrinter) VisitNode(i int, sz int, header interface{}, file interface{}) {
//			fi := file.(os.FileInfo)
//			if i == sz-1 {
//				fmt.Printf("└── ")
//			} else {
//				fmt.Printf("├── ")
//			}
//			fmt.Printf("%s (%v)\n", fi.Name(), fi.ModTime())
//		}
type GroupVisitor interface {
	VisitHeader(interface{})
	VisitNode(int, int, interface{}, interface{})
}

// GroupWalker is a group walker implementation into which you provide a visitor.
type GroupWalker struct {
}

// Walk will walk a group given your custom visitor.
// For example:
//  printer := &GroupWalker{}
//  printer.Walk(grouped, &GroupPrinter{})
func (g *GroupWalker) Walk(grouped map[int64][]interface{}, visitor GroupVisitor) {
	for _, group := range grouped {
		visitor.VisitHeader(group[0])
		groupSize := len(group)
		for i, file := range group {
			visitor.VisitNode(i, groupSize, group[0], file)
		}
	}
}
