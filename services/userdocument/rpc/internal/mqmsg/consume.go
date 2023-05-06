package mqmsg

type Consumer interface {
	ReadDocAnalysisMessage(data []byte) error
}
