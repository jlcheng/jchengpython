#!/usr/bin/python

"""
A collection of functions to assist with shell scripting.

This module is distributed under the MIT license. 
"""

from subprocess import Popen, PIPE, call
import re
import shlex
import os, signal
from os import getuid, chdir
from pwd import getpwuid

def call_list(cmd, **keyargs):
    """
    Executes a command and return its stdout as a list.
    """
    keyargs.update({
            'shell'   : True,
            'stdout'  : PIPE
            })
    o = Popen(cmd, **keyargs).stdout
    if ( o != None ):
        retval = o.readlines()
    else:
        retval = []
    retval = map(lambda line:line.strip(), retval)
    return retval

def pidgrep(*args, **keyargs):
    """
    Similar to ps grep, but return pids only
    keyargs['u']  = username or uid
    args          = list of keywords to match on
    """
    procs = psgrep(*args, **keyargs)
    retval = []
    for proc in procs:
        arr = proc.split()
        if len(arr) > 1:
            retval.append(arr[1])
    return retval

def psgrep(*args, **keyargs):
    """
    Executes ps in a shell and return results as a list.

    keyargs['u']  = username or uid
    args          = list of keywords to match on
    """
    if keyargs.has_key('u'):
        res = call_list('ps -u %s -f' % keyargs['u'])
    else:
        res = call_list('ps -ef')
    for arg in args:
        res = filter(lambda line: line.find(arg) != -1, res)
    return res

def sudo_call(cmd, **keyargs):
    """
    Executes a command as the root user or a specified user, returns the retval from the process.

    keyargs['u']  = username or uid
    """
    args1 = ['/usr/bin/sudo']
    if keyargs.has_key('u'):
        args1 = shlex.split('/usr/bin/sudo -u %s' % keyargs['u'])
    args2 = shlex.split(cmd)
    args = args1 + args2
    return call(args)

def service_help_start(cmd, sname, pskey=None, **keyargs):
    """
    Starts a service using a shell command. Optionally specifying a user.

    keyargs['u']  = username or uid
    """
    if pskey == None:
        pskey = sname

    if len(psgrep(pskey)) == 0:
        use_sudo = (keyargs.has_key('u') and getpwuid(getuid()) != keyargs['u'])
        if ( use_sudo ):
            retval = sudo_call(cmd, **keyargs)
        else:
            retval = call(cmd, shell=True)
        if retval == 0:
            print '%s started' % sname
        else:
            print 'error starting %s, error code = %s' % (sname, retval)
        return retval
    else:
        print 'process "%s" already started' %sname
        return 1

