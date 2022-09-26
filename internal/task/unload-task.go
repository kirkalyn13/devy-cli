package task

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	clearBuilder = "docker builder prune -f"
	clearVolume  = "docker volume prune -f"
)

// ========== Docker Cleanup Task ==========
func UnloadDocker() {
	clearDocker(clearBuilder, "builder")
	clearDocker(clearVolume, "volume")
	log.Print("Waiting for next schedule...")
}

// ========== Docker Clear Command Run ==========
func clearDocker(query string, subject string) {
	args := strings.Split(query, " ")

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Cleanup failed: %v", err)
	}

	log.Printf("%s\n", output)

	now := time.Now()
	dtCurrent := fmt.Sprintf("%v-%v-%v %v:%v:%v", now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute(), now.Second())
	log.Printf("Done %v cleanup as of: %v", subject, dtCurrent)
}
