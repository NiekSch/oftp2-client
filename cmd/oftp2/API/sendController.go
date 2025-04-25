package API

import (
	"fmt"
	"net/http"
	"oftp2-client/internal/liboftp2/client"
	"os"
)

func SendFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	odetteIdSelf := query.Get("SSIDSelf")
	odetteIdServer := query.Get("SSIDServer")
	server := "cert-oftp.daf.com" // Replace with actual server address
	port := 6619                  // Replace with actual port number
	filePath := query.Get("filePath")
	datasetName := query.Get("datasetName")
	verbose := false // Set to true for verbose output

	sendFile(odetteIdSelf, odetteIdServer, filePath, datasetName, server, port, verbose)
}

func sendFile(odetteIdSelf, odetteIdServer, filePath, datasetName string, server string, port int, verbose bool) {

	s := client.OFTP2Client{
		ServerHost: server,
		ServerPort: port,
		Verbose:    verbose,
		OdetteId:   odetteIdSelf,
	}

	err := s.Connect()
	if err != nil {
		panic(err)
	}

	err = s.StartSession("", false, false, false)
	if err != nil {
		fmt.Printf("start session failed: %v\n", err)
		os.Exit(1)
	}

	err = s.SendFile(datasetName,
		filePath,
		client.FileFormatUnstructured,
		//"O2010CUSTOMER",
		odetteIdServer,
		client.SecurityLevelNone,
		false,
		false,
		false,
		false)
	if err != nil {
		fmt.Printf("send file failed: %v\n", err)
		os.Exit(1)
	}

	err = s.EndSession()
	if err != nil {
		fmt.Printf("end session failed: %v\n", err)
		os.Exit(1)
	}

	s.Close()
}
