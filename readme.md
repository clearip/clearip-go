# clear ip golang library

golang client library for clear ip

## Installation

```bash
go get -u github.com/clearip/clearip-go
```

## usage

Get ip info:

```go
package main

import (
	clearip "clearip"
	"fmt"
)

func main() {

	clearipClient, err := clearip.NewClient("API key here")
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := clearipClient.GetIpinfo("ip here")
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
