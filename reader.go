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

            mode := info.Mode()
            if mode&os.ModeSocket != 0 {
                fmt.Printf("%s is a socket, skipping...\n", path)
                return nil
            }

            if _, err := os.Stat(path); os.IsNotExist(err) {
                fmt.Printf("file does not exist: %s, skipping...\n", path)
                return nil
            }


		body, err := os.ReadFile(path)
		if err != nil {
			log.Printf("unable to read file: %v", err)
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