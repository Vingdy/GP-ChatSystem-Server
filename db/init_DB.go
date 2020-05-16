package db

import (
	"errors"
	"log"
)

func InitTable() {
	//用户数据库
	err := execSQL(`CREATE TABLE IF NOT EXISTS gp.user(
	id SERIAL NOT NULL,
	username varchar(30) NOT NULL,
	password varchar(30) NOT NULL,
	nickname varchar(30) NOT NULL,
	role varchar(30) NOT NULL,
	phone varchar(30),
	label varchar(30),
	fonttype varchar(30),
	fontcolor varchar(30),
	isban varchar(5) NOT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY (username)
	);`)
	if err != nil {
		log.Panicln("init table gp.user failed " + err.Error())
	} else {
		log.Println("table gp.user has been created")
	}

	err = execSQL(`CREATE TABLE IF NOT EXISTS gp.friend(
	id SERIAL NOT NULL,
	username1 varchar(30),
	nickname1 varchar(30),
	username2 varchar(30),
	nickname2 varchar(30),
	ischeck varchar(5),
	PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panicln("init table gp.friend failed " + err.Error())
	} else {
		log.Println("table gp.friend has been created")
	}

	err = execSQL(`CREATE TABLE IF NOT EXISTS gp.room(
	id SERIAL NOT NULL,
	roomname varchar(30) NOT NULL,
	isban varchar(5) NOT NULL,
	PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panicln("init table gp.room failed " + err.Error())
	} else {
		log.Println("table gp.room has been created")
	}

	err = execSQL(`CREATE TABLE IF NOT EXISTS gp.history(
	id SERIAL NOT NULL,
	username varchar(30) NOT NULL,
	chat varchar(200) NOT NULL,
	fonttype varchar(30) NOT NULL,
	fontcolor varchar(30) NOT NULL,
	time varchar(30) NOT NULL,
	PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panicln("init table gp.history failed " + err.Error())
	} else {
		log.Println("table gp.history has been created")
	}

	err = execSQL(`CREATE TABLE IF NOT EXISTS gp.comment(
	id SERIAL NOT NULL,
	username varchar(30) NOT NULL,
	comment varchar(200) NOT NULL,
	fromusername varchar(30) NOT NULL,
	fromnickname varchar(30) NOT NULL,
	time varchar(30) NOT NULL,
	PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panicln("init table gp.comment failed " + err.Error())
	} else {
		log.Println("table gp.comment has been created")
	}

	err = execSQL(`INSERT INTO gp.room(roomname, isban) SELECT '公共房间', '0' FROM dual WHERE NOT EXISTS(SELECT * FROM gp.room WHERE roomname = '公共房间');`)
	if err != nil {
		log.Panicln("init table room data failed " + err.Error())
	} else {
		log.Println("table room data has been created")
	}

	err = execSQL(`INSERT INTO gp.user(username, password, nickname, role, phone, label, fonttype, fontcolor, isban) SELECT 'admin','admin','admin','admin','', '','宋体','#000000','0' FROM dual WHERE NOT EXISTS(SELECT * FROM gp.user WHERE username = 'admin');`)
	if err != nil {
		log.Panicln("init table user data failed " + err.Error())
	} else {
		log.Println("table user data has been created")
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
