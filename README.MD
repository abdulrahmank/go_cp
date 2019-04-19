**GoCP (Go copy)**

**GoCP** is a command line tool to copy files from given source to destination.

It uses the concept of memory mapping and concurrency to achieve the speed of linux cp.

Current benchmarking is as follows: (Refer copy_test.go for bench mark test)

| Command  |      Time taken for one operation      |  Size of files copied |
|----------|:--------------------------------------:|----------------------:|
|go_cp|10.190916786 s|8GB|
|cp|9.650145525 s|8GB|