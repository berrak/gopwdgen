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
	"testing"
)

func TestSCryptSalt(t *testing.T) {
	password := "mysecret"
	salt := "mysalt"
	want := "f3f944d9f5ddd0617abb78da9b267dddfbcfd3329cb447c4ff83ac7b6b584979"
	got := SCryptSalt(password, salt)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
	// Run again -- should not change
	got = SCryptSalt(password, salt)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
	// salt and password equal
	password = "mysalt"
	want = "c3d9267dc79093fac9591dbbedb413910efe181d7ef669732642cc55e1ed95c4"
	got = SCryptSalt(password, salt)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
	// a very short salt and password
	salt = "1"
	password = "0"
	want = "07cc1a2a1192178929f38591ca80aaa6fcf6b5fb9efcf258270072023d90ad71"
	got = SCryptSalt(password, salt)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
	// salt and password equal
	password = "1"
	want = "33146a9ae28af8a2bdeab85b7b331bab1f3b0e79fb47035125c53aa2eb5f6a70"
	got = SCryptSalt(password, salt)
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}

}
