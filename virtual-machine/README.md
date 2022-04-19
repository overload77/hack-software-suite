# Hack virtual machine
JVM like virtual machine implementation for Jack compiled vm code for Hack platform<br>
It's a stack machine

## Build & use
```bash
go build -o jackvm
./jackvm DirectoryOfVmFiles
```
This will create a DirectoryOfVmFiles.asm file in the same directory

## Stack machine design principle
![alt text](https://github.com/overload77/hack-software-suite/blob/main/virtual-machine/stack-machine.png?raw=true)
