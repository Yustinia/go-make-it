binName := "gmk"

# list recipes
default:
    just --list

# run
run:
    go run .

# build go-make-it
build:
    go build -o {{ binName }} .

# clean build artifs
clean:
    rm -rv {{ binName }}

# stable install
install: build
    mv -v {{ binName }} ~/.local/bin/{{ binName }}

# install via symlink (development)
dev-install: build
    ln -sf $(pwd)/{{ binName }} ~/.local/bin/{{ binName }}
