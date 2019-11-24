package cert_handling

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"io"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type LookupResult struct {
	Hostname string
	Port     int
	DaysLeft int
	Expiry   time.Time
	Err      error
}

const (
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

var Dialer = &net.Dialer{Timeout: 7 * time.Second}

// Given a hostname and port combo do a TLS connection
// and return a LookupResult struct with the results
// of the query
func DoLookup(givenHostname string, givenPort int) LookupResult {
	var res LookupResult

	res.Hostname = givenHostname
	res.Port = givenPort
	conn, err := tls.DialWithDialer(Dialer, "tcp", givenHostname + ":" + strconv.Itoa(givenPort), nil)

	if err != nil {
		res.Err = err
		return res
	}

	defer conn.Close()

	validErr := conn.VerifyHostname(givenHostname)

	if validErr != nil {
		res.Err = validErr
		return res
	}

	if len(conn.ConnectionState().PeerCertificates) > 0 {
		res.Expiry = conn.ConnectionState().PeerCertificates[0].NotAfter
	}

	res.DaysLeft = DaysLeft(res.Expiry)

	return res
}

// Given a filename, return a slice of strings of all the lines.
func ReadFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("can't open ")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')

		if err != nil && err != io.EOF {
			return lines, err
		}

		if err == nil || err == io.EOF {
			lines = append(lines, strings.Trim(line, "\n"))
		}

		if err == io.EOF {
			break
		}
	}

	return lines, nil
}


// Given a filename of domains, parse URLs and do lookups for
// everything - return a slice of results.
func DoLookupsForFile(givenFilename string) (results []LookupResult) {

	lines, _ := ReadFile(givenFilename)
	for _, line := range lines {
		if len(line) > 0 && !strings.HasPrefix(line,"#") {
			urlObject, _ := url.Parse(line)
			var port int

			if urlObject.Port() == "" {
				port = 443
			} else {
				port, _ = strconv.Atoi(urlObject.Port())
			}

			thisResult := DoLookup(urlObject.Hostname(), port)
			results = append(results, thisResult)
		}
	}

	return results
}

// Given a time.Time date, return the number of days difference
// between that and today.
func DaysLeft(givenDate time.Time) (daysLeft int) {
	now := time.Now()
	days := givenDate.Sub(now).Hours() / 24
	return int(days)
}

// given a filename (text file of all domains) and a threshold, output a certificate
// validity report to screen.
func OutputCertificateValidityReport(givenFilename string, warningThreshold int) {
	theSpinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	theSpinner.Start()
	finishedLookups := DoLookupsForFile(givenFilename)
	theSpinner.Stop()
	for _, thisResult := range finishedLookups {
		if thisResult.Err == nil {
			fmt.Print("https://" + thisResult.Hostname + ":" + strconv.Itoa(thisResult.Port) + " ⟶ " + thisResult.Expiry.Format("January 02, 2006") + "\n")
			if thisResult.DaysLeft < warningThreshold {
				fmt.Printf(WarningColor,"\tdays left: " + strconv.Itoa(thisResult.DaysLeft))
				fmt.Println()
			} else {
				fmt.Printf(NoticeColor,"\tdays left: " + strconv.Itoa(thisResult.DaysLeft))
				fmt.Println()
			}
		} else {
			fmt.Printf(ErrorColor, "https://" + thisResult.Hostname + ":" + strconv.Itoa(thisResult.Port) + "\n")
			fmt.Printf(ErrorColor, "\t" +  thisResult.Err.Error() + "\n")
		}
		fmt.Println()
	}
}
