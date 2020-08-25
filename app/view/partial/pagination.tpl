{{ define "pagination" }}
<ul>
    {{ range $i, $_ := .PaginationPageCount }}
    {{ if gt $i 0 }}
    <li><a href="?page={{ $i }}&limit={{ $.PaginationLimit }}">{{ $i }}</a></li>
    {{ end }}
    {{ end }}
</ul>
{{ end }}