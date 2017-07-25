package main

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	logger := log.New(stdout, "", log.Lshortfile)

	certBytes, err := ioutil.ReadFile("./test.crt")
	if err != nil {
		logger.Print("ERROR:", err)
		return
	}
	logger.Print("INFO:", "Successfully read the cert!")
	logger.Printf("DEBUG: Certificate:\n%s", certBytes)

	certBlock, _ := pem.Decode(certBytes)
	if certBlock == nil {
		logger.Print("ERROR:", "No PEM data could be decoded!")
		return
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		logger.Print("ERROR:", err)
	}

	hostname := "os-us.osaas-lab.rcsops.com"
	err = cert.VerifyHostname(hostname)
	if err != nil {
		logger.Print(err)
	} else {
		logger.Print("INFO:", "SUCCESS!!! Cert was successfully read and decrypted")
	}
}
