#!/usr/bin/env python

import sys
import subprocess

def call_echo(arg):
    executable=sys.argv[0]
    subprocess.call([executable, "echo", arg])

if len(sys.argv) == 1:
    call_echo("{")
    call_echo("{}")
    call_echo("@{}")
elif len(sys.argv) > 2 and sys.argv[1] == "echo":
    print(sys.argv[2])
