entrypt := "main.go"
bin_name := "gmk"

# list recipes
default:
    just --list

# run
run:
    go run {{ entrypt }}

# build go-make-it
build:
    go build -o {{ bin_name }} {{ entrypt }}

# clean build artifs
clean:
    rm -rv {{ bin_name }}
