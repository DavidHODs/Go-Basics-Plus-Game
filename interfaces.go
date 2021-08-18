package main

import "fmt"

func main() {
	fb := new(Facebook)
	twtr := new(Twitter)
	lnkdin := new(Linkedln)

	ScrollFeeds(twtr, fb, lnkdin)
}

func ScrollFeeds(platforms ...SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("************************************")
	}
}

type SocialMedia interface {
	Feed() []string
	Fame() int
}

type Facebook struct {
	URL     string
	Name    string
	Friends int
}

type Twitter struct {
	URL     string
	Name    string
	Friends int
}

type Linkedln struct {
	URL     string
	Name    string
	Friends int
}

func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (f *Facebook) Fame() int {
	return f.Friends
}

func (l *Linkedln) Feed() []string {
	return []string{
		"Linkedln feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (l *Linkedln) Fame() int {
	return l.Friends
}

func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (t *Twitter) Fame() int {
	return t.Friends
}
