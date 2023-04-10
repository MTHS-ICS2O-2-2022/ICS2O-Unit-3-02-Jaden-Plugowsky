// Created by: Jaden Plugowsky
// Created on: April 2023
//
// This program finds the volume of a right rectangular prism

package main

import "fmt"

func main() {
	// This function finds the volume of a right rectangular prism
	var baseLength float64
	var baseWidth float64
	var height float64
	var volume float64

	// Input
	fmt.Println("\nThis program finds the volume of a Right Rectangular Prism using given values.")
	fmt.Println()
	fmt.Print("Enter the Base's Length value (cm): ")
	fmt.Scanln(&baseLength)
	fmt.Print("Enter the Base's Width value (cm): ")
	fmt.Scanln(&baseWidth)
	fmt.Print("Enter the Height value (cm): ")
	fmt.Scanln(&height)

	// Process
	volume = (baseLength * baseWidth * height) / 3

	// Output
	fmt.Println()
	fmt.Println("The volume of the Right Rectangular Prism is:", volume, "cm³.")

	fmt.Println("\nDone.")
}