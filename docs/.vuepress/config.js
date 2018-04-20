module.exports = {
  title: 'DingTalk Golang SDK',
  description: '简单，易用，稳定',
  themeConfig: {
    sidebar:  [
      ['/guide/', '介绍'],
      ['/guide/getting_started', '起步'],
      ['/guide/open_api_auth', '授权'],
      ['/guide/open_api_user', '用户相关'],
      ['guide/top_api', 'TOP']
    ],
    // 假定 GitHub。也可以是一个完整的 GitLab 网址
    repo: 'icepy/go-dingtalk',
    // 如果你的文档不在仓库的根部
    docsDir: 'docs',
    // 可选，默认为 master
    docsBranch: 'master',
    // 默认为 true，设置为 false 来禁用
    editLinks: true
  }
}