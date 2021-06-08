{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">Categories</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable">
    <div class="row">
        <div class="col">
            <ul>
                {{ range $i, $v := .Categories }}
                <li>
                    <a href="/posts/categories/{{ $v.Name }}">{{ $v.Name }}</a>
                </li>
                {{ end }}
            </ul>
        </div>
    </div>
</div>
{{ template "pagination" .Pagination }}
{{ end }}