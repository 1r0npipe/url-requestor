package server_errors

import "errors"

var (
	ErrReadJson        = errors.New("can't read the JSON input")
	ErrFileRead        = errors.New("can't read config file")
	ErrDecodeYAML      = errors.New("can't decode yaml file")
	ErrStartHTTPServer = errors.New("can't start HTTP server, due to error")
)
