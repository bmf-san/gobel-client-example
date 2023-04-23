{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">タグ</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    <div class="row">
        <div class="col">
        {{ range $i, $v := .Tags.Tags }}
        <a class="tag margin-1rem" href="/posts/tags/{{ $v.Name }}">{{ $v.Name }}</a>
        {{ end }}
        </div>
    </div>
    <div class="row">
        <div class="col">
            <!-- Google Adsense ディスプレイ -->
            <!-- Your adsense codes -->
        </div>
    </div>
</div>
{{ template "pagination" .Tags.Pagination }}
{{ end }}