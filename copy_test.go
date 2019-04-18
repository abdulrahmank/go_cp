package main

import (
	"github.com/abdulrahmank/go_cp/internal/mock"
	"github.com/golang/mock/gomock"
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
	mockWriterImpl.EXPECT().Write(content, "/Users/kabdul/go/src/github.com/abdulrahmank/go_cp/test_resource/sample.txt")

	loaderImpl = mockloaderImpl
	writerImpl = mockWriterImpl

	copy([]string{file, dfile})
}

func TestCopyDirWithOnlyFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockloaderImpl := mock.NewMockLoader(ctrl)
	mockWriterImpl := mock.NewMockCpWriter(ctrl)
	file := "./test_resource/"
	dfile := "./test_resource"
	content := []byte("sample content")

	mockloaderImpl.EXPECT().Load("test_resource/sample.txt").Return(content, nil)
	mockWriterImpl.EXPECT().Write(content, "/Users/kabdul/go/src/github.com/abdulrahmank/go_cp/test_resource/sample.txt")

	loaderImpl = mockloaderImpl
	writerImpl = mockWriterImpl

	copy([]string{file, dfile})
}
