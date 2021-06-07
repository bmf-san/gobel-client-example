{{ define "body" }}
<div class="container">
	<div class="row">
		<div class="col">
			<h1>Error</h1>
			<p>Error Code:{{ .Code }}</p>
			<p>Error Message:{{ .Message }}</p>
		</div>
	</div>
</div>
{{ end }}