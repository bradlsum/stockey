all: getfishy platter
getfishy:
	go build -o getfishy ./cmd/getfishy
platter:
	go build -o platter ./cmd/platter
clean:
	rm getfishy platter
