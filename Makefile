.PHONY: main

main: clean
	set -x \
		&& GOOS=darwin GOGC=off GOARCH=amd64 \
			go build \
				-ldflags "-extldflags \"-static\"" \
				-gcflags '-m' \
				-gccgoflags '-O8' \
				-o ./bin/task

run: main
	GOGC=off ./bin/task

clean:
	rm -rf ./bin/*
