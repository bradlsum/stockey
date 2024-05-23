all: getfishy platter
getfishy:
	go build ./cmd/getfishy
platter:
	go build ./cmd/platter
clean:
	rm getfishy platter
