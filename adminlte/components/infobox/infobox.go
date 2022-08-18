/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-19 05:50:40
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-19 05:53:29
 * @FilePath: /themes/adminlte/components/infobox/infobox.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package infobox

import (
	"html/template"
	"strings"

	adminTemplate "github.com/huyongchao98/go-admin/template"
)

type InfoBox struct {
	*adminTemplate.BaseComponent

	Icon       template.HTML
	Text       template.HTML
	Number     template.HTML
	Content    template.HTML
	Color      template.HTML
	IsHexColor bool
	IsSvg      bool
}

func New() InfoBox {
	return InfoBox{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "infobox",
			HTMLData: List["infobox"],
		},
	}
}

func (i InfoBox) SetIcon(value template.HTML) InfoBox {
	i.Icon = value
	if strings.Contains(string(value), "svg") {
		i.IsSvg = true
	}
	return i
}

func (i InfoBox) SetText(value template.HTML) InfoBox {
	i.Text = value
	return i
}

func (i InfoBox) SetNumber(value template.HTML) InfoBox {
	i.Number = value
	return i
}

func (i InfoBox) SetContent(value template.HTML) InfoBox {
	i.Content = value
	return i
}

func (i InfoBox) SetColor(value template.HTML) InfoBox {
	i.Color = value
	if strings.Contains(string(value), "#") {
		i.IsHexColor = true
	}
	return i
}

func (i InfoBox) GetContent() template.HTML { return i.GetContentWithData(i) }
