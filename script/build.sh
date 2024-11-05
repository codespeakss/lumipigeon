set -e

current_dir="$(cd "$(dirname "$0")" && pwd)"
echo "$current_dir：$current_dir"
verFile=$current_dir"/../version"
Ver=$(cat $verFile) ;
BuildT=$(date -u +'%Y%m%dT%H%M%SZ')
GitBranch=$(git rev-parse --abbrev-ref HEAD)
GitCommit=$(git rev-parse --short HEAD)

CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o out/lumilet -ldflags "-X main.Version=$Ver -X main.BuildTime=$BuildT -X main.GitBranch=$GitBranch -X main.GitCommit=$GitCommit" cmd/version.go cmd/lumilet.go ;
echo "module 'lumilet' # build end with code: "$?

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o out/lumipigeon -ldflags "-X main.Version=$Ver -X main.BuildTime=$BuildT -X main.GitBranch=$GitBranch -X main.GitCommit=$GitCommit" cmd/version.go cmd/lumipigeon.go ;
echo "module 'lumipigon'  # build end with code: "$?
