package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"
	Bold   = "\033[1m"
	Underlined = "\033[4m"
)


func help() {
	fmt.Println(Red,"Usage: gohunter -u <https://example.com> -d <delay_time> -o <Output> -w <Wordlist Path>",Reset)
	fmt.Println("-u --url      Target URL")
	fmt.Println("-d --delay    Delay between Requests")
	fmt.Println("-o --Output   Save results to File")
	fmt.Println("-h --help     Show help Menu")
	fmt.Println("-w --wordlist Wordlist Path")
}

//==============================================================================================================================================================================================================

func main() {
	var wg sync.WaitGroup

	if len(os.Args) < 2 {
		fmt.Println(Red,"Usage: gohunter -h",Reset)
		fmt.Println(Red,"Usage: gohunter -u <https://example.com> -w <Wordlist>",Reset)
		os.Exit(0)
	}
	if len(os.Args) <= 3 {
		fmt.Println(Red,"Usage: gohunter -u <https://example.com> -w <Wordlist>",Reset)
		os.Exit(0)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		help()
		return
	}

	if len(os.Args) <= 5 {
		if os.Args[1] == "-u" || os.Args[1] == "--url" {
			if len(os.Args) <= 2 || os.Args[2] == "" {
				fmt.Println(Yellow,"URL Required OR Invalid URL",Reset)
				fmt.Println(Red,"Usage: gohunter -u <https://example.com>",Reset)
				os.Exit(0)
			}
			if os.Args[3] == "-w" || os.Args[3] == "--wordlist"{
				if len(os.Args) <= 4 || os.Args[4] == ""{
					fmt.Println(Red,"NO WORDLIST DETECTED",Reset)
					os.Exit(0)
				}
			}
				wg.Add(1)
				go func(){
					defer wg.Done()
					fmt.Println("")
					simple(&wg)
				}()
		}
	}else if len(os.Args) <=  7 {
		if os.Args[1] == "-u" || os.Args[1] == "--url" {
			if os.Args[3] == "-o" || os.Args[3] == "--output"{
				if len(os.Args) <= 5 || os.Args[4] == ""{
					fmt.Println(Red,"EMPTY OUTPUT FILENAME",Reset)
					fmt.Println(Red,"NO WORDLIST DETECTED",Reset)
					fmt.Println(Red,"Usage: gohunter -u <https://example.com> -o <Output> -w <Wordlist Path>",Reset)
					os.Exit(0)
				}
			if os.Args[5] == "-w" || os.Args[5] == "--wordlist"{
				if len(os.Args) <= 6 || os.Args[6] == "" {
					fmt.Println(Red,"INVALID COMMAND FORMAT",Reset)
					fmt.Println(Red,"Usage: gohunter -u <https://example.com> -o <Output> -w <Wordlist Path>",Reset)
					os.Exit(0)
				}
				wg.Add(1)
				go func(){
					defer wg.Done()
					fmt.Println("")
					simpleoutput(&wg)
					
				}()
			}
		}
	}
		if os.Args[3] == "-d" || os.Args[3] == "--delay"{
			if len(os.Args) <= 5 || os.Args[4] == ""{
				fmt.Println(Red,"INVALID COMMAND FORMAT",Reset)
				fmt.Println(Red,"Usage: gohunter -u <https://example.com> -d <delay_time> -w <Wordlist Path>",Reset)
				os.Exit(0)
			}
				if os.Args[5] == "-w" || os.Args[5] == "--wordlist"{
					if len(os.Args) <= 6 || os.Args[6] == "" {
						fmt.Println(Red,"NO WORDLIST DETECTED",Reset)
						fmt.Println(Red,"Usage: gohunter -u <https://example.com> -d <delay_time> -w <Wordlist Path>",Reset)
						os.Exit(0)
			}
				wg.Add(1)
				go func(){
					defer wg.Done()
					fmt.Println("")
					delay(&wg)
				}()
		}
	}
	}else if len(os.Args) <= 9 {
		if os.Args[1] == "-u" || os.Args[1] == "--url" {
			if os.Args[3] == "-d" || os.Args[3] == "--delay" {
				if os.Args[5] == "-o" || os.Args[5] == "--output" {
					if len(os.Args) <= 7 || os.Args[6] == ""{
						fmt.Println(Red,"EMPTY OUTPUT FILENAME",Reset)
						fmt.Println(Red,"NO WORDLIST DETECTED",Reset)
						os.Exit(0)
					}
				if os.Args[7] == "-w" || os.Args[7] == "--wordlist"{
					if len(os.Args) <= 8 || os.Args[8] == "" {
						fmt.Println(Red,"NO WORDLIST DETECTED",Reset)
						fmt.Println(Red,"Usage: gohunter -u <https://example.com> -d <delay_time> -o <Output> -w <Wordlist Path>",Reset)
						os.Exit(0)
					}
				wg.Add(1)
				go func(){
					defer wg.Done()
					fmt.Println("")
					output(&wg)
				}()
				}
			}
		}	
	}
	}else{
		fmt.Println(Red,"INVALID COMMAND FORMAT",Reset)
		fmt.Println("Usage: gohunter -u <https://example.com> -d <delay_time> -o <Output> -w <Wordlist Path>")
		os.Exit(1)
	}

	wg.Wait()
}


//==============================================================================================================================================================================================================

func simple(wg *sync.WaitGroup) {
	wordlist := os.Args[4]
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Yellow,Bold,"				DIRECTORY FINDER				  ",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("")
	fmt.Println(Blue,"NORMAL SCAN!",Reset)
	fmt.Println("------------------------------------------------------------------------------------")
	read := bufio.NewScanner(file)

	for read.Scan() {
		reade := read.Text()
		base := os.Args[2]

		if base == "" {
			fmt.Println("Invalid or Empty URL")
			break
		}

		baseURL := base + "/" + reade

		create, err := http.Get(baseURL)
		if err != nil {
			fmt.Println("Sending Request Error")
			continue
		}
		defer create.Body.Close()

		blue := "\x1b[34m"
		red := "\x1b[31m"
		green := "\x1b[32m"
		reset := "\x1b[0m"

		var statusMessage string

		if create.StatusCode == 200 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", green, create.Request.URL, create.StatusCode, reset)
		} else if create.StatusCode == 404 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", red, create.Request.URL, create.StatusCode, reset)
		} else {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", blue, create.Request.URL, create.StatusCode, reset)
		}
		fmt.Printf("%s",statusMessage)
	}

	fmt.Println("===========================================================")
	fmt.Println(Cyan,"    		Done Scanning Directories!",Reset)
	fmt.Println("===========================================================")
}

//==============================================================================================================================================================================================================

func delay(wg *sync.WaitGroup) {
	wordlist := os.Args[6]
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Yellow,Bold,"				DIRECTORY FINDER				  ",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("")
	fmt.Println(Blue,"DELAY NORMAL SCAN!",Reset)
	fmt.Println("------------------------------------------------------------------------------------")
	read := bufio.NewScanner(file)
	delayTime, _ := strconv.Atoi(os.Args[4])

	for read.Scan() {
		reade := read.Text()
		base := os.Args[2]

		if base == "" {
			fmt.Println("Invalid or Empty URL")
			break
		}

		baseURL := base + "/" + reade

		create, err := http.Get(baseURL)
		if err != nil {
			fmt.Println("Error while sending request")
			continue
		}
		defer create.Body.Close()

		blue := "\x1b[34m"
		red := "\x1b[31m"
		green := "\x1b[32m"
		reset := "\x1b[0m"

		var statusMessage string

		if create.StatusCode == 200 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", green, create.Request.URL, create.StatusCode, reset)
		} else if create.StatusCode == 404 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", red, create.Request.URL, create.StatusCode, reset)
		} else {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", blue, create.Request.URL, create.StatusCode, reset)
		}
		fmt.Printf("%s",statusMessage)

		time.Sleep(time.Duration(delayTime) * time.Second)
	}

	fmt.Println("===========================================================")
	fmt.Println(Cyan,"		Done Scanning Directories!",Reset)
	fmt.Println("===========================================================")
}

//==============================================================================================================================================================================================================


func simpleoutput(wg *sync.WaitGroup) {
	arg := os.Args[4] + ".txt"
	createFile, err := os.OpenFile(arg, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer createFile.Close()

	wordlist := os.Args[6]
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Yellow,Bold,"				DIRECTORY FINDER				  ",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("")
	fmt.Println(Blue,"NORMAL SCAN WITH OUTPUT OPTION!",Reset)
	fmt.Println("------------------------------------------------------------------------------------")
	
	read := bufio.NewScanner(file)

	for read.Scan() {
		reade := read.Text()
		base := os.Args[2]

		if base == "" {
			fmt.Println("Invalid or Empty URL")
			break
		}

		baseURL := base + "/" + reade

		create, err := http.Get(baseURL)
		if err != nil {
			fmt.Println("Error while sending request")
			continue
		}
		defer create.Body.Close()

		blue := "\x1b[34m"
		red := "\x1b[31m"
		green := "\x1b[32m"
		reset := "\x1b[0m"

		var statusMessage string
		var result string

		if create.StatusCode == 200 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", green, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		} else if create.StatusCode == 404 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", red, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		} else {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", blue, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		}

		fmt.Printf("%s",statusMessage)

		// Write to file
		writer := bufio.NewWriter(createFile)
		
		writer.WriteString(result)
		writer.Flush()
	}
	fmt.Println("===========================================================")
	fmt.Println(Cyan,"	  Done Scanning Directories!",Reset)
	fmt.Println(Cyan,"		RESULT SAVED",Reset)
	fmt.Println("===========================================================")

	if err := read.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

//==============================================================================================================================================================================================================


func output(wg *sync.WaitGroup) {
	arg := os.Args[6] + ".txt"
	createFile, err := os.OpenFile(arg, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer createFile.Close()

	wordlist := os.Args[8]
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Yellow,Bold,"				DIRECTORY FINDER				  ",Reset)
	fmt.Println("===================================================================================")
	fmt.Println(Red,"///////////////////////////////////////////////////////////////////////////////////",Reset)
	fmt.Println("")
	fmt.Println(Blue,"SCAN WITH DELAY & OUTPUT OPTION!",Reset)
	fmt.Println("------------------------------------------------------------------------------------")
	
	
	read := bufio.NewScanner(file)
	delayTime, _ := strconv.Atoi(os.Args[4])

	for read.Scan() {
		reade := read.Text()
		base := os.Args[2]

		if base == "" {
			fmt.Println("Invalid or Empty URL")
			break
		}

		baseURL := base + "/" + reade

		create, err := http.Get(baseURL)
		if err != nil {
			fmt.Println("Error while sending request")
			continue
		}
		defer create.Body.Close()

		blue := "\x1b[34m"
		red := "\x1b[31m"
		green := "\x1b[32m"
		reset := "\x1b[0m"

		var statusMessage string
		var result string

		if create.StatusCode == 200 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", green, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		} else if create.StatusCode == 404 {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", red, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		} else {
			statusMessage = fmt.Sprintf("%v%s | [%d]%v\n", blue, create.Request.URL, create.StatusCode, reset)
			result = fmt.Sprintln(create.Request.URL,"|", "[",create.StatusCode,"]")
		}

		fmt.Printf("%s",statusMessage)

		time.Sleep(time.Duration(delayTime) * time.Second)

		// Write to file
		writer := bufio.NewWriter(createFile)
		
		writer.WriteString(result)
		writer.Flush()
	}
	fmt.Println("===========================================================")
	fmt.Println(Cyan,"	  Done Scanning Directories!",Reset)
	fmt.Println(Cyan,"		RESULT SAVED",Reset)
	fmt.Println("===========================================================")

	if err := read.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}


