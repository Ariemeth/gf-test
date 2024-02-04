package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Starting Frame Assault Test")

	// Ask the user for the frame type to create
	userFrameType, err := getFrameSizeFromUser()
	if err != nil {
		fmt.Println(err.Error(), "Exiting...")
		return
	}

	userFrame, err := NewFrame(userFrameType, "User")
	if err != nil {
		fmt.Println(err.Error(), "Exiting...")
		return
	}

	userFrame.PrintDebugInfo()

	fmt.Println("--------------------") // Separator line

	randomFrame, err := generateRandomFrame("Random Pilot")
	if err != nil {
		fmt.Println(err.Error(), "Exiting...")
		return
	}

	randomFrame.PrintDebugInfo()
}

func getFrameSizeFromUser() (FrameType, error) {
	var input int
	fmt.Println("Enter the frame type (0 for Light, 1 for Medium, 2 for Heavy):")
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, err
	}

	switch input {
	case 0:
		return Light, nil
	case 1:
		return Medium, nil
	case 2:
		return Heavy, nil
	default:
		return 0, fmt.Errorf("invalid frame type")
	}
}

func generateRandomFrame(pilot string) (frame, error) {
	randomFrameType := FrameType(rand.Intn(3))
	randomFrame, err := NewFrame(randomFrameType, pilot)
	if err != nil {
		return frame{}, err
	}
	return randomFrame, nil
}
