# build
Because of https://github.com/getlantern/systray it seems to not be compatible with cross-compiling. See this repo for prerequisites.

# Building requirements
`go get -u github.com/gobuffalo/packr/packr`
## Windows only
`go get github.com/akavel/rsrc`
 
# Building
## Linux
1. put Website in dist/ directory 
2. `packr`
3. `go build -o photon_explorer .`

## Windows (also crosscompilation)
1. put Website in dist/ directory 
2. `packr`
3. `rsrc -manifest photon_explorer.manifest`
4. `go build -ldflags="-H windowsgui" -o photon_explorer.exe .`
    * for cross-compilation: `env GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o photon_explorer.exe .` 