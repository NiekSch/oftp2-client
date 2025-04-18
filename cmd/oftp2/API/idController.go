package api

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/NiekSch/oftp2-client/internal/liboftp2/client"
)



type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if params["id"] == string(rune(user.ID)) {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}


func determineId() {

	r := client.OFTP2Client{
		ServerHost: activeOptions.Server,
		ServerPort: activeOptions.Port,
		OdetteId:   activeOptions.OdetteId,
		Verbose:    activeOptions.Verbose,
	}

	ssid, err := r.QueryServerCapabilities()
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	r.Close()

	fmt.Printf("Server's id is: '%s'\n", ssid.Id)
}
