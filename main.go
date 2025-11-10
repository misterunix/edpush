package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
)

//go:embed template1.tpl
var template1 string

//go:embed template2.tpl
var template2 string

func main() {
	fmt.Println("Hello, edpush!")

	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}
	fmt.Println("User Home Directory:", userHome)

	imageDir = path.Join(userHome, "Pictures")
	os.MkdirAll(imageDir, 0755)
	fmt.Println("Image Directory:", imageDir)

	workingDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	fmt.Println("Working Directory:", workingDir)

	t1 = path.Join(workingDir, "template1.py")
	t2 = path.Join(workingDir, "template2.py")

	os.Remove(t1)
	os.Remove(t2)

}
