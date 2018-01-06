.PHONY: main

main: clean
	set -x \
		&& export CGO_ENABLED="0" \
		&& export CGO_CFLAGS="-g -O9" \
        && export CGO_CXXFLAGS="-g -O9" \
        && export CGO_FFLAGS="-g -O9" \
        && export CGO_LDFLAGS="-g -O9" \
		&& go build \
			-ldflags "-w -s -extldflags \"-static\"" \
			-gcflags '-m' \
			-gccgoflags '-O9' \
			-v \
			-o ./bin/task
asm: clean
	set -x \
	&& go tool compile \
		-E  \
		-S internal/levenshtein.go

run: main
	GOGC=off ./bin/task

clean:
	rm -rf ./bin/*
