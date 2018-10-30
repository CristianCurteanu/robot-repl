# Super Robot

This is a project written in Golang from scratch for a toy robot. It is an interpreter for a DSL which takes commands as input,  process it and makes some calculations for a robot on a board.

# Installation

You can download this repository, and if you have the Golang compiler installed just run:
```bash
$ go run main.go
```

This will launch the interactive shell.

# Instructions

  - `PLACE` - this instruction places robot on a spific location on board. This would be defined by `x` and `y` points, which come as arguments, and the orientation of the robot, by third argument of this instruction. Ex: `PLACE 1, 2, LEFT` will place robot on `x` = 1, `y` = 2 location, with the orientation on `LEFT`;
  - `MOVE` - this instruction will make the robot to move forward
  - `LEFT` - this instruction will turn the robot to 90 degrees on left
  - `RIGHT` - this instruction will turn the robot to 90 degrees on right
  - `REPORT` - this instruction will print the robot's current location and orientation

# Launch from a file

# API
