package main

import "fmt"
// Homework:
// 1. Create new struct vehicle with some values (engineStatus = string)
// 2. Create method fixEngine which will update the status
// 3. Create a function which changes one engine of one vehicle to another (kind of like a swap)
// 4. Create a function which finds the best vehicle in a list of vehicles

type Vehicle struct{
    Brand string
    Year int
    EngineStatus string
    Miles int
}



func (v *Vehicle) fixEngine( fix string)  {

    if v.EngineStatus != "Fixed"{
        fix = v.EngineStatus
        fmt.Println("New engine status is", fix)
    } else{
        fmt.Println("Engine status is already", v.EngineStatus)
    }

}

func changeEngine(vehicleOne, vehicleTwo *Vehicle) {
    vehicleThree:= vehicleOne.EngineStatus
    vehicleOne.EngineStatus = vehicleTwo.EngineStatus
    vehicleTwo.EngineStatus = vehicleThree


}

func findBestCar(vehicle []Vehicle) Vehicle{
   
    bestCar:= vehicle[0]
    for _, v := range vehicle{
        if v.Year > bestCar.Year{
            bestCar = v
        }
    }

 
    return bestCar
}

func main() {

    car1:= Vehicle{"Toyota", 2025, "Broken", 20000}
    car2:= Vehicle{"Ford", 2024, "Fixed", 30000}

    car1.fixEngine("Broken")
    changeEngine(&car1, &car2)
    fmt.Println("New engine status for Car 1 is", car1.EngineStatus)
    fmt.Println("New engine status for Car 2 is", car2.EngineStatus)

 
    testCars:= []Vehicle{car1,car2}
    result:= findBestCar(testCars)
    fmt.Printf("The best car is %s as it is the most recent", result.Brand)
    
}
// Homework:
// 1. Create new struct vehicle with some values (engineStatus = string)
// 2. Create method fixEngine which will update the status
// 3. Create a function which changes one engine of one vehicle to another (kind of like a swap)
// 4. Create a function which finds the best vehicle in a list of vehicles