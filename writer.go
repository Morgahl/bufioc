package bufioc

import (
	"bufio"
	"io"
)

// Writer implements buffering for an io.WriteCloser object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.WriteCloser.
type Writer struct {
	wc io.WriteCloser
	w  *bufio.Writer
}

// NewWriter returns a new Writer whose buffer has the default size.
func NewWriter(wc io.WriteCloser) *Writer {
	return &Writer{
		wc: wc,
		w:  bufio.NewWriter(wc),
	}
}

// NewWriterSize returns a new Writer whose buffer has at least the specified size.
// If the argument io.WriteCloser is already a io.WriteCloser with large enough size,
// it uses the underlying io.WriteCloser.
func NewWriterSize(wc io.WriteCloser, size int) *Writer {
	return &Writer{
		wc: wc,
		w:  bufio.NewWriterSize(wc, size),
	}
}

// Available returns how many bytes are unused in the buffer.
func (w *Writer) Available() int {
	return w.w.Available()
}

// Buffered returns the number of bytes that have been written into the current buffer.
func (w *Writer) Buffered() int {
	return w.w.Buffered()
}

// Close writes any buffered data to the underlying io.WriteCloser then closes the Writer,
// rendering it unusable for I/O. It returns an error, if any.
func (w *Writer) Close() error {
	if err := w.w.Flush(); err != nil {
		w.wc.Close()
		return err
	}

	return w.wc.Close()
}

// Flush writes any buffered data to the underlying io.WriteCloser.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

// ReadFrom implements io.ReaderFrom.
func (w *Writer) ReadFrom(r io.Reader) (int64, error) {
	return w.w.ReadFrom(r)
}

// Reset discards any unflushed buffered data, clears any error, and
// resets w to write its output to wc.
func (w *Writer) Reset(wc io.WriteCloser) error {
	if err := w.wc.Close(); err != nil {
		return err
	}

	w.wc = wc
	w.w.Reset(wc)

	return nil
}

// Size returns the size of the underlying buffer in bytes.
func (w *Writer) Size() int {
	return w.w.Size()
}

// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (w *Writer) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

// WriteByte writes a single byte.
func (w *Writer) WriteByte(b byte) error {
	return w.w.WriteByte(b)
}

// WriteRune writes a single Unicode code point, returning
// the number of bytes written and any error.
func (w *Writer) WriteRune(r rune) (int, error) {
	return w.w.WriteRune(r)
}

// WriteString writes a string.
// It returns the number of bytes written.
// If the count is less than len(s), it also returns an error explaining
// why the write is short.
func (w *Writer) WriteString(s string) (int, error) {
	return w.w.WriteString(s)
}
