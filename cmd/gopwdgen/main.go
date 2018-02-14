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

package main

import (
	"fmt"
	"os"
	"strings"

	gopwdgen "github.com/berrak/gopwdgen/pkg/lib"
)

const (
	successExitCode = 0
	errorExitCode   = 1
)

// Version to be added by build chain (linker option)
var Version = "Nothing Provided"

func main() {

	cmdline := os.Args
	var ps gopwdgen.ParsedStruct

	ps = gopwdgen.Init(cmdline, ps, Version)

	// first argument is file path for hashing
	if ps.FirstArg != "" && ps.FirstArgIsFile {
		// Generate md5, sha1 and  sha256 file hashes
		md5sum := gopwdgen.HashFileMD5(ps.FirstArg)
		fmt.Printf("md5sum: %s %s\n", ps.FirstArg, md5sum)
		hash := gopwdgen.HashFile(ps.FirstArg)
		fmt.Printf("  sha1: %s %s\n", ps.FirstArg, hash)
		hash256 := gopwdgen.HashFile256(ps.FirstArg)
		fmt.Printf("sha256: %s %s\n", ps.FirstArg, hash256)
		gopwdgen.ExitApp(successExitCode)
	}

	// no command line argument given --> generate random password
	if ps.FirstArg == "" {
		pwd := gopwdgen.RandomPassword()
		fmt.Println(pwd)
		gopwdgen.ExitApp(successExitCode)
	}

	if ps.OracleString != "" && ps.SaltString != "" {
		fmt.Printf("Contradicting salts (ORACLE_SID=%s <-> GOPWDGEN_SALT=%s) in the environment. Remove one of them.\n", ps.OracleString, ps.SaltString)
		gopwdgen.ExitApp(errorExitCode)
	}

	// Use environment ORACLE_SID for salt scrypt-hashing
	if ps.OracleString != "" && ps.FirstArg != "" {
		// Generate password on supplied string
		pwd := gopwdgen.SCryptSalt(ps.FirstArg, ps.OracleString)
		if ps.IsUpper {
			pwd = strings.ToUpper(pwd)
		}
		fmt.Println(pwd)
		gopwdgen.ExitApp(successExitCode)
	}

	// Use environment GOPWDGEN_SALT for salt scrypt-hashing
	if ps.SaltString != "" && ps.FirstArg != "" {
		// Generate password on supplied string
		pwd := gopwdgen.SCryptSalt(ps.FirstArg, ps.SaltString)
		if ps.IsUpper {
			pwd = strings.ToUpper(pwd)
		}
		fmt.Println(pwd)
		gopwdgen.ExitApp(successExitCode)
	}

	gopwdgen.ExitApp(successExitCode)

}
