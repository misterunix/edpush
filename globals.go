package main

var userHome string
var imageDir string
var workingDir string
var t1 string
var t2 string

type tmpl1 struct {
	RandomNumber int
	Prompt       string
	NPrompt      string
	Model        string
	Steps        int
	Width        int
	Height       int
	ScaleUp      bool
	RemoveSmall  bool
	Sampler      string
	Count        int
	SmallImage   string
	LargeImage   string
	ModelsPath   string
}
