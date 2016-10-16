# govdotype

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/govdotype)
[![Build Status](https://travis-ci.org/shamsher31/govdotype.svg)](https://travis-ci.org/shamsher31/govdotype)
[![GitHub release](http://img.shields.io/github/release/shamsher31/govdotype.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)

Return type of video based on mime type which is useful when you are writing video uploading apps.

### How to install
```go
go get github.com/shamsher31/govdotype
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/govdotype"
)

func main() {
  fmt.Println(vdotype.Get("./golang.mp4"))
  // video/mp4
}
```

### Related
[goimgtype](https://github.com/shamsher31/goimgtype)<br>
[goimgext](https://github.com/shamsher31/goimgext)<br>
[govdoext](https://github.com/shamsher31/govdoext)<br>

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
