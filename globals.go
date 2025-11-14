package main

var userHome string
var imageDir string
var workingDir string
var t1 string
var t2 string

type mainData struct {
	RandomNumber  int
	Prompt        string
	NPrompt       string
	ModelLocation string
	Model         string
	LoraLocation  string
	Lora          string
	Steps         int
	Width         int
	Height        int
	ScaleUp       bool
	RemoveSmall   bool
	Sampler       string
	Count         int
	SmallPath     string
	SmallImage    string
	LargePath     string
	LargeImage    string
	ModelsPath    string
}

var infile string

var md mainData
