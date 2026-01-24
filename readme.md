<h1 align="center" style="margin: 30px 0 30px; font-weight: bold; font-size: 30px">RuoYi-Go</h1>
<h4 align="center">基于Golang+Gin+Gorm实现的若依服务端脚手架</h4>

## 平台简介
* 本仓库为后端技术栈 [Gin](https://gin-gonic.com/zh-cn/docs) + [Gorm](https://gorm.io/zh_CN/docs/index.html) 的 `golang` 版本。
* 配套前端代码仓库地址 [RuoYi-Vue3](https://github.com/yangzongzhuan/RuoYi-Vue3) 或使用 [RuoYi-Vue3-ts](https://github.com/zzh948498/RuoYi-Vue3-ts)
* 其他生态组件请访问 **[若依官网](http://ruoyi.vip/)**

## 后端运行
> **提示：** 运行前请先安装好 `go` 环境，版本 `1.18` 以上。

    # 克隆项目
    git clone https://github.com/fivepmcoder/ruoyi-go.git

    # 进入项目目录
    cd ruoyi-go

    # 修改配置文件
    cp application-example.yaml application.yaml

    # 安装依赖
    go mod tidy

    # 启动服务
    go run main.go

## 前端运行
    # 调整 .env 文件
    VUE_APP_BASE_API = '/dev-api' 改为 VITE_APP_BASE_API = '/api'

    # 调整 vite.config.js 文件
    server: {
      port: 8000,
      open: false,
      proxy: {
        // https://cn.vitejs.dev/config/#server-proxy
        '/api': {
          target: 'http://localhost:3000',
          changeOrigin: true,
          rewrite: (p) => p.replace(/^\/api/, '/api')
        }
      }
    },

    # 安装依赖
    npm install

    # 启动服务
    npm run dev

## 后端打包
    # 打包
    go build main.go

## 内置功能
1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9.  登录日志：系统登录日志记录查询包含登录异常。

## 🤝 贡献指南
欢迎所有形式的贡献！无论是报告 Bug、提出新功能建议，还是提交代码改进。

### 如何贡献
1. **Fork 本仓库**
    点击右上角的 "Fork" 按钮，将项目复制到你的 GitHub 账户
2. **克隆到本地**
    ```
    git clone https://github.com/fivepmcoder/ruoyi-go.git
    cd ruoyi-go
    ```
3. **创建特性分支**
    ```
    # 功能开发
    git checkout -b feature/your-feature-name

    # Bug 修复
    git checkout -b fix/your-bug-fix

    # 文档更新
    git checkout -b docs/your-doc-update
    ```
4. **进行开发**
    - 遵循项目现有的代码风格
    - 添加必要的注释和文档
    - 如果是新功能，请更新 README
5. **提交更改**
    ```
    git add .
    git commit -m "feat: 添加新功能描述"

    # 提交信息规范：
    # feat: 新功能
    # fix: Bug 修复
    # docs: 文档更新
    # style: 代码格式调整
    # refactor: 代码重构
    # test: 测试相关
    # chore: 构建/工具链相关
    ```
6. **推送到 GitHub**
    ```
    git push origin feature/your-feature-name
    ```
7. **创建 Pull Request**
    - 在 GitHub 上打开你的 Fork 仓库
    - 点击 "New Pull Request" 按钮
    - 填写 PR 标题和详细描述：
        - 说明改动的目的和内容
        - 如果修复了 Issue，请关联对应的 Issue 编号
        - 附上测试截图或日志（如适用）
    - 等待维护者审核

### 报告问题
如果你发现了 Bug 或有功能建议，请：
1. 在 [Issues](https://github.com/fivepmcoder/ruoyi-go/issues) 中搜索是否已有相关问题
2. 如果没有，创建新的 Issue，并提供：
    - 清晰的标题
    - 详细的问题描述
    - 复现步骤（如果是 Bug）
    - 期望的行为
    - 环境信息（Go 版本、操作系统等）

## ⚠️ 免责声明
> **重要提示**：本项目为个人学习和研究项目，**未经过企业级生产环境验证**。
### 使用须知
-   🔸 本项目仅供学习、研究和参考使用
-   🔸 不建议直接用于生产环境，除非经过充分测试和评估
-   🔸 使用本项目代码所产生的任何问题，作者不承担责任
-   🔸 建议在使用前进行全面的安全审计和性能测试

### 生产环境使用建议
如果你计划在生产环境中使用本项目，请务必：
1. **安全审计**
    - 审查所有安全相关代码
    - 检查依赖包的安全漏洞
    - 实施额外的安全措施（如 WAF、限流等）
2. **性能测试**
    - 进行压力测试和负载测试
    - 优化数据库查询和索引
    - 配置合理的连接池大小
3. **监控和日志**
    - 添加完善的日志系统
    - 配置应用监控和告警
    - 实施错误追踪机制
4. **备份和恢复**
    - 制定数据备份策略
    - 测试灾难恢复流程
    - 准备应急预案

### 已知限制
-   ⚠️ 缺少完整的单元测试和集成测试
-   ⚠️ 未进行大规模并发测试
-   ⚠️ 部分功能可能需要根据实际业务调整
-   ⚠️ 文档可能存在不完善之处

## 📄 开源协议
本项目采用 [MIT License](LICENSE) 开源协议。
这意味着你可以：
-   ✅ 自由使用、复制、修改、合并、发布、分发本软件
-   ✅ 用于商业或非商业目的
-   ✅ 在遵守协议的前提下，可以闭源使用
但需要：
-   📌 在软件和文档中保留版权声明和许可声明
-   📌 软件按"原样"提供，不提供任何形式的担保

## 📮 联系方式
-   🐛 Issues: [GitHub Issues](https://github.com/fivepmcoder/ruoyi-go/issues)

**如果这个项目对你有帮助，请给个 ⭐️ Star 支持一下！**

Made with ❤️ by [Zero](https://github.com/fivepmcoder)

## Star History
<a href="https://www.star-history.com/#fivepmcoder/ruoyi-go&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=fivepmcoder/ruoyi-go&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=fivepmcoder/ruoyi-go&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=fivepmcoder/ruoyi-go&type=Date" />
 </picture>
</a>
