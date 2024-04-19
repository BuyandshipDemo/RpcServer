package mysql

import (
	"RpcServer/model"
	"log"
)

func Query(Id int64) (string, error) {
	db := DB.Model(model.Helloworld{})
	var res *model.Helloworld
	db = db.Where("id = ?", Id).Debug().Find(&res)
	if db.Error != nil {
		return "", db.Error
	}
	log.Println(res)
	return res.Msg, nil
}