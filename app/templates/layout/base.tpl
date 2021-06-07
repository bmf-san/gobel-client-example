{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">
<head>
    <title>gobel-client-example</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="" />
    <meta name="og:title" content="" />
    <meta name="og:description" content="" />
    <meta property="og:url" content="" />
    <meta property="og:type" content="">
    <meta property="og:image" content="" />
    <meta property="og:site_name" content="" />
    <meta property="og:locale" content="ja_JP" />
    <meta name="twitter:card" content="" />
    <meta name="twitter:site" content="" />
	<link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
</head>
<body>
    <header>
        <div>
            <nav class="nav sp-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link-logo"><b>gobel-client-example</b></a>
                    </div>
                    <div class="nav-right">
                        <a class="nav-link" href="/">Home</a>
                        <a class="nav-link" href="/posts">Posts</a>
                        <a class="nav-link" href="/categories">Categories</a>
                        <a class="nav-link" href="/tags">Tags</a>
                    </div>
                </div>
            </nav>
            <nav class="nav pc-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link text-decoration-none">gobel-client-example</a>
                    </div>
                    <div class="nav-right">
                        <div id="nav-sp-drawer">
                            <input id="nav-sp-input" type="checkbox" class="sp-hidden">
                            <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
                            <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"><span></span></label>
                            <div id="nav-sp-content">
                                <h1>Navigation</h1>
                                <a class="nav-link" href="/">Home</a>
                                <a class="nav-link" href="/posts">Posts</a>
                                <a class="nav-link" href="/categories">Categories</a>
                                <a class="nav-link" href="/tags">Tags</a>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </header>

    {{ template "headline" . }}

    {{ template "body" . }}

    <footer class="sticky-footer margin-top-2rem">
        <div class="container">
            <div class="row">
                <div class="col">
                    <h1>Footer</h1>
                </div>
            </div>
            <div class="row text-align-center">
                <div class="col">
                    <a class="color-text-reverse" href="https://twitter.com/bmf_san">Twitter</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="https://github.com/bmf-san">Github</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/sitemap">RSS</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/feed">Feed</a>
                </div>
            </div>
            <hr>
            <div class="row text-align-center">
                <div class="col">
                    <small>gobel-client-example is licensed MIT.</small>
                </div>
            </div>
        </div>
    </footer>
</body>
</html>
{{ end }}


