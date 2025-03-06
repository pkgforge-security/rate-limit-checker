### ℹ️ About
This is a fork of [Sh1Yo/rate-limit-checker](https://github.com/Sh1Yo/rate-limit-checker) with this PR Added: https://github.com/Sh1Yo/rate-limit-checker/pull/6

### 🖳 Installation
Use [soar](https://github.com/pkgforge/soar) & Run:
```bash
soar add 'rate-limit-checker#github.com.pkgforge-security.rate-limit-checker'
```

### 🧰 Usage
```mathematica
❯ rate-limit-checker --help
Check whether a domain has a rate limit enabled

Usage:
  rate-limit-checker [flags]

Flags:
  -h, --help                 help for rate-limit-checker
  -i, --ignore-code-change   Continue after the code changing
  -X, --method string        HTTP method to use (default "GET")
  -o, --output string        Output file for logs
  -c, --requests-count int   Number of requests to send (default 1000)
  -t, --threads int          Number of threads to use (default 10)
  -u, --url string           URL to send requests to
```
