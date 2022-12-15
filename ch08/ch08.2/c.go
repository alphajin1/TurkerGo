package main

import "fmt"

// BitFlag

type Room uint8

const (
	MasterRoom Room = 1 << iota
	BathRoom
	Kitchen
	LivingRoom
)

func SetLight(rooms, room Room) Room {
	return rooms | room
}

func ResetLight(rooms, room Room) Room {
	return rooms &^ room
}

func IsTurnOn(rooms, room Room) bool {
	return rooms&room == room
}

func TurnOnLights(rooms Room) {
	if IsTurnOn(rooms, MasterRoom) {
		fmt.Println("Turn on MasterRoom Light")
	}

	if IsTurnOn(rooms, BathRoom) {
		fmt.Println("Turn on BathRoomLight")
	}

	if IsTurnOn(rooms, Kitchen) {
		fmt.Println("Turn on Kitchen Light")
	}

	if IsTurnOn(rooms, LivingRoom) {
		fmt.Println("Turn on LivingRoom Light")
	}
}

func main() {
	var rooms Room = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	TurnOnLights(rooms)

	//===result===
	//Turn on MasterRoom Light
	//Turn on LivingRoom Light
}
