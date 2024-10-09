package main

import (
    "fmt"
    "os"
    "path/filepath"
	"log"
)

func main() {
    err := filepath.Walk("/mnt/data", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
            return err
        }
        fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)

		if !info.IsDir() {
		body, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
			return err
		}
		fmt.Printf("read %d bytes from file\n",len(body))
	}

        return nil

    })
    if err != nil {
        fmt.Println(err)
    }
}