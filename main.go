package main

import (
	"certdates/cert_handling"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	domainsTextFilePtr := flag.String("domains", "", "text file of domains that you want to check")
	thresholdPtr := flag.Int("threshold", 30, "threshold of warnings")
	flag.Parse()

	if flag.NArg() > 0 {
		domain := flag.Arg(0)
		proc, res := cert_handling.DoLookupForString(domain)
		if proc {
			cert_handling.PrintResult(res, *thresholdPtr)
		}
	} else {
		if *domainsTextFilePtr == "" {
			fmt.Printf("usage: %s --domains=[domains txt file] --t=[threshold int]", filepath.Base(os.Args[0]))
			os.Exit(1)
		} else {
			cert_handling.OutputCertificateValidityReport(*domainsTextFilePtr, *thresholdPtr)
		}
	}
	}


