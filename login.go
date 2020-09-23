package events_api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type loginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func loginUser(username string , password string) (bool , error){return true , nil }

func loginHandler(w http.ResponseWriter , r *http.Request) {

	if r.Method !=http.MethodPost {
		w.WriteHeader(405)
		return
	}

	//Read request

	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body erro %v" ,err)
		w.WriteHeader(500)
		return
	}

	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	ok, err := loginUser(schema.Username, schema.Password)
	if err != nil {
		log.Printf("Login user DB error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}


	if !ok {
		log.Printf("Unauthorized access for user: %v", schema.Username)
		w.WriteHeader(401) // Wrong password or username, Return 401.
		return
	}
	w.WriteHeader(200) // Successfully logged in.
}