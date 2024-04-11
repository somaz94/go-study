module musicapp

go 1.21.5

replace 24lab.net/testlib => ./src/24lab.net/testlib

require (
	24lab.net/testlib v0.0.0
	github.com/gorilla/mux v1.8.1
)
