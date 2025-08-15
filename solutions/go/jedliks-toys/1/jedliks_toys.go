package jedlik

import "fmt"

// type Car struct {
//     speed, batteryDrain, battery, distance int
// }

// TODO: define the 'Drive()' method
func (r *Car) Drive() {
    if r.battery < r.batteryDrain {
        return
    }
    r.distance += r.speed
    r.battery -= r.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (r *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", r.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (r *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", r.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (r *Car) CanFinish(trackDistance int) bool {
    numDrivesRemaining := r.battery / r.batteryDrain
    distancePossible := numDrivesRemaining * r.speed
    return distancePossible >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
