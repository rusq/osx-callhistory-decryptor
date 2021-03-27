# MacOS X Call history decryptor/converter to CSV

[![Build Status](https://travis-ci.org/rusq/osx-callhistory-decryptor.svg?branch=master)](https://travis-ci.org/rusq/osx-callhistory-decryptor)

Converts the MacOS X call history to CSV file format.

This is the Golang implementation of the [n0fates'][1] [Call History
Decryptor][2], and is based on [n0fate's][1] presentation descibing the
internals of the database:
https://papers.put.as/papers/macosx/2014/Forensic-artifacts-for-Yosemite-call-history-and-sms-anlaysis-ENG.pdf

Motivation for this implementation was:

* to improve the usability by having just one binary executable;
* increase the execution speed by using the standard library functions;
* provide the convenient output format (CSV); and
* describe the usage to make it more accessible to those who require getting the
  call history from MacOS X for any reason, but lacking the time or the
  technical knowledge required to set up the Python interpreter and packages
  needed for the [ogirinal implementation][2].

All credit for the decryption logic goes to [n0fate][1].

## Purpose
Decrypt and save the macOS call history to a CSV file.

## Download
Downloads are available on the [Releases page][5].

## How this works

The program creates a copy of the original database in a temporary directory and
operates on that copy.  After the Call History has been printed out, the
temporary file is deleted.

The original database is not changed during the execution.

For reference:  macOS stores the Call History data in the following location:

    "$HOME/Library/Application Support/CallHistoryDB/CallHistory.storedata"

## Usage
Start the program with `-h` command line flag to see the usage help.

Simple usage:

    $ ./osx-callhistory-decryptor [flags] [database_file]

Where `database_file` is optional os macOS (on Windows you'd have to provide the
filename).

## macOS

Open the Terminal.app. ([How?][3])

1. Start the call history decryptor:

        $ ./osx-callhistory-decryptor

   It will try to locate the default call history file, make a temporary copy
   and open it.

2. You will be prompted for your user's logon password - this allows the program
   to fetch the callhistory encryption key from the OS X keychain.  You can also
   provide the call history encryption key manually using the `-k` command line
   flag. Example:

        $ ./osx-callhistory-decryptor -k YSBzZWNyZXQga2V5IDEyCg==

3. The output will be printed onto the terminal by default.  You can specify an
   output file by providing the `-o` command line flag:

        $ ./osx-callhistory-decryptor -o output.csv

### Opening a database from a non-default location
If, for any reason, you wish to open a different file than the default, the
first command line parameter should contain the filename location:

    $ ./osx-callhistory-decryptor -o output.csv Calls.db

### Specifying the custom time format
By default the time format is RFC3339 without the "T" time/date separator
(`"2006-01-02 15:04:05Z07:00"`).  Optionally, one can change that behaviour with
the `-time-format` flag by passing a [different format][6].  For example, if is
is required to have just a date and time, invoke program like so:

    $ ./osx-callhistory-decryptor -time-format="2006-01-02 15:04"

The formatting is described in depth in the [Go time package documentation][6].

## Linux, Windows, etc.

You will need to obtain the database and the encryption key from the original
macOS system.

1. Get the copy of the `CallHistory.storedata` from source OS X machine.  The file is stored in this location:
        
        $HOME/Library/Application Support/CallHistoryDB/CallHistory.storedata

    with `$HOME` being the user's home directory.

    Copy it to the same directory where you've unpacked the 'callhistory'

    If you get the "Operation not permitted" on latest MacOSes:

      1. Go into "System Preferences";
      2. Choose "Security and Privacy";
      3. Go to "Privacy" tab, select "Full Disk Access" item;
      4. Add the Utilities/Terminal.app — or whatever you're using — to the list.

2. Get the key from the source macOS X keychain:
    
    1. search the macOS X keychain for the *Call History User Data Key*
    2. double-click the entry, and put the checkmark opposite the "show password" field.
    3. Enter your user's account password and copy the key value to the clipboard.

3. Open the terminal or cmd.exe prompt on Windows ([How?][4]).  Start the
   callhistory decryptor on your machine:

        C:>osx-callhistory-decryptor.exe -k <key value from step 2> <filename from step 1>

4. The output will be printed onto the terminal by default.  You can specify an
   output file by providing the `-o` command line flag:

        C:>osx-callhistory-decryptor.exe -o your_ex_callhistory_lol.csv <filename from step 1>

## Licence 
OS X Call history decryptor

Copyright (C) 2016  n0fate (GPL2 license)

Copyright (C) 2018-2021  rusq (golang implementation, GPL3)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.


[1]: https://github.com/n0fate/
[2]: https://github.com/n0fate/OS-X-Continuity
[3]: http://blog.teamtreehouse.com/introduction-to-the-mac-os-x-command-line
[4]: https://www.wikihow.com/Open-the-Command-Prompt-in-Windows
[5]: https://github.com/rusq/osx-callhistory-decryptor/releases
[6]: https://golang.org/pkg/time/#pkg-constants
