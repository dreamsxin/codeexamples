package start

import (
	"fmt"
	"html/template"
	"log"
	"mtlogin/common/middleware"
	"net/http"
	"path/filepath"
	"strings"

	coreCfg "github.com/baowk/dilu-core/config"
	"github.com/gin-gonic/gin"

	"github.com/baowk/dilu-core/core"

	strip "github.com/grokify/html-strip-tags-go" // => strip
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	//AppRouters = make([]func(), 0)
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "dilu start -c config.dev.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config.yml", "Start server with provided configuration file")
}

func ShowLang(c *gin.Context) string {
	lang, exist := c.Get("lang")
	if !exist {
		return ""
	}
	arr := strings.Split(lang.(string), "_")
	s := strings.ToLower(arr[0])
	switch s {
	case "zh-tw":
		return "ZH"
	case "zh", "zh-cn":
		return "ZH"
	case "en", "en-us":
		return "EN"
	case "ru", "ru-ru":
		return "RU"
	}
	return s
}

func GenerateURL(c *gin.Context) string {

	path, exits := c.Get("path")
	core.Log.Sugar().Debugf("GenerateURL %v", path)
	if exits {
		return path.(string)
	}
	return ""
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}
	v := viper.New()
	v.SetConfigFile(configYml)
	//v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}

	var cfg coreCfg.AppCfg

	if err = v.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}

	mergeCfg(&cfg, nil)

	core.Init()

	//初始化gin
	r := core.GetGinEngine()

	// html模板
	templ := template.Must(template.New("").Funcs(template.FuncMap{
		"safehtml": func(s string) template.HTML {
			return template.HTML(s)
		},
		"upper": strings.ToUpper,
		"strip": strip.StripTags,
		"mod": func(i, j int) bool {
			return i%j == 0
		},
	}).ParseGlob("./*.html"))
	r.SetHTMLTemplate(templ)

	r.StaticFS("/images", http.Dir("./images"))
	r.StaticFS("/js", http.Dir("./js"))
	r.StaticFS("/css", http.Dir("./css"))

	r.Group("/zh").Use(middleware.SetLangHandler("zh")).GET("/*name", func(c *gin.Context) {
		name := c.Param("name")
		log.Println("---------name------------", name)
		file, err := filepath.Abs(c.Request.URL.Path)
		if err != nil {
			log.Println(err)
			return
		}
		base := filepath.Base(file)
		//log.Println(base)
		//ext := filepath.Ext(file)
		//log.Println(ext)

		i18n := middleware.GetI18n(c)
		c.HTML(http.StatusOK, base, gin.H{
			"t":    i18n,
			"lang": ShowLang(c),
			"URL":  GenerateURL(c),
		})
	})
	r.Group("/en").Use(middleware.SetLangHandler("en")).GET("/*name", func(c *gin.Context) {
		name := c.Param("name")
		log.Println("---------name------------", name)
		file, err := filepath.Abs(c.Request.URL.Path)
		if err != nil {
			log.Println(err)
			return
		}
		base := filepath.Base(file)
		//log.Println(base)
		//ext := filepath.Ext(file)
		//log.Println(ext)

		i18n := middleware.GetI18n(c)
		c.HTML(http.StatusOK, base, gin.H{
			"t":    i18n,
			"lang": ShowLang(c),
			"URL":  GenerateURL(c),
		})
	})
	r.Group("/ru").Use(middleware.SetLangHandler("ru")).GET("/*name", func(c *gin.Context) {
		name := c.Param("name")
		log.Println("---------name------------", name)
		file, err := filepath.Abs(c.Request.URL.Path)
		if err != nil {
			log.Println(err)
			return
		}
		base := filepath.Base(file)
		//log.Println(base)
		//ext := filepath.Ext(file)
		//log.Println(ext)

		i18n := middleware.GetI18n(c)
		c.HTML(http.StatusOK, base, gin.H{
			"t":    i18n,
			"lang": ShowLang(c),
			"URL":  GenerateURL(c),
		})
	})
	r.NoRoute(func(c *gin.Context) {
		log.Println("---------c.Request.URL.Path------------", c.Request.URL.Path)
		file, err := filepath.Abs(c.Request.URL.Path)
		if err != nil {
			log.Println(err)
			return
		}
		base := filepath.Base(file)
		//log.Println(base)
		//ext := filepath.Ext(file)
		//log.Println(ext)

		i18n := middleware.GetI18n(c)
		c.HTML(http.StatusOK, base, gin.H{
			"t":    i18n,
			"lang": ShowLang(c),
			"URL":  GenerateURL(c),
		})
	})

	//初始化路由
	for _, f := range AppRouters {
		f()
	}

	// 注册服务
	go func() { //主服务启动后回调
		<-core.Started
		startedInit()
	}()

	go func() { //服务关闭释放资源
		<-core.ToClose
		toClose()

	}()
	core.Run()
}

// 服务启动后要初始化的资源
func startedInit() {
	core.Log.Debug("服务启动，初始化执行完成")
}

// 服务关闭要释放的资源
func toClose() {
	core.Log.Debug("服务关闭需要释放的资源")
}

func mergeCfg(local, remote *coreCfg.AppCfg) {
	if remote != nil {
		core.Cfg = *local
		core.Cfg = *remote
		core.Cfg.Server.Mode = local.Server.Mode
		core.Cfg.Server.RemoteEnable = local.Server.RemoteEnable
		core.Cfg.Remote = local.Remote
		core.Cfg.Server.Name = local.Server.Name
		core.Cfg.Server.Port = local.Server.Port
		core.Cfg.Server.Host = local.Server.Host
	} else {
		core.Cfg = *local
	}
}
