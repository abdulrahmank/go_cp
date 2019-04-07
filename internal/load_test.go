package go_cp__test

import (
	"github.com/abdulrahmank/go_cp"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("Should be able to load a file from disk to memory", func(t *testing.T) {
		if byteArr, err := go_cp.Load("../test_resource/sample.txt"); err != nil {
			t.Errorf("Error occured %v\n", err)
		} else {
			if strings.TrimSpace(string(byteArr)) != "This is a sample file" {
				t.Errorf("Expected %s but was %s\n", "This is a sample file", string(byteArr))
			}
		}
	})
}
