package writer_test

import (
	"github.com/abdulrahmank/go_cp/writer"
	"io/ioutil"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Run("Should be able to write the given byte Array to disk", func(t *testing.T) {
		file := "../test_resource/write_sample.txt"
		content := "Writing sample content in file using munmap"
		writerImpl := writer.CpMMapWriterImpl{}
		if e := writerImpl.Write([]byte(content), file);
			e != nil {
			t.Errorf("Error occurred %v\n", e)
		} else {
			bytes, _ := ioutil.ReadFile(file)
			if content != string(bytes) {
				t.Errorf("Expected %s but was %s\n", content, string(bytes))
			}
		}
		os.Remove(file)
	})
}
