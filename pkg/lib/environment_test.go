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
	"os"
	"testing"
)

func TestGetEnvString(t *testing.T) {

	os.Clearenv()
	os.Setenv("GOPWDGEN_SALT", "test")

	envString := "GOPWDGEN_SALT"
	want := "test"
	got := getEnvString(envString)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
	// unset environment
	os.Clearenv()
	envString = "GOPWDGEN_SALT"
	want = ""
	got = getEnvString(envString)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}

	os.Setenv("ORACLE_SID", "orcl")
	envString = "ORACLE_SID"
	want = "orcl"
	got = getEnvString(envString)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}

}
