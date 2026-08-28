package datasource

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DataSource 定义数据源接口。
//
// 资源所有权规则：如果数据源自己打开的资源（如 FilePathSource 打开的文件），
// ByteData 在读取后负责关闭；如果资源由调用方提供（如 ReaderSource 的流），
// 调用方保留所有权，ByteData 不会关闭它。
type DataSource interface {
	ByteData() ([]byte, error)
}

// DataSink 定义数据目标接口
type DataSink interface {
	Write() (io.WriteCloser, error)
}

// --- FilePathSource ---
type FilePathSource string

func (p FilePathSource) Open() (io.ReadCloser, error) {
	return os.Open(string(p))
}

func (p FilePathSource) ByteData() ([]byte, error) {
	reader, err := p.Open()
	if err != nil {
		return nil, err
	}
	// 本数据源自己打开的文件，读取后负责关闭。
	defer reader.Close()
	return io.ReadAll(reader)
}

// --- BytesSource ---
type BytesSource []byte

func (b BytesSource) Open() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(b)), nil
}

func (b BytesSource) ByteData() ([]byte, error) {
	return b, nil
}

// --- ReaderSource ---
type ReaderSource struct {
	reader io.ReadCloser
}

// NewReaderSource creates a ReaderSource from an io.ReadCloser.
func NewReaderSource(r io.ReadCloser) *ReaderSource {
	return &ReaderSource{reader: r}
}

func (r ReaderSource) Open() (io.ReadCloser, error) {
	return r.reader, nil
}

func (r ReaderSource) ByteData() ([]byte, error) {
	// 流由调用方提供，所有权在调用方，这里只读取、不关闭。
	return io.ReadAll(r.reader)
}

// --- FilePathSink ---
type FilePathSink string

func (p FilePathSink) Write() (io.WriteCloser, error) {
	return os.Create(string(p))
}

// --- BytesSink ---
type BytesSink struct {
	buf bytes.Buffer
}

func (b *BytesSink) Write() (io.WriteCloser, error) {
	// ✅ 移除 Reset()，避免数据竞争。
	// 如果需要复用 BytesSink，应该在外部重新 New 一个
	return nopWriteCloser{&b.buf}, nil
}

func (b *BytesSink) Bytes() []byte {
	return b.buf.Bytes()
}

// nopWriteCloser 包装 io.Writer
type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error { return nil }

// UrlPathSink is an implementation of a data sink based on HTTP/HTTPS URLs.
// It implements the DataSink interface and is used to upload data to a remote server
// via an HTTP PUT request. This is useful for pushing converted files directly to
// your own backend, or to pre-signed cloud storage URLs (e.g., AWS S3, Aliyun OSS).
type UrlPathSink string

// Write initiates an HTTP PUT request to the specified URL and returns a WriteCloser.
// The caller should write data to this WriteCloser and then call Close().
// Upon Close(), the buffered data is flushed to the remote server.
func (u UrlPathSink) Write() (io.WriteCloser, error) {
	pr, pw := io.Pipe()

	// 启动一个后台协程，将写入 Pipe 的数据流式发送到 HTTP 请求中
	go func() {
		defer pr.Close()

		req, err := http.NewRequest(http.MethodPut, string(u), pr)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		// 设置通用的流式上传 Header
		req.Header.Set("Content-Type", "application/octet-stream")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		defer resp.Body.Close()

		// 检查 HTTP 响应状态码，非 2xx 视为上传失败
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			pw.CloseWithError(fmt.Errorf("upload to URL failed with status %d", resp.StatusCode))
			return
		}
	}()

	return pw, nil
}
