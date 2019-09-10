package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/entity"
	"log"
)

//noinspection GoNameStartsWithPackageName
type FishDao struct {
	db *sql.DB
}

func NewFishDao() *FishDao {
	dao := &FishDao{
		db: config.GetDB(),
	}
	return dao
}

func (d *FishDao) Create(f *entity.FishEntity) {
	sqlStr := "insert into t_fish(`id`,`title`,`url`,`order`,`type`,`view`,`create_time`) values (?,?,?,?,?,?,?)"
	result, e := d.db.Exec(sqlStr, f.Id, f.Title, f.Url, f.Order, f.Type, f.View, f.CreateTime)
	if e != nil {
		log.Println(e)
	}
	if i, _ := result.RowsAffected(); i != 1 {
		log.Println(f.Title, " create error...")
	}
}

func (d *FishDao) All(fishType string) (*[]entity.FishEntity, error) {
	sqlStr := "select `id`,`title`,`url`,`order`,`view`,`create_time` from t_fish where `type`=? order by `order` asc"
	rows, e := d.db.Query(sqlStr, fishType)
	defer rows.Close()
	if e != nil {
		log.Println(e)
		return nil, e
	}
	fishAll := make([]entity.FishEntity, 0)
	var fishEntity entity.FishEntity
	for rows.Next() {
		_ = rows.Scan(&fishEntity.Id, &fishEntity.Title, &fishEntity.Url, &fishEntity.Order, &fishEntity.View, &fishEntity.CreateTime)
		fishAll = append(fishAll, fishEntity)
	}
	return &fishAll, nil
}

func (d *FishDao) DeleteByType(fishType string) {
	sqlStr := "delete from t_fish where `type`=?"
	_, e := d.db.Exec(sqlStr, fishType)
	if e != nil {
		log.Println(e)
	}
}
