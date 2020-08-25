package response

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// PostIndex is a data for index view.
type PostIndex struct {
	Posts      *model.Posts
	Pagination *model.Pagination
}

// PostIndexByCategory is a data for index view by category.
type PostIndexByCategory struct {
	Posts      *model.Posts
	Pagination *model.Pagination
}

// PostIndexByTag is a data for index view by tag.
type PostIndexByTag struct {
	Posts      *model.Posts
	Pagination *model.Pagination
}

// PostShow is a data for show view.
type PostShow struct {
	Post *model.Post
}

// ExecutePostIndex responses a index view.
func (r *Response) ExecutePostIndex(w http.ResponseWriter, p *PostIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/index.tpl", "view/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostIndexByCategory responses a index view by category.
func (r *Response) ExecutePostIndexByCategory(w http.ResponseWriter, p *PostIndexByCategory) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/category.tpl", "view/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostIndexByTag responses a index view by tag.
func (r *Response) ExecutePostIndexByTag(w http.ResponseWriter, p *PostIndexByTag) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/tag.tpl", "view/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostShow responses a show view by tag.
func (r *Response) ExecutePostShow(w http.ResponseWriter, p *PostShow) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/show.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}
