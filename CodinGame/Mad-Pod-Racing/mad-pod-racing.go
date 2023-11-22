package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var laps int
	fmt.Scan(&laps)

	var checkpointCount int
	fmt.Scan(&checkpointCount)

	for i := 0; i < checkpointCount; i++ {
		var checkpointX, checkpointY int
		fmt.Scan(&checkpointX, &checkpointY)
	}
	for {
		for i := 0; i < 2; i++ {
			// x: x position of your pod
			// y: y position of your pod
			// vx: x speed of your pod
			// vy: y speed of your pod
			// angle: angle of your pod
			// nextCheckPointId: next check point id of your pod
			var x, y, vx, vy, angle, nextCheckPointId int
			fmt.Scan(&x, &y, &vx, &vy, &angle, &nextCheckPointId)
		}
		for i := 0; i < 2; i++ {
			// x2: x position of the opponent's pod
			// y2: y position of the opponent's pod
			// vx2: x speed of the opponent's pod
			// vy2: y speed of the opponent's pod
			// angle2: angle of the opponent's pod
			// nextCheckPointId2: next check point id of the opponent's pod
			var x2, y2, vx2, vy2, angle2, nextCheckPointId2 int
			fmt.Scan(&x2, &y2, &vx2, &vy2, &angle2, &nextCheckPointId2)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// You have to output the target position
		// followed by the power (0 <= power <= 200)
		// i.e.: "x y power"
		fmt.Println("8000 4500 100")
		fmt.Println("8000 4500 100")
	}
}
