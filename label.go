package yalzo

type LabelList []string

func (l *LabelList) GetListLength() int {
	return len(*l)
}

func (l *LabelList) GetList(width int) []string {
	return *l
}

func (l *LabelList) IsInAryRange(i int) bool {
	if i < 0 || len(*l)-1 < i {
		return false
	}

	return true
}

func (l *LabelList) Contains(label string) (int, bool) {
	for i, v := range *l {
		if label == v {
			return i, true
		}
	}

	return 0, false
}

func (l *LabelList) Add(label string) int {
	if _, ok := l.Contains(label); !ok && label != "" {
		*l = append(*l, label)
	}

	return len(*l)
}

func (l *LabelList) Remove(i int) interface{} {
	if !l.IsInAryRange(i) {
		return ""
	}

	delValue := (*l)[i]
	*l = append((*l)[:i], (*l)[i+1:]...)

	return delValue
}

func (l *LabelList) GetPresentName(i int) string {
	if !l.IsInAryRange(i) {
		return ""
	}

	return (*l)[i]
}

func (l *LabelList) Rename(i int, name string) {
	if !l.IsInAryRange(i) || name == "" {
		return
	}

	(*l)[i] = name
}

func (l *LabelList) Exchange(i, j int) {
	if !l.IsInAryRange(i) || !l.IsInAryRange(j) {
		return
	}

	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *LabelList) ChangeLabel(i int, name string) {
	l.Rename(i, name)
}
