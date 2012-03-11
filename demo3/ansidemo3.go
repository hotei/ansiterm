// demo3.go
// (c) 2012 David Rook 

package main

import (
	"ansiterm"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var license = "(c) 2012 David Rook - released under Simplified FreeBSD license"

var sensors []chan int
var fields []*Field

var runfor = 1 // minutes to run demo

type Field struct {
	Tag       string
	Data      string
	Row, Col  int
	IsVisible bool
}

func init() {
	sensors = make([]chan int, 0)
}

func tempSensor(t chan int) {
	tempreading := 32
	for {
		sleepytime := rand.Int31n(10)
		time.Sleep(time.Duration(sleepytime) * time.Second)
		newtemp := int(rand.Int31n(100) + 32)
		tempreading = (tempreading*95 + newtemp*5) / 100
		t <- tempreading
	}
}

func seekerrSensor(s chan int) {
	totalerrs := 0
	for {
		sleepytime := 1
		time.Sleep(time.Duration(sleepytime) * time.Second)
		if rand.Int31n(2) == 1 {
			totalerrs++
		}
		s <- totalerrs
	}
}

func clockSensor(c chan int) {
	ticks := 0
	for {
		sleepytime := 1
		time.Sleep(time.Duration(sleepytime) * time.Second)
		ticks++
		c <- ticks
	}
}

func startSensors() {
	t := make(chan int)
	sensors = append(sensors, t)
	go tempSensor(t)

	s := make(chan int)
	sensors = append(sensors, s)
	go seekerrSensor(s)

	c := make(chan int)
	sensors = append(sensors, c)
	go clockSensor(c)

	fmt.Printf("startSensors() fini\n")
}

func initFields() {
	var temp = Field{"Drive Temp(F):", "", 10, 10, true}
	fields = append(fields, &temp)

	var seekerrs = Field{"Count seek errors:", "", 11, 10, true}
	fields = append(fields, &seekerrs)

	var clock = Field{"TickTock:", "", 12, 10, true}
	fields = append(fields, &clock)

	fmt.Printf("initFields() fini\n")
}

// leaves cursor at end of last field updated
func (f *Field) Show(s string) {
	ansiterm.MoveToRC(f.Row, f.Col)
	ansiterm.Erase(len(f.Tag) + len(f.Data))
	f.Data = s
	fmt.Printf("%s%s", f.Tag, f.Data)
}

func main() {
	ansiterm.ResetTerm(0)
	ansiterm.ClearPage()
	initFields()
	startSensors()
	runfor = 1
	done := time.After(time.Duration(runfor) * time.Minute)
	fmt.Printf("This demo will stop after %d minutes\n", runfor)

L1:
	for {
		select {
		case t := <-sensors[0]:
			fields[0].Show(fmt.Sprintf("%d", t))
		case t := <-sensors[1]:
			fields[1].Show(fmt.Sprintf("%d", t))
		case t := <-sensors[2]:
			fields[2].Show(fmt.Sprintf("%d", t))
		// note: label is required else it 'breaks' the select, not the for
		case _ = <-done:
			break L1
		}
	}
	if false {
		os.Exit(0)
	}
	ansiterm.ClearPage()
}
