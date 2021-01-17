package main

import "fmt"

// You throw a ball vertically upwards with an initial speed v (in km per hour).
// The height h of the ball at each time t is given by h = v*t - 0.5*g*t*t where g is Earth's gravity (g ~ 9.81 m/s**2).
// A device is recording at every tenth of second the height of the ball. For example with v = 15 km/h the device gets something of the following form:
// (0, 0.0), (1, 0.367...), (2, 0.637...), (3, 0.808...), (4, 0.881..) ... where the first number is the time in tenth of second and the second number the height in meter.
//
// Task
// Write a function max_ball with parameter v (in km per hour) that returns the time in tenth of second of the maximum height recorded by the device.

const Gravity = 9.81

func main() {
	fmt.Println(MaxBall(37))
}

func MaxBall(v int) int {
	for i := 0.0; ; i++ {
		if formula(float64(v), i) > formula(float64(v), i+1) {
			return int(i)
		}
	}
}

func formula(v, t float64) float64 {
	return 5*v/18*t/10 - 0.5*Gravity*t*t/100
}


//4.17×0, 3 −0.5×9.81×0.3×0.1    0, 3679
//4.17×0,3 −0.5×9.81×0.3×0.3     0, 80955
//4.17×0,4 −0.5×9.81×0.4×0.4     0, 8832
//4.17×0,5 −0.5×9.81×0.5×0.5     0, 85875
//4.17×0, 7 −0.5×9.81×0.7×0.7    0, 51555
