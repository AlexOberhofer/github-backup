## Github-Backup

A utility to backup github repositories

Written in Go

## Status

This project only really exists to fit my personal use case. Feel free to fork or use the code as you please. 

## Building

```
$ go build github-backup
```

## Setup

Requires github personal access token to use.

```
echo "MY_GITHUB_ACCESS_TOKEN" >> .token
```

## Usage

Clone single repository:

```
$ ./git-backup -cs <url>
```

Clone all public repositories for user

```
$ ./git-backup -cu <username>
```

Clone all repositories for personal access token user

```
$ ./git-backup -cu PERSONAL_ACCESS_TOKEN
```

Delete local backup directory

```
$ ./git-backup -r <directory>
```

## Info
License: GPLv3

by Alex Oberhofer
