package meta

import "time"

const baseFormat = "2006-01-02 15:04:05"

type ByUploadTime []FileMeta

func (b ByUploadTime) Len() int {
	return len(b)
}

func (b ByUploadTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByUploadTime) Less(i, j int) bool {
	//通过文件上传的时间，对文件进行排序
	iTime, _ := time.Parse(baseFormat, b[i].UploadAt)
	jTime, _ := time.Parse(baseFormat, b[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
