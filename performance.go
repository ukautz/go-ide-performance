package performance

var foo = 0

func Foo() int {
	if foo < 8 {
		foo ++
	} else {
		foo = 0
	}
	return foo
}