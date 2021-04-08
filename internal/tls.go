/*
Copyright © 2021 Wilson Husin

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

import (
	"fmt"

	"github.com/wilsonehusin/confiar/internal/cryptographer"
)

var cryptoBackend cryptographer.Cryptographer

func NewTLSSelfAuthority(backendType string, names []string, ips []string) error {
	switch backendType {
	case "gostd":
		cryptoBackend = &cryptographer.GoStd{}
	default:
		return fmt.Errorf("unknown cryptographer backend type: %s", backendType)
	}
	return cryptoBackend.NewTLSSelfAuthority(names, ips)
}
