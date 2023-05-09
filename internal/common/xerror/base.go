package xerror

import "fmt"

type codeType uint32

var baseErrMap = make(map[codeType]struct{})

type baseErr struct {
	code codeType
	msg  string
	ext  string
}

func (b *baseErr) GetExt() string {
	return b.ext
}

func (b *baseErr) GetMsg() string {
	return b.msg
}

func (b *baseErr) GetCode() codeType {
	return b.code
}

func (b *baseErr) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, ext: %s", b.GetCode(), b.GetMsg(), b.GetExt())
}

func newErr(code codeType, msg string) SearchErr {
	if _, ok := baseErrMap[code]; ok {
		panic(fmt.Sprintf("error code %d is duplicated.\n", code))
	}

	baseErrMap[code] = struct{}{}
	return &baseErr{
		code: code,
		msg:  msg,
	}
}
