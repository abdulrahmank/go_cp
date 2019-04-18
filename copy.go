package main

import (
	"github.com/abdulrahmank/go_cp/loader"
	"github.com/abdulrahmank/go_cp/writer"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	if !file.IsDir() {
		wg.Add(1)
		go copyFile(srcPath, destPath+"/"+file.Name(), &wg)
	} else {
		infos, _ := ioutil.ReadDir(srcPath)
		for _, info := range infos {
			if !info.IsDir() {
				wg.Add(1)
				go copyFile(srcPath+"/"+info.Name(), destPath+"/"+info.Name(), &wg)
			}
		}
	}
	wg.Wait()
}

func copyFile(srcPath string, destPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	if bytes, e := loaderImpl.Load(srcPath); e != nil {
		log.Fatal(e)
	} else {
		e = writerImpl.Write(bytes, destPath)
		if e != nil {
			log.Fatal(e)
		}
	}
}
