# windows
export GOOS=windows; export GOARCH=amd64; go build -o shrun_${GOOS}_${GOARCH}.exe; mv shrun_${GOOS}_${GOARCH}.exe binaries;
export GOOS=windows; export GOARCH=386; go build -o shrun_${GOOS}_${GOARCH}.exe; mv shrun_${GOOS}_${GOARCH}.exe binaries;

# linux
export GOOS=linux; export GOARCH=amd64; go build -o shrun_${GOOS}_${GOARCH}; mv shrun_${GOOS}_${GOARCH} binaries;
export GOOS=linux; export GOARCH=386; go build -o shrun_${GOOS}_${GOARCH}; mv shrun_${GOOS}_${GOARCH} binaries;
export GOOS=linux; export GOARCH=arm64; go build -o shrun_${GOOS}_${GOARCH}; mv shrun_${GOOS}_${GOARCH} binaries;

# mac
export GOOS=darwin; export GOARCH=arm64; go build -o shrun_${GOOS}_${GOARCH}; mv shrun_${GOOS}_${GOARCH} binaries;
export GOOS=darwin; export GOARCH=amd64; go build -o shrun_${GOOS}_${GOARCH}; mv shrun_${GOOS}_${GOARCH} binaries;

# for more combinations see - https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63