{{define "base"}}
<!DOCTYPE html>
<html lang="ja">

<head>
    <title></title>
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
    <meta name="twitter:card" content="summary" />
    <meta name="twitter:site" content="@bmf_san" />
    <link rel="canonical" href="">
    <link rel="alternate" hreflang="ja" href="">
</head>

<body>
    <section>
    {{template "body" .}}
    </section>
</body>

</html>
{{end}}
