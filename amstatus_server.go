package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
)

type AmstatusInfo struct {
	Host    string `json:"host"`
	Disk    string `json:"disk"`
	Dumping string `json:"dumping"`
	Dumped  string `json:"dumped"`
}

func amstatus() ([]byte, error) {
	cmd := exec.Command("/usr/sbin/amstatus", "rolling", "--dumpingtape")
	return cmd.Output()
}

func amandaServer(w http.ResponseWriter, req *http.Request) {
	output, err := amstatus()
	if err != nil {
		/*
		   TODO: Handle the return codes
		     0  = success
		     1  = error
		     4  = a dle failed
		     8  = Don't know the status of a dle (RESULT_MISSING in the report)
		     16 = tape error or no more tape
		*/
		//	w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		//	return
	}

	// Convert the raw amstatus output to an AmstatusInfo object.
	ui, err := parseAmstatusInfo(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Create the JSON representation of the system amstatus.
	data, err := json.MarshalIndent(ui, " ", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Write the HTTP response headers and body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}
