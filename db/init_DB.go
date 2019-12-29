package db

import (
	"errors"
	"log"
)

func InitTable() {
	//用户数据库
	err := execSQL(`CREATE TABLE IF NOT EXISTS gp.user(
	id SERIAL NOT NULL,
	username varchar(32) NOT NULL,
	password varchar(32) NOT NULL,
	nickname varchar(32) NOT NULL,
	role varchar(32) NOT NULL,
	PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panicln("init table gp.user failed " + err.Error())
	} else {
		log.Println("table gp.user has been created")
	}
}
func execSQL(sql string) error {
	//检测sql语句长度
	if len(sql) <= 0 {
		return errors.New("execSQL sql empty")
	}
	//sql语句准备
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	//事务执行
	_, err = stmt.Exec()
	//关闭事务
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
