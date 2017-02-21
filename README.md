#golang
Plaground for learning/experimenting with the Go language

#setup
sudo dnf install golang

add to bashrc:
export GOPATH=~/Repos/personal/golang/
export PATH=$PATH:$GOPATH/bin

##Hello
build: > go install Hello
run: > hello
