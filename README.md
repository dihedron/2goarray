# 2goarray

A simple utility to encode a file (or any other data) into a Go byte slice.

This utility is forked from https://github.com/cratonica/2goarray.

## Installation

Having [set up your Go environment](http://golang.org/doc/install), simply run
```bash
$> go get github.com/cratonica/2goarray
```

## Usage

To create a Golang source file containing your resource as a `[]byte`, pipe it into the utility and capture the output. You can provide a name for the generated slice symbol (default: `Data`) and optionally a package name. For example:
```bash
    $> $GOPATH/bin/2gorray -array MyArray -package mypackage < myimage.png > myimage.go
```
This will output something like:

```golang
    package mypackage

    var MyArray []byte = []byte {
      0x49, 0x20, 0x63, 0x61, 0x6e, 0x27, 0x74, 0x20, 0x62, 0x65, 0x6c, 0x69,
      0x65, 0x76, 0x65, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x61, 0x63, 0x74, 0x75,
      0x61, 0x6c, 0x6c, 0x79, 0x20, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64,
      0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x20, 0x4b, 0x75, 0x64, 0x6f, 0x73,
      0x20, 0x66, 0x6f, 0x72, 0x20, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x20, 0x74,
      0x68, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x2e, 0x0a,
    }
```

In oder to create multiple resources within a single Go source file, you can chain multiple invocations via a simple bash sript along the lines of the following:
```bash
#/bin/sh

OUTPUT=resources.go
PACKAGE=mypackage

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/dihedron/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/dihedron/2goarray
        exit
    fi
fi

if [ $# -eq 0 ]; then
    echo Please specify at least an input file
    exit
fi

echo Generating $OUTPUT
if [ -z "$PACKAGE" ]; then
    cat <<EOF > $OUTPUT
//+build linux darwin
EOF
else 
    cat <<EOF > $OUTPUT
//+build linux darwin
package $PACKAGE
EOF
fi

for resource in "$@"; do
    if [ ! -f "$resource" ]; then
        echo $resource is not a valid file
        exit
    fi    

    # turn dashes and underscores into CamelCase
    arrayname=$(basename -- "$resource")
    arrayname="${arrayname%.*}"
    arrayname="$(echo $arrayname | sed -r 's/(^|[-_])(\w)/\U\2/g')"

    echo >> $OUTPUT 
    cat "$resource" | $GOPATH/bin/2goarray -array $arrayname >> $OUTPUT
    if [ $? -ne 0 ]; then
        echo Failure generating $OUTPUT
        exit
    fi

done

echo Finished
```

## Contributors

The utility was originally developed by:
- [Clint Caywood](https://github.com/cratonica)
- [Paul Vollmer](https://github.com/paulvollmer)

This version provides more fine grained control over the generated output via additional command line switches.