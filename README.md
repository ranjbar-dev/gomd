# gomd

golang tool for creating doc for golang projects


### Build binary 

`go build -o ./build ./cmd/gomd.go`


### Usage 

`gomd <action> <...>`

#### Usage parse struct 

command: `gomd parse-struct <path_to_go_folder> <path_to_output_folder>`

example: `gomd parse-struct ./types ./docs`