package cmd

import (
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
	genNode.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	genNode.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNode.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNode.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
}

func generateNodeJS(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("directory")
	createDirectory(dir)
	checkNPMInstallation()

	changeDirectory(dir)

	initNodeProject(cmd)

	installLibraries(cmd, "libs")
	installLibraries(cmd, "dev-libs")
}
