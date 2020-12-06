# MacOS X Call history decryptor/converter to CSV

[![Build Status](https://travis-ci.org/rusq/osx-callhistory-decryptor.svg?branch=master)](https://travis-ci.org/rusq/osx-callhistory-decryptor)

Converts the MacOS X call history to CSV file format.

This is a Golang implementation of the [n0fates'][1] [Call History Decryptor][2], and is based on [n0fates'][1] presentation descibing the internals of the database: https://papers.put.as/papers/macosx/2014/Forensic-artifacts-for-Yosemite-call-history-and-sms-anlaysis-ENG.pdf

Motivation for different implementation is:

* to improve usability by having just one binary executable;
* increase the execution speed by using standard library functions;
* providing more convenient output format (CSV); and
* describe the usage to make it more accessible to those who require to get the call history from MacOS X for any reason, but lacking the time or the technical knowledge required to set up the Python interpreter and packages needed for the [ogirinal implementation][2].

All credit for the decryption logic goes to [n0fate][1].

## Purpose
Decrypt and save the call history of the OS X Yosemite+ to a CSV file.

## Download
Downloads are available on [Releases page][5].

## Usage
Start the program with `-h` command line flag to see the usage help.  Available options will differ depending on the OS the program being started on.

## MacOS

Open the Terminal.app. ([How?][3])

1. Get the copy of the `CallHistory.storedata` from source OS X machine.  The file is stored in this location:
        
        "$HOME/Library/Application Support/CallHistoryDB/CallHistory.storedata"

    with `$HOME` being the user's home directory.

    Copy it to the same directory where you've unpacked the 'callhistory':

       $ cp "$HOME/Library/Application Support/CallHistoryDB/CallHistory.storedata" .

2. Start the callhistory decryptor:

        $ ./callhistory

3. You will be prompted for your user's logon password, this allows the program to fetch the callhistory key from the OS X keychain.  You can also provide the call history key manually using the `-k` command line flag.  Example:

        $ ./callhistory -k YSBzZWNyZXQga2V5IDEyCg==

4. The output will be printed onto the terminal by default.  You can specify an output file by providing the `-o` command line flag:

        $ ./callhistory -o output.csv

If the database file is called differently than `CallHistory.storedata`, then use `-f` command line flag to provide the filename:

        $ callhistory -o output.csv -f Calls.db


## Linux, Windows, etc.

You will still to obtain the database and the encryption key from the MacOS system.

1. Get the copy of the `CallHistory.storedata` from source OS X machine.  The file is stored in this location:
        
        $HOME/Library/Application Support/CallHistoryDB/CallHistory.storedata

    with `$HOME` being the user's home directory.

    Copy it to the same directory where you've unpacked the 'callhistory'

    If you get the "Operation not permitted" on latest MacOSes:

      1. Go into "System Preferences";
      2. Choose "Security and Privacy";
      3. Go to "Privacy" tab, select "Full Disk Access" item;
      4. Add the Utilities/Terminal.app — or whatever you're using — to the list.

2. Get the key from the source MacOS X keychain:
    
    1. search the MacOS X keychain for the *Call History User Data Key*
    2. double-click the entry, and put the checkmark opposite the "show password" field.
    3. Enter your user's account password and copy the key value to the clipboard.

3. Open the terminal or cmd.exe prompt on Windows ([How?][4]).  Start the callhistory decryptor on your machine:

        C:>callhistory.exe -k <key value from step 2>

4. The output will be printed onto the terminal by default.  You can specify an output file by providing the `-o` command line flag:

        C:>callhistory.exe -o your_ex_callhistory_lol.csv

If the database file is called differently than `CallHistory.storedata`, then use `-f` command line flag to provide the filename:

        $ callhistory -o your_ex_callhistory_lol.csv -f Calls.db

## Licence 
OS X Call history decryptor

Copyright (C) 2016  n0fate (GPL2 license)

Copyright (C) 2018,2019  rusq (golang implementation, GPL3)

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
