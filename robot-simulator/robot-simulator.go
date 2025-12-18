package robot

import "fmt"

// See defs.go for other definitions

// Step 1
// Define Dir type here.
// Define N, E, S, W here.
const (
	N = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	return fmt.Sprintf("%d", d)
}

// Step 2
// Define Action type here.
type Action byte

func StartRobot(command chan Command, action chan Action) {
	for value := range command {
		switch value {
		case 'R':
			action <- 'R'
		case 'L':
			action <- 'L'
		case 'A':
			action <- 'A'
		}
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		robot, _ = actionRobot(byte(act), &extent, robot)
	}
	report <- robot
}

func actionRobot(act byte, extent *Rect, robot Step2Robot) (Step2Robot, error) {
	var err error
	switch act {
	case 'R':
		robot.Dir = (robot.Dir + 1) % 4
	case 'L':
		robot.Dir = (robot.Dir + 3) % 4
	case 'A':
		err = advance(extent, &robot)
	}
	return robot, err
}

func advance(extent *Rect, robot *Step2Robot) error {
	switch robot.Dir {
	case N:
		if extent.Max.Northing > robot.Northing {
			robot.Northing++
			return nil
		}
	case E:
		if extent.Max.Easting > robot.Easting {
			robot.Easting++
			return nil
		}
	case S:
		if extent.Min.Northing < robot.Northing {
			robot.Northing--
			return nil
		}
	case W:
		if extent.Min.Easting < robot.Easting {
			robot.Easting--
			return nil
		}
	}
	return fmt.Errorf("could not advance")
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	Name string
	Move byte
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() { action <- Action3{Name: name, Move: 'F'} }()
	if name == "" {
		log <- "invalid name"
		return
	}
	for _, v := range script {
		ch := byte(v)
		if ch != 'R' && ch != 'L' && ch != 'A' {
			log <- "invalid command"
			return
		}
		action <- Action3{Name: name, Move: ch}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer func() { rep <- robots }()
	if checkDuplicateNames(robots) {
		log <- "duplicate robot names"
		return
	}
	if checkSamePositions(&extent, robots) {
		log <- "same position"
		return
	}
	counter := 0
	for counter != len(robots) {
		acc := <-action
		if acc.Move == 'F' {
			counter++
			continue
		}
		found := false
		for i, robot := range robots {
			if robot.Name == acc.Name {
				step2Robot, err := actionRobot(acc.Move, &extent, robot.Step2Robot)
				if err != nil {
					log <- "rect collision"
				} else if robotCollision(robots, robot.Name, step2Robot) {
					log <- "robot collision"
				} else {
					robots[i].Step2Robot = step2Robot
				}
				found = true
				break
			}
		}
		if !found {
			log <- "robot not found"
			return
		}
	}
}

func checkDuplicateNames(robots []Step3Robot) bool {
	counter := make(map[string]bool, len(robots))
	for _, value := range robots {
		if counter[value.Name] {
			return true
		}
		counter[value.Name] = true
	}
	return false
}

func checkSamePositions(extent *Rect, robots []Step3Robot) bool {
	counter := make(map[Pos]bool, len(robots))
	for _, value := range robots {
		if counter[value.Pos] {
			return true
		}
		if value.Northing > extent.Max.Northing || value.Northing < extent.Min.Northing ||
			value.Easting > extent.Max.Easting || value.Easting < extent.Min.Easting {
			return true
		}
		counter[value.Pos] = true
	}
	return false
}

func robotCollision(robots []Step3Robot, name string, robot Step2Robot) bool {
	for _, value := range robots {
		if value.Name != name && value.Step2Robot.Pos == robot.Pos {
			return true
		}
	}
	return false
}
