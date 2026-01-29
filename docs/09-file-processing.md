# File Processing in Go

Coming from Node.js `fs` module or Python's file handling? Here's how Go does it.

## Reading Files

### Read Entire File

```javascript
// Node.js
const data = fs.readFileSync('file.txt', 'utf8');
// or async
const data = await fs.promises.readFile('file.txt', 'utf8');
```

```go
// Go
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}
content := string(data) // []byte to string
```

### Read Line by Line

```javascript
// Node.js
const rl = readline.createInterface({ input: fs.createReadStream('file.txt') });
for await (const line of rl) {
    console.log(line);
}
```

```go
// Go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close() // Always close files!

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

## Writing Files

### Write Entire File

```javascript
// Node.js
fs.writeFileSync('output.txt', 'Hello, World!');
```

```go
// Go
err := os.WriteFile("output.txt", []byte("Hello, World!"), 0644)
if err != nil {
    log.Fatal(err)
}
```

### Write with More Control

```go
file, err := os.Create("output.txt") // Creates or truncates
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Write string
file.WriteString("Line 1\n")

// Write bytes
file.Write([]byte("Line 2\n"))

// Buffered writing (better performance)
writer := bufio.NewWriter(file)
writer.WriteString("Line 3\n")
writer.Flush() // Don't forget to flush!
```

### Append to File

```go
file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("Appended line\n")
```

## CSV Processing

Go has a built-in `encoding/csv` package - no external dependencies needed!

### Reading CSV

```javascript
// Node.js (with csv-parse)
import { parse } from 'csv-parse/sync';
const records = parse(csvContent, { columns: true });
```

```go
// Go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    // Read all at once
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // records is [][]string
    for _, row := range records {
        fmt.Println(row[0], row[1]) // Access by index
    }
}
```

### Reading CSV Row by Row (Memory Efficient)

```go
reader := csv.NewReader(file)

// Skip header
header, _ := reader.Read()

for {
    row, err := reader.Read()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }

    // Process row
    fmt.Println(row)
}
```

### CSV to Structs

```go
type Person struct {
    Name  string
    Age   int
    Email string
}

func parseCSVToStructs(filename string) ([]Person, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var people []Person
    for i, row := range records {
        if i == 0 {
            continue // Skip header
        }

        age, _ := strconv.Atoi(row[1])
        people = append(people, Person{
            Name:  row[0],
            Age:   age,
            Email: row[2],
        })
    }

    return people, nil
}
```

### Writing CSV

```go
file, err := os.Create("output.csv")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

// Write header
writer.Write([]string{"Name", "Age", "Email"})

// Write rows
writer.Write([]string{"Alice", "30", "alice@example.com"})
writer.Write([]string{"Bob", "25", "bob@example.com"})

// Or write all at once
data := [][]string{
    {"Name", "Age", "Email"},
    {"Alice", "30", "alice@example.com"},
    {"Bob", "25", "bob@example.com"},
}
writer.WriteAll(data)
```

### CSV Options

```go
reader := csv.NewReader(file)

// Custom delimiter (e.g., TSV)
reader.Comma = '\t'

// Allow variable number of fields
reader.FieldsPerRecord = -1

// Skip comments
reader.Comment = '#'
```

## JSON File Processing

```go
// Read JSON file to struct
type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

func loadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err
    }

    return &config, nil
}

// Write struct to JSON file
func saveConfig(filename string, config *Config) error {
    data, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filename, data, 0644)
}
```

## Working with Paths

```go
import "path/filepath"

// Join paths (cross-platform)
path := filepath.Join("dir", "subdir", "file.txt")

// Get directory and filename
dir := filepath.Dir(path)   // "dir/subdir"
file := filepath.Base(path) // "file.txt"
ext := filepath.Ext(path)   // ".txt"

// Get absolute path
abs, _ := filepath.Abs("relative/path")

// Walk directory
filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.IsDir())
    return nil
})
```

## File Info and Checks

```go
// Check if file exists
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    fmt.Println("File does not exist")
}

// Get file info
info, err := os.Stat("file.txt")
if err == nil {
    fmt.Println("Size:", info.Size())
    fmt.Println("ModTime:", info.ModTime())
    fmt.Println("IsDir:", info.IsDir())
}
```

## Temporary Files

```go
// Create temp file
tmpFile, err := os.CreateTemp("", "prefix-*.txt")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(tmpFile.Name()) // Clean up
defer tmpFile.Close()

tmpFile.WriteString("temporary content")
fmt.Println("Temp file:", tmpFile.Name())
```

## Key Differences from JS/Python

| Concept | JS/Python | Go |
|---------|-----------|-----|
| Close files | Often automatic (GC) | Always use `defer file.Close()` |
| Error handling | try/catch | Check every error return |
| Encoding | Usually UTF-8 default | Bytes by default, explicit string conversion |
| CSV parsing | External packages | Built-in `encoding/csv` |
| Path handling | `path` module | `path/filepath` (cross-platform) |

## Common Patterns

### Process Large File Line by Line

```go
func processLargeFile(filename string, process func(string) error) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Handle long lines
    buf := make([]byte, 0, 64*1024)
    scanner.Buffer(buf, 1024*1024)

    lineNum := 0
    for scanner.Scan() {
        lineNum++
        if err := process(scanner.Text()); err != nil {
            return fmt.Errorf("line %d: %w", lineNum, err)
        }
    }

    return scanner.Err()
}
```

### Copy File

```go
func copyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}
```
