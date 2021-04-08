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
	"regexp"
)

// courtesy of https://www.socketloop.com/tutorials/golang-use-regular-expression-to-validate-domain-name
// but mostly a hack until https://github.com/golang/go/issues/31671
var fqdnRegExp = regexp.MustCompile(`^((([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.)*([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)

func ValidFQDN(fqdn string) bool {
	return fqdnRegExp.MatchString(fqdn)
}

var ipv4RegExp = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)

func ValidIPAddr(ipaddr string) bool {
	return ipv4RegExp.MatchString(ipaddr)
}
