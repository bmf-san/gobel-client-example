{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">Tags</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable">
    <div class="row">
        <div class="col">
        {{ range $i, $v := .Tags }}
        <a class="tag margin-1rem" href="/posts/tags/{{ $v.Name }}">{{ $v.Name }}</a>
        {{ end }}
        </div>
    </div>
</div>
{{ template "pagination" .Pagination }}
{{ end }}