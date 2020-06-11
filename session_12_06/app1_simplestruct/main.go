package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type car struct {
	cleanness   int  // 0..100
	speed       int  // 0..200
	position    int  //
	isStarted   bool //
	areLightsOn bool //
	isLocked    bool //
}

// move
// accelerate <value>
// brake <value>
// clean
// start_stop
// lights_on
// lights_off

func main() {
	var c *car

	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter commands to execute")

	for s.Scan() {
		cmd := s.Text()

		if cmd == "quit" {
			break
		}

		if cmd == "new" {
			c = newCar()
			outputCar(c)
			continue
		}

		if c == nil {
			fmt.Println("car was not initialized")
			continue
		}

		tokens := strings.Split(cmd, " ")

		var result bool

		switch tokens[0] {
		case "move":
			result = move(c)
		case "clean":
			result = clean(c)
		case "start_stop":
			result = startStopEngine(c)
		case "lights_on":
			result = lightsOn(c)
		case "lights_off":
			result = lightsOff(c)

		case "accelerate", "brake":
			if len(tokens) < 2 {
				fmt.Println("some of parameters is missing:")
				break
			}

			value, err := strconv.Atoi(tokens[1])

			if err != nil {
				fmt.Println("some of parameters cannot be parsed:", err)
				break
			}

			if tokens[0] == "accelerate" {
				result = accelerate(c, value)
			} else {
				result = brake(c, value)
			}

		default:
			fmt.Println("unrecognized command")
		}

		if result {
			outputCar(c)
		} else {
			fmt.Println("command failed")
		}
	}

	// reference type - default value is nil
	// pointers
	// slices - partially
	// maps
	//
	// functions
}

func newCar() *car {
	return &car{
		cleanness: 100,
	}
}

// requirements gathering and analysis
// design and architecture
// development
// testing
// deployment
// operation

func move(c *car) bool {
	if !c.isStarted {
		return false
	}

	if c.speed == 0 {
		return true
	}

	c.position += c.speed

	if c.cleanness > 0 {
		c.cleanness -= 5
	}

	return true
}

func accelerate(c *car, value int) bool {
	if !c.isStarted {
		return false
	}

	c.speed += value

	return true
}

func clean(c *car) bool {
	c.cleanness = 100
	return true
}

func brake(c *car, value int) bool {
	if !c.isStarted {
		return true
	}

	if c.speed > 0 {
		c.speed -= value
		if c.speed < 0 {
			c.speed = 0
		}
	}

	return true
}

func startStopEngine(c *car) bool {
	if !c.isStarted {
		c.isStarted = true
		return true
	}

	if c.speed > 0 {
		return false
	}

	c.isStarted = false

	return true
}

func lightsOn(c *car) bool {
	c.areLightsOn = true
	return true
}

func lightsOff(c *car) bool {
	c.areLightsOn = false
	return true
}

func carToString(c *car) string {
	return fmt.Sprintf(
		"{cleanness: %d, speed: %d, position: %d, isStarted: %t, areLightsOn: %t, islocked: %t}",
		c.cleanness, c.speed, c.position, c.isStarted, c.areLightsOn, c.isLocked)
}

func outputCar(c *car) {
	fmt.Println(carToString(c))
}

type doner struct {
	isMeat   bool
	isBread  bool
	isBig    bool
	progress int // 0..100
}

func newDoner(isMeat bool, isBread bool, isBig bool) *doner {
	return &doner{
		isMeat:   isMeat,
		isBread:  isBread,
		isBig:    isBig,
		progress: 100,
	}
}

func copyDoner(d *doner) *doner {
	copy := newDoner(d.isMeat, d.isBread, d.isBig)
	copy.progress = d.progress
	return copy
}

func eatDoner(d *doner) {
	if d.progress == 0 {
		return
	}

	d.progress -= 20
}

type oopDoner struct {
	isMeat   bool
	isBread  bool
	isBig    bool
	progress int
}

func newOopDoner(isMeat bool, isBread bool, isBig bool) *oopDoner {
	return &oopDoner{
		isMeat:   isMeat,
		isBread:  isBread,
		isBig:    isBig,
		progress: 100,
	}
}

func (d *oopDoner) copy() *oopDoner {
	copy := newOopDoner(d.isMeat, d.isBread, d.isBig)
	copy.progress = d.progress
	return copy
}

func (d *oopDoner) eat() {
	if d.progress == 0 {
		return
	}

	d.progress -= 20
}

func test() {
	d := newOopDoner(true, true, true)

	d.eat() // eatDoner(d)
}
