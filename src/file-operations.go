/* ********************************************************************************
*    File Operations
*    Copyright (C) 2020  Alex Oberhofer
*
*    This program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*
*    This program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*
*    You should have received a copy of the GNU General Public License
*    along with this program.  If not, see <https://www.gnu.org/licenses/>.
********************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// CreateDirIfNotExist -
//
// Creates a directory if it doesn't exist
////////////////////////////////////////////////////////////////////////
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// DeleteBackupFromDir -
//
// Deletes a backup from an unzipped directory
////////////////////////////////////////////////////////////////////////
func DeleteBackupFromDir(dir string) {

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Printf("You are about to remove the following directory : %s\n", dir)

	fmt.Printf("Proceed??? (y/n) ->")
	proceed, _ := consoleReader.ReadString('\n')

	if proceed == "y" || proceed == "Y" {
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Will not remove the following directory : %s\n", dir)
		fmt.Printf("Exiting...\n")
		os.Exit(1)
	}

}
