package logger

import (
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

func getCallerFrame(skip int) (frame runtime.Frame, ok bool) {
	const skipOffset = 2 // skip getCallerFrame and Callers

	pc := make([]uintptr, 1)
	numFrames := runtime.Callers(skip+skipOffset, pc[:])
	if numFrames < 1 {
		return
	}

	frame, _ = runtime.CallersFrames(pc).Next()
	return frame, frame.PC != 0
}

func getFolderFile(s string) string {
	const pathCount = 0

	b := 0
	a := 0
	for i := len(s) - 2; i > 0; i-- {
		if os.IsPathSeparator(s[i]) {
			a = i + 1
			if b > pathCount {
				break
			}
			b++
		}
	}

	return s[a:]
}

func getStack(skip int) []string {
	const seperator2 = ": "
	stack := string(debug.Stack())
	// + 2 -> skip: "runtime/debug.stack" and "daneshvar/sesame/logger.(*logger).stack"
	// * 2 -> 2 lines per a call
	// + 1 skip: "goroutine 1 [running]:"
	stacks := make([]string, 0)
	if skip > 0 {
		skip = (skip+2)*2 + 1
	}
	prev := 0
	line := ""
	for i := range stack {
		if stack[i] == '\n' {
			if skip <= 0 {
				s := stack[prev:i]
				if skip < 0 {
					if strings.Contains(s, "runtime/panic") {
						skip = 0
					}
				} else {
					k := 2
					for j := len(s) - 1; j >= 0; j-- {
						if s[j] == '/' {
							k--
							if k == 0 {
								if p := strings.Index(s, " +0x"); p > 0 {
									stacks = append(stacks, line+getFolderFile(strings.TrimSpace(s[:p])))
									line = ""
								} else {
									stacks = append(stacks, line+strings.TrimSpace(s))
									line = ""
								}
								break
							}
						}
					}
					if k != 0 {
						line += strings.TrimSpace(s) + seperator2
					}
				}
			} else {
				skip--
			}
			prev = i + 1
		}
	}
	return stacks
}
