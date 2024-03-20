// Author: Moisés Adame-Aguilar
// Creation Date: March 18th, 2024

package lib

import(
	"bytes"
)

// InputBuffer represents a buffer for storing input data.
type InputBuffer struct {
    buffer bytes.Buffer
}

// NewInputBuffer creates a new InputBuffer instance.
func NewInputBuffer() *InputBuffer {
    return &InputBuffer{}
}

// ReadFrom reads input from the given reader and stores it in the buffer.
func (ib *InputBuffer) ReadFromString(inputString string) (int64, error) {
    ib.buffer.Reset()
    reader := bytes.NewBufferString(inputString)
    return ib.buffer.ReadFrom(reader)
}

// GetContents returns the contents of the buffer as a byte slice.
func (ib *InputBuffer) GetContents() []byte {
    return ib.buffer.Bytes()
}

// ToString returns the contents of the buffer as a string.
func (ib *InputBuffer) ToString() string {
    return string(ib.GetContents())
}
