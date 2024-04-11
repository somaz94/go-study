## Module Method
Given the issues you're encountering and starting from the beginning, here’s a simplified step-by-step guide to help you set up and run your Go project with local packages using Go modules.

### Step 1: Initialize Your Main Module
Navigate to the root directory of your project. This is where your main.go and music_app.go files are located. We will initialize a Go module here.
```go
// root
go mod init musicapp

// src/24lab.net/testlib
cd src/24lab.net/testlib
go mod init 24lab.net/testlib
```

### Step 2: Adjust Import Path in Your Go Files
In your music_app.go file, make sure you are importing the testlib package using its full path as defined by the module system, not relative paths. Since you've initialized the testlib package as its own module with the path 24lab.net/testlib, use that as the import path.
```go
// music_app.go
package main

import "24lab.net/testlib"

func music_app() {
	song := testlib.GetMusic("Alicia Keys")
	println(song)
}

// main.go
package main

func main() {
	music_app()
}
```

### Step 3: Replace Directive in `go.mod`
Back in the root of your project, where the go.mod file is, you need to add a replace directive to tell Go how to find the testlib package locally.
Open your go.mod file and add a replace directive that points to the local directory of your testlib module relative to the go.mod file.
```go
module musicapp

go 1.16

replace 24lab.net/testlib => ./src/24lab.net/testlib

require 24lab.net/testlib v0.0.0
```

### Step 4: Run go mod tidy
Run go mod tidy to tidy up the module dependencies, including adding any missing module requirements for imports in your code and removing unused modules.
```go
go mod tidy
```

### Step 5: Run Your Application
Now, you should be able to run your application without encountering the previous errors.

```go
go run main.go music_app.go
...
Fallin
```
- This command compiles and runs your program, executing the code defined in `main.go` and `music_app.go`.