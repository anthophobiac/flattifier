# Flattifier

A minimal CLI tool for converting JSON to CSV written in Go using:

- [Cobra](https://github.com/spf13/cobra) – for CLI structure
- [fatih/color](https://github.com/fatih/color) – for colourful terminal output

---

### How to use

#### 1. Clone the repo

```bash
git clone https://github.com/anthophobiac/flattifier.git
cd flattifier
go install
```

#### 2. Run the convert command
```bash
flattifier convert -i input.json -o output.csv
```

### Project Structure
```
flattifier/
├── cmd/                 # Cobra commands
├── internal/converter   # Contains converting logic
├── pkg/logger/          # Logger that uses fatih/color
├── main.go              # Application entrypoint
```
