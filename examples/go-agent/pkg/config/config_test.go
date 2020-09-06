package config

import (
	"testing"
)

func Test_DetectSlaveVersion_01(t *testing.T) {
	t.Log("DetectSlaveVersion: ", DetectWorkerVersion())
}
