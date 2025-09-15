package main

<<<<<<< HEAD
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
=======
import (
    "fmt"
)

func main() {
    stringList := []string{"spaceship", "sugar", "football", "sun", "paper"}
    fmt.Println(filterWords(stringList, "su"))
}

func filterWords(stringList []string, prefix string) []string {
    newStringList := []string{}

    // Outer loop: go through each word in the list
    for _, word := range stringList {

        // 2. [KEY TO PREVENT PANIC]
        // Make sure prefix is not longer than the word
        // (otherwise word[i] would panic because we go out of range)
        if len(prefix) > len(word) {
            continue
        }

        match := true

        // 1. Nested for loop: go through each character of the prefix
        for i := 0; i < len(prefix); i++ {

            // 1.2. Compare each character of prefix with the word
            if word[i] != prefix[i] {
                match = false
                break // stop checking this word
            }
        }

        // If all prefix characters matched, add word to results
        if match {
            newStringList = append(newStringList, word)
        }
    }

    return newStringList
}


>>>>>>> 5541719133f409ba1a5d98d4a1c5c44ea027da6c
