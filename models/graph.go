package models

import (
	"fmt"
)

type Tree struct {
	Name     string
	Children []*Tree
}

func (t *Tree) String() string {
	return fmt.Sprintf("%s\n%s", t.Name, t.dfs(0, 4, ""))
}

func (t *Tree) dfs(depth, maxdepth int, prefix string) (out string) {
	if depth > maxdepth {
		return ""
	}
	for i, child := range t.Children {
		if i == len(t.Children)-1 {
			out = out + fmt.Sprintf("%s%s\n", prefix+"└─ ", child.Name)
			out = out + child.dfs(depth+1, maxdepth, prefix+"   ")
		} else {
			out = out + fmt.Sprintf("%s%s\n", prefix+"├─ ", child.Name)
			out = out + child.dfs(depth+1, maxdepth, prefix+"│  ")
		}
	}
	return out
}
