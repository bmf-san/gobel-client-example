{{ define "pagination" }}
<div class="pagination">
    <ul class="pagination-list">
        {{ range $i, $_ := .Pager.PaginationPageCount }}
            {{ if gt $i 0 }}
            {{ if $.QueryParams }}
                <li><a {{if eq $.Pager.PaginationPage $i}} class="pagination-active" {{end}} href="?{{ $.QueryParams }}&page={{ $i }}&limit={{ $.Pager.PaginationLimit }}"><span>{{ $i }}</span></a></li>
            {{ else }}
                <li><a {{if eq $.Pager.PaginationPage $i}} class="pagination-active" {{end}} href="?page={{ $i }}&limit={{ $.Pager.PaginationLimit }}"><span>{{ $i }}</span></a></li>
            {{ end }}
            {{ end }}
        {{ end }}
    </ul>
</div>
{{ end }}