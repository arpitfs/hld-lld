![plugin](../../assets/plugin.gif)

The plugin in go can be used to add a new feature or dynamically load a new library to the existing service.
This is supported in linux and macOS.

To create a library use below command:

```
go build -buildmode=plugin -o <file-name>.so plugin.go
```

The above command creates a complies the so file which is converted into machine code and it can be loaded in the CPU when called from the main program.

The plugin provides methods to relink and relocate the symbols which on runtime can be loaded to the memory and cpu get the address of the instruction during execution.
