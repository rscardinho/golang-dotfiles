package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rscardinho/golang-dotfiles/internal/config"
	"github.com/rscardinho/golang-dotfiles/internal/logger"
	"github.com/rscardinho/golang-dotfiles/internal/tasks"
)

var (
	installConfigFilename = "install.toml"
)

func main() {
	if err := requestSudo(); err != nil {
		log.Fatal(err)
	}
	keepSudoAlive()

	logFile := logger.Load()

	parsedFile, err := config.Load(installConfigFilename)
	if err != nil {
		log.Fatal(err)
	}

	tasks.ExecuteAll("🔧  Installing packages", parsedFile.Packages, logFile)

	fmt.Printf("\n✅  Done. Logs saved to %s.\n", logger.Filename(*logFile))

	confirmAndReboot()
}

func requestSudo() error {
	cmd := exec.Command("sudo", "-v")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func keepSudoAlive() {
	go func() {
		for {
			time.Sleep(60 * time.Second)
			_ = exec.Command("sudo", "-v").Run()
		}
	}()
}

func confirmAndReboot() {
	fmt.Print("\nDo you want to reboot now? [y/N]: ")
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer == "y" || answer == "yes" {
		fmt.Println("\nRebooting...")

		cmd := exec.Command("sudo", "reboot")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to reboot: %s", err)
		}
	} else {
		fmt.Println("Reboot skipped.")
	}
}
