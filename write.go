package go_cp

import (
	"os"
	"syscall"
)

func Write(content []byte, fileName string) error {
	if file, e := os.Create(fileName); e != nil {
		return e
	} else {
		file.Seek(int64(len(content)-1), 0)
		file.Write([]byte(" "))
		if mmap, err := syscall.Mmap(int(file.Fd()), 0, 100, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED); err != nil {
			return err
		} else {
			copy(mmap, content)
			if err := syscall.Munmap(mmap); err != nil {
				return err
			}
		}
	}
	return nil
}
