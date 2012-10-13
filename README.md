# Basics of Lisp In Go

[GoDoc](http://go.pkgdoc.org/github.com/howeyc/ligo)

Very early stages!!

Example:
```go
func IntAdd(vals ...interface{}) interface{} {
	var result int64
	for _, val := range vals {
		v := reflect.ValueOf(val)
		intVal := v.Int()
		result += intVal
	}
	return result
}

func main() {
	list1 := ligo.List(1, 2, 3, 5)
	vec1 := ligo.Vector(1, 2, 3, 5)
	fmt.Printf("Add: %d\n", ligo.Reduce(IntAdd, list1))
	fmt.Printf("Add: %d\n", ligo.Reduce(IntAdd, vec1))
}
```

[![Build Status](https://secure.travis-ci.org/howeyc/ligo.png?branch=master)](http://travis-ci.org/howeyc/ligo)

