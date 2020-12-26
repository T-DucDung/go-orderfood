package models

import "time"

type Comment struct {
	IDStore string `json:"id_store"`
	IDUser  string `json:"id_user"`
	Comment string `json:"comment"`

	CreateDate int64 `json:"create_date" `
	UpdateDate int64 `json:"update_date" `
	ID         int   `json:"id"`
}

func (this *Comment) CreateCom() error {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	data, err := db.Prepare("INSERT INTO comment (created_date,updated_date,comment,user_id,store_id) values(?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = data.Exec(t, t, this.Comment, this.IDUser, this.IDStore)
	if err != nil {
		return err
	}
	return nil
}
