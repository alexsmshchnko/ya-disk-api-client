## Installation

```
go get -u github.com/alexsmshchnko/ya-disk-api-client
```

## Description
Implementation of https://yandex.ru/dev/disk/api/concepts/about.html

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	yaDisk "github.com/alexsmshchnko/ya-disk-api-client"
)

var (
	authToken string
)

func main() {
	if authToken = os.Getenv("YA_DISK_AUTH_TOKEN"); authToken == "" {
		log.Fatal(fmt.Errorf("failed to load env variable %s", "YA_DISK_AUTH_TOKEN"))
	}

	client, err := yaDisk.NewClient(authToken, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	disk, _, err := client.GetDiskInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(disk)
}
```