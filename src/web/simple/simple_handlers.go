package simple

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

// Stuff struct describe stuff
type Stuff struct {
	ID    int
	name  string
	price int
}

var stuffs []Stuff

func initStuffList() {
	stuffs = append(stuffs, Stuff{1, "Cup", 5})
}

func (stuff Stuff) toString() string {
	toReturn := "Stuff\n\n"
	toReturn += "ID: " + strconv.Itoa(stuff.ID) + "\n"
	toReturn += "Name: " + stuff.name + "\n"
	toReturn += "Price: " + strconv.Itoa(stuff.price) + "\n"
	return toReturn
}

// returns stuff, is stuff found and provided message
func getStaffWithSpecifiedID(strID string) (Stuff, bool, string) {
	ID, err := strconv.Atoi(strID)
	if err != nil {
		return Stuff{}, false, "Bad ID specified"
	}
	for index, value := range stuffs {
		if value.ID == ID {
			return stuffs[index], true, "Success"
		}
	}
	return Stuff{}, false, "We dont have stuff with such ID"
}

func getStaffWithSpecifiedName(name string) (Stuff, bool, string) {
	for index, value := range stuffs {
		if value.name == name {
			return stuffs[index], true, "Success"
		}
	}
	return Stuff{}, false, "We dont have stuff with such name"
}

func stuffEndPoint(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		stuffID := req.URL.Query().Get("id")
		if stuffID == "" {
			fmt.Fprint(res, "Stuff ID not specified!")
			res.WriteHeader(400)
			return
		}
		stuff, ok, message := getStaffWithSpecifiedID(stuffID)
		if !ok {
			fmt.Fprint(res, message)
		} else {
			fmt.Fprint(res, stuff.toString())
		}
	} else if req.Method == "POST" {

	} else if req.Method == "PUT" {

	} else if req.Method == "DELETE" {
		stuffID := req.URL.Query().Get("id")
		if stuffID == "" {
			fmt.Fprint(res, "Stuff ID not specified!")
			res.WriteHeader(400)
			return
		}
		message, ok := removeStuffByID(stuffID)
		if !ok {
			res.WriteHeader(400)
			fmt.Fprint(res, message)
			return
		}
		res.WriteHeader(200)
		fmt.Fprint(res, message, " , Stuff with ID: ", stuffID, " has been removed")
	}
}

func removeStuffByID(id string) (string, bool) {
	ID, error := strconv.Atoi(id)
	if error != nil {
		return "Bad ID specified!", false
	}
	for index, value := range stuffs {
		if value.ID == ID {
			removeStuffByPosition(index)

			return "Success", true
		}
	}
	return "We dont have stuff with such ID", false
}

func removeStuffByPosition(index int) {
	stuffs = append(stuffs[:index], stuffs[index+1:]...)
}

// StartSimpleServer start
func StartSimpleServer(port int) {
	initStuffList()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	Port := ":" + strconv.Itoa(port)
	server := &http.Server{Addr: Port}
	go listenForShutDown(server, ch)
	http.HandleFunc("/stuff", stuffEndPoint)
	server.ListenAndServe()
}

func listenForShutDown(server *http.Server, ch <-chan os.Signal) {
	<-ch
	fmt.Println()
	fmt.Printf("Go routine: server that working on port %s has succesfully stoped!", server.Addr)
	fmt.Println()
	server.Shutdown(nil)
}
