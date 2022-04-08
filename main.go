package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bndr/gojenkins"
)

func main() {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "yuval", "0312")

	_, err := jenkins.Init(ctx)

	if err != nil {
		panic(err)
	}

	names, err := jenkins.GetAllJobNames(ctx)
	if err != nil {
		panic(err)
	}

	num, err := jenkins.BuildJob(ctx, names[0].Name, map[string]string{})

	if err != nil {
		panic(err)
	}

	build, err := jenkins.GetBuildFromQueueID(ctx, num)

	if err != nil {
		panic(err)
	}

	for build.IsRunning(ctx) {
		time.Sleep(time.Second * 10)
		fmt.Println("still running waiting another 10 seconds")
	}

	if strings.Contains(build.GetConsoleOutput(ctx), "SUCCESS") {

		fmt.Println("build done Successfuly")
	}

	// fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetConsoleOutput(ctx))
}
