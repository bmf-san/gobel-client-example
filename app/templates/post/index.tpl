{{ define "headline" }}
    <div class="headline background-color-main color-text-reverse">
        <div class="container">
            <div class="row">
                <div class="col">
                    <h1 class="text-align-center color-text-reverse">Posts</h1>
                </div>
            </div>
        </div>
    </div>
{{ end }}
{{define "body"}}
<div class="container-readable">
    {{ range $i, $v := .Posts }}
    <div class="row">
        <div class="col">
            <article>
                <span class="article-date">{{ $v.CreatedAt.Format "2006 Jan 02" }}</span>
                <h1 class="margin-0rem"><a href="/posts/{{ $v.Title }}"><b>{{ $v.Title }}</b></a></h1>
                <div class="text-align-right">
                    <a href="/posts/categories/{{ $v.Category.Name }}">{{ $v.Category.Name }}</a>
                </div>
                <div class="text-align-right">
                    {{ range $v.Tags }}
                    <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                    {{ end }}
                </div>
            </article>
        </div>
    </div>
    {{ end }}
    {{ template "pagination" .Pagination }}
</div>
{{ end }}