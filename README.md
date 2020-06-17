# PHOTONAI Explorer - build instructions
Because of https://github.com/getlantern/systray it seems to not be compatible with cross-compiling. See this repo for prerequisites.

# Building requirements
`go get -u github.com/gobuffalo/packr/packr`
## Windows only
`go get github.com/akavel/rsrc`
## Linux only
`sudo apt-get install gcc libgtk-3-dev libappindicator3-dev`
On Linux Mint, `libxapp-dev` is also required.
 
# Building
## Linux
1. put Website in dist/ directory 
2. `packr`
3. `go build -o photonai_explorer .`

## Windows (also crosscompilation)
1. put Website in dist/ directory 
2. `packr`
3. `rsrc -manifest photonai_explorer.manifest`
4. `go build -ldflags="-H windowsgui" -o photonai_explorer.exe .`
    * for cross-compilation: `env GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o photonai_explorer.exe .` 
    * Be advised when cross-compiling: the create rsrc.syso can interfer with the building process when creating a linux binary
