package main

import (
	"fmt"
	"sort"
)

type priceSt  struct {
	p float64
	t    int64
}

type priceStList []priceSt
func (p priceStList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p priceStList) Len() int           { return len(p) }
func (p priceStList) Less(i, j int) bool { return p[i].p < p[j].p }

func main() {
	a := priceSt{p:1.65,t:111}
	b := priceSt{p:1.66,t:112}
	c := priceStList{}
	c = append(c,a)
	c = append(c,b)

	sort.Sort(c)
	fmt.Println(c[0])
}

