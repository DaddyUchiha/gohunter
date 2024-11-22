package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Purple     = "\033[35m"
	Cyan       = "\033[36m"
	Bold       = "\033[1m"
	Underlined = "\033[4m"
)

func printHelp() {
	fmt.Println(Red, "Usage: gohunter -u <https://example.com> -d <delay_time> -o <Output> -w <Wordlist Path>", Reset)
	fmt.Println("-u, --url      Target URL")
	fmt.Println("-d, --delay    Delay between requests (seconds)")
	fmt.Println("-o, --output   Save results to a file")
	fmt.Println("-h, --help     Show this help menu")
	fmt.Println("-w, --wordlist Path to the wordlist")
}

func validateArgs(requiredArgs int, errorMessage string) {
	if len(os.Args) < requiredArgs {
		fmt.Println(Red, errorMessage, Reset)
		os.Exit(1)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
        fmt.Println("===================================================================================")
        fmt.Println(Yellow,Bold,"                               DIRECTORY FINDER                                  ",Reset)
        fmt.Println("===================================================================================")
        fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
        fmt.Println("")

	if len(os.Args) < 2 {
		fmt.Println(Red, "Usage: gohunter -h", Reset)
		os.Exit(1)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	validateArgs(5, "Usage: gohunter -u <https://example.com> -w <Wordlist>")

	baseURL := ""
	wordlistPath := ""
	delayTime := 0
	outputFile := ""

	// Parse arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-u", "--url":
			validateArgs(i+1, "URL is required.")
			baseURL = os.Args[i+1]
			i++
		case "-w", "--wordlist":
			validateArgs(i+1, "Wordlist path is required.")
			wordlistPath = os.Args[i+1]
			i++
		case "-d", "--delay":
			validateArgs(i+1, "Delay time is required.")
			delayTime, _ = strconv.Atoi(os.Args[i+1])
			i++
		case "-o", "--output":
			validateArgs(i+1, "Output file name is required.")
			outputFile = os.Args[i+1]
			i++
		}
	}

	// Validate base URL
	if _, err := url.ParseRequestURI(baseURL); err != nil {
		fmt.Println(Red, "Invalid URL provided.", Reset)
		os.Exit(1)
	}

	// Validate wordlist file
	if _, err := os.Stat(wordlistPath); os.IsNotExist(err) {
		fmt.Println(Red, "Wordlist file not found.", Reset)
		os.Exit(1)
	}

	// Start scanning
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanDirectories(baseURL, wordlistPath, delayTime, outputFile)
	}()
	wg.Wait()
}

func scanDirectories(baseURL, wordlistPath string, delay int, outputFile string) {
	wordlist, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println(Red, "Error opening wordlist:", err, Reset)
		return
	}
	defer wordlist.Close()

	var fileWriter *bufio.Writer
	if outputFile != "" {
		output, err := os.OpenFile(outputFile+".txt", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(Red, "Error creating output file:", err, Reset)
			return
		}
		defer output.Close()
		fileWriter = bufio.NewWriter(output)
		defer fileWriter.Flush()
	}

	scanner := bufio.NewScanner(wordlist)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(Yellow, Bold, "Starting Directory Finder Scan", Reset)
	fmt.Println("")
	fmt.Println("Target:", baseURL)
	fmt.Println("--------------------------------------------------------------------------------")

	for scanner.Scan() {
		path := scanner.Text()
		fullURL := fmt.Sprintf("%s/%s", baseURL, path)

		resp, err := http.Head(fullURL)
		if err != nil {
			fmt.Printf("%sError requesting %s: %v%s\n", Red, fullURL, err, Reset)
			continue
		}
		defer resp.Body.Close()

		statusColor := getStatusColor(resp.StatusCode)
		fmt.Printf("%s%s | [%d]%s\n", statusColor, fullURL, resp.StatusCode, Reset)

		if fileWriter != nil {
			fileWriter.WriteString(fmt.Sprintf("%s | [%d]\n", fullURL, resp.StatusCode))
		}

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(Red, "Error reading wordlist:", err, Reset)
	}
	fmt.Println(Cyan, "Scan completed.", Reset)
}

func getStatusColor(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return Green
	case statusCode >= 300 && statusCode < 400:
		return Blue
	case statusCode >= 400 && statusCode < 500:
		return Purple
	case statusCode >= 500:
		return Red
	default:
		return Cyan
	}
}
