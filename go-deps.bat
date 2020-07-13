::Script to setup project dependencies

ECHO OFF

ECHO Deploying Go Dependencies...

go get github.com/go-git/go-git

go get github.com/google/go-github/github

ECHO DONE!

EXIT \b