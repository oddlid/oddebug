/*
   Copyright 2017 Odd Eivind Ebbesen

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

package oddebug

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func DebugParts() (string, string, int) {
	pc, fn, line, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name(), fn, line
}

func BaseName(path string) string {
	return filepath.Base(path)
}

func Strip(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func DebugInfo() string {
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)
}

func DebugInfoMedium(strip string) string {
	var funcname string
	pc, fn, line, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	if strip != "" {
		funcname = Strip(fname, strip)
	} else {
		funcname = fname
	}
	return fmt.Sprintf("%s[%s:%d]", funcname, filepath.Base(fn), line)
}

func DebugInfoShort() string {
	_, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("[%s:%d]", filepath.Base(fn), line)
}
