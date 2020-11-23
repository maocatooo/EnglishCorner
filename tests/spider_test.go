package tests

import (
	"EnglishCorner/spider/youdict"
	"fmt"
	"sync"
	"testing"
)

var s sync.WaitGroup

func TestGetBody(t *testing.T) {
	got, err := youdict.GetBody("hi")
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(got)
}

func TestFind(t *testing.T) {
	m, e := youdict.Find("hi")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("%#v \n", m)
	fmt.Print(m.MSounds, m.YSounds)
}

func TestFindMany(t *testing.T) {

	for _, v := range []string{
		"Spelling", "beleeve", "is", "common", "till",
	} {
		s.Add(1)
		go func() {
			m, _ := youdict.Find(v)
			fmt.Printf("%#v \n", m)
			s.Done()
		}()

	}
	s.Wait()
}
