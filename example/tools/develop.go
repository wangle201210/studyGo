package main

import (
	"github.com/wangle201210/studyGo/tools/devlop"
	"os"
)

func sql2struct() {
	sql := "CREATE TABLE `broker_user` (\n  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,\n  `broker_user_id` int(11) unsigned NOT NULL COMMENT '经纪人用户ID',\n  `broker_user_type` tinyint(2) NOT NULL COMMENT '经纪人类型：1-直销 2-共建 3-兼职',\n  `mobile` varchar(11) NOT NULL COMMENT '手机号码',\n  `name` varchar(255) NOT NULL COMMENT '姓名',\n  `id_card_number` varchar(255) NOT NULL COMMENT '身份证号码',\n  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像URL',\n  `status` tinyint(1) NOT NULL COMMENT '账号状态：1-启用 2-禁用',\n  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  `deleted_at` datetime DEFAULT NULL,\n  PRIMARY KEY (`id`),\n  KEY `idx_broker_user_id` (`broker_user_id`) USING BTREE,\n  KEY `idx_idcard` (`id_card_number`(191)) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='经纪人用户'"
	s := devlop.Sql2struct(sql)
	str := s.String()
	var (
		fileName = "struct.txt"
		f        *os.File
	)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	} else {
		f, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0655)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
	f.WriteString("\n=====================================\n")
	if _, err := f.WriteString(str); err != nil {
		panic(err)
	}
}

func modCheck() {
	filename := "/Users/med/work/git/med-doctor-workstation/med-doctor-operation/go.mod"
	d := &devlop.Mod{
		FilePath: filename,
	}
	d.ModCheck()
}
