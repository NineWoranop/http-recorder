package main

import (
	"fmt"
	"http-recorder/cmd"
	"http-recorder/httpclient"
	"os"
)

func work(myargs cmd.CmdArguments) (err error) {
	var filepath string = myargs.NextFilePath()
	var data []byte
	data, err = httpclient.FetchData(myargs)
	if err == nil {
		os.WriteFile(filepath, data, 0644)
		fmt.Printf("%s have been written\n", filepath)
	} else {
		fmt.Printf("Failed to reached.%s\n", err)
	}
	return err
}

func main() {
	var err error
	var myargs *cmd.CmdArguments
	myargs, err = cmd.ReadArguments()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var args = *myargs
	fmt.Printf("arguments are %s\n", args.String())

	for {
		work(args)
		args.Sleep()
	}
}
