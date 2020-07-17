## Github-Backup

A utility to backup github repositories

Written in Go

## Status

This project only really exists to fit my personal use case. Feel free to fork or use the code as you please. (Within the terms of the GPL)

TODO:
- Tar/Zipping archives
- pwd/key authorizaiton

## Building

```
$ go build github-backup
```

## Usage

Query public repositories:

```
$ ./git-backup -q <username>
```

Clone single repository:

```
$ ./git-backup -cs <url>
```

Clone all public repositories for user

```
$ ./git-backup -cu <username>
```

Delete local backup directory

```
$ ./git-backup -r <directory>
```

## Info
License: GPLv3

by Alex Oberhofer
