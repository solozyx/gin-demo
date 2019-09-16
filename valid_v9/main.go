package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"

	// v9 支持多语言处理
	"gopkg.in/go-playground/validator.v9"

	// 基于语言做错误显示
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type User struct {
	// 验证标签 v8 binding  v9 validate
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

func main() {
	// 支持语言
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	// 创建 v9 翻译器
	Uni = ut.New(en, zh, zh_tw)
	// 创建 v9 验证器
	Validate = validator.New()

	r := gin.Default()
	r.GET("/test", startPage)
	r.POST("/test", startPage)
	r.Run(":8080")
}

func startPage(c *gin.Context) {
	// 这部分应放到中间件中
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := Uni.GetTranslator(locale)
	switch locale {
	case "zh":
		// 把翻译器注册到验证器中
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		// break
	case "en":
		en_translations.RegisterDefaultTranslations(Validate, trans)
		// break
	case "zh_tw":
		zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
		// break
	default:
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		// break
	}

	// 自定义错误内容
	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// 这块应该放到公共验证方法中
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		c.Abort()
		return
	}

	if err := Validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		c.String(http.StatusInternalServerError, fmt.Sprintf("%v", sliceErrs))
		c.Abort()
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%v", user))
}
