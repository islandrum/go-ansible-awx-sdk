# go-ansible-awx-sdk

Ansible AWX SDK for the Go programming language.


This SDK has been developed against AWX `14.0.0`.

## Installing

```
go get -u github.com/islandrum/go-ansible-awx-sdk
```

## Example

We can simply import `goawx` and call its services, such as the PingService:

```
import (
    "log"
    awx "github.com/islandrum/go-ansible-awx-sdk"
)

func main() {
    client := awx.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := client.PingService.Ping()
    if err != nil {
        log.Fatalf("Ping awx err: %s", err)
    }

    log.Println("Ping awx: ", result)
}
```

More examples can be found at [here](https://github.com/islandrum/go-ansible-awx-sdk/tree/main/examples).

## Roadmap

go-ansible-awx-sdk is still in development, and its roadmap could be found [here](https://github.com/islandrum/go-ansible-awx-sdk/blob/main/ROADMAP.md).

## Contribute

* Submit a [pull request](https://github.com/islandrum/go-ansible-awx-sdk/pulls) for fixes or features;

## awx sdk

* this sdk is forked from awx https://github.com/Kaginari/ansible-tower-sdk
