package iomanager

type IoManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
