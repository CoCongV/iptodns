set -eu


VERSION=$(git describe --abbrev=0 --tags)
BUILDTIME=$(date -u +%Y/%m/%d-%H:%M:%S)

go build -o dist/iptodns-"$VERSION" main.go
