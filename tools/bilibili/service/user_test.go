package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
	"testing"
)

var (
	s   *Service
	ctx = context.Background()
)

func init() {
	s = New()
	log.Init(nil)
}

func TestUser(t *testing.T) {
	list := []string{
		"啊吗粽", "阿特警官", "布锅锅", "宝剑嫂", "拜托了小翔哥", "翠花不太脆", "才疏学浅的才浅", "DarkCarrot", "朵朵花林", "多多poi丶",
		"大漠叔叔", "大碗拿铁", "倒悬的橘子", "大象放映室", "大祥哥来了", "盗月社食遇记", "导演小策", "EdmundDZhang", "二二酸酸",
		"凤凰传奇", "附魔星Fms", "泛式", "火柴人AlanBecker", "黑猫厨房", "和猫住", "花少北丶", "好运的鱼", "浑元Rysn", "画渣花小烙",
		"幾加乘", "极客湾Geekerwan", "记录生活的蛋黄派", "嘉然今天吃什么", "九三的耳朵不是特别好", "极速拍档", "狂阿弥_",
		"老番茄", "刘谦", "柳青瑶本尊", "老师好我叫何同学", "利维坦mY", "罗翔说刑法", "泠鸢yousa", "MayTree五月树", "魔法Zc目录", "某幻君", "M木糖M",
		"绵羊料理", "木鱼水心", "哦呼w", "oooooohmygosh", "帕梅拉PamelaReif", "枪弹轨迹", "取景框看世界", "青莲又失利了", "RAY的模型世界", "日食记",
		"叔贵", "塑料叉FOKU", "森纳映画", "沙盘上的战争", "帅soserious", "水无月菌", "苏星河牛通", "兔叭咯", "田野上的繁荣", "Vicky宣宣", "汪品先院士",
		"无穷小亮的科普日常", "我是郭杰瑞", "小潮院长", "-小D-biu", "熊乐乐大魔王", "小片片说大片", "小透明明TM", "小文哥吃吃吃",
		"香香软软的小泡芙", "小约翰可汗", "逍遥散人", "星有野", "o小庄o", "远古黑金", "硬核的半佛仙人", "杨可爱Ukulele", "一鹿有车", "=咬人猫=",
		"影视飓风", "鹦鹉梨", "zettaranc", "稚晖君", "自来卷三木", "在下铁头阿彪", "朝烟今天唱歌了没",
		// 再来核实下名单
		"hanser", "电影最TOP", "机智的Kason", "老爸评测", "先看评测", "中国BOY超级大猩猩",
	}
	_ = s.UpdateUerList(ctx, list)
}

func TestUpstat(t *testing.T) {
	s.UpdateUpstat(ctx)
}
