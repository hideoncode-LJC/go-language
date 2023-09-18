package const 

const (
	n1 = 1 + iota
	n2 = 2 + iota
	n3
	n4
	n5
)


func PrintIota() {
	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)
	fmt.Println("n4", n4)
	fmt.Println("n5", n5)
}