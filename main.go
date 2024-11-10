package main

import (
	"fmt"
	"os"
	"log"
	"github.com/abhishek-kamat-nutanix/Diskreader/reader"
)

func checktype(path string) {
	info, err := os.Stat(path)
	if err != nil {
        if os.IsNotExist(err) {
            log.Printf("%s does not exist.", path)
        } else {
            log.Fatalf("Error checking path: %v", err)
        }
        return
    }

	switch mode := info.Mode(); {
    case mode.IsDir():
        fmt.Printf("%s is a directory\n", path)
        reader.HandleFile(path)
    case mode.IsRegular():
        fmt.Printf("%s is a regular file\n", path)
    case mode&os.ModeDevice != 0:
        fmt.Printf("%s is a device (block/char)\n", path)
        /*err:= reader.HandleBlock("/dev/sda", os.Stdout)
        if err != nil {
            fmt.Println("Error:", err)
        } */
        reader.HandleBlockSys(path,os.Stdout)
    default:
        fmt.Printf("%s is of an unknown type\n", path)
    }
}
func main(){
	checktype("/dev/loop1")
}