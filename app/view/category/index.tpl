{{define "body"}}
<div>
    <h1>/categories</h1>
    {{ range $i, $v := .Categories }}
    <p><span>Category ID</span>:{{ $v.ID }}</p>
    <p><span>Category Name</span>:{{ $v.Name }}</p>
    {{ end }}
</div>
{{end}}