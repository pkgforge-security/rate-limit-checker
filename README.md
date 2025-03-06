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
  --method, -X string
      HTTP Method to use for the request (default "GET")
  --threads, -t int
      Number of threads to use (default 10)
  --requests-count -c int
      Number of requests to send (default 500)
  --ignore-code-change, -i
      Continue after the rate limiter code was changed
  --url, -u string
      URL to check
  --output, -o string
      Output file to save the results
```