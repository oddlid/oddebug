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

const VERSION string = "2017-07-24"

var (
	NOOP     bool   = false // set to true if not debugging and you want to speed up funcs that will not be used anyway
	NOOP_MSG string = "oddebug.NOOP=true"
)

func DebugParts() (string, string, int) {
	if NOOP {
		return NOOP_MSG, "", -1
	}
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
	if NOOP {
		return NOOP_MSG
	}
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)
}

func DebugInfoMedium(strip string) string {
	if NOOP {
		return NOOP_MSG
	}
	//fmt.Println("oddebug.DebugInfoMedium() called")
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
	if NOOP {
		return NOOP_MSG
	}
	_, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("[%s:%d]", filepath.Base(fn), line)
}

// DebugForWraps is intended to be used in a wrapper in the calling package, for greater flexibility 
func DebugForWraps(noop bool, callchain_lvl int, pkg_strip_prefix string, file_basename bool) (funcname, filename string, line int) {
	if noop {
		return
	}
	pc, filename, line, _ := runtime.Caller(callchain_lvl)
	funcname = runtime.FuncForPC(pc).Name()
	if pkg_strip_prefix != "" {
		funcname = Strip(funcname, pkg_strip_prefix)
	}
	if file_basename {
		filename = filepath.Base(filename)
	}
	return
}
