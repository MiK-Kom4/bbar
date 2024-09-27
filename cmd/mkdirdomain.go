/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var domain string = "initial"

// mkdirdomainCmd represents the mkdirdomain command
var mkdirdomainCmd = &cobra.Command{
	Use:   "mkdirdomain",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// 現在のディレクトリを所得
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(domain, currentDir)
		folderPath := filepath.Join(currentDir, domain)
		fmt.Println(folderPath)

		// ディレクトリが存在しないことを確認して作成
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			// 権限はsubfinderに合わせている
			err := os.Mkdir(folderPath, 0744)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Created folder %s\n", folderPath)
		} else {
			fmt.Println("Folder already exists")
		}
	},
}

func init() {
	rootCmd.AddCommand(mkdirdomainCmd)

	// Here you will define your flags and configuration settings.
	// -d フラグでドメイン名を受け取るオプションを追加
	mkdirdomainCmd.Flags().StringVarP(&domain, "domain", "d", "", "create a directory")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkdirdomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mkdirdomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
