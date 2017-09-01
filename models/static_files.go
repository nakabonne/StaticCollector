package models

type StaticFiles []*StaticFile

func (p StaticFiles) Len() int {
	return len(p)
}

func (p StaticFiles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p StaticFiles) Less(i, j int) bool {
	return !p[i].TargetDay.After(p[j].TargetDay)
}
