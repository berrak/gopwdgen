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
	"log"
)

// ParsedStruct (exported) contains result of parsing
type ParsedStruct struct {
	OracleString, SaltString string
	FirstArg                 string
	FirstArgIsFile, IsUpper  bool
}

/////////////////////////
// "Globals" within pkg
/////////////////////////
const (
	successExitCode = 0
	oracleEnvSid    = "ORACLE_SID"
	gopwdgenEnvSalt = "GOPWDGEN_SALT"
	gopwdgenEnvUp   = "GOPWDGEN_UP"
)

var (
	versionLetter = "v"
)

// Init parse command line and environment
func Init(arg []string, p ParsedStruct, gopwdgenVersion string) ParsedStruct {

	var parsed ParsedStruct
	args := parseArgsFlags()

	// --version
	if args.version {
		version := versionLetter + gopwdgenVersion
		printVersion(version)
		ExitApp(successExitCode)
	}

	// --help
	if args.pwdhelp {
		printHelp()
		ExitApp(successExitCode)
	}

	// retrieve the environment variables
	parsed.OracleString = getEnvString(oracleEnvSid)
	parsed.SaltString = getEnvString(gopwdgenEnvSalt)
	parsed.IsUpper = getEnvBoolean(gopwdgenEnvUp)

	// no command argument or salts found in environment -> generate random password
	if flag.Arg(0) == "" {
		if parsed.OracleString == "" || parsed.SaltString == "" {
			return parsed
		}
	}

	// any given command line argument?
	if flag.Arg(0) != "" {
		parsed.FirstArg = arg[1]
		// is given argument a string or a file?
		fileExists, err := fileExists(parsed.FirstArg)
		if err != nil {
			log.Fatalf("Error in fileExist(%s)", parsed.FirstArg)
		}
		if fileExists {
			// prepared to hash given file with md5, sha1 and sha256
			parsed.FirstArgIsFile = true
			return parsed
		}
	}

	// done, hash given string with environment salts
	return parsed
}
