package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

//go:embed template1.tpl
var template1 string

//go:embed template2.tpl
var template2 string

func main() {
	fmt.Println("Hello, edpush!")

	md := mainData{
		RandomNumber:  42,
		Prompt:        "A beautiful landscape",
		NPrompt:       "low quality, blurry",
		ModelLocation: "models/stable-diffusion",
		Model:         "default-model",
		LoraLocation:  "models/lora",
		Steps:         50,
		Width:         512,
		Height:        512,
		ScaleUp:       true,
		RemoveSmall:   false,
		Sampler:       "dpmpp_2s_a",
		Count:         1,
		SmallImage:    "small.png",
		LargeImage:    "large.png",
		ModelsPath:    "/path/to/models",
	}

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

	tmpl, err := template.New("passOne").Parse(template1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "s1 Parse:", err)
		log.Fatalln(err)
	}

	small, err := os.OpenFile(t1, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "small OpenFile:", err)
		log.Fatalln(err)
	}

	defer small.Close()

	err = tmpl.Execute(small, md)
	if err != nil {
		fmt.Fprintln(os.Stderr, "tmpl.Execute(small, md):", err)
		log.Fatalln(err)
	}

}
