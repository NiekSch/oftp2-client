package API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"oftp2-client/internal/liboftp2/client"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == strconv.Itoa(users[0].ID) {
		json.NewEncoder(w).Encode(users[0])
	}
}

func DetermineId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	server := "cert-oftp.daf.com" // Replace with actual server address
	port := 6619                  // Replace with actual port number
	odetteId := params["SSID"]
	verbose := false // Set to true for verbose output

	fmt.Printf("Server: %s\n", server)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Odette ID: %s\n", odetteId)
	fmt.Printf("Verbose: %t\n", verbose)

	var id string = determineId(server, port, odetteId, verbose)

	w.Header().Set("Content-Type", "plaintext")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Server's id is: '%s'\n", id)))
	fmt.Printf("Server's id is: '%s'\n", id)
}

func determineId(server string, port int, odetteId string, verbose bool) string {

	r := client.OFTP2Client{
		ServerHost: server,
		ServerPort: port,
		OdetteId:   odetteId,
		Verbose:    verbose,
	}

	ssid, err := r.QueryServerCapabilities()
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	r.Close()

	fmt.Printf("Server's id is: '%s'\n", ssid.Id)

	return ssid.Id
}
