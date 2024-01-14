# go-huawei-hg659

Go client to interact with a Huawei HG659 router

## Usage

```bash
go get -u github.com/axatol/go-huawei-hg659
```

## Example

```go
package main

import huaweihg659 "github.com/axatol/go-huawei-hg659"

func main() {
  ctx := context.Background()
  client := huaweihg659.NewClient("http://192.168.1.1")

  if err := client.Login("username", "password"); err != nil {
    panic(err)
  }

  devices, err := client.ListKnownLANDevices(ctx)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%#v\n", devices)
}
```
