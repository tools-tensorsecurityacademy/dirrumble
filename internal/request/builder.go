package request

import (
        "bufio"
        "io"
        "net"
        "net/http"
        "net/url"
        "os"
        "strings"

        "github.com/tools-tensorsecurityacademy/dirrumble/internal/models"
)

func BuildAndSend(target string, payload string, opts *models.Options) (*http.Response, error) {
        if opts.RawMethod || opts.RequestFile != "" {
                return sendRawRequest(target, payload, opts)
        }

        
        fullURL := target
        if !strings.HasSuffix(target, "/") && payload != "" {
                fullURL += "/"
        }
        fullURL += payload

        req, err := http.NewRequest(opts.Method, fullURL, nil)
        if err != nil {
                return nil, err
        }

        
        for _, h := range opts.Headers {
                parts := strings.SplitN(h, ":", 2)
                if len(parts) == 2 {
                        key := strings.TrimSpace(parts[0])
                        val := strings.TrimSpace(parts[1])
                        req.Header.Set(key, val)
                }
        }

       
        req.Header.Set("User-Agent", "DirRumble/1.0 (Tensor Security Academy - Team Alpha)")
        if opts.RequestKeepAlive {
                req.Header.Set("Connection", "keep-alive")
        }
        if opts.Opaque {
                req.RequestURI = fullURL
        }

      
        tr := &http.Transport{
                DisableCompression: true,
                ForceAttemptHTTP2:  false,
                MaxIdleConns:       opts.Threads,
                MaxConnsPerHost:    opts.Threads,
        }

        client := &http.Client{Transport: tr}
        return client.Do(req)
}

func sendRawRequest(target string, payload string, opts *models.Options) (*http.Response, error) {
        var rawReq string

        if opts.RequestFile != "" {
                data, err := os.ReadFile(opts.RequestFile)
                if err != nil {
                        return nil, err
                }
                rawReq = strings.ReplaceAll(string(data), "FUZZ", payload)
        } else {
                rawReq = strings.ReplaceAll(opts.Method, "FUZZ", payload)
        }

        host := parseHostPort(target)
        conn, err := net.Dial("tcp", host)
        if err != nil {
                return nil, err
        }
        defer conn.Close()

        if _, err = io.WriteString(conn, rawReq); err != nil {
                return nil, err
        }

        return http.ReadResponse(bufio.NewReader(conn), nil)
}

func parseHostPort(u string) string {
        parsed, _ := url.Parse(u)
        host := parsed.Host

        if strings.Contains(host, ":") {
                return host // already has port
        }
        if parsed.Scheme == "https" {
                return host + ":443"
        }
        return host + ":80"
}
