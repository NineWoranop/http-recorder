package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"
)

var usage = fmt.Sprintf(`Usage: %s -url= [-scrape-interval=10s]|[-path=]|[-total-dat-file=1]|[-cert=CERT_PATH -key=KEY_PATH | -accept-invalid-cert]]`, os.Args[0])
var counter uint = 0
var maxTotalFiles uint = 999999

type CmdArguments struct {
	Url                 string
	Certificate         string
	Key                 string
	SkipServerCertCheck bool
	ScrapeInternal      time.Duration
	Path                string
	TotalFiles          uint
}

func ReadArguments() (args *CmdArguments, err error) {
	certPtr := flag.String("cert", "", "client certificate file")
	keyPtr := flag.String("key", "", "client certificate's key file")
	skipServerCertCheckPtr := flag.Bool("accept-invalid-cert", false, "Accept any certificate during TLS handshake. Insecure, use only for testing.")
	scrapeInternalPtr := flag.Duration("scrape-interval", time.Second*10, "Scrape interval to fetch metrics and write dat file")
	pathPtr := flag.String("path", "./", "Path for write dat file")
	totalFilesPtr := flag.Uint("total-dat-file", 1, "Number of dat files to write")
	urlPtr := flag.String("url", "http://localhost:9090/metrics", "Url endpoint to fetch metrics")

	flag.Parse()

	if urlPtr == nil {
		return nil, fmt.Errorf("need a url.\n%s", usage)
	} else if url, urlErr := url.Parse(*urlPtr); urlErr != nil || url.Scheme == "" {
		// Open file since arg appears not to be a valid URL (parsing error occurred or the scheme is missing).
		if _, err = os.Open(*urlPtr); err != nil {
			return nil, fmt.Errorf("error opening file %s", err)
		}
	} else {
		// Validate Client SSL arguments since arg appears to be a valid URL.
		if (*certPtr != "" && *keyPtr == "") || (*certPtr == "" && *keyPtr != "") {
			return nil, fmt.Errorf("%s\n with TLS client authentication: %s --cert /path/to/certificate --key /path/to/key METRICS_URL", usage, os.Args[0])
		}
	}
	if *totalFilesPtr > maxTotalFiles {
		fmt.Printf("total-dat-file is more than 999999. It will reduce to be 999999")
	}

	return &CmdArguments{*urlPtr, *certPtr, *keyPtr, *skipServerCertCheckPtr, *scrapeInternalPtr, *pathPtr, *totalFilesPtr}, nil
}

func (args CmdArguments) String() string {
	jsonText, _ := json.Marshal(args)
	return string(jsonText)
}

func (args CmdArguments) NextFilePath() string {
	num := counter + 1
	var filename string = fmt.Sprintf("%06d.dat", num)
	var filepath string = args.Path + filename
	args.increaseCounter()
	return filepath
}

func (args CmdArguments) increaseCounter() {
	counter += 1
	counter %= args.TotalFiles
}

func (args CmdArguments) Sleep() bool {
	fmt.Printf("sleep for %s\n", args.ScrapeInternal)
	time.Sleep(args.ScrapeInternal)
	return true
}
