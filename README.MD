## Github-Backup

A utility to backup github repositories

Written in Go

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