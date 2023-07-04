package res

// 上传文件返回信息
type FileUploadInfo struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}
