{{define "navbar"}}
 <div class="navbar navbar-default navbar-fixed-top">
        <div class="container">
              <a class="navbar-brand" href="/">我的博客</a>
              <ul class="nav navbar-nav">
                  <li {{if .IsHome}} class="active" {{end}}><a href="/">首页</a></li>
                  <li {{if .IsCategory}} class="active" {{end}}><a href="/category">分类</a></li>
                  <li {{if .IsArticle}} class="active" {{end}}><a href="/article">文章</a></li>
              </ul>

              <div class="pull-right">
                  <ul class="nav navbar-nav">
                      {{if .IsLogin}}
                      <li><a href="/login?exit=true">退出</a></li>
                      {{else}}
                      <li><a href="/login">管理员登录</a></li>
                      {{end}}
                  </ul>
              </div>
        </div>

        <!-- <div class="page-header">
              <h1>我的第一篇博客</h1>
              <h6 class="text-muted">文章发表于2016年</h6>
              <p>
                  大家好，这是我第一篇博客，谢谢大家的支持。
              </p>
        </div> -->
</div>


{{end}}