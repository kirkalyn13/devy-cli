package task

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (
	gitEmail = "git config --global user.email"
)

// ========== Switch Git Email ==========
func SwitchEmail(user string) error {
	args := strings.Split(gitEmail, " ")
	args = append(args, fmt.Sprintf("%v", user))

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Unable to switch email: %v", err)
		return err
	}

	log.Printf("Switching... %s\n", output)

	log.Printf("Successfully switched to: %v", user)

	return nil
}
