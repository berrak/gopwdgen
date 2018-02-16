// Package lib contains only internal functions to cmd/gopwdgen
/*
Copyright 2018 The gopwdgen Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package lib

import (
	"flag"
	"fmt"
	"os"
)

type pwdArgs struct {
	pwdhelp bool
	version bool
}

func parseArgsFlags() *pwdArgs {
	versionPtr := flag.Bool("version", false, "Print version")
	pwdhelpPtr := flag.Bool("help", false, "Print help")

	flag.Parse()

	return &pwdArgs{
		version: *versionPtr,
		pwdhelp: *pwdhelpPtr,
	}
}

func printVersion(version string) {
	fmt.Printf("gopwdgen version: %s \n", version)
}

func printHelp() {
	fmt.Printf("Usage:\n\tgopwdgen\n")
	fmt.Printf("\tgopwdgen <filename>|password\n")
	fmt.Printf("\tgopwdgen --help|--version\n")
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
