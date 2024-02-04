package main

import "fmt"

type FrameType int

func (ft FrameType) GetName() string {
	switch ft {
	case Light:
		return "Light"
	case Medium:
		return "Medium"
	case Heavy:
		return "Heavy"
	default:
		return "Unknown"
	}
}

const (
	Light FrameType = iota
	Medium
	Heavy
)

type frame struct {
	frameType FrameType
	structure int
	hull      int
	speed     float64
	actions   int
	pilot     string
}

func NewFrame(frameType FrameType, pilot string) (frame, error) {
	switch frameType {
	case Light:
		return frame{
			frameType: frameType,
			structure: 50,
			hull:      25,
			speed:     10.0,
			actions:   2,
			pilot:     pilot,
		}, nil
	case Medium:
		return frame{
			frameType: frameType,
			structure: 100,
			hull:      50,
			speed:     7.5,
			actions:   3,
			pilot:     pilot,
		}, nil
	case Heavy:
		return frame{
			frameType: frameType,
			structure: 150,
			hull:      75,
			speed:     5.0,
			actions:   4,
			pilot:     pilot,
		}, nil
	default:
		return frame{}, fmt.Errorf("invalid frame type")
	}
}

func (f frame) PrintDebugInfo() {
	fmt.Println("Frame Type:", f.frameType.GetName())
	fmt.Println("Pilot:", f.pilot)
	fmt.Println("Structure:", f.structure)
	fmt.Println("Hull:", f.hull)
	fmt.Println("Speed:", f.speed)
	fmt.Println("Actions:", f.actions)
}
