{{ define "body" }}
<div>
    <h1>/categories</h1>
    {{ range $i, $v := .Categories }}
    <p><span>Category ID</span>:{{ $v.ID }}</p>
    <a href="/posts/categories/{{ $v.Name }}"><span>Category Name</span>:{{ $v.Name }}</a>
    {{ end }}
</div>
{{ template "pagination" .Pagination }}
{{ end }}