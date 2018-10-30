package robot

import (
	"fmt"
	"strconv"
	"strings"

	"../board"
)

type Robot struct {
	X           int64
	Y           int64
	Orientation string
	Map         board.Board
	Command     []string
}

var validOrientations []string
var maneuvers = make(map[string]map[string]string)

func init() {
	validOrientations = []string{"UP", "DOWN", "LEFT", "RIGHT"}
	// I know that it could be not so elegant, but it's efficient though
	maneuvers["UP"] = make(map[string]string)
	maneuvers["UP"]["LEFT"] = "LEFT"
	maneuvers["UP"]["RIGHT"] = "RIGHT"

	maneuvers["LEFT"] = make(map[string]string)
	maneuvers["LEFT"]["LEFT"] = "DOWN"
	maneuvers["LEFT"]["RIGHT"] = "UP"

	maneuvers["DOWN"] = make(map[string]string)
	maneuvers["DOWN"]["LEFT"] = "RIGHT"
	maneuvers["DOWN"]["RIGHT"] = "LEFT"

	maneuvers["RIGHT"] = make(map[string]string)
	maneuvers["RIGHT"]["LEFT"] = "UP"
	maneuvers["RIGHT"]["RIGHT"] = "DOWN"
	fmt.Println(maneuvers)
}

func NewRobot(name string) Robot {
	return Robot{X: 0, Y: 0, Orientation: "RIGHT"}
}

func (robot *Robot) LandOn(board board.Board) {
	robot.Map = board
}

func (robot *Robot) Execute(command string) error {
	command = strings.Split(strings.ToUpper(command), "--")[0] // Ensure to skip comments, by `--` syntax
	robot.Command = strings.Split(command, " ")
	err := robot.ValidateCommand()
	if err != nil {
		return err
	}
	robot.TakeAction()
	return nil
}

type WrongCommandError struct {
	Command string
}

func (ex *WrongCommandError) Error() string {
	return fmt.Sprintf("Wrong Typed Command: %s", ex.Command)
}

type ValidationError struct {
	Message string
}

func (ex *ValidationError) Error() string {
	return fmt.Sprintf("%s", ex.Message)
}

func (robot *Robot) ValidateCommand() error {
	if robot.Command[0] == "PLACE" {
		x, err := strconv.ParseInt(strings.Trim(robot.Command[1], ","), 10, 64)
		if err != nil {
			return &ValidationError{"Cannot parse X coordinate from: " + robot.Command[1]}
		}

		y, err := strconv.ParseInt(strings.Trim(robot.Command[2], ","), 10, 64)
		if err != nil {
			return &ValidationError{"Cannot parse X coordinate from: " + robot.Command[2]}
		}

		if robot.Map.Width < y || robot.Map.Height < x {
			return &ValidationError{"Wrong coordinates introduced!"}
		}

		if !validateOrientation(robot.Command[3]) {
			return &ValidationError{"Wrong orientation introduced!"}
		}

		robot.X = x
		robot.Y = y
		robot.Orientation = robot.Command[3]
		return nil
	} else if !validCommand(robot.Command[0]) {
		return &ValidationError{"Unknown command: " + robot.Command[0]}
	}
	return nil
}

func (robot *Robot) TakeAction() {
	if robot.Command[0] == "PLACE" {
		fmt.Println("command:", robot.Command)
		x, _ := strconv.ParseInt(strings.Trim(robot.Command[1], ","), 10, 64)
		y, _ := strconv.ParseInt(strings.Trim(robot.Command[2], ","), 10, 64)

		robot.X = x
		robot.Y = y
		robot.Orientation = robot.Command[3]
		return
	} else if robot.Command[0] == "MOVE" {
		if robot.Orientation == "UP" && robot.X < robot.Map.Height {
			robot.X++
		} else if robot.Orientation == "DOWN" && robot.X > 0 {
			robot.X--
		} else if robot.Orientation == "RIGHT" && robot.Y < robot.Map.Width {
			robot.Y++
		} else if robot.Orientation == "LEFT" && robot.Y < 0 {
			robot.Y--
		}
		return
	} else if robot.Command[0] == "LEFT" || robot.Command[0] == "RIGHT" {
		robot.Orientation = maneuvers[robot.Orientation][robot.Command[0]]
		return
	} else if robot.Command[0] == "EXIT" {

	}
}

func (robot *Robot) IsPrintingCommand() bool {
	return robot.Command[0] == "REPORT"
}

func (robot Robot) PrintPosition() string {
	return fmt.Sprintf("X: %d, Y: %d, Orientation: %s", robot.X, robot.Y, robot.Orientation)
}

func validCommand(command string) bool {
	for _, e := range []string{"MOVE", "LEFT", "RIGHT", "REPORT", "EXIT", ""} {
		if command == e {
			return true
		}
	}
	return false
}

func validateOrientation(target string) bool {
	for _, e := range validOrientations {
		if e == target {
			return true
		}
	}
	return false
}
