package template

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"log"
	"strconv"
	"strings"

	"dilu/common/config"
	"dilu/resources"
)

type Pagination struct {
	PerPage     int
	CurrentPage int
	TotalPage   int
	baseUrl     string

	// render parts
	FirstPart  []int
	MiddlePart []int
	LastPart   []int
}

// constructor
func New(totalPage, currentPage int, baseUrl string) *Pagination {
	if currentPage == 0 {
		currentPage = 1
	}

	if currentPage > totalPage {
		currentPage = totalPage
	}

	return &Pagination{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		baseUrl:     baseUrl,
	}
}

// 总共几页
func (p *Pagination) TotalPages() int {
	return p.TotalPage
}

// 是否要显示分页
func (p *Pagination) HasPages() bool {
	return p.TotalPages() > 1
}

const onEachSide int = 3

func (p *Pagination) generate() {
	if !p.HasPages() {
		return
	}
	if p.TotalPages() < (onEachSide*2 + 6) { // 11页以内
		p.FirstPart = p.GetPageRange(1, p.TotalPages())
	} else {
		window := onEachSide * 2
		lastPage := p.TotalPages()
		if p.CurrentPage < window { // 靠近开头
			p.FirstPart = p.GetPageRange(1, window+2)
			p.LastPart = p.GetPageRange(lastPage-1, lastPage)
		} else if p.CurrentPage > (lastPage - window) { // 靠近结尾
			p.FirstPart = p.GetPageRange(1, 2)
			p.LastPart = p.GetPageRange(lastPage-(window+2), lastPage)
		} else { // 在中间
			p.FirstPart = p.GetPageRange(1, 2)
			p.MiddlePart = p.GetPageRange(p.CurrentPage-onEachSide, p.CurrentPage+onEachSide)
			p.LastPart = p.GetPageRange(lastPage-1, lastPage)
		}
	}
}

func (p *Pagination) GetPageRange(start, end int) []int {
	var ret []int
	for i := start; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}

func (p *Pagination) GetUrl(page int) string {
	strPage := strconv.Itoa(page)

	// baseUrl, _ := url.Parse(p.baseUrl)
	// params := baseUrl.Query()
	// delete(params, "page")
	// strParam := ""
	// for k, v := range params {
	// 	strParam = strParam + "&" + k + "=" + v[0] // TODO
	// }

	// href := baseUrl.Path + "?page=" + strPage + strParam
	return strings.Replace(p.baseUrl, "{page}", strPage, -1)
}

func (p *Pagination) IsActive(page int) bool {
	if p.CurrentPage == page {
		return true
	} else {
		return false
	}
}

func (p *Pagination) GetPrevious() int {
	if p.CurrentPage <= 1 {
		return p.CurrentPage
	}
	return p.CurrentPage - 1
}

func (p *Pagination) GetNext() int {
	if p.CurrentPage == p.TotalPages() {
		return p.CurrentPage
	}
	return p.CurrentPage + 1
}

func (p *Pagination) Render() template.HTML {
	p.generate()

	tmplpath := "html/" + config.Ext.Template.Path + "/pagination.tmpl"

	var out bytes.Buffer

	t := template.Must(template.New("pagination.tmpl").ParseFS(resources.TemplateFs, tmplpath))

	err := t.Execute(&out, p)
	if err != nil {
		return template.HTML(fmt.Sprintf("Error executing template: %s", err))
	}
	return template.HTML(html.UnescapeString(out.String()))
}
