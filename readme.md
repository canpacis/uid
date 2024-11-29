# UID

## Get

```bash
go get github.com/CanPacis/uid
```

## Usage

```go
import (
  "github.com/CanPacis/uid"
)

func main() {
  // create new id
  id := uid.New()

  // parse a string id
  id, err := uid.Parse("0123456789abcdef")
  // alternatively, if you know it is a valid id
  id := uid.MustParse("0123456789abcdef")

  // validate an id
  if err := uid.Validate("invalid"); err != nil {
    // handle error
  }
}
```

> UID implements the `encoding.BinaryMarshaler`, `encoding.BinaryUnmarshaler`, `encoding.TextMarshaler`, `encoding.TextUnmarshaler`, `json.Marshaler`, `json.Unmarshaler`, `driver.Valuer`, `sql.Scanner` for ease of use with reading from and writing to other mediums.