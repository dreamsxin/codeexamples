package middleware

import (
	"sort"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/baowk/dilu-core/core"
	t "github.com/dreamsxin/go-i18n"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

func SetLangHandler(lang string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if lang != "" {
			c.Set("path", "/"+lang)
			core.Log.Sugar().Debugln("------------------SetLangHandler-------------------------", lang)
		}
		firstlang := lang
		if firstlang == "" {
			firstlang = c.Query("lang")
		}

		if firstlang == "" {
			firstlang = c.GetHeader("Accept-Language")
		}
		i18n, firstlang := NewI18n(firstlang)
		c.Set("i18n", i18n)
		c.Set("lang", firstlang)
		c.Next()
	}
}

func init() {
	t.Load("./langs")
}

func GetI18n(c *gin.Context) *t.Translations {
	i, exist := c.Get("i18n")
	if exist {
		return i.(*t.Translations)
	}
	lang := c.Query("lang")
	if lang == "" {
		lang = c.GetHeader("Accept-Language")
	}
	i18n, lang := NewI18n(lang)
	c.Set("i18n", i18n)
	c.Set("lang", lang)
	return i18n
}

type Langs struct {
	Id         int     `json:"id" gorm:"type:int;primaryKey;autoIncrement;comment:主键"` //主键
	Lang       string  `json:"lang" gorm:"type:varchar(20);comment:Lang"`              //
	Name       string  `json:"name" gorm:"type:varchar(40);comment:语言名称"`              //语言名称
	Desc       string  `json:"desc" gorm:"type:varchar(255);comment:简短说明"`             //简短说明
	Status     int     `json:"status" gorm:"type:int unsigned;comment:1隐藏 2正常"`        //1隐藏 2正常
	Sort       int     `json:"sort" gorm:"type:int unsigned;comment:排序"`               //排序
	Similarity float64 `json:"-" gorm:"-"`                                             //相似度
}

func NewI18n(lang string) (*t.Translations, string) {
	supported := t.Locales()
	core.Log.Sugar().Debugf("GetI18n lang %v supported %v", lang, supported)

	firstlang := lang // "zh-cn"
	tags, _, _ := language.ParseAcceptLanguage(lang)
	core.Log.Debug("GetI18n", zap.Any("tags", tags))

	langs := []string{}
	for _, tag := range tags {
		langs = append(langs, strings.ToLower(strings.Replace(tag.String(), "-", "_", -1)))
	}
	if len(tags) <= 0 {
		langs = append(langs, firstlang)
	}

	tempList := make([]*Langs, len(supported))
	for i, s := range supported {
		tempList[i] = &Langs{Lang: s, Similarity: 0}
	}
	if len(tempList) <= 0 {
		return nil, lang
	}
	for pos, userlang := range langs {
		for i := range tempList {
			//取相似度最高的
			similarity := strutil.Similarity(tempList[i].Lang, userlang, metrics.NewLevenshtein()) - (0.1 * float64(pos))
			if tempList[i].Similarity < (similarity) {
				tempList[i].Similarity = similarity
			}
			if similarity >= 0.8 {
				trans := t.L(tempList[i].Lang)
				trans.Remove = true
				core.Log.Sugar().Debugf("GetI18n 计算语言相似度 Similarity %v Lang %v userlang %v", tempList[i].Similarity, tempList[i].Lang, userlang)
				return trans, tempList[i].Lang
			}
			core.Log.Sugar().Debugf("GetI18n 计算语言相似度 Similarity %v Lang %v userlang %v", tempList[i].Similarity, tempList[i].Lang, userlang)
		}
	}
	sort.Slice(tempList, func(i, j int) bool {
		return tempList[i].Similarity >= tempList[j].Similarity
	})
	core.Log.Sugar().Debugf("GetI18n 计算语言相似度 tempList %v", tempList[0])

	trans := t.L(tempList[0].Lang)
	trans.Remove = true
	return trans, tempList[0].Lang
}
