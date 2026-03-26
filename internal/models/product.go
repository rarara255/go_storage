package models

import(
	"errors"
	"regexp"
//	"strings"
	"time"
)

var sbinRegex = regexp.MustCompile(`^[0-9]{12,20}$`)

type Product struct{
	Name string
	SBIN string
	ExpiryDate time.Time
}

func NewProduct(name, sbin, date string) (*Product, error){
	if !sbinRegex.MatchString(sbin){
		return nil, errors.New("Некорректная длина SBIN")
	}

	layout := "02.01.2006"
	expire, error := time.Parse(layout, date)
	if error != nil{
		return nil, errors.New("Не получилось спарсить дату!")
	}

	if expire.Before(time.Now()){
		return nil, errors.New("Товар с истёкшим сроком годности не может быть добавлен")
	}

	return &Product{
		Name: name,
		SBIN: sbin,
		ExpiryDate: expire,
	}, nil
}

func (p *Product) DateToString() string{
	return p.ExpiryDate.Format("02.01.2006")
}

