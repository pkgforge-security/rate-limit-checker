package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	threads    int
	method     string
	link       string
	statusCode int
	cont       bool
	count      int
	output     string
	wg         sync.WaitGroup
	logWriter  io.Writer
	cmd        = &cobra.Command{
		Use:   "rate-limit-checker",
		Short: "Check whether a domain has a rate limit enabled",
		Run:   runLoadTest,
	}
)

func main() {
	setupFlags(cmd)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setupFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method to use")
	cmd.Flags().IntVarP(&threads, "threads", "t", 10, "Number of threads to use")
	cmd.Flags().IntVarP(&count, "requests-count", "c", 1000, "Number of requests to send")
	cmd.Flags().BoolVarP(&cont, "ignore-code-change", "i", false, "Continue after the code changing")
	cmd.Flags().StringVarP(&link, "url", "u", "", "URL to send requests to")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file for logs")
}

func runLoadTest(cmd *cobra.Command, args []string) {
	if link == "" {
		log.Fatalf("URL was not provided")
		return
	}

	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			log.Fatalf("Failed to create log file: %s\n", err.Error())
		}
		defer file.Close()
		logWriter = file
	} else {
		logWriter = os.Stdout
	}

	// Use sendRequest for the initial request
	initialRequest, err := sendRequest(method, link, 0)
	if err != nil {
		fmt.Fprintf(logWriter, "URL is not available: %s\n", err.Error())
		return
	}
	statusCode = initialRequest.StatusCode
	fmt.Fprintf(logWriter, "Initial request status code: %d\n", statusCode)

	var channel = make(chan int, count)
	for threadCount := 0; threadCount < threads; threadCount++ {
		go func() {
			for requestCount := range channel {
				sendRequest(method, link, requestCount)
				wg.Done()
			}
		}()
	}

	for requestCount := 0; requestCount < count; requestCount++ {
		wg.Add(1)
		channel <- requestCount
	}

	close(channel)
	wg.Wait()
}

func sendRequest(method, link string, requestCount int) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest(method, link, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintln(logWriter, err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(logWriter, err.Error())
		return nil, err
	}

	fmt.Fprintf(logWriter, "Request %d: status code %d, body length %d\n", requestCount, response.StatusCode, len(body))
	if requestCount == 0 {
		statusCode = response.StatusCode
	}

	if statusCode != response.StatusCode && !cont {
		fmt.Fprintf(logWriter, "Status code mismatch: expected %d, got %d\n", statusCode, response.StatusCode)
		os.Exit(5)
	}

	return response, nil
}
