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
	"archive/tar"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	fmt.Printf("Proceed??? Press any key to continue or Ctrl + C to quit. ->")

	proceed := ""
	proceed, _ = consoleReader.ReadString('\n')

	if proceed != "" {
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("\nWill not remove the following directory : %s\n", dir)
		fmt.Printf("Exiting...\n")
		os.Exit(1)
	}

}

func getAuthToken() string {
	data, err := ioutil.ReadFile(".token")

	if err != nil {
		fmt.Println("Error reading .token file", err)
		return ""
	}

	return string(data)
}

// Tar -
//
// Creates a tar file from the backup directory
////////////////////////////////////////////////////////////////////////
func Tar(src string, writers ...io.Writer) error {

	// ensure the src actually exists before trying to tar it
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("Unable to tar files - %v", err.Error())
	}

	mw := io.MultiWriter(writers...)

	tw := tar.NewWriter(mw)
	defer tw.Close()

	// walk path
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {

		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files (thanks to [kumo](https://medium.com/@komuw/just-like-you-did-fbdd7df829d3) for this suggested update)
		if !fi.Mode().IsRegular() {
			return nil
		}

		/* Since we are creating a backup and not active clones - I am electing to ignore the .git folder */
		if strings.Contains(file, ".git") {
			return nil
		}

		// create a new dir/file header
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// update the name to correctly reflect the desired destination when untaring
		header.Name = strings.TrimPrefix(strings.Replace(file, src, "", -1), string(filepath.Separator))

		// write the header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()

		return nil
	})

}
