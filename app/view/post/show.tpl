{{ define "body" }}
<div>
    <h1>/posts/:title</h1>
    <p><span>Post ID</span>:{{ .Post.ID }}</p>
    <p><span>Admin ID</span>:{{ .Post.Admin.ID }}</p>
    <p><span>Category ID</span>:{{ .Post.Category.ID }}</p>
    <p><span>Category Name</span>:{{ .Post.Category.Name }}</p>
    {{ range .Post.Tags }}
    <p><span>Tag ID</span>:{{ .ID }}</p>
    <p><span>Tag Name</span>:{{ .Name }}</p>
    {{ end }}
    <p><span>Title</span>:{{ .Post.Title }}</p>
    <p><span>MD Body</span>:{{ .Post.MDBody }}</p>
    <p><span>HTML Body</span>:{{ .Post.HTMLBody }}</p>
    <p><span>Status</span>:{{ .Post.Status }}</p>
    {{ range .Post.Comments }}
    <p><span>Comment ID</span>:{{ .ID }}</p>
    <p><span>Comment Name</span>:{{ .Body }}</p>
    {{ end }}
    <p><span>CreateAt</span>:{{ .Post.CreatedAt }}</p>
    <p><span>UpdatedAt</span>:{{ .Post.UpdatedAt }}</p>
</div>

<!-- TODO: implement later -->
<!-- <form action="/posts/{{ .Title }}/comments" method="post">
    <textarea name="" id="" cols="30" rows="10"></textarea>
    <button type="submit">Submit</button>
</form> -->
{{ end }}