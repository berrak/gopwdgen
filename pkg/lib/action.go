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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// HashFile256 calculates sha256 on fileName
func HashFile256(fileName string) (sha256Hash string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	sha256Hash = hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}

// HashFile calculates sha1 on fileName
func HashFile(fileName string) (sha1Hash string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sha1Hash = hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

// HashFileMD5 calculates md5sum on fileName
func HashFileMD5(fileName string) (md5Hash string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	md5Hash = hex.EncodeToString(h.Sum(nil))
	return md5Hash
}

// RandomPassword generates a password that is 32 characters long with 10 digits, 10 symbols.
// Allowing upper and lower case letters (allowing only lowercase set to false).
// Allowing (set to 'true') repeat characters.
func RandomPassword() (pwd string) {

	pwd = MustGenerate(32, 10, 10, false, true)
	return (pwd)
}
