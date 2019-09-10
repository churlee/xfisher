package dao

import (
	"database/sql"
	"lilith/common"
	"lilith/config"
	"lilith/model/dto"
	"lilith/model/entity"
	"strings"
)

type CommunicationDao struct {
	db *sql.DB
}

func NewCommunicationDao() *CommunicationDao {
	dao := &CommunicationDao{
		db: config.GetDB(),
	}
	return dao
}

func (c *CommunicationDao) ListByNode(node, title string, start, size int) ([]dto.ListCommunicationDto, error) {
	sqlStr := "select c.id,c.title,c.view,c.create_time,u.nick,u.photo,u.position from t_communication as c left join t_user as u on c.user_id=u.id where c.node=if(?='',c.node,?) and c.title like if(?,c.title,?) order by c.create_time desc limit ?,?"
	bTitle := title == ""
	list := make([]dto.ListCommunicationDto, 0)
	if bTitle {
		if strings.ContainsAny(title, common.SqlStr) {
			return list, nil
		}
	}
	rows, e := c.db.Query(sqlStr, node, node, bTitle, "%"+title+"%", start, size)
	defer rows.Close()
	if e != nil {
		return list, e
	}
	var d dto.ListCommunicationDto
	for rows.Next() {
		_ = rows.Scan(&d.Id, &d.Title, &d.View, &d.CreateTime, &d.Nick, &d.Photo, &d.Position)
		list = append(list, d)
	}
	return list, nil
}

func (c *CommunicationDao) CountByNode(node, title string) int {
	sqlStr := "select count(1) from t_communication as c where c.node=if(?='',c.node,?) and c.title like if(?,c.title,?)"
	bTitle := title == ""
	var i = 0
	if bTitle {
		if strings.ContainsAny(title, common.SqlStr) {
			return i
		}
	}
	row := c.db.QueryRow(sqlStr, node, node, bTitle, "%"+title+"%")
	_ = row.Scan(&i)
	return i
}

func (c *CommunicationDao) Create(e entity.CommunicationEntity) int {
	sqlStr := "insert into t_communication(id,title,content,node,user_id,view,create_time) values (?,?,?,?,?,?,?)"
	result, err := c.db.Exec(sqlStr, e.Id, e.Title, e.Content, e.Node, e.UserId, e.View, e.CreateTime)
	if err != nil {
		return 0
	}
	i, _ := result.RowsAffected()
	return int(i)
}

func (c *CommunicationDao) FindById(id string) dto.CommunicationInfoDto {
	sqlStr := "SELECT t.id,t.title,t.content,t.node,t.`view`,u.nick,u.photo,u.position,t.create_time FROM t_communication AS t LEFT JOIN t_user as u ON t.user_id=u.id WHERE t.id=?"
	row := c.db.QueryRow(sqlStr, id)
	var d dto.CommunicationInfoDto
	_ = row.Scan(&d.Id, &d.Title, &d.Content, &d.Node, &d.View, &d.Nick, &d.Photo, &d.Position, &d.CreateTime)
	return d
}

func (r *CommunicationDao) AddView(id string) int {
	sqlStr := "update t_communication as t set t.view=t.view+1 where t.id=?"
	result, e := r.db.Exec(sqlStr, id)
	if e != nil {
		return 0
	}
	i, _ := result.RowsAffected()
	return int(i)
}
