

package fileutil

import "testing"

func Test_CopyFile_01(t *testing.T) {
	if _, err := CopyFile("d:\\a.conf", "d:\\b.conf", true); err != nil {
		t.Error("failed", err)
	}
}

func Test_Md5_01(t *testing.T) {
	md5, err := GetFileMd5("d:\\time.exe")
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("md5: " + md5)
}

func Test_SetExecutable_01(t *testing.T) {
	md5, err := GetFileMd5("d:\\time.exe")
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("md5: " + md5)
}
