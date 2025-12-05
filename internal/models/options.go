package models

import (
        "flag"
        "fmt"
        "os" // ← ADD THIS
)

type Options struct {
        TargetURL        string
        Wordlist         string
        Method           string
        Headers          []string
        Threads          int
        Opaque           bool
        RawMethod        bool
        NoContentLen     bool
        RequestFile      string
        RequestKeepAlive bool
        Debug            bool
}

type multiStringFlag []string

func (f *multiStringFlag) String() string { return "" }
func (f *multiStringFlag) Set(val string) error {
        *f = append(*f, val)
        return nil
}

func ParseOptions() *Options {
        opts := &Options{}

        flag.StringVar(&opts.TargetURL, "u", "", "Target URL (required)")
        flag.StringVar(&opts.Wordlist, "w", "", "Wordlist file (required)")
        flag.StringVar(&opts.Method, "X", "GET", "HTTP method or raw request template")
        flag.StringVar(&opts.RequestFile, "request", "", "Request file to send verbatim")
        flag.BoolVar(&opts.RawMethod, "raw-method", false, "Treat -X as raw request template")
        flag.BoolVar(&opts.Opaque, "opaque", false, "Allow absolute URI fuzzing")
        flag.BoolVar(&opts.NoContentLen, "no-content-length", false, "Omit Content-Length header")
        flag.BoolVar(&opts.RequestKeepAlive, "request-keepalive", false, "Add Connection: keep-alive")
        flag.IntVar(&opts.Threads, "t", 200, "Number of threads")
        flag.BoolVar(&opts.Debug, "debug", false, "Enable debug output")

        var headers multiStringFlag
        flag.Var(&headers, "H", "Add header (e.g., -H \"Key: Value\")")
        opts.Headers = headers

        flag.Parse()

        if opts.TargetURL == "" || opts.Wordlist == "" {
                fmt.Println("Error: -u and -w are required.")
                flag.Usage()
                os.Exit(1) // ← Now os is defined
        }

        if opts.Debug {
                fmt.Printf("Debug: Target=%s, Threads=%d, Raw=%v\n", opts.TargetURL, opts.Threads, opts.RawMethod)
        }

        return opts
}
