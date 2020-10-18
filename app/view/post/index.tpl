{{define "body"}}
<div>
    <h1>/posts</h1>
    {{ range $i, $v := .Posts }}
    <p><span>Post ID</span>:{{ $v.ID }}</p>
    <p><span>Admin ID</span>:{{ $v.Admin.ID }}</p>
    <p><span>Category ID</span>:{{ $v.Category.ID }}</p>
    <p><span>Category Name</span>:{{ $v.Category.Name }}</p>
    {{ range $v.Tags }}
    <p><span>Tag ID</span>:{{ .ID }}</p>
    <p><span>Tag Name</span>:{{ .Name }}</p>
    {{ end }}
    <a href="/posts/{{ $v.Title }}"><span>Title</span>:{{ $v.Title }}</a>
    <p><span>MD Body</span>:{{ $v.MDBody }}</p>
    <p><span>HTML Body</span>:{{ $v.HTMLBody }}</p>
    <p><span>Status</span>:{{ $v.Status }}</p>
    {{ range $v.Comments }}
    <p><span>Comment ID</span>:{{ .ID }}</p>
    <p><span>Comment Body</span>:{{ .Body }}</p>
    {{ end }}
    <p><span>CreateAt</span>:{{ $v.CreatedAt }}</p>
    <p><span>UpdatedAt</span>:{{ $v.UpdatedAt }}</p>
    <hr>
    {{ end }}
</div>
{{ template "pagination" .Pagination }}
{{ end }}
