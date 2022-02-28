package main

import (
	"GO_MTS/lesson7/mycrypt"
	"flag"
	"fmt"
	"log"
)

func main() {

	var fileSource, hashFile, outFile string

	flag.StringVar(&fileSource, "source-file", "", "File source")
	flag.StringVar(&hashFile, "hash-file", "", "File hash")
	flag.StringVar(&outFile, "out-file", "sign.txt", "FIle output")

	flag.Parse()
	action := flag.Args()[0]

	switch action{
	case "enc":
		encoder, err := mycrypt.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(encoder)

		err = encoder.EncryptSha256()


		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}
	case "dec":
		decoder, err := mycrypt.NewDecryptor(hashFile, outFile, fileSource)
		if err != nil{
			log.Fatal(err)
			return
		}
		fmt.Println(decoder)

		err = decoder.ValidateSign()
		if err != nil {
			panic(err)
		}

	default:
		log.Fatal("Use enc or dec param")
	}
}
