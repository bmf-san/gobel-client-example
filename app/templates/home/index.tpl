{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">ホーム</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    {{ template "search" }}
    {{ template "posts" .Posts }}
    <div class="row">
        <div class="col text-align-center margin-top-1rem margin-bottom-1rem">
            <a class="color-text" href="/posts">記事をもっと見る</a>
        </div>
    </div>
</div>
{{ end }}