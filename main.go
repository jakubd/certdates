package main

import (
	"certdates/cert_handling"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	domainsTextFilePtr := flag.String("domains", "notset", "text file of domains that you want to check")
	thresholdPtr := flag.Int("threshold", 30, "threshold of warnings")
	flag.Parse()

	if *domainsTextFilePtr == "notset" {
		fmt.Printf("usage: %s --domains=[domains txt file] --t=[threshold int]", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {
		cert_handling.OutputCertificateValidityReport(*domainsTextFilePtr, *thresholdPtr)
	}
}
