package main

import "fmt"

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

type ScoreUpdater struct {
	score        int
	observerList []Observer
}

type Observer interface {
	Update(value int)
}

type Gamer struct {
	name  string
	score int
	Pro   Subject
}

func NewGamer(ss Subject, name string) *Gamer {
	newGamer := &Gamer{
		name: name,
		Pro:  ss,
	}
	newGamer.Pro.RegisterObserver(newGamer)
	return newGamer
}

func (u *Gamer) Update(score int) {
	u.score = score
	u.display()

}

func (u *Gamer) display() {
	fmt.Printf("I believe you can do more! %s received score updates:%d\n", u.name, u.score)
}

func NewScoreUpdater() *ScoreUpdater {
	return &ScoreUpdater{
		score:        0,
		observerList: make([]Observer, 0),
	}
}

func (su *ScoreUpdater) RegisterObserver(o Observer) {
	su.observerList = append(su.observerList, o)
}

func (su *ScoreUpdater) RemoveObserver(o Observer) {
	found := false
	i := 0
	for ; i < len(su.observerList); i++ {
		if su.observerList[i] == o {
			found = true
			break
		}
	}
	if found {
		su.observerList = append(su.observerList[:i], su.observerList[i+1:]...)
	}
}

func (su *ScoreUpdater) NotifyObserver() {
	for _, observer := range su.observerList {
		observer.Update(su.score)
	}
}

func (su *ScoreUpdater) SetValue(value int) {
	su.score = value
	su.NotifyObserver()
}

func main() {
	simpleSubject := NewScoreUpdater()
	NewGamer(simpleSubject, "Nurym")
	NewGamer(simpleSubject, "Adil")

	c := NewGamer(simpleSubject, "Gaziz")

	simpleSubject.SetValue(50)
	simpleSubject.RemoveObserver(c)

	simpleSubject.SetValue(60)
}
