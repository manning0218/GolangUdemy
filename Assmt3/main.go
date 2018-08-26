package main

import (
    "os"
    "log"
    "io"
);

func main () {
    fileName := os.Args[1] // Gets the first argument passed when executing 

    file := openFile(fileName)

    io.Copy(os.Stdout, file)

    file.Close()
}

func openFile(fileName string) *os.File {
    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    return file
}
