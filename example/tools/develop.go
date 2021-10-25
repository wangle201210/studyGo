package main

import (
	"github.com/wangle201210/studyGo/tools/devlop"
	"os"
)

func sql2struct() {
	sql := "CREATE TABLE `user_phone_auth` (\n  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,\n  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',\n  `med_privacy_agreement` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '医联app隐私协议：0 未读、1已读',\n  `created_at` int(10) unsigned NOT NULL DEFAULT '0',\n  `updated_at` int(10) unsigned DEFAULT '0',\n  PRIMARY KEY (`id`),\n  KEY `idx_phone` (`phone`)\n) ENGINE=InnoDB AUTO_INCREMENT=48560 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='手机号对应的一些权限';"
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
