package dao

import (
	"database/sql"
	"lilith/common"
	"lilith/config"
	"lilith/model/dto"
	"lilith/model/entity"
	"strings"
)

type BookmarkDao struct {
	db *sql.DB
}

func NewBookmarkDao() *BookmarkDao {
	dao := &BookmarkDao{
		db: config.GetDB(),
	}
	return dao
}

func (b *BookmarkDao) Create(bookmark entity.BookmarkEntity) (int, error) {
	sqlStr := "insert into t_bookmark(id,title,url,user_id,create_time,share,tags) values (?,?,?,?,?,?,?)"
	result, e := b.db.Exec(sqlStr, bookmark.Id, bookmark.Title, bookmark.Url, bookmark.UserId, bookmark.CreateTime, bookmark.Share, bookmark.Tags)
	if e != nil {
		return 0, e
	}
	i, _ := result.RowsAffected()
	return int(i), nil
}

func (b *BookmarkDao) List(title, uid string, start, size int) ([]dto.BookmarkListDto, error) {
	bTitle := title == ""
	bUid := uid == ""
	sqlStr := "select b.id,b.title,b.url,b.create_time,u.nick,b.tags from t_bookmark b left join t_user u on b.user_id=u.id where b.user_id=if(?,b.user_id,?) and b.title like if(?,b.title,?) order by b.create_time desc limit ?,?"
	rows, e := b.db.Query(sqlStr, bUid, uid, bTitle, "%"+title+"%", start, size)
	defer rows.Close()
	list := make([]dto.BookmarkListDto, 0)
	if e != nil {
		return list, e
	}
	for rows.Next() {
		var bDto dto.BookmarkListDto
		_ = rows.Scan(&bDto.Id, &bDto.Title, &bDto.Url, &bDto.CreateTime, &bDto.Nick, &bDto.Tags)
		list = append(list, bDto)
	}
	return list, nil
}

func (b *BookmarkDao) Count(title, uid string) int {
	bTitle := title == ""
	var i = 0
	if bTitle {
		if strings.ContainsAny(title, common.SqlStr) {
			return i
		}
	}
	bUid := uid == ""
	sqlStr := "select count(1) from t_bookmark b left join t_user u on b.user_id=u.id where b.user_id=if(?,b.user_id,?) and b.title like if(?,b.title,?)"
	row := b.db.QueryRow(sqlStr, bUid, uid, bTitle, "%"+title+"%")
	_ = row.Scan(&i)
	return i
}
