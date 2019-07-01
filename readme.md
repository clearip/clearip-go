# iptrace.io Golang client

Golang client for iptrace web-service API to determine location of visitors based on their IP. 

## Installation

```bash
go get -u github.com/getiptrace/iptrace-go
```

## usage

Get ip info:

```go
package main

import (
	iptrace "github.com/getiptrace/iptrace-go"
	"fmt"
)

func main() {

	iptraceClient, err := iptrace.NewClient("API Key HERE")
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := iptraceClient.IPRepo.GetAllDataByIP("IP HERE")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}



```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
