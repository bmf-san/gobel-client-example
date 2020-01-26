{{define "body"}}
<div>
    <h1>/tags</h1>
    {{ range $i, $v := .Tags }}
    <p><span>Tag ID</span>:{{ $v.ID }}</p>
    <p><span>Tag Name</span>:{{ $v.Name }}</p>
    {{ end }}
</div>
{{end}}