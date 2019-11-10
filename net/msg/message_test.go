package msg

import "testing"

func TestNsq2(t *testing.T) {
	nsqPing()
	wait()
	publicTopicAndMessage()
}
