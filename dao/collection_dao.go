package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/entity"
)

type CollectionDao struct {
	db *sql.DB
}

func NewCollectionDao() *CollectionDao {
	dao := &CollectionDao{
		db: config.GetDB(),
	}
	return dao
}

func (c *CollectionDao) Create(cEntity entity.CollectionEntity) (int, error) {
	sqlStr := "insert into t_collection(id,user_id,nav_id) values (?,?,?)"
	result, e := c.db.Exec(sqlStr, cEntity.Id, cEntity.UserId, cEntity.NavId)
	if e != nil {
		return 0, e
	}
	i, _ := result.RowsAffected()
	return int(i), nil
}

func (c *CollectionDao) Exist(userId, navId string) int {
	sqlStr := "SELECT COUNT(1) FROM t_collection WHERE user_id=? and nav_id=?"
	row := c.db.QueryRow(sqlStr, userId, navId)
	total := 0
	_ = row.Scan(&total)
	return total
}

func (c *CollectionDao) All(userId string) ([]entity.NavEntity, error) {
	navAll := make([]entity.NavEntity, 0)
	sqlStr := "select t.id,t.title,t.desc,t.icon,t.like,t.view,t.url,t.collection from t_nav as t left join t_collection as c on t.id=c.nav_id where c.user_id=? order by t.view desc"
	rows, e := c.db.Query(sqlStr, userId)
	defer rows.Close()
	if e != nil {
		return navAll, e
	}
	var navEntity entity.NavEntity
	for rows.Next() {
		_ = rows.Scan(&navEntity.Id, &navEntity.Title, &navEntity.Desc, &navEntity.Icon, &navEntity.Like, &navEntity.View, &navEntity.Url, &navEntity.Collection)
		navAll = append(navAll, navEntity)
	}
	return navAll, nil
}
