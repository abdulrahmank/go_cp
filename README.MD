**GoCP (Go copy)**

This is a WIP experimental repo to create a command line tool (_gocp_) faster than CP. Trying to achieve it using 
memory mapped files and concurrency.

It has various methods of copying files:
1. using memory mapped files. (gocp command)
2. using golang io.Copy. (cpio command)

**Usage**
```
gocp {source_dir|source_file} {destination_dir}
cpio {source_dir|source_file} {destination_dir}

Note: source_dir, destination_dir should be given till last '/'

```
**BenchMark** <sub>Refer: copy_test.go</sub>

| Command  |      Time taken for one operation      |  Size of files copied |
|----------|:--------------------------------------:|----------------------:|
|gocp|35.755392714 s|8GB|
|cp|11.453348709 s|8GB|