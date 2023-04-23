{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">{{ .Posts.CategoryName }}</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{define "body"}}
<div class="container-readable margin-top-2rem">
    {{ template "posts" .Posts }}
    {{ template "pagination" .Posts.Pagination }}
</div>
{{ end }}