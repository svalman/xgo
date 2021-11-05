package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	NoType = ErrorType(iota)
	XmlInvalidToken
	XmlUnknownTag
	XmlCannotDecodeTag
	XmlWrongToken
	BadRequest
	NotFound
	Config_NoDefaultDataSource
)

type (
	ErrorType uint

	customError struct {
		errorType     ErrorType
		originalError error
		contextInfo   map[string]string
	}
)

// Error извлекает сообщение из customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New создает новую customError
func (t ErrorType) New(msg string) error {
	return customError{errorType: t, originalError: errors.New(msg)}
}

// New создает новую customError с форматированным сообщением
func (t ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return customError{errorType: t, originalError: err}
}

// Wrapf создает новую ошибку которую можно обернуть и
// добавляет к ней форматированную строку
func (t ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return customError{errorType: t, originalError: newErr}
}

// Wrap создает новую ошибку которую можно обернуть
func (t ErrorType) Wrap(err error, msg string) error {
	return t.Wrapf(err, msg)
}

// AddErrorContext добавляет к ошибке значение в контекст
func AddErrorContext(err error, field, message string) error {

	if customErr, ok := err.(customError); ok {
		customErr.contextInfo[field] = message
		return customError{
			errorType:     customErr.errorType,
			originalError: customErr.originalError,
			contextInfo:   customErr.contextInfo}
	}

	return customError{errorType: NoType, originalError: err, contextInfo: make(map[string]string)}
}

// GetErrorContext отдает контекст ошибки
func GetErrorContext(err error) map[string]string {
	emptyContext := make(map[string]string)
	if customErr, ok := err.(customError); ok || len(customErr.contextInfo) > 0 {
		return customErr.contextInfo
	}
	return emptyContext
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}

// Создает новую нетипизированную ошибку
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf Создает новую нетипизированную ошибку с текстом - форматированной строкой
func Newf(msg string, args ...interface{}) error {
	return &customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap Оборачивает ошибку строкой с текстом
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Wrapf Оборачивает ошибку в форматированную строку
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			contextInfo:   customErr.contextInfo,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// Cause Выдает изначальную причину ошибки
func Cause(err error) error {
	return errors.Cause(err)
}
