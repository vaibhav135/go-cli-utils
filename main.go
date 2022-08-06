/*
Copyright © 2022 Vaibhav Singh Bisht lazygy49@gmail.com

*/
package main

import (
	"github.com/vaibhav135/go-cli-utils/cmd"
	"github.com/vaibhav135/go-cli-utils/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
