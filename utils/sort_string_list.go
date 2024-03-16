package utils

type SortStringList []string

func (s SortStringList) Len() int {
	return len(s)
}

func (s SortStringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortStringList) Less(i, j int) bool {
	return s[i] < s[j]
}
