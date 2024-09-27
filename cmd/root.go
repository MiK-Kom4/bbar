/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var domain string = "initial"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bbar",
	Short: "A CLI tool for automating reconnaissance tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if domain != "" {
			if err := createDomainDirectory(domain); err != nil {
				fmt.Println(err)
			}

			// subfinderを実行するためのコマンドを作成
			cmd := exec.Command("sh", "-c", fmt.Sprintf("subfinder -d %s -all -recursive | sort -u | tee subdomain.txt", domain))

			// 実行するディレクトリをドメイン名のディレクトリに設定
			cmd.Dir = filepath.Join(".", domain)
			fmt.Printf("%s", cmd.Dir)

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(output))
			return

		} else {
			fmt.Println("You must provide a domain via -d flag")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// -d フラグをrootコマンドに追加
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain name to create a directory for")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createDomainDirectory(domain string) error {
	// 現在のディレクトリを所得
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	folderPath := filepath.Join(currentDir, domain)

	// ディレクトリが存在しないことを確認して作成
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 権限はsubfinderに合わせている
		err := os.Mkdir(folderPath, 0744)
		if err != nil {
			return fmt.Errorf("failed to create directory %v", err)
		}
		fmt.Printf("Created folder %s\n", folderPath)
	} else {
		fmt.Println("Folder already exists")
	}

	return nil
}
