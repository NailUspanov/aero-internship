package files

type FileDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Ext        string `json:"ext"`
	Bytes      []byte `json:"bytes"`
	DateCreate int64  `json:"date_create"`
	UserId     string `json:"user_id"`
	FileSize   int64  `json:"file_size"`
}
