# DirRumble - The fastest path to the other side
### Tensor Security Academy • Developed by Team Alpha
```
   _____  _        _____             _     
  |  __ \| |      |  __ \           | |    
  | |  | | |      | |__) |___  _ __ | | __ 
  | |  | | |      |  _  // _ \| '_ \| |/ / 
  | |__| | |____  | | \ \ (_) | |_) |   <  
  |_____/|______| |_|  \_\___/| .__/|_|\_\ 
                               | |         
                               |_|
```      

***Fast • Raw • Unfiltered***
A modern, high-performance HTTP directory and fuzzing tool written in Go. Inspired by FFUF and UFF, DirRumble removes all automatic normalization from the Go HTTP client — giving you full control over every byte sent to the target.

Perfect for penetration testing, bug bounty, red teaming, and WAF bypass research. 

### Features (v0.1)
- Extremely fast concurrent scanning (200+ threads default)
- Real-time colored output with status code, size, word count, and response time
- Raw HTTP request mode (no normalization, no auto-headers)
- FUZZ keyword replacement in raw request files or -X payload
- Absolute URI support (--opaque)
- Disable Content-Length, custom Connection header, malformed header support
- Built-in Tensor Security Academy User-Agent and branding
- Zero external dependencies — single static binary

### Installation
#### Option 1: Install directly (recommended)
```
go install github.com/tools-tensorsecurityacademy/dirrumble/cmd/dirrumble@latest
```
#### Option 2: Build from source
```
git clone https://github.com/tools-tensorsecurityacademy/dirrumble.git
cd dirrumble
go build -ldflags="-s -w" -o dirrumble ./cmd/dirrumble
sudo mv dirrumble /usr/local/bin/
```

#### Usage Examples
```
# Basic directory brute-force
dirrumble -u https://target.com -w wordlist.txt

# High-performance scan
dirrumble -u https://example.com -w /usr/share/wordlists/dirbuster/directory-list-2.3-medium.txt -t 400

# Raw request mode (UFF-style)
dirrumble -u https://target.com -w payloads.txt -request raw.req

# WAF bypass techniques
dirrumble -u https://site.com/FUZZ \
  -w common.txt \
  -H "X-Forwarded-For: 127.0.0.1" \
  -H "User-Agent: Mozilla/5.0" \
  --opaque \
  --no-content-length
```

#### Command Line Options
```
Flag,Description,Default
-u,Target URL (required),—
-w,Path to wordlist (required),—
-t,Number of concurrent threads,200
-X,HTTP method (or raw template in raw mode),GET
-H,Add custom header (can be used multiple times),—
-request,Load raw request from file (FUZZ placeholder),—
--raw-method,Treat -X as raw request template,false
--opaque,Use absolute URI in request line,false
--no-content-length,Omit Content-Length header,false
--request-keepalive,Force Connection: keep-alive,false
--debug,Enable verbose debug output,false
```

#### Sample Output
```
[32m200]   3741 bytes |   89 w |  124 ms | https://target.com/admin
[33m403]     12 bytes |    4 w |   67 ms | https://target.com/server-status
[36m301]      0 bytes |    0 w |   54 ms | https://target.com/login → /login/
[32m200]   9381 bytes | 1420 w |  112 ms | https://target.com/.git/HEAD
```

#### Roadmap (Next Releases)
```
Version,Planned Features
v0.2,"JSON/CSV output, match/filter codes, proxy support"
v0.3,"Recursion, rate limiting, auto-calibration"
v1.0,"Virtual host fuzzing, extensions, plugins"
```

#### License
```
MIT License
Copyright © 2025 Tensor Security Academy – Team Alpha
```

Permission is hereby granted, free of charge, to any person obtaining a copy of this software...

#### Credits & Thanks:
- Inspired by ffuf by @ffuf and uff by @sw33tLie
- Built with love by Team Alpha at Tensor Security Academy
