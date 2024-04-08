package main

import (
	"fmt"
	"sync"
)

// LS simulates the Login Server
func LoginServerRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Login Server started...")
	LoginServer()
	fmt.Println("Login Server done.")
}

// PS simulates the Patch Server
func PatchServerRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Patch Server started...")
	
	fmt.Println("Patch Server done.")
}

// GS simulates the Game Server
func GameServerRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Game Server started...")
	
	fmt.Println("Game Server done.")
}

