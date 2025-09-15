package main

import "log"

///
/// STUDENT
///

type Student struct {
	Name      string
	Age       int
	Grade     int
	ID        int
	Classroom int
	Marks     []int
}

func (s Student) calculateGrade() int {
	return (s.Marks[0] + s.Marks[1]) / 2
}

func (s *Student) updateGrade() {
	s.Grade = s.calculateGrade()
}

///
/// OTHER
///

func assignClassroom(s *Student) {
	s.Classroom = 2
}

///
/// MAIN
///

func main() {
	jack := Student{Name: "jack"}
	jack.updateGrade()
	log.Println(jack.Classroom)

	assignClassroom(jack)
	//jack.Classroom = assignClassroom(jack)

	// var anna *Student
	// anna = &Student{}
	// anna.Name = "Anna"
}

// Homework:
// 1. Create new struct vehicle with some values (engineStatus = string)
// 2. Create method fixEngine which will update the status
// 3. Create a function which changes one engine of one vehicle to another (kind of like a swap)
// 4. Create a function which finds the best vehicle in a list of vehicles
