{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">カテゴリ</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    <div class="row">
        <div class="col">
            <ul>
            {{ range $i, $v := .Categories.Categories }}
            <li>
                <a href="/posts/categories/{{ $v.Name }}">{{ $v.Name }}</a>
            </li>
            {{ end }}
            </ul>
        </div>
    </div>
    <div class="row">
        <div class="col">
            <!-- Google Adsense ディスプレイ -->
            <!-- Your adsense codes -->
        </div>
    </div>
</div>
{{ template "pagination" .Categories.Pagination }}
{{ end }}