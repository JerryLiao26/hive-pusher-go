# hive-pusher-go
A code pusher for hive written in Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/JerryLiao26/hive-pusher-go)](https://goreportcard.com/report/github.com/JerryLiao26/hive-pusher-go)
[![codebeat badge](https://codebeat.co/badges/a5075de0-34eb-4840-abe0-e4d04f7dcbdb)](https://codebeat.co/projects/github-com-jerryliao26-hive-pusher-go-master)
[![License](https://img.shields.io/github/license/JerryLiao26/hive-pusher-go.svg)](https://opensource.org/licenses/MIT)

## Usage
Get the package with:
```
go get -u github.com/JerryLiao26/hive-pusher-go
```
Then, in your code, use pusher by:
```go
package main

import "github.com/JerryLiao26/hive-pusher-go"

func main() {
	pusher := hive.Pusher{
    	Path: "Your Hive Server Path",
    	Token: "Your Pusher Token",
    }
    
    pusher.Send("Your Message")
}
```