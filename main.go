/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-20
	* This file calculates the price of a car 
	*/
	
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// assign constants 
	const basePrice float64 = 28000
	const floorMat float64 = 800
	const navSystem float64 = 1000
	const heatedSeats float64 = 800
	const warrantyCost float64 = 2800

	// total cost
	total := basePrice

	// base price
	fmt.Println("Base Price: $", basePrice)

	// INPUT
	// promts for an input, then converts to lowercase
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Would you like floor mats? (yes/no): ")
	matsInput, _ := reader.ReadString('\n')
	mats := strings.ToLower(strings.TrimSpace(matsInput))

	fmt.Print("Would you like a navigation system? (yes/no): ")
	navInput, _ := reader.ReadString('\n')
	nav := strings.ToLower(strings.TrimSpace(navInput))

	fmt.Print("Would you like heated leather seats? (yes/no): ")
	seatsInput, _ := reader.ReadString('\n')
	seats := strings.ToLower(strings.TrimSpace(seatsInput))

	fmt.Print("Would you like a 5-year extended warranty? (yes/no): ")
	warrantyInput, _ := reader.ReadString('\n')
	warranty := strings.ToLower(strings.TrimSpace(warrantyInput))

	// PROCESS
	if mats == "yes" {
		fmt.Println("Floor mats: $", floorMat)
		total += floorMat
	}

	if nav == "yes" {
		fmt.Println("Navigation system: $", navSystem)
		total += navSystem
	}

	if seats == "yes" {
		fmt.Println("Heated leather seats: $", heatedSeats)
		total += heatedSeats
	}

	if warranty == "yes" {
		fmt.Println("5-Year extended warranty: $", warrantyCost)
		total += warrantyCost
	}

	// ax and final cost
	tax := total * 0.13
	finalCost := total + tax

	fmt.Println("13% Taxes: $", tax)
	fmt.Println("Final cost of car: $", finalCost)

	fmt.Println("\nDone.")
}