package main

import "fmt"

type Item struct {
	Name string
	Qty  int
}

type Cart struct {
	Items []Item
	Owner *string
}

func addItem(c Cart, it Item) {
	c.Items = append(c.Items, it)
	c.Items[0].Qty = 9
}

func main() {
	owner := "tom"
	c1 := Cart{Items: make([]Item, 1, 2), Owner: &owner}
	c1.Items[0] = Item{"a", 1}

	p := &c1.Items[0]

	addItem(c1, Item{"b", 2})
	fmt.Println(c1.Items, len(c1.Items))

	c2 := c1
	c2.Items = append(c2.Items, Item{"c", 3})
	c2.Items = append(c2.Items, Item{"d", 4})
	c2.Items[0].Qty = 7
	*c2.Owner = "jerry"
	fmt.Println(c1.Items[0].Qty, p.Qty, *c1.Owner)

	fmt.Println(p == &c1.Items[0], p == &c2.Items[0], c1.Owner == c2.Owner)

	fmt.Println(c1.Items[:2])
}
