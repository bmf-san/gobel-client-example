{{ define "search" }}
<div class="row">
	<div class="col">
		<form class="search" action="/posts/search" method="GET">
			<input class="margin-bottom-0rem" type="search" name="keyword" placeholder="キーワードを入力">
			<button class="search-button" type="submit">検索</button>
		</form>
	</div>
</div>
{{ end }}