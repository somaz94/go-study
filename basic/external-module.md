## Module Method
Using an external module in your Go project involves a few straightforward steps. These steps help manage dependencies effectively with Go Modules, which is the dependency management system used by Go. Here's how you can include and use an external module in your project:

### Step 1: Initialize Your Project as a Module (If Not Already Done)
If your project is not already initialized as a module, navigate to the root directory of your project and run:
```go
// root
go mod init <module name>
```
- Replace `<module name>` with the name of your project, which typically follows the format of an internet domain you own, in reverse order (e.g., com.example.mymodule). 
- If you've already done this step (as indicated by the presence of a go.mod file in your project's root directory), you can skip to the next step.

### Step 2: Add the External Module
To use an external module, you simply need to import it in your Go files just like you would with any other package. For example, if you want to use the popular gorilla/mux package for handling HTTP requests, you would add the following import statement in one of your external-module.go files:
```go
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func external_module() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
```

### Step 3: Download the Module
After adding the import statement for the external module, run the following command from your project's root directory to download the module and add it to your project's dependencies:
```go
go mod tidy
...
go: finding module for package github.com/gorilla/mux
go: downloading github.com/gorilla/mux v1.8.1
go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.1
```
- This command cleans up your project's dependencies, adding any missing modules and removing unused ones. It will automatically download the `gorilla/mux` package (or any other packages you've imported) and update your `go.mod` and `go.sum` files accordingly.
- 

### Step 4: Build and Run Your Project
With the external module now used in your project, you can build and run your project as usual with:
```go
go build
```

And then run your built executable, or directly run:
```go
go run main.go external_module.go
...

http://localhost:8080/
Hello, My Name is Somaz!
```
