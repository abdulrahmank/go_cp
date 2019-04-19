**GoCP (Go copy)**

**GoCP** is a command line tool to copy files from given source to destination.

It uses the concept of memory mapping and concurrency to achieve the speed of linux cp.

**Usage**
```
gocp {source_dir|source_file} {destination_dir}

Note: source_dir, destination_dir should be given till last '/'

```
**BenchMark** <sub>Refer: copy_test.go</sub>

| Command  |      Time taken for one operation      |  Size of files copied |
|----------|:--------------------------------------:|----------------------:|
|gocp|35.755392714 s|8GB|
|cp|11.453348709 s|8GB|