package main

import (
	"fmt"
	"path/filepath"
)

// ZSH is a singleton instance of ZSH_T
type zsh int

var ZSH zsh

const ZSH_HOOK = `
direnv_hook() { eval "$(%s export zsh)"; };
[[ -z $precmd_functions ]] && precmd_functions=();
precmd_functions=($precmd_functions direnv_hook)
`

func (z zsh) Hook(direnv string) string {
	return fmt.Sprintf(ZSH_HOOK, direnv)
}

func (z zsh) Escape(str string) string {
	return ShellEscape(str)
}

func (z zsh) Export(key, value string) string {
	return "export " + z.Escape(key) + "=" + z.Escape(value) + ";"
}

func (z zsh) Unset(key string) string {
	return "unset " + z.Escape(key) + ";"
}

func (z zsh) AbsPath() (string, error) {
	return filepath.Abs("zsh")
}
