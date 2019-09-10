package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/entity"
)

type FeedbackDao struct {
	db *sql.DB
}

func NewFeedbackDao() *FeedbackDao {
	dao := &FeedbackDao{
		db: config.GetDB(),
	}
	return dao
}

func (f *FeedbackDao) Create(fn entity.FeedbackEntity) int {
	sqlStr := "insert into t_feedback(id,user_id,content,create_time) values (?,?,?,?)"
	result, e := f.db.Exec(sqlStr, fn.Id, fn.UserId, fn.Content, fn.CreateTime)
	if e != nil {
		return 0
	}
	i, _ := result.RowsAffected()
	return int(i)
}
