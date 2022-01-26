package httpclient

import (
	"crypto/tls"
	"fmt"
	"http-recorder/cmd"
	"io/ioutil"
	"net/http"
	"time"
)

func makeTransport(args cmd.CmdArguments) (*http.Transport, error) {
	// Start with the DefaultTransport for sane defaults.
	transport := http.DefaultTransport.(*http.Transport).Clone()
	// Conservatively disable HTTP keep-alives as this program will only
	// ever need a single HTTP request.
	transport.DisableKeepAlives = true
	// Timeout early if the server doesn't even return the headers.
	transport.ResponseHeaderTimeout = time.Minute
	tlsConfig := &tls.Config{InsecureSkipVerify: args.SkipServerCertCheck}
	if args.Certificate != "" && args.Key != "" {
		cert, err := tls.LoadX509KeyPair(args.Certificate, args.Key)
		if err != nil {
			return nil, err
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}
	transport.TLSClientConfig = tlsConfig
	return transport, nil
}

func FetchData(args cmd.CmdArguments) ([]byte, error) {
	transport, err := makeTransport(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", args.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GET request for URL %q failed: %v", args.Url, err)
	}
	client := http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing GET request for URL %q failed: %v", args.Url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("GET request for URL %q returned HTTP status %s", args.Url, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
