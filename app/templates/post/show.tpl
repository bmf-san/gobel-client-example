{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">{{ .Post.Title }}</h1>
                <div class="text-align-center">
                    {{ range .Post.Tags }}
                    <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                    {{ end }}
                </div>
                <div class="text-align-center">
                    <a href="/posts/categories/{{ .Post.Category.Name }}">{{ .Post.Category.Name }}</a>
                </div>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<section>
    <div class="container-readable">
        <div class="row">
            <div class="col">
                <p><span class="article-date">{{ .Post.CreatedAt.Format "2006 Jan 02" }}</span>・<span class="article-date">Updated On{{ .Post.UpdatedAt.Format "2006 Jan 02" }}</span></p>
                {{ .Post.HTMLBody | unescape }}
            </div>
        </div>
    </div>
</section>
{{ end }}