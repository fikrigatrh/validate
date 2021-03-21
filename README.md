# Package validate

Package validate implements

# Installation

Just use go get.

```go
go get github.com/fikrigatrh/validate
```

And then just import the package into your own code.

```go
import (
	"github.com/fikrigatrh/validate"
)
```

# Usage

If you check value include number and want return error

```
res := gtr_validate.Num("1234","5678")
if res {
    //return error
}
```

If you check value include space and want return error
```
res := gtr_validate.Space("nama tengah")
if res {
    //return error
}
```
