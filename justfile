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
