package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"
)

//go:embed template1.tpl
var template1 string

//go:embed template2.tpl
var template2 string

func init() {

	var err error

	flag.StringVar(&infile, "infile", "config.json", "Input file path")
	flag.Parse()

	userHome, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}

	imageDir = path.Join(userHome, "Pictures")
	os.MkdirAll(imageDir, 0755)

	workingDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

}

func main() {
	fmt.Println("Hello, edpush!")

	fmt.Println("User Home Directory:", userHome)

	fmt.Println("Image Directory:", imageDir)

	fmt.Println("Working Directory:", workingDir)

	/*
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
			SmallPath:     "~/images/",
			SmallImage:    "small.png",
			LargePath:     "~/images/",
			LargeImage:    "large.png",
			ModelsPath:    "/path/to/models",
		}
	*/
	cfile := path.Join(workingDir, infile)
	status, err := createBlankConfig(cfile)
	if err != nil {
		log.Fatalln("Error creating blank config:", err)
	}
	if status {
		fmt.Println("Config file is ready at", cfile)
		os.Exit(0)
	}

	tfile, err := os.ReadFile(cfile)
	if err != nil {
		log.Fatalln("Error reading config file:", err)
	}
	err = json.Unmarshal(tfile, &md)
	if err != nil {
		log.Fatalln("Error unmarshalling config file:", err)
	}

	// Generate the two python scripts from templates

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

	tmpl, err = template.New("passTwo").Parse(template2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "s2 Parse:", err)
		log.Fatalln(err)
	}

	large, err := os.OpenFile(t2, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "large OpenFile:", err)
		log.Fatalln(err)
	}

	defer large.Close()

	err = tmpl.Execute(large, md)
	if err != nil {
		fmt.Fprintln(os.Stderr, "tmpl.Execute(large, md):", err)
		log.Fatalln(err)
	}

	cmd := exec.Command("./installer_files/env/bin/python", "t1")
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Error executing first python script:", err)
	}

	cmd = exec.Command("./installer_files/env/bin/python", "t2")
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Error executing second python script:", err)
	}

	if md.RemoveSmall {
		err = os.Remove(path.Join(md.SmallPath, md.SmallImage))
		if err != nil {
			log.Fatalln("Error removing small image:", err)
		}
	}

}

// Create blank jason config file if it does not exist
func createBlankConfig(configPath string) (bool, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		blankConfig := mainData{}
		file, err := os.Create(configPath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return false, err
		}
		defer file.Close()

		j, err := json.MarshalIndent(blankConfig, " ", "  ")
		if err != nil {
			fmt.Println("Error marshalling blank config to JSON:", err)
			return false, err
		}

		os.WriteFile(configPath, j, 0644)
		fmt.Println("Created blank config file at", configPath)
		return true, nil
	}
	return true, nil
}
