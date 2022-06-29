package validator

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Validator struct {
	instance *validator.Validate
	uni      *ut.UniversalTranslator
}

func New() (*Validator, error) {
	lang := zh.New()
	uni := ut.New(lang, lang)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	err := translations.RegisterDefaultTranslations(validate, trans)
	return &Validator{
		instance: validate,
		uni: uni,
	}, err
}

func (v *Validator) Validate(s interface{}) error{
	errs := v.instance.Struct(s)
	if errs != nil {
		err := errs.(validator.ValidationErrors)
		trans, _ := v.uni.GetTranslator("zh")
		return errors.New(err[0].Translate(trans))
	}
	return nil
}
