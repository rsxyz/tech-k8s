package main

import (
	"fmt"
	"os"
	"os/exec"
)

type ChartInfo struct {
	ReleaseName string
	ChartPath   string
	ValuesFile  string
	Message     string
}

func directoryExists(dirName string) bool {
	info, err := os.Stat(dirName)
	return !os.IsNotExist(err) && info.IsDir()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func runHelm(args []string) error {
	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func helmInstall(chart ChartInfo) {
	fmt.Println("Helm Chart Install START: " + chart.Message)
	if !directoryExists(chart.ChartPath) {
		fmt.Println(chart.ChartPath, "does not exists.")
		os.Exit(1)
	}
	if !fileExists(chart.ValuesFile) {
		fmt.Println(chart.ValuesFile, "does not exists.")
		os.Exit(1)
	}
	args := []string{"install", chart.ReleaseName, chart.ChartPath, "-f", chart.ValuesFile}
	fmt.Println(args)
	if err := runHelm(args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Helm Chart Install END: " + chart.Message)
}

func main() {
	chart := ChartInfo{
		ReleaseName: "v1",
		ChartPath:   "charts/ingress-nginx",
		ValuesFile:  "overrides/ingress-nginx.values.yaml",
		Message:     "Ingress-Nginx",
	}

	helmInstall(chart)
}
