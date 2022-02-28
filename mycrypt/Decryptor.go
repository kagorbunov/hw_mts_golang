package mycrypt

import (
	"GO_MTS/lesson7/signature"
	"io/ioutil"
	"os"
)

type Decryptor struct {
	fileHash string
	hashString string
	fileString string
	fileSource string
}

func NewDecryptor(fileHash string, fileString string, fileSource string) (dec *Decryptor, err error)  {
	hashString, err := ioutil.ReadFile(fileHash)
	if err != nil {
		return
	}

	dec = &Decryptor{fileHash: fileHash, hashString: string(hashString), fileString: fileString, fileSource: fileSource}
	return
}

//add validate()

//gen signature
//build signature from signFile.xt

func (dec Decryptor) ValidateSign() (err error){
	file, err := os.Open(dec.fileString)
	if err!=nil {
		return err
	}
	defer file.Close()
	// add new method in signature and build sign from signFile.txt
	buildSign, err := signature.NewFileSourceFromSignatureSha256(file)
	if err != nil {
		return
	}
	//gen signature
	genFile, err := os.Open(dec.fileSource)
	if err!=nil {
		return err
	}
	defer file.Close()
	genSign, err := signature.NewSignatureSha256FromFileSource(genFile, dec.hashString)
	if err != nil {
		return
	}
	// compare signature
	buildSign.Equals(genSign)
	return
}

//signature.Equals(signature)