package meta

import "sort"

//文件元信息
type FileMeta struct {
	FileSha1 string //文件hash
	FileName string //文件名称
	FileSize int64  //文件大小
	Location string //地址
	UploadAt string //更新时间
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UploadFileMeta 新增/更新文件元信息
func UploadFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// 通过文件hash查找文件
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// 获取批量文件信息列表
func GetListFileMetas(count int) []FileMeta {
	var fMetaArray []FileMeta
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	if count > len(fMetaArray) {
		return fMetaArray
	}
	return fMetaArray[0:count]
}

// 删除文件元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
