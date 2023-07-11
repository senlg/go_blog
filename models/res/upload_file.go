package res

// 上传文件返回信息
type FileUploadInfo struct {
	FileUrl   string `json:"file_url"`
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}
