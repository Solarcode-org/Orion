#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish powershell; do
	if [ $sh = "powershell" ]; then
		go run . completion powershell >"completions/orion.ps1"
	else
		go run . completion "$sh" >"completions/orion.$sh"
	fi
done
