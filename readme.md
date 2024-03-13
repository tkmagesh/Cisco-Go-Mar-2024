# Go Foundation

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 3:00 PM (20 mins)
- Wind Up       : 5:00 PM

## Methodology
- No powerpoint presentation
- Discuss & Code
- No dedicated time for Q&A

## Repository
- https://github.com/tkmagesh/cisco-go-mar-2024

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for VS Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)
    - Configuring Go Extension
        - cmd + shift + p
        - Go: install/update tools
        - Hit ENTER

### Verification
> go version

## Why Go?
- Simplicity
    - ONLY 25 keywords 
    - NO access modifiers (public / private etc)
    - NO reference types (everything is a "value" by default)
    - NO Pointer arithmatic
    - NO classes (only structs)
    - NO inheritance (only composition)
    - NO exceptions (only errors)
    - NO "try-catch-finally" constructs
    - NO implicit type conversion
    - Language Constructs
        - var, :=
        - if else, switch case
        - for
        - function
        - struct
        - interface
- Performance
    - Close to hardware (No JVM/CLR)
    - Build for each platform
        - Tooling support for cross-compilation
    - On par with c++
- Concurrency
    - Managed Concurrency (built in scheduler)
    - A concurrent operation is represented as a goroutine (lightweight - 4KB)
    - Language support for concurrency
        - go keyword, channel data type, channel operator ( <- ), for-range & select-case constructs
    - API support
        - sync package
        - sync/atomic package
    ![image](./images/go-concurrency.png)

## To execute a go file
> go run <file_name.go>
## To create a build
> go build <file_name.go>

> go build -o <binary_name> <file_name.go>
## To get the list of env variables used by go tool
> go env

> go env <var_1> <var_2>
## To set the values of env variables
> go env -w <var_1>=<val_1> <var_2>=<val_2> ....
## The environment variables for Cross compilation
- GOOS
- GOARCH
## To get the list of supported platforms (OS/processor)
> go tool dist list
## To cross compile
> GOOS=<target_os> GOARCH=<target_arch> go build  <file_name.go>

ex:
> GOOS=windows GOARCH=386 go build  01-hello-world.go