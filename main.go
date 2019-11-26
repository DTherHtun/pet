package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const dataFile string = "/var/data/hola.txt"

func pet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(dataFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	switch r.Method {
	case "GET":
		file, err := os.Stat(dataFile)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadFile(dataFile)
		if err != nil {
			log.Fatal(err)
		}
		filedata := string(data)
		if file.Size() == 0 {
			fmt.Fprintf(w, "Current Pod Name => %s \n", hostName)
			fmt.Fprintf(w, "No data posted yet.\n")
		} else {
			fmt.Fprintf(w, "Current Pod Name => %s \n", hostName)
			fmt.Fprintf(w, filedata)
		}

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyData := string(body)
		datawriter := bufio.NewWriter(file)
		datawriter.WriteString(bodyData + "\n\r")
		fmt.Println("New data has been received and stored.")
		datawriter.Flush()
		fmt.Fprintf(w, "Data stored on pod %s \n", hostName)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.\n")
	}
}

func main() {
	http.HandleFunc("/", pet)
	fmt.Printf("Starting pet server at 8080.......... \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
