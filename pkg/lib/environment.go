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
	"log"

	"github.com/kelseyhightower/envconfig"
)

type specification struct {
	Schema string `envconfig:"ORACLE_SID"`
	Salt   string
	Up     bool `default:"false"`
}

/*
export ORACLE_SID
export GOPWDGEN_SALT
export GOPWDGEN_UP
*/

// GetOracleEnv retrieves environment strings
func getEnvString(envVar string) string {

	var s specification

	envconfig.Process("gopwdgen", &s)
	err := envconfig.Process("gopwdgen", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	switch envVar {
	case oracleEnvSid:
		return s.Schema
	case gopwdgenEnvSalt:
		return s.Salt
	default:
		log.Fatal(err.Error())
	}
	return ""
}

func getEnvBoolean(envVar string) bool {

	var s specification

	envconfig.Process("gopwdgen", &s)
	err := envconfig.Process("gopwdgen", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	switch envVar {
	case gopwdgenEnvUp:
		return s.Up
	default:
		log.Fatal(err.Error())
	}

	return false
}
