package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

var appId = 147
var appSecret = "e88af5e013786589a60="

func main() {
	baseEncodingAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", strconv.Itoa(appId), appSecret)))
	fmt.Println("base encoding :", baseEncodingAuth)
}
