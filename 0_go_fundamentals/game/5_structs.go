package game

import (
	"fmt"
	"slices"
)

const (
	maxX = 600
	maxY = 800
)

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

func Game() {
	var it Item
	fmt.Println("i: %#v\n", it)

	fmt.Println(NewItem(10, 12))
	fmt.Println(NewItem(10, 12000))
	it = Item{10, 12} // Not flexible; Provide in order and no optionals
	fmt.Println("i: %#v\n", it)

	it = Item{
		Y: 23,
	}
	fmt.Println("i: %#v\n", it)

	it.Move(10, 20)
	fmt.Printf("i Move: %#v\n", it)

	p1 := Player{
		Name: "Rahul",
	}
	fmt.Printf("P1: %+v\n", p1)
	// Embedded struct's fields on accessible at the top level
	fmt.Println("P1.X", p1.X)
	p1.Move(100, 200)
	fmt.Printf("p1 item Move: %#v\n", p1.Item)

	// Exercise: Add new field of type slice, allow only 3 keys, no duplicates
	fmt.Println(p1.Found(Copper))
	fmt.Println(p1.Found(Copper))
	fmt.Println(p1.Found(Key(7)))
	fmt.Println("keys", p1.Keys)

	// Use %#v for debugging/logging
	// a, b := 1, "1"
	// fmt.Printf("a=%v, b=%v\n", a, b)
	// fmt.Printf("a=%#v, b=%#v\n", a, b)

	fmt.Println("Interfaces: Moving all movables for 100x and 1y")
	movables := []Mover{
		&it,
		&p1,
	}
	moveAll(movables, 100, 1)
	fmt.Println("i: %#v\n", it)
	fmt.Println("i: %#v\n", p1)

	fmt.Println("Empty interface")
	emptyInterface()

}

func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

/*
Factory functions
func NewItem(x, y int) Item
func NewItem(x, y int) *Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)
*/

func NewItem(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d out of the bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return i, nil
}

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item
	Keys []Key
}

/*
Interfaces
Set of methods signatures & types
No implements keyword. No inheritance. Behavior based typing
Structural typing - Go checks for all methods required by the interface in compile time
We define interfaces as "what you need", not "what you provide"
Why they exists,
 - Decoupling
 - Dependency inversion
*/

type Mover interface {
	Move(int, int)
}

/*
In many languages -> Pointers can always be passed down. Never up

For ex., c example
int* make_number() {
    int x = 42;
    return &x;
}
int main() {
	int* p = make_number();
	printf("%d\n", *p);
}
Compilation suceeds; But behavior is invalid
x will be a stack variable
	------
In Go, it's allowed. Go compiler does ESCAPE ANALYSIS - x escapes to heap
Pointer will be added in heap
* Memory outlives the function
* Pointer remains valid
Garbage Collector cleans it up later
*/

/*
time package works with VALUE semantics
Value semantics: everyone has their own copy
Pointer semantics: everone share the same copy (heap, requires locking)
*/

/*
Value vs Pointer receiver
Use value semantics in general
  - Try to keep same semantics on all methods

When you must use pointer receiver
  - If you have lock field
  - If you need to mutate the struct
  - when decoding & unmarshalling
*/
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
		// OK
	default:
		return fmt.Errorf("unknown key: %q", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}

func emptyInterface() {
	var a any

	a = 7
	fmt.Println("1. a ", a)

	a = "seven"
	fmt.Println("2. a ", a)

	/*
		Don't use any.
		Exceptions:
		- Serialization
		- Printing
	*/

	// type assertion
	s := a.(string)
	fmt.Println("s ", s)

	// comma ok expressions
	i, ok := a.(int32)
	if !ok {
		fmt.Printf("Not an integer (%T)!\n", a)
	} else {
		fmt.Println("i ", i)
	}

	switch a.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("Other type - %T \n", a)
	}
}
