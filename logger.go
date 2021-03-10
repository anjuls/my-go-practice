package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	Logger  *log.Logger
}

func(j *Job) Print(p string){
	j.Logger.Print(p)
}

func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	// same as
	// job := &Job{Command: "demo",
	//            Logger: log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Logger.Print("test")
	job.Print("starting now")
}