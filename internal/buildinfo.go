/*
Copyright © 2021 Wilson Husin <wilsonehusin@gmail.com>

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

package internal

var Version = "v0.2.0"
var GitSHA = "000000"
var Go = "420.69"

func BuildInfo() *map[string]string {
	return &map[string]string{
		"Version": Version,
		"GitSHA":  GitSHA,
		"Go":      Go,
	}
}
