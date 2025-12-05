package fuzzer

import (
        "bufio"
        "fmt"
        "io"
        "os"
        "strings"
        "sync"
        "time"

        "github.com/tools-tensorsecurityacademy/dirrumble/internal/models"
        "github.com/tools-tensorsecurityacademy/dirrumble/internal/output"
        "github.com/tools-tensorsecurityacademy/dirrumble/internal/request"
)

func Run(opts *models.Options) {
        output.PrintBanner()

        file, err := os.Open(opts.Wordlist)
        if err != nil {
                fmt.Printf("Error opening wordlist: %v\n", err)
                os.Exit(1)
        }
        defer file.Close()

        words := make(chan string, opts.Threads)
        results := make(chan *models.Result, opts.Threads)
        var wg sync.WaitGroup

        // Producer
        go func() {
                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                        words <- scanner.Text()
                }
                close(words)
        }()

        // Workers
        for i := 0; i < opts.Threads; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        for word := range words {
                                start := time.Now()

                                resp, err := request.BuildAndSend(opts.TargetURL, word, opts)

                                status := 0
                                length := int64(0)
                                wordCount := 0
                                lineCount := 0
                                var body []byte

                                if err == nil && resp != nil {
                                        body, _ = io.ReadAll(resp.Body)
                                        resp.Body.Close()

                                        content := strings.TrimSpace(string(body))
                                        wordCount = len(strings.Fields(content))
                                        lineCount = len(strings.Split(content, "\n"))
                                        if lineCount > 0 && content != "" {
                                                lineCount++ // account for last line without \n
                                        }

                                        status = resp.StatusCode
                                        length = resp.ContentLength
                                        if length < 0 {
                                                length = int64(len(body))
                                        }
                                }

                                results <- &models.Result{
                                        URL:     opts.TargetURL + "/" + word,
                                        Status:  status,
                                        Length:  length,
                                        Words:   wordCount,
                                        Lines:   lineCount,
                                        Time:    time.Since(start),
                                        Payload: word,
                                        Error:   err,
                                }
                        }
                }()
        }

        // Collector
        go func() {
                wg.Wait()
                close(results)
        }()

        // Output
        for res := range results {
                output.PrintResult(res)
        }
}
