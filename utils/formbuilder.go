package utils

import (
	"io"
	"mime/multipart"
	"os"
)

type FormBuilder interface {
	CreateFormFile(fieldName string, file *os.File) error
	WriteField(fieldName, value string) error
	Close() error
	FormDataContentType() string
}

type Form struct {
	writer *multipart.Writer
}

func NewFormBuilder(body io.Writer) *Form {
	return &Form{
		writer: multipart.NewWriter(body),
	}
}

func (f *Form) CreateFormFile(fieldName string, file *os.File) error {
	fieldWriter, err := f.writer.CreateFormFile(fieldName, file.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(fieldWriter, file)
	if err != nil {
		return err
	}
	return nil
}

func (f *Form) WriteField(fieldName, value string) error {
	return f.writer.WriteField(fieldName, value)
}

func (f *Form) Close() error {
	return f.writer.Close()
}

func (f *Form) FormDataContentType() string {
	return f.writer.FormDataContentType()
}
