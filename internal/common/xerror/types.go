package xerror

type SearchErr interface {
	GetCode() codeType
	GetMsg() string
	GetExt() string
	error
}

func NewErrWithExt(err SearchErr, ext string) SearchErr {
	return &baseErr{
		code: err.GetCode(),
		msg:  err.GetMsg(),
		ext:  ext,
	}
}

func NewSearchErr(code uint32, msg string, ext string) SearchErr {
	return &baseErr{
		code: codeType(code),
		msg:  msg,
		ext:  ext,
	}
}
