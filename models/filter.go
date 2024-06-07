package models

 const FilterStatus = "status"
 const FilterPriority = "priority"

type TaskFilter struct {
	Type  string
	Value string
}

func IsTaskExists(f []Task, t Task) bool {
	for _, item := range f {
		if item.Id == t.Id {
			return true
		}
	}
	return false
}

func (f TaskFilter) MatchTask(t Task) bool {
	switch f.Type {
	case FilterStatus:
		if t.Status == f.Value {
			return true
		}
	case FilterPriority:
		if t.Priority == f.Value {
			return true
		}
	}
	return false
}

func FilterTaskSlice(t []Task, f []TaskFilter) []Task {
	FilterTask := make([]Task, 0)
	for _, item := range t {
		for _, fitem := range f {
			if fitem.MatchTask(item) == true {
				if IsTaskExists(FilterTask, item) == false {
					FilterTask = append(FilterTask, item)
				}
			}
		}
	}
	return FilterTask
}