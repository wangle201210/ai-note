# ai-note 记账小助手
基于 dify + goframe 的记账小助手   
将消费数据自动分类到对应类型，并且可以进行数据分析   
如果需要添加分类，可以在数据库中添加即可  
接入到公众号后，每个openID就对应一个用户  
汇总数据时可以说“2025年1月的数据”，“上月**类型的消费”，“前天数据”，“昨天上午的餐饮消费”等

# 效果如下

![image.png](https://github.com/wangle201210/ai-note/blob/dd02dc19daa0faf8094d3d55c5c1585b4039feb0/app/ai-note/resource/public/resource/image/img.png?raw=true)
![image_1.png](https://github.com/wangle201210/ai-note/blob/dd02dc19daa0faf8094d3d55c5c1585b4039feb0/app/ai-note/resource/public/resource/image/img_1.png?raw=true)

# 运行服务端
导入mysql数据库文件 `note.sql`
```bash
cd app/ai-note
go run main.go
```
# 运行 dify
[本地部署 dify](https://docs.dify.ai/zh-hans/getting-started/install-self-hosted/docker-compose).  
或者[云服务部署 dify](https://docs.dify.ai/zh-hans/getting-started/cloud).  
在"工作室"菜单点击导入DSL，导入[note.yml](https://github.com/wangle201210/ai-note/blob/dd02dc19daa0faf8094d3d55c5c1585b4039feb0/note_20250115.yml).  
在"编排"中修改"会话变量"。 

# 接入微信公众号
修改 app/ai-note/internal/controller/wechat/wechat.go 和 app/ai-note/internal/controller/dify/dify.go 中的配置即可   
