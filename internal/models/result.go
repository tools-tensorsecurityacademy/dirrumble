package models

import "time"


type Result struct {
    URL     string    // Fuzzed URL
    Status  int       // HTTP status code
    Length  int64     // Response body length
    Words   int       // Approximate word count (for filtering)
    Lines   int       // Number of lines in response
    Time    time.Duration // Response time in milliseconds
    Payload string    // The fuzzed payload
    Error   error     // Any request error (nil if successful)
}
