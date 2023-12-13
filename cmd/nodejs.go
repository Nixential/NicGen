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

var genNode = &cobra.Command{
	Use:   "node-gen",
	Short: "Generate nodejs project",
	Long: `Generate nodejs project
	Example usage:
	nixjs node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"`,
	Run: generateNodeJS,
}

func init() {
	rootCmd.AddCommand(genNode)
	genNode.Flags().StringP("directory", "D", "./myApp", "Outpur directory for the project")
	genNode.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNode.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNode.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
}

func generateNodeJS(cmd *cobra.Command, args []string) {
	var stderr bytes.Buffer
	//Create the directory
	dir, errDir := cmd.Flags().GetString("directory")
	if errDir != nil {
		log.Fatal("Fatal error:", errDir)
	}

	if dir != "" {
		cmd := exec.Command("mkdir", dir)

		cmd.Stderr = &stderr // Capture standard error

		if err := cmd.Run(); err != nil {
			fmt.Println("Error executing command:", err)
			fmt.Println("Standard error output:", stderr.String())
			return
		}

		fmt.Printf("Directory %v created successfully\n", dir)
	}

	//Check if npm is installed
	checkNPM := exec.Command("npm", "--version")

	npmVerOut, errNPM := checkNPM.Output()
	if errNPM != nil {
		fmt.Println("npm is not installed on your system or not fount in PATH")
		log.Fatal(errNPM)
	}
	fmt.Printf("npm version %v was detected on your system\n", string(npmVerOut))

	fmt.Printf("Changing to directory %v ...\n", dir)

	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	def, errDefault := cmd.Flags().GetBool("yes")

	if errDefault != nil {
		log.Fatal("Fatal error", errDefault)
	}
	if errDefault != nil {
		log.Fatal("Fatal error", errDefault)
	}

	cmdInit := exec.Command("npm", "init")

	cmdInit.Stdin = os.Stdin
	cmdInit.Stdout = os.Stdout
	cmdInit.Stderr = os.Stderr

	if def {
		cmdInit.Args = append(cmdInit.Args, "-y")
	}

	if err := cmdInit.Run(); err != nil {
		log.Fatal("Error executing command:", err)
	}
	fmt.Println("Node JS project initialized...")

	//Install libs
	librariesStr, err := cmd.Flags().GetString("libs")
	if err != nil {
		log.Fatal("Error retrieving libraries:", err)
	}

	// Split the libraries string into a slice of individual libraries
	libraries := strings.Fields(librariesStr)

	if len(libraries) > 0 {
		fmt.Println("Installing node modules:", strings.Join(libraries, ", "))

		// Create a slice with "npm" and "install" as the first two elements
		npmArgs := []string{"install"}
		// Append each library as a separate element in the slice
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

	//Install dev libs
	devLibsStr, err := cmd.Flags().GetString("dev-libs")
	if err != nil {
		log.Fatal("Error retrieving libraries:", err)
	}

	// Split the libraries string into a slice of individual libraries
	devLibraries := strings.Fields(devLibsStr)

	if len(devLibraries) > 0 {
		fmt.Println("Installing node modules:", strings.Join(devLibraries, ", "))

		// Create a slice with "npm" and "install" as the first two elements
		npmArgs := []string{"install"}
		// Append each library as a separate element in the slice
		npmArgs = append(npmArgs, "--save-dev")
		npmArgs = append(npmArgs, devLibraries...)

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
