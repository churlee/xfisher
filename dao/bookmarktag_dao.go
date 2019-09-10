package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/entity"
)

type BookmarkTagDao struct {
	db *sql.DB
}

func NewBookmarkTagDao() *BookmarkTagDao {
	dao := &BookmarkTagDao{
		db: config.GetDB(),
	}
	return dao
}

func (b *BookmarkTagDao) Create(tag entity.BookmarkTagEntity) (int, error) {
	sqlStr := "insert into t_bookmark_tag(id,bookmark_id,tag,create_time) values (?,?,?,?)"
	result, e := b.db.Exec(sqlStr, tag.Id, tag.BookmarkId, tag.Tag, tag.CreateTime)
	if e != nil {
		return 0, e
	}
	i, _ := result.RowsAffected()
	return int(i), nil
}
