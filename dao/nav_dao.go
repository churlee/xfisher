package dao

import (
	"database/sql"
	"lilith/common"
	"lilith/config"
	"lilith/model/entity"
	"log"
	"strings"
)

//noinspection GoNameStartsWithPackageName
type NavDao struct {
	db *sql.DB
}

func NewNavDao() *NavDao {
	dao := &NavDao{
		db: config.GetDB(),
	}
	return dao
}

func (s *NavDao) Create(nav entity.NavEntity) int {
	sqlStr := "insert into t_nav(`id`, `title`, `type`, `desc`, `icon`, `like`, `view`, `url`) values (?,?,?,?,?,?,?,?)"
	result, e := s.db.Exec(sqlStr, nav.Id, nav.Title, nav.Type, nav.Desc, nav.Icon, nav.Like, nav.View, nav.Url)
	if e != nil {
		log.Println(e)
		return 0
	}
	i, _ := result.RowsAffected()
	return int(i)
}

func (s *NavDao) AddNavView(id string) {
	sqlStr := "update t_nav as t set t.view=t.view+1 where t.id=?"
	result, e := s.db.Exec(sqlStr, id)
	if e != nil {
		log.Println(e)
		return
	}
	if i, _ := result.RowsAffected(); i != 1 {
		return
	}
}

func (s *NavDao) AddLike(id string) {
	sqlStr := "update t_nav as t set t.like=t.like+1 where t.id=?"
	result, e := s.db.Exec(sqlStr, id)
	if e != nil {
		log.Println(e)
		return
	}
	if i, _ := result.RowsAffected(); i != 1 {
		return
	}
}

func (s *NavDao) AddCollection(id string) {
	sqlStr := "update t_nav as t set t.collection=t.collection+1 where t.id=?"
	result, e := s.db.Exec(sqlStr, id)
	if e != nil {
		log.Println(e)
		return
	}
	if i, _ := result.RowsAffected(); i != 1 {
		return
	}
}

func (s *NavDao) NavListByType(navType string) ([]entity.NavEntity, error) {
	navAll := make([]entity.NavEntity, 0)
	sqlStr := "select t.id,t.title,t.desc,t.icon,t.like,t.view,t.url,t.collection from t_nav as t where t.type=? order by t.view desc"
	rows, e := s.db.Query(sqlStr, navType)
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

func (s *NavDao) NavHot() ([]entity.NavEntity, error) {
	navAll := make([]entity.NavEntity, 0)
	sqlStr := "select t.id,t.title,t.desc,t.icon,t.like,t.view,t.url from t_nav as t order by t.view desc limit 1,10"
	rows, e := s.db.Query(sqlStr)
	defer rows.Close()
	if e != nil {
		return navAll, e
	}
	var navEntity entity.NavEntity
	for rows.Next() {
		_ = rows.Scan(&navEntity.Id, &navEntity.Title, &navEntity.Desc, &navEntity.Icon, &navEntity.Like, &navEntity.View, &navEntity.Url)
		navAll = append(navAll, navEntity)
	}
	return navAll, nil
}

func (s *NavDao) FindByTitle(title string) ([]entity.NavEntity, error) {
	navAll := make([]entity.NavEntity, 0)
	sqlStr := "select t.id,t.title,t.desc,t.icon,t.like,t.view,t.url,t.collection from t_nav as t where t.title like ? order by t.view desc"
	if strings.ContainsAny(title, common.SqlStr) {
		return navAll, nil
	}
	rows, e := s.db.Query(sqlStr, "%"+title+"%")
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
