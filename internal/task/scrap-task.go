// Automation Script to delete generated dev files on Temp Dir whih are not yet cleared every 10min.
package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// ========== Delete Files in Temp Dir ==========
func DeleteTempFiles() {
	tempDir := os.TempDir()

	log.Print("Deleting Files...")
	os.Chdir(tempDir)
	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// Verify file if temp tars
		_, err := strconv.Atoi(f.Name())
		if err != nil {
			continue
		} else {
			err := os.RemoveAll(f.Name())

			if err != nil {
				log.Fatal(err)
				continue
			}

			log.Print("Deleted File: ", f.Name())
		}
	}

	now := time.Now()
	dtCurrent := fmt.Sprintf("%v-%v-%v %v:%v:%v", now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute(), now.Second())
	log.Print("Deleted files as of: ", dtCurrent)
	log.Print("Waiting for next schedule...")
}
