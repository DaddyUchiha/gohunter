GoHunter
-----------------------------------------------------------------------------------------------------------------------------

The gohunter is a simple command-line application written in Go that scans a target URL for various directories using a wordlist. This tool is useful for penetration testers and security professionals who want to discover hidden directories or endpoints on web servers. The tool supports various options, including:

1. Target URL: Specify the URL to scan.

2. Delay between Requests: Introduce a delay between requests to avoid overwhelming the server.
3. Output File: Save the results of the scan to a specified output file.
4. Wordlist Path: Provide a custom wordlist for the scan.

# Features

- Normal Scan: Performs a standard scan using the provided wordlist.
- Scan with Delay: Introduces a delay between requests to reduce server load.
- Output to File: Saves the results of the scan to a specified output file.
- Colorful Output: Provides a user-friendly interface with colored output for better readability.



## Installation

### Install using Go 

```bash
  go install https://github.com/DaddyUchiha/gohunter/cmd@latest
```

#### [ OR ]

###  Install using Git Clone 

```bash
git clone https://github.com/DaddyUchiha/gohunter
cd gohunter
```
#### Build the project

```bash
go build -o gohunter main.go
```

#### Run the tool:

```bash
gohunter -u <https://example.com> -d <delay_time> -o <Output> -w <Wordlist Path>  
```


## Usage/Examples

- #### *help*
```bash
gohunter -h
```
```bash
-u --url      Target URL
-d --delay    Delay between Requests
-o --Output   Save results to File
-h --help     Show help Menu
-w --wordlist Wordlist Path
```

- #### *Basic Scan*
```bash
gohunter -u https://example.com
```

- #### *Scan with Delay*
```bash
gohunter -u https://example.com -d <delay_time_in_seconds> -w <path_to_wordlist>
```
- #### *Save Output to File*
```bash
gohunter -u https://example.com -o <output_filename> -w <path_to_wordlist>
```
- #### *Complete Example*
```bash
gohunter -u https://example.com -d 2 -o result -w ./wordlist.txt
```

