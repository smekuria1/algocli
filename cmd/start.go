/*
Copyright © 2023 Solomon Mekuria Solmek18@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	multiinput "github.com/smekuria1/algocli/cmd/multiInput"
	"github.com/smekuria1/algocli/cmd/program"
	"github.com/smekuria1/algocli/cmd/steps"
	"github.com/spf13/cobra"
)

const logo = `
     ___       __        _______   ______     ______  __       __  
    /   \     |  |      /  _____| /  __  \   /      ||  |     |  | 
   /  ^  \    |  |     |  |  __  |  |  |  | |  ,----'|  |     |  | 
  /  /_\  \   |  |     |  | |_ | |  |  |  | |  |     |  |     |  | 
 /  _____  \  |   ----.|  |__| | |   --   | |  ---- || ------||  |
/__/     \__\ |_______| \______|  \______/   \______||_______||__| 
`

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the program and displays list of available DS and Algos",
	Long:  `Start initializes a bubble tea program that displays a list of available data structures and algorithms.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program
		options := steps.Options{
			AlgorithmType: "Binary Search",
		}
		//flagAlgorithm := cmd.Flag("algorithm").Value.String()
		// if flagAlgorithm != "" {
		// 	options = steps.Options{
		// 		AlgorithmType: flagAlgorithm,
		// 	}
		// }
		algorithm := &program.Algorithm{
			Name: options.AlgorithmType,
		}

		steps := steps.InitSteps(&options)
		fmt.Println(logoStyle.Render(logo))
		for _, step := range steps.Steps {
			s := &multiinput.Selection{}
			tprogram = tea.NewProgram(multiinput.InitialModel(step.Options, s, step.Headers, algorithm))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(err)
			}
			algorithm.ExitCli(tprogram)
			*step.Field = s.Choice
			algorithm.Name = s.Choice

		}
		fmt.Println(endingMsgStyle.Render("You have selected: " + options.AlgorithmType))
		fmt.Println(endingMsgStyle.Render("Thanks for using algocli!"))
		algorithm.Run()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.Flags().StringP("algorithm", "a", "", "Select an algorithm")

}
