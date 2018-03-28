// Package bufioc implements buffered I/O. It wraps an io.ReadCloser or
// io.WriteCloser object, creating another object (Reader or Writer) that
// also implements the interface but provides buffering and some help for
// textual I/O.
//
// Basically a convinience wrapper of Package bufio
package bufioc

import (
	"bufio"
	"io"
)

// Reader implements buffering for an io.ReadCloser object.
type Reader struct {
	r  *bufio.Reader
	rc io.ReadCloser
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader(rc io.ReadCloser) *Reader {
	return &Reader{
		rc: rc,
		r:  bufio.NewReader(rc),
	}
}

// NewReaderSize returns a new Reader whose buffer has at least the specified size.
// If the argument io.ReadCloser is already a io.ReadCloser with large enough size,
// it uses the underlying io.ReadCloser.
func NewReaderSize(rc io.ReadCloser, size int) *Reader {
	return &Reader{
		rc: rc,
		r:  bufio.NewReaderSize(rc, size),
	}
}

// Buffered returns the number of bytes that can be read from the current buffer.
func (r *Reader) Buffered() int {
	return r.r.Buffered()
}

// Close closes the Reader, rendering it unusable for I/O. It returns an error, if any.
func (r *Reader) Close() error {
	return r.rc.Close()
}

// Discard skips the next n bytes, returning the number of bytes discarded.
//
// If Discard skips fewer than n bytes, it also returns an error.
// If 0 <= n <= r.Buffered(), Discard is guaranteed to succeed without
// reading from the underlying io.ReadCloser.
func (r *Reader) Discard(n int) (int, error) {
	return r.r.Discard(n)
}

// Peek returns the next n bytes without advancing the reader. The bytes stop
// being valid at the next read call. If Peek returns fewer than n bytes, it
// also returns an error explaining why the read is short. The error is
// ErrBufferFull if n is larger than b's buffer size.
func (r *Reader) Peek(n int) ([]byte, error) {
	return r.r.Peek(n)
}

// Read reads data into p.
// It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying Reader,
// hence bytes may be less than len(p).
// At EOF, the count will be zero and err will be io.EOF.
func (r *Reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

// ReadByte reads and returns a single byte.
// If no byte is available, returns an error.
func (r *Reader) ReadByte() (byte, error) {
	return r.r.ReadByte()
}

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadBytes returns err != nil if and only if the returned data does not end in
// delim. For simple uses, a Scanner may be more convenient.
func (r *Reader) ReadBytes(delim byte) ([]byte, error) {
	return r.r.ReadBytes(delim)
}

// ReadLine is a low-level line-reading primitive. Most callers should use
// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
//
// ReadLine tries to return a single line, not including the end-of-line bytes.
// If the line was too long for the buffer then isPrefix is set and the
// beginning of the line is returned. The rest of the line will be returned
// from future calls. isPrefix will be false when returning the last fragment
// of the line. The returned buffer is only valid until the next call to
// ReadLine. ReadLine either returns a non-nil line or it returns an error,
// never both.
//
// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
// No indication or error is given if the input ends without a final line end.
// Calling UnreadByte after ReadLine will always unread the last byte read
// (possibly a character belonging to the line end) even if that byte is not
// part of the line returned by ReadLine.
func (r *Reader) ReadLine() ([]byte, bool, error) {
	return r.r.ReadLine()
}

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
// rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
// and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
func (r *Reader) ReadRune() (rune, int, error) {
	return r.r.ReadRune()
}

// ReadSlice reads until the first occurrence of delim in the input,
// returning a slice pointing at the bytes in the buffer.
// The bytes stop being valid at the next read.
// If ReadSlice encounters an error before finding a delimiter,
// it returns all the data in the buffer and the error itself (often io.EOF).
// ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
// Because the data returned from ReadSlice will be overwritten
// by the next I/O operation, most clients should use
// ReadBytes or ReadString instead.
// ReadSlice returns err != nil if and only if line does not end in delim.
func (r *Reader) ReadSlice(delim byte) ([]byte, error) {
	return r.r.ReadSlice(delim)
}

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end in
// delim. For simple uses, a Scanner may be more convenient.
func (r *Reader) ReadString(delim byte) (string, error) {
	return r.r.ReadString(delim)
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from rc.
func (r *Reader) Reset(rc io.ReadCloser) error {
	if err := r.rc.Close(); err != nil {
		return err
	}

	r.rc = rc
	r.r.Reset(rc)

	return nil
}

// Size returns the size of the underlying buffer in bytes.
func (r *Reader) Size() int {
	return r.r.Size()
}

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
func (r *Reader) UnreadByte() error {
	return r.r.UnreadByte()
}

// UnreadRune unreads the last rune. If the most recent read operation on
// the buffer was not a ReadRune, UnreadRune returns an error.  (In this
// regard it is stricter than UnreadByte, which will unread the last byte
// from any read operation.)
func (r *Reader) UnreadRune() error {
	return r.r.UnreadRune()
}

// WriteTo implements io.WriterTo.
// This may make multiple calls to the Read method of the underlying Reader.
func (r *Reader) WriteTo(w io.Writer) (int64, error) {
	return r.r.WriteTo(w)
}
