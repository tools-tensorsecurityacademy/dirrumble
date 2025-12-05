# DirRumble v0.1
### Tensor Security Academy • Developed by Team Alpha
   _____  _        _____             _     
  |  __ \| |      |  __ \           | |    
  | |  | | |      | |__) |___  _ __ | | __ 
  | |  | | |      |  _  // _ \| '_ \| |/ / 
  | |__| | |____  | | \ \ (_) | |_) |   <  
  |_____/|______| |_|  \_\___/| .__/|_|\_\ 
                               | |         
                               |_|         

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
