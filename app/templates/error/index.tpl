{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
	<div class="container">
		<div class="row">
			<div class="col">
				<h1 class="text-align-center color-text-reverse">エラー</h1>
			</div>
		</div>
	</div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable">
	<div class="row">
		<div class="col">
			<h1>{{ .ErrorData.Code }}</h1>
			<p>{{ .ErrorData.Message }}</p>
		</div>
	</div>
</div>
{{ end }}