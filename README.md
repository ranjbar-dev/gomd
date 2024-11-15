# gomd

golang tool for creating doc for golang projects


### Install 

`go install github.com/ranjbar-dev/gomd@v0.0.4`


### Usage 

`gomd <action> <...>`

#### Usage parse struct 

command: `gomd parse-struct <path_to_go_folder> <path_to_output_folder>`

example: `gomd parse-struct ./types ./docs`


### TODOS 

- handle enums document in parse-structures 

- new parser for gateway handlers like websocket 