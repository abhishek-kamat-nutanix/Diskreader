package reader

import (
    "encoding/hex"
    "fmt"
    "os"
    "path/filepath"
	"log"
    "io"
    "syscall"
)

func HandleFile(pathway string) {
    err := filepath.Walk(pathway, func(path string, info os.FileInfo, err error) error {
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

func HandleBlock(devicePath string, writer io.Writer) error {
    device, err := os.Open(devicePath)
    if err != nil {
        return fmt.Errorf("unable to open block device %s: %w", devicePath, err)
    }
    defer device.Close()

    // Buffer to hold data read from the block device
    buffer := make([]byte, 4096) // Adjust buffer size as needed

    for i:=0;i<1;i++ {
        n, err := device.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break // End of file reached
            }
            return fmt.Errorf("error reading from block device: %w", err)
        }
        if n > 0 {
            // Write the read data to the provided writer (stdout or another writer)
            hexData := hex.EncodeToString(buffer[:n])
            if _, err := writer.Write([]byte(hexData)); err != nil {
                return fmt.Errorf("error writing to writer: %w", err)
            }
        }
    }
    return nil
}

func HandleBlockSys(devicePath string,writer io.Writer ) {
    disk := devicePath
    var fd, numread int
    var err error

    fd, err = syscall.Open(disk, syscall.O_RDONLY, 0777)

    if err != nil {
        fmt.Print(err.Error(), "\n")
        return
    }

    buffer := make([]byte, 4096)

    numread, err = syscall.Read(fd, buffer)

    if numread > 0 {
        // Write the read data to the provided writer (stdout or another writer)
        hexData := hex.EncodeToString(buffer[:numread])
        if _, err := writer.Write([]byte(hexData)); err != nil {
            fmt.Print("error writing to writer: %w", err)
        }
    }

    if err != nil {
        fmt.Print(err.Error(), "\n")
    }

    fmt.Printf("Numbytes read: %d\n", numread)
    // fmt.Printf("Buffer: %b\n", buffer)

    err = syscall.Close(fd)

    if err != nil {
        fmt.Print(err.Error(), "\n")
    }
}

