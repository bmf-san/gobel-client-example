{{ define "body" }}
<div>
    <h1>/tags</h1>
    {{ range $i, $v := .Tags }}
    <p><span>Tag ID</span>:{{ $v.ID }}</p>
    <a href="/posts/tags/{{ $v.Name }}"><span>Tag Name</span>:{{ $v.Name }}</a>
    {{ end }}
</div>
{{ template "pagination" .Pagination }}
{{ end }}