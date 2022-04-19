# Jack virtual machine
JVM like virtual machine implementation for Jack compiled vm code(like Java's bytecode) for Hack platform<br>
It works as a stack machine

## Build & use
```bash
go build -o jackvm
./jackvm DirectoryOfVmFiles
```
This will create a DirectoryOfVmFiles.asm file inside of DirectoryOfVmFiles

## Stack machine design principle
![alt text](https://github.com/overload77/hack-software-suite/blob/main/virtual-machine/stack-machine.png?raw=true)
