# YafGo - Yet Another Framework for Go

封装提取一个满足自己日常需求的 Golang **重型框架** <img width="24" src="data:image/png;base64,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" />

参考项目：

-   [goravel](https://github.com/goravel/goravel)
-   [gohub](https://github.com/summerblue/gohub)

framework 代码基于 goravel 修改
