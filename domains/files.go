package domain

type Files struct {
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	FileName string      `firestore:"file_name" json:"file_name"`
	UserId   string      `firestore:"user_id" json:"user_id"`
	FileType string      `firestore:"file_type" json:"file_type"`
	Id       string      `firestore:"id" json:"id"`
	URI      string      `firestore:"uri" json:"uri"`
}
