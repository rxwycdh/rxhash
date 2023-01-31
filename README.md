# rxhash [![GoDoc](https://godoc.org/github.com/rxwycdh/rxhash?status.svg)](https://godoc.org/github.com/rxwycdh/rxhash)

rxhash is a Go library for creating a unique hash value for struct in Go,  but **data consistency**.
> waht is data consistency?
>
> ```golang
> import "github.com/rxwycdh/rxhash"
> type A struct {
>   Values []int
> }
> 
> a1 := A{Values: []int{1,2,3}}
> a2 := A{Values: []int{3,2,1}} // unorder but content is equal a1
> 
> fmt.Println(HashStruct(a) == HashStruct(b)) // true
> ```
>
> **So This library is useful for detecting whether the contents of a structure have changed.**


## Installation

Standard `go get`:

```
$ go get github.com/rxwycdh/rxhash
```

## Qucik Start

A quick code example is shown below:

```go
type Student struct {
    Name    string
    Address []string
    School  School
}

type School struct {
    Labels   map[string]any
    Teachers []Teacher
}

type Teacher struct {
    Subjects []string
}

// TEST1: hash equal if you change the order of the array in struct
student1 := Student{
    Name:    "xiaoming",
    Address: []string{"mumbai", "london", "tokyo", "seattle"},
    School: School{
        Labels: map[string]any{
            "phone":   "123456",
            "country": "China",
        },
    Teachers: []Teacher{{Subjects: []string{"math", "chinese", "art"}}},
    },
}

student1UnOrder := Student{
    Name:    "xiaoming",
    Address: []string{"mumbai", "london", "seattle", "tokyo"}, // **change this order!!**
    School: School{
        Labels: map[string]any{
            "phone":   "123456",
            "country": "China",
        },
    Teachers: []Teacher{{Subjects: []string{"math", "chinese", "art"}}},
    },
}

s1, _ := HashStruct(student1)
s2, _ := HashStruct(student1UnOrder)
fmt.Printf("student1 hash: %s, student2 hash: %s, student1 == student2 ? -> %t \n", s1, s2, s1 == s2)
// Output:
// student1 hash: 744398b55ba132754147289b30955aa4, student2 hash: 744398b55ba132754147289b30955aa4, student1 == student2 ? -> true

// TEST2: different attr in struct to calc hash
student3 := Student{
    // Name is different from student1, student1UnOrder
    Name:    "xiaohong",
    Address: []string{"mumbai", "london", "seattle", "tokyo"},
    School: School{
        Labels: map[string]any{
            "phone":   "123456",
            "country": "China",
        },
    Teachers: []Teacher{{Subjects: []string{"math", "chinese", "art"}}},
    },
}

s3, _ := HashStruct(student3)
fmt.Printf("student3 hash: %s", s3)
// Output:
// student3 hash: 3f3bbdb3dcc6e7645ce30a7ef58c2e58
```
