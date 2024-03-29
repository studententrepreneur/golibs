package domain

import "time"

type AbstractLog struct {
	CreatedBy        string      `firestore:"createdBy" json:"createdBy"`
	CreatedDate      interface{} `firestore:"createdDate" json:"createdDate"`
	LastModifiedBy   string      `firestore:"lastModifiedBy" json:"lastModifiedBy"`
	LastModifiedDate interface{} `firestore:"lastModifiedDate" json:"lastModifiedDate"`
}

func CreateLogs(createdBy string) AbstractLog {
	log := AbstractLog{}
	log.CreatedBy = createdBy
	log.CreatedDate = time.Now()

	log.LastModifiedBy = createdBy
	log.LastModifiedDate = time.Now()
	return log
}

func UpdateLogs(createdBy string, log AbstractLog) AbstractLog {
	log.LastModifiedBy = createdBy
	log.LastModifiedDate = time.Now()
	return log
}
