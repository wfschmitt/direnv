package main

import (
	"fmt"
	"path/filepath"
)

type bash int

var BASH bash

const BASH_HOOK = `PROMPT_COMMAND="eval \"\$(%s export bash)\";$PROMPT_COMMAND"`

func (b bash) Hook(direnv string) string {
	return fmt.Sprintf(BASH_HOOK, direnv)
}

func (b bash) Escape(str string) string {
	return ShellEscape(str)
}

func (b bash) Export(key, value string) string {
	return "export " + b.Escape(key) + "=" + b.Escape(value) + ";"
}

func (b bash) Unset(key string) string {
	return "unset " + b.Escape(key) + ";"
}

func (b bash) AbsPath() (string, error) {
	return filepath.Abs("bash")
}
