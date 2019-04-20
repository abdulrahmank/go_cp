package main

import (
	"github.com/abdulrahmank/go_cp/internal/mock"
	"github.com/abdulrahmank/go_cp/loader"
	"github.com/abdulrahmank/go_cp/writer"
	"github.com/golang/mock/gomock"
	"os"
	"os/exec"
	"testing"
)

func TestCopyFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockloaderImpl := mock.NewMockLoader(ctrl)
	mockWriterImpl := mock.NewMockCpWriter(ctrl)
	file := "./test_resource/sample.txt"
	dfile := "./test_resource"
	content := []byte("sample content")

	mockloaderImpl.EXPECT().Load(file).Return(content, nil)
	mockWriterImpl.EXPECT().Write(content, "./test_resource/sample.txt")

	loaderImpl = mockloaderImpl
	writerImpl = mockWriterImpl

	copy([]string{file, dfile})
}

func TestCopyDirWithOnlyFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockloaderImpl := mock.NewMockLoader(ctrl)
	mockWriterImpl := mock.NewMockCpWriter(ctrl)
	file := "./test_resource"
	dfile := "./test_resource"
	content := []byte("sample content")

	mockloaderImpl.EXPECT().Load("./test_resource/sample.txt").Return(content, nil)
	mockWriterImpl.EXPECT().Write(content, "./test_resource/sample.txt")

	loaderImpl = mockloaderImpl
	writerImpl = mockWriterImpl

	copy([]string{file, dfile})
}

func BenchmarkGoCp(b *testing.B) {
	os.Mkdir("./dest_cp", os.ModePerm)
	os.Mkdir("./dest_gocp", os.ModePerm)
	os.Mkdir("./benchmark_resources", os.ModePerm)
	os.Mkdir("./benchmark_resources/dir", os.ModePerm)
	command := exec.Command("mkfile", "-n", "4g", "./benchmark_resources/dir/file1.txt")
	if _, e := command.Output(); e != nil {
		b.Error(e)
	}
	command = exec.Command("mkfile", "-n", "4g", "./benchmark_resources/dir/file2.txt")
	if _, e := command.Output(); e != nil {
		b.Error(e)
	}

	b.Run("Copy using gocp", func(b *testing.B) {
		loaderImpl = &loader.MMapLoaderImpl{}
		writerImpl = &writer.CpMMapWriterImpl{}
		copy([]string{"./benchmark_resources/dir", "./dest_gocp"})
	})

	b.Run("Copy using cpio", func(b *testing.B) {
		copyIOCopy(nil, []string{"./benchmark_resources/dir", "./dest_cp"})
	})

	os.RemoveAll("./dest_cp")
	os.RemoveAll("./dest_gocp")
	os.RemoveAll("./benchmark_resources")
}
