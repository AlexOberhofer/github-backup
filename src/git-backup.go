/* ********************************************************************************
*    Git Backup Utility Driver
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
	"flag"
	"fmt"
	"time"
)

//TODO (actually make some flags...)
func Usage() {
	fmt.Printf("Usage: -q <username> \n")
}

//Return current timestamp (to be used for backup filename)
func GetCurrentTimeStamp() string {
	return time.Now().Format(time.RFC850)
}

func main() {

	username := flag.String("q", "", "Github Username to query")
	url := flag.String("cs", "", "Single Repository URL to clone")
	flag.Parse()

	fmt.Printf("Backup Utility started: %s\n", GetCurrentTimeStamp())

	//directory := "./my-repo"

	//fmt.Printf("The following repositories will be cloned: \n")

	//ListRepositories()

	if *username != "" {
		GetStats(*username)
	}

	if *url != "" {
		//DoAClone(*url)
		DoACloneDir(*url, "backup/")
	}

	fmt.Printf("Backup Utility completed: %s\n", GetCurrentTimeStamp())
}
