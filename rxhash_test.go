package rxhash

import (
	"fmt"
	"testing"
)

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

func TestHashStruct(t *testing.T) {
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
		Address: []string{"mumbai", "london", "seattle", "tokyo"},
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
}
