package unittest_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"asposecellscloud/datasource"
)

func TestBytesSource_ByteData(t *testing.T) {
	src := datasource.BytesSource([]byte("hello"))
	data, err := src.ByteData()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("expected 'hello', got %s", string(data))
	}
}

func TestBytesSource_Open(t *testing.T) {
	src := datasource.BytesSource([]byte("test"))
	rc, err := src.Open()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer rc.Close()
	data, _ := io.ReadAll(rc)
	if string(data) != "test" {
		t.Errorf("expected 'test', got %s", string(data))
	}
}

func TestFilePathSource_ByteData(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	os.WriteFile(path, []byte("file content"), 0644)

	src := datasource.FilePathSource(path)
	data, err := src.ByteData()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != "file content" {
		t.Errorf("expected 'file content', got %s", string(data))
	}
}

func TestFilePathSource_NonExistent(t *testing.T) {
	src := datasource.FilePathSource("/nonexistent/file.txt")
	_, err := src.ByteData()
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestReaderSource_ByteData(t *testing.T) {
	r := io.NopCloser(bytes.NewReader([]byte("reader data")))
	src := datasource.NewReaderSource(r)
	data, err := src.ByteData()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != "reader data" {
		t.Errorf("expected 'reader data', got %s", string(data))
	}
}

func TestBytesSink_Write(t *testing.T) {
	sink := &datasource.BytesSink{}
	w, err := sink.Write()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	w.Write([]byte("output"))
	w.Close()

	if string(sink.Bytes()) != "output" {
		t.Errorf("expected 'output', got %s", string(sink.Bytes()))
	}
}

func TestFilePathSink_Write(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	sink := datasource.FilePathSink(path)

	w, err := sink.Write()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	w.Write([]byte("written"))
	w.Close()

	data, _ := os.ReadFile(path)
	if string(data) != "written" {
		t.Errorf("expected 'written', got %s", string(data))
	}
}
