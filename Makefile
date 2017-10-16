.PHONY: main

main: clean
	set -x \
		&& GOOS=darwin GOGC=off GOARCH=amd64 \
			go build \
				-ldflags "-extldflags \"-static\"" \
				-gcflags '-m' \
				-o ./bin/task

clean:
	rm -rf ./bin/*
