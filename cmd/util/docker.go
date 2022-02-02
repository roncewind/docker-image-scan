package util

import (
	"bytes"

	"encoding/json"
	"fmt"
	"log"
	"os/exec"

)

func doesExecutibleExist(exe string) bool {

    path, err := exec.LookPath(exe)
    if err != nil {
        log.Printf("didn't find '%s' executable\n", exe)
		return false
    }
	log.Printf("'%s' executable is in '%s'\n", exe, path)
	return true
}

func parseJSON(data []byte) {
	var record map[string]json.RawMessage
	err := json.Unmarshal(data, &record)
	var vulns []json.RawMessage
	err2 := json.Unmarshal(record["vulnerabilities"], &vulns)

	if err == nil && err2 == nil {
		for key, value := range vulns {
			// fmt.Printf("[%s]: %s\n\n",key, value)
			// fmt.Printf("[%d] %s\n", key, reflect.TypeOf(value))
			var vuln map[string]json.RawMessage
			if nil == json.Unmarshal(value, &vuln) {
				var title string
				if nil == json.Unmarshal(vuln["title"], &title) {
					fmt.Printf("[%d] %s\n\n", key, title)
				}
			}
		}
		var count int
		if nil == json.Unmarshal(record["dependencyCount"], &count) {
			fmt.Printf("Dependencies: %d\n", count)
		}
		if nil == json.Unmarshal(record["uniqueCount"], &count) {
			fmt.Printf("Vulnerabilities: %d\n", count)
		}
	} else {
		log.Println("Err:", err)
		log.Println("Err2:", err2)
	}
}

func runDockerScan(image string) {
	if !doesExecutibleExist("docker") {
		log.Fatalf("Please install 'docker' or add it to the PATH.\n")
	}
	osCmd := exec.Command("docker", "scan", "--json", image)
	var stdout, stderr bytes.Buffer
	osCmd.Stdout = &stdout
	osCmd.Stderr = &stderr
	log.Printf("Scanning %s ...\n", image)
	err := osCmd.Run()
	if err != nil {
		//FIXME: for some reason Run() returns 'exit status 1' even though it runs fine.
		if err.Error() != "exit status 1" {
			log.Printf("Error running exe command: %s\n", err)
		}
		// log.Fatalf("'docker scan --json %s' failed.\nerr: %s\nAre you using DockerDesktop?", image, err)
	}
	// outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	// fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	log.Println("Parsing scan...")
	parseJSON(stdout.Bytes())
}

