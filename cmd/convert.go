/*
Copyright © 2024 Achnologia <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Achno/gowall/internal/image"
	"github.com/Achno/gowall/utils"
	"github.com/spf13/cobra"
)

var (
	theme string
	batchFiles []string
)


var convertCmd = &cobra.Command{
	Use:   "convert [image path / batch flag]",
	Short: "convert an img's color shceme",
	Long: `convert an img's color shceme`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		switch {

		case len(batchFiles) > 0:
			fmt.Println("Processing batch files...")
			expandedFiles := utils.ExpandHomeDirectory(batchFiles)
			image.ProcessBatchImgs(expandedFiles,theme)

		case len(args) > 0:
			fmt.Println("Processing single image...")
			expandFile := utils.ExpandHomeDirectory(args)
			image.ProcessImg(expandFile[0], theme)
			
		default:
			fmt.Println("Error: requires at least 1 arg(s), only received 0")
			_ = cmd.Usage()
		}


	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&theme,"theme","t","catpuccin","Usage : --theme [ThemeName-Lowercase]")
	convertCmd.Flags().StringSliceVarP(&batchFiles,"batch","b",nil,"Usage: --batch [file1.png,file2.png ...]")

	// Here you will define your flags and configuration settings.

}
