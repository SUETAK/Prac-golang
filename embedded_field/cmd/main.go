package main

import "fmt"

func main() {
	mom := newMom("サザエ")
	kid := newKid("タラオ", "タラちゃん")
	fmt.Println(mom.Hello())
	fmt.Println(kid.Hello())
	fmt.Println(kid.Unchi())

	hs := []Helloer{
		newMom("母"),
		newKid("たらこ", "おかめ"),
	}

	for _, h := range hs {
		fmt.Println(Sing(h))
	}
}

type Helloer interface {
	Hello() string
}

func Sing(h Helloer) string {
	return h.Hello() + " 歌っている表現を入れる"
}

type Mom struct {
	name string
}

func newMom(n string) *Mom {
	m := new(Mom)
	m.name = n
	return m
}

func (m *Mom) Hello() string {
	return m.name + "っていう名前です"
}

func (m *Mom) Unchi() string {
	return m.name + "が苗字です"
}

type Kid struct {
	Mom      //embedded field
	nickname string
}

func newKid(n, m string) *Kid {
	k := new(Kid)
	k.name = n
	k.nickname = m
	return k
}
