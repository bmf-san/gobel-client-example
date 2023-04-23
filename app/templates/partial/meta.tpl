{{ define "meta" }}
	<title>{{ .Title }}</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    {{ if .NoIndex }}
    <meta name="robots" content="noindex">
    {{ else }}
    <meta name="description" content="{{ .Description }}" />
    <meta name="og:title" content="{{ .OGTitle }}" />
    <meta name="og:description" content="{{ .OGDescription }}" />
    <meta property="og:url" content="{{ .OGURL }}" />
    <meta property="og:type" content="{{ .OGType }}">
    <meta property="og:image" content="{{ .OGImage }}" />
    <meta property="og:site_name" content="{{ .OGSiteName }}" />
    <meta property="og:locale" content="{{ .OGLocale }}" />
    <meta name="twitter:card" content="{{ .TwitterCard }}" />
    <meta name="twitter:site" content="{{ .TwitterSite }}" />
	<link rel="canonical" href="{{ .Canonical }}">
	{{ end }}
{{ end }}