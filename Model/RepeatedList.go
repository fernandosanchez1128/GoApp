package Model

/**
	Object to keep in shared memory repeated number during sort
*/
type RepeatedList struct {
	List []int
}

func NewRepeatedList() *RepeatedList {
	return &RepeatedList{List: make([]int, 0)}
}

func (r *RepeatedList) Add(value int) {
	r.List = append(r.List, value)
}
