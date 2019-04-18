package main

import (
	"github.com/abdulrahmank/go_cp/loader"
	"github.com/abdulrahmank/go_cp/writer"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var CmdCp = &cobra.Command{
	Use:   "go_cp",
	Short: "To copy files from source to destination",
	Args:  cobra.MinimumNArgs(2),
	Run:   CmdCpFn,
}

var loaderImpl loader.Loader
var writerImpl writer.CpWriter

func main() {
	loaderImpl = &loader.MMapLoaderImpl{}
	writerImpl = &writer.CpMMapWriterImpl{}
	if err := CmdCp.Execute(); err != nil {
		log.Fatal(err)
	}
}

func CmdCpFn(cmd *cobra.Command, args []string) {
	copy(args)
}

func copy(args []string) {
	srcPath := args[0]
	destPath := args[1]
	file, _ := os.Stat(srcPath)
	destDir, _ := filepath.Abs(destPath)
	if !file.IsDir() {
		copyFile(srcPath, destDir+"/"+file.Name())
	} else {
		filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				copyFile(path, destDir+"/"+info.Name())
			}
			return nil
		})
	}
}

func copyFile(srcPath string, destPath string) {
	if bytes, e := loaderImpl.Load(srcPath); e != nil {
		log.Fatal(e)
	} else {
		e = writerImpl.Write(bytes, destPath)
		if e != nil {
			log.Fatal(e)
		}
	}
}
