package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Random distance between Earth and Mars is ")

	num := rand.Intn(401000000-56000000+1) + 56000000
	fmt.Print(num)
	fmt.Println(" kms.")
	// how long does it take to get to Mars?
	const lightSpeed = 299792
	const hrPerDay = 24

	var distance = 56000000
	time := int(travelTime(float64(distance), float64(lightSpeed)))

	fmt.Print("From Earth to Mars in: ")
	fmt.Println(time, "seconds with speed light")
	fmt.Print("When Earth and Mars are on the opposite sides of the Sun it will take: ")
	distance = 401000000
	time = int(travelTime(float64(distance), float64(lightSpeed)))
	fmt.Println(time, "seconds with speed light")
	const speed = 100800
	const spaceXDist = 96300000
	time = int(travelTime(float64(spaceXDist), float64(speed)))
	fmt.Println("The SpaceX Interplanetary Transport System")
	fmt.Print("It will take ")
	fmt.Println(time, "hours")
	fmt.Println(time/hrPerDay, "days")

	fmt.Print("To go to Mars within 28 days, the speed must be ")
	fmt.Print(int(bestSpeed(56000000, 28)))
	fmt.Println(" kms per hour.")

}

func travelTime(a, b float64) float64 {
	return a / b
}

func bestSpeed(a float64, b int) float64 {
	var hours = b * 24
	return a / float64(hours)
}
