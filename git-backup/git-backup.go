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
	"os"
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

	//parse program arguments
	username := flag.String("q", "", "Github Username to query")
	url := flag.String("cs", "", "Single Repository URL to clone")
	userPublicClone := flag.String("cu", "", "Clone all public repositories for GitHub user")
	userPublicCloneAndTar := flag.String("cut", "", "Clone and Tar all public repositories for GitHub user")
	userPublicCloneTarGZ := flag.String("tgz", "", "Clone, Tar, and Gzip all public repositories for GitHub user")
	removeBackupFromDir := flag.String("r", "", "Remove backup directory")

	flag.Parse()

	start := time.Now()
	fmt.Printf("###################################################################################\n")
	fmt.Printf("# Backup Utility started: %s\n", GetCurrentTimeStamp())
	fmt.Printf("###################################################################################\n")

	if *username != "" {
		GetStats(*username)
	}

	if *url != "" {
		DoAClone(*url)
	}

	if *userPublicClone != "" {
		CloneAllPublicRepos(*userPublicClone)
	}

	if *userPublicCloneAndTar != "" {

		CloneAllPublicRepos(*userPublicCloneAndTar)

		fmt.Printf("###################################################################################\n")
		fmt.Printf("# Tar started: %s\n", GetCurrentTimeStamp())
		fmt.Printf("###################################################################################\n")

		writer, writerErr := os.Create("./backup.tar")

		if writerErr != nil {
			panic(writerErr)
		}

		err := Tar("./backup", writer)

		if err != nil {
			panic(err)
		}

		fmt.Printf("###################################################################################\n")
		fmt.Printf("# Tar finished: %s\n", GetCurrentTimeStamp())
		fmt.Printf("###################################################################################\n")
	}

	if *userPublicCloneTarGZ != "" {

		CloneAllPublicRepos(*userPublicCloneTarGZ)

		fmt.Printf("###################################################################################\n")
		fmt.Printf("# Tar + Gzip started: %s\n", GetCurrentTimeStamp())
		fmt.Printf("###################################################################################\n")

		writer, writerErr := os.Create("./backup.tar.gz")

		if writerErr != nil {
			panic(writerErr)
		}

		err := TarGZ("./backup", writer)

		if err != nil {
			panic(err)
		}

		fmt.Printf("###################################################################################\n")
		fmt.Printf("# Tar + Gzip finished: %s\n", GetCurrentTimeStamp())
		fmt.Printf("###################################################################################\n")

	}

	if *removeBackupFromDir != "" {
		DeleteBackupFromDir(*removeBackupFromDir)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("###################################################################################\n")
	fmt.Printf("# Backup Utility completed: %s\n", GetCurrentTimeStamp())
	fmt.Printf("# Backup finished in %s\n ", elapsed)
	fmt.Printf("###################################################################################\n")

}
