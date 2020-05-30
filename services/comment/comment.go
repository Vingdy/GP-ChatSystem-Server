package comment

import (
	"GP/db"
	"log"
)

type CommentInfo struct {
	UserName     string `json:"username"`
	FromUserName string `json:"fromusername"`
	FromNickName string `json:"fromnickname"`
	Comment      string `json:"comment"`
	Time         string `json:"time"`
}

func NewComment(username, fromusername, fromnickname, comment, time string) (err error) {
	insertSql := "insert into gp.comment(username, fromusername, fromnickname, comment, time) values (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		log.Println("NewComment insertSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, fromusername, fromnickname, comment, time)
	if err != nil {
		log.Println("NewComment exec fail")
		return err
	}
	return nil
}

func GetCommentList(username string) (commentInfo []*CommentInfo, err error) {
	commentInfo = []*CommentInfo{}
	querySql := "select username, fromusername, fromnickname, comment, time from gp.comment where username = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetCommentList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("GetCommentList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var comment CommentInfo
		err := rows.Scan(&comment.UserName, &comment.FromUserName, &comment.FromNickName, &comment.Comment, &comment.Time)
		if err != nil {
			return nil, err
		}
		commentInfo = append(commentInfo, &comment)
	}
	return commentInfo, nil
}
