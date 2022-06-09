version=""

if [ -f "VERSION" ]; then
    version=`cat VERSION`
fi

if [[ -z $version ]]; then
    if [ -d ".git" ]; then
        version=`git symbolic-ref HEAD | cut -b 12-`-`git rev-parse HEAD`
    else
        version="unknown"
    fi
fi

go build -ldflags "-X main.Version=$version" main.go