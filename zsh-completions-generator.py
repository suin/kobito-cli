#!/usr/bin/env python
# -*- coding: utf-8 -*-

import commands
import os

kobito_completion_file = os.path.dirname(__file__) + "/zsh-completions/_kobito"
kobito_command = os.path.dirname(__file__) + "/bin/kobito"

def main():
	help = get_kobito_help()
	subcommands = parse_help(help)
	code = to_zsh_completion_code(subcommands)
	script = get_completion_script()
	new_script = embed_code_to_script(code, script)
	put_completion_script(new_script)

def get_kobito_help():
	status, help = commands.getstatusoutput(kobito_command)

	if status != 0:
		raise Exception("command error")

	return help

def parse_help(help):
	def is_command_line(line): return "::" in line
	def to_command_and_description(line):
		command_and_arguments, description = line.split("::")
		command = command_and_arguments.strip().split(" ")[0]
		return command, description.strip()

	lines = filter(is_command_line, help.splitlines())
	subcommands = map(to_command_and_description, lines)

	return subcommands

def to_zsh_completion_code(subcommands):
	def compile(command):
		return (" " * 4) + "'%s:%s'" % command

	return "\n".join(map(compile, subcommands)) + "\n"

def get_completion_script():
	return open(kobito_completion_file, 'r').read()

def embed_code_to_script(code, script):
	head, tail = script.split("subcommands=(", 1)
	body, foot = tail.split(")", 1)
	new_script = head + "subcommands=(\n" + code + "  )" + foot
	return new_script

def put_completion_script(script):
	open(kobito_completion_file, "w").write(script)

if __name__ == "__main__":
	main()
