# Slugify

[![Go Report Card](https://goreportcard.com/badge/github.com/anthoz69/go-slugify)](https://goreportcard.com/report/github.com/anthoz69/go-slugify)
[![GoDoc](https://godoc.org/github.com/metal3d/go-slugify?status.svg)](https://godoc.org/github.com/metal3d/go-slugify)

This is a simple package that handle slugify.Marshal() function that returns slugified string. The Slugify is generate slug from string to URL-friendly.

It replaces accentuated chars to non-accentuated and spaces by minus sign. All other chars (non-alphanumeric) are removed.

## New Feature

* Support Thai language
* Trim space from string before marshal

## Installation

    go get -u github.com/anthoz69/go-slugify

## Usage

See http://godoc.org/github.com/metal3d/go-slugify

Example:

```go
slugify.Marshal(string, bool) // bool when set true all strings to lower case.
```

```go
package main

import (
    "fmt"
    "github.com/metal3d/go-slugify"
)

func main(){
    text1 := "Être en été est à mi-chemin de noël"
    slug1 := slugify.Marshal(text1)
    fmt.Println(slug1)
    // print: etre-en-ete-est-a-mi-chemin-de-noel

    text2 := "เมาท์ก่อนหน้าอัลไซเมอร์จัมโบ้เอเซีย เทรลเลอร์ช็อป ปิ้งคอมเพล็กซ์แฟรี่ บลูเบอร์รี"
    slug2 := slugify.Marshal(text2)
    fmt.Println(slug2)
    // print: เมาท์ก่อนหน้าอัลไซเมอร์จัมโบ้เอเซีย-เทรลเลอร์ช็อป-ปิ้งคอมเพล็กซ์แฟรี่-บลูเบอร์รี
}
```

# Misc

A big thanks to the author of the Javascript function that is the base of this package:
http://irz.fr/slugme-permalien-javascript-slug/

