package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func createDirectory(dir string) {
	if dir == "" {
		return
	}

	var stderr bytes.Buffer
	cmd := exec.Command("mkdir", dir)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Standard error output:", stderr.String())
		return
	}

	fmt.Printf("Directory %v created successfully\n", dir)
}

func checkNPMInstallation() {
	checkNPM := exec.Command("npm", "--version")
	npmVerOut, err := checkNPM.Output()
	if err != nil {
		fmt.Println("npm is not installed on your system or not found in PATH")
		log.Fatal(err)
	}
	fmt.Printf("npm version %v was detected on your system\n", string(npmVerOut))
}

func changeDirectory(dir string) {
	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}
}

func initNodeProject(cmd *cobra.Command) {
	def, _ := cmd.Flags().GetBool("yes")
	cmdInit := exec.Command("npm", "init")
	if def {
		cmdInit.Args = append(cmdInit.Args, "-y")
	}

	cmdInit.Stdin = os.Stdin
	cmdInit.Stdout = os.Stdout
	cmdInit.Stderr = os.Stderr

	if err := cmdInit.Run(); err != nil {
		log.Fatal("Error executing command:", err)
	}

	fmt.Println("Node JS project initialized...")
}

func installLibraries(cmd *cobra.Command, flag string) {
	librariesStr, _ := cmd.Flags().GetString(flag)
	libraries := strings.Fields(librariesStr)
	if len(libraries) > 0 {
		fmt.Println("Installing node modules:", strings.Join(libraries, ", "))

		npmArgs := []string{"install"}
		if flag == "dev-libs" {
			npmArgs = append(npmArgs, "--save-dev")
		}
		npmArgs = append(npmArgs, libraries...)

		installCmd := exec.Command("npm", npmArgs...)
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			log.Fatalf("Error installing node modules: %v", err)
		}
		fmt.Println("Node modules installed successfully.")
	}
}
