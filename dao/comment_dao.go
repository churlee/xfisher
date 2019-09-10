package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/dto"
	"lilith/model/entity"
)

type CommentDao struct {
	db *sql.DB
}

func NewCommentDao() *CommentDao {
	dao := &CommentDao{
		db: config.GetDB(),
	}
	return dao
}

func (c *CommentDao) CreateCommnet(entity entity.CommentEntity) (int, error) {
	sqlStr := "insert into t_comment(id,resource_id,user_id,content,create_time) values (?,?,?,?,?)"
	result, e := c.db.Exec(sqlStr, entity.Id, entity.ResourceId, entity.UserId, entity.Content, entity.CreateTime)
	if e != nil {
		return 0, e
	}
	i, _ := result.RowsAffected()
	return int(i), nil
}

func (c *CommentDao) CommentList(resourceId string, start, size int) ([]dto.CommentListDto, error) {
	sqlStr := "SELECT c.id,c.content,c.create_time,u.nick,u.photo,u.position FROM t_comment AS c LEFT JOIN t_user AS u ON c.user_id=u.id WHERE c.resource_id=? ORDER BY c.create_time ASC LIMIT ?,?"
	rows, e := c.db.Query(sqlStr, resourceId, start, size)
	defer rows.Close()
	list := make([]dto.CommentListDto, 0)
	if e != nil {
		return list, e
	}
	var commentListDto dto.CommentListDto
	for rows.Next() {
		_ = rows.Scan(&commentListDto.Id, &commentListDto.Content, &commentListDto.CreateTime, &commentListDto.Nick, &commentListDto.Photo, &commentListDto.Position)
		list = append(list, commentListDto)
	}
	return list, nil
}

func (c *CommentDao) CountCommentByResourceId(resourceId string) int {
	sqlStr := "SELECT COUNT(1) FROM t_comment WHERE resource_id=?"
	row := c.db.QueryRow(sqlStr, resourceId)
	total := 0
	_ = row.Scan(&total)
	return total
}

func (c *CommentDao) IsCommentByUserId(resourceId, userId string) int {
	sqlStr := "SELECT COUNT(1) FROM t_comment WHERE resource_id=? and user_id=?"
	row := c.db.QueryRow(sqlStr, resourceId, userId)
	total := 0
	_ = row.Scan(&total)
	return total
}
