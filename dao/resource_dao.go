package dao

import (
	"database/sql"
	"lilith/common"
	"lilith/config"
	"lilith/model/dto"
	"lilith/model/entity"
	"strings"
)

type ResourceDao struct {
	db *sql.DB
}

func NewResourceDao() *ResourceDao {
	dao := &ResourceDao{
		db: config.GetDB(),
	}
	return dao
}

func (r *ResourceDao) Create(entity entity.ResourceEntity) (int, error) {
	sqlStr := "insert into t_resource(id,title,content,type,link_url,link_pwd,view,user_id,create_time) values (?,?,?,?,?,?,?,?,?)"
	result, e := r.db.Exec(sqlStr, entity.Id, entity.Title, entity.Content, entity.Type, entity.LinkUrl, entity.LinkPwd, entity.View, entity.UserId, entity.CreateTime)
	if e != nil {
		return 0, nil
	}
	i, e := result.RowsAffected()
	if e != nil {
		return 0, nil
	}
	return int(i), nil
}

func (r *ResourceDao) ResourceCount(resourceType, title string) int {
	sqlStr := "SELECT COUNT(1) FROM t_resource AS t WHERE t.type=if(?,t.type,?) AND t.title LIKE if(?,t.title,?)"
	bType := resourceType == ""
	bTitle := title == ""
	i := 0
	if !bTitle {
		if strings.ContainsAny(title, common.SqlStr) {
			return 0
		}
	}
	row := r.db.QueryRow(sqlStr, bType, resourceType, bTitle, "%"+title+"%")
	_ = row.Scan(&i)
	return i
}

func (r *ResourceDao) ResourceList(resourceType, title string, start, size int) ([]dto.ResourceListDto, error) {
	sqlStr := "SELECT t.id,t.title,t.type,u.nick,t.create_time FROM t_resource AS t LEFT JOIN t_user as u ON t.user_id=u.id WHERE t.type=if(?,t.type,?) AND t.title LIKE if(?,t.title,?) ORDER BY t.create_time DESC LIMIT ?,?"
	bType := resourceType == ""
	bTitle := title == ""
	resourceList := make([]dto.ResourceListDto, 0)
	if !bTitle {
		if strings.ContainsAny(title, common.SqlStr) {
			return resourceList, nil
		}
	}
	rows, e := r.db.Query(sqlStr, bType, resourceType, bTitle, "%"+title+"%", start, size)
	defer rows.Close()
	if e != nil {
		return nil, e
	}
	var resource dto.ResourceListDto
	for rows.Next() {
		_ = rows.Scan(&resource.Id, &resource.Title, &resource.Type, &resource.Nick, &resource.CreateTime)
		resourceList = append(resourceList, resource)
	}
	return resourceList, nil
}

func (r *ResourceDao) ResourceInfo(id string) dto.ResourceInfoDto {
	sqlStr := "SELECT t.id,t.title,t.type,t.content,t.`view`,t.link_url,t.link_pwd,u.nick,u.photo,u.position,t.create_time FROM t_resource AS t LEFT JOIN t_user as u ON t.user_id=u.id WHERE t.id=?"
	row := r.db.QueryRow(sqlStr, id)
	var d dto.ResourceInfoDto
	_ = row.Scan(&d.Id, &d.Title, &d.Type, &d.Content, &d.View, &d.LinkUrl, &d.LinkPwd, &d.Nick, &d.Photo, &d.Position, &d.CreateTime)
	return d
}

func (r *ResourceDao) AddView(id string) int {
	sqlStr := "update t_resource as t set t.view=t.view+1 where t.id=?"
	result, e := r.db.Exec(sqlStr, id)
	if e != nil {
		return 0
	}
	i, _ := result.RowsAffected()
	return int(i)
}
