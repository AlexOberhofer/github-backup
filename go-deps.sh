#!/usr/bin/bash

#Script to setup project dependencies

echo "Deploying Go Dependencies..."

go get github.com/go-git/go-git

go get github.com/google/go-github/github

echo "DONE!"