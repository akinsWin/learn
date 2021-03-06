// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// https://github.com/coreos/etcd/blob/master/pkg/runtime/fds_linux.go
package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {
	fl, err := FDLimit()
	if err != nil {
		panic(err)
	}

	fu, err := FDUsage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("FDLimit: %d\n", fl)
	fmt.Printf("FDUsage: %d\n", fu)
}

func FDLimit() (uint64, error) {
	var rlimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return 0, err
	}
	return rlimit.Cur, nil
}

func FDUsage() (uint64, error) {
	fds, err := ioutil.ReadDir("/proc/self/fd")
	if err != nil {
		return 0, err
	}
	return uint64(len(fds)), nil
}
