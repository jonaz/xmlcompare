package xmlcompare

type byName []Node

func (l byName) Less(i, j int) bool {
	return l[i].XMLName.Space+l[i].XMLName.Local < l[j].XMLName.Space+l[j].XMLName.Local
}
func (l byName) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l byName) Len() int      { return len(l) }
