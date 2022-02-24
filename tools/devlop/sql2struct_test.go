package devlop

import (
	"fmt"
	"testing"
)

func TestSql2struct(t *testing.T) {
	s := "CREATE TABLE `callback_event` (\n  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,\n  `appId` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '当前消息来源appId',\n  `business` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务来源 场景二维码才有，默认为空',\n  `openId` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户在微信公众号的openId',\n  `userId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'openId注册的临时用户id',\n  `scene_qrcode_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'med_wechat.scene_qrcode',\n  `scene_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '场景值',\n  `scene_type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '场景类型',\n  `content` text COLLATE utf8mb4_unicode_ci COMMENT '整个数据的json',\n  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  PRIMARY KEY (`id`),\n  KEY `appIdIdx` (`appId`),\n  KEY `openIdIdx` (`openId`),\n  KEY `userIdIdx` (`userId`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='微信回调事件记录'"
	s2s := Sql2struct(s)
	fmt.Println(s2s.String())
}
