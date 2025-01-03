一、概述
本接口文档描述了基于 Gin 框架搭建的 Web 应用所提供的一系列 API 接口，主要涵盖用户注册、登录、登出以及消息相关操作功能。
二、接口详情
1.用户注册接口
- 接口路径：POSTregister 
- 功能描述：用于创建新用户账号，接收用户提交的注册信息并将其插入到数据库中。
- 请求参数：
  - 请求体参数（JSON 格式）：
    - nickname（字符串类型）：用户昵称，要求长度大于等于 3。
    - username（字符串类型）：用户名。
    - password（字符串类型）：用户密码，要求长度大于等于 3。
- 响应结果：
  - 成功（状态码 200）：
    - message（字符串类型）：返回 User created，表示用户创建成功。
  - 失败（状态码 400）：
    - error（字符串类型）：若请求体参数格式不正确，返回对应的错误提示信息，如 Nickname is too short 或 Password is too short 等；若数据绑定出现问题，返回具体的数据绑定错误信息。
  - 服务器内部错误（状态码 500）：
    - error（字符串类型）：若在数据库插入操作过程中出现错误，返回具体的数据库操作错误信息。
- 请求示例：
  - JSON 格式：
{
    nickname testuser,
    username test_username,
    password testpass
}
  - curl 命令格式 （假设服务运行在本地 8080 端口）：
curl -X POST -H Content-Type applicationjson -d '{nicknametestuser, usernametest_username, passwordtestpass}' httplocalhost8080register
2、用户登录接口
- 接口路径：POSTLogin
- 功能描述：用于验证用户登录信息，根据用户名和密码进行匹配，若匹配成功则设置登录状态的 Cookie。
- 请求参数：
  - 请求体参数（JSON 格式）：
    - username（字符串类型）：用户名，需与注册时的用户名一致。
    - password（字符串类型）：用户密码，需与注册时的密码匹配。
  - 查询参数：
    - password（字符串类型）：通过查询字符串再次传递用户密码，用于和请求体中的密码进行对比验证。
- 响应结果：
  - 成功（状态码 200）：
    - message（字符串类型）：返回 User created，表示用户登录成功。
  - 失败（状态码 400）：
    - error（字符串类型）：若请求体参数格式不正确，返回具体的数据绑定错误信息；若密码不匹配，返回 Password does not match。
- 请求示例：
  - JSON 格式：
{
    username test_username,
    password testpass
}
  - curl 命令格式 （假设服务运行在本地 8080 端口）：
curl -X POST -H Content-Type applicationjson -d '{usernametest_username, passwordtestpass}' httplocalhost8080loginpassword=testpass
3、用户登出接口
- 接口路径：POSTLogout
- 功能描述：用于清除用户登录状态，通过设置登录状态的 Cookie 为失效状态来实现登出功能。
- 响应结果：
  - 成功（状态码 200）：
    - message（字符串类型）：返回 User logged out，表示用户登出成功。
- 请求示例
  - curl 命令格式 （假设服务运行在本地 8080 端口）：
curl -X POST httplocalhost8080logout
4、留言接口（需登录后访问）
- 接口路径：POSTLeaveM
- 功能描述：在用户登录状态下，接收用户提交的消息内容并插入到数据库的 message 表中。
- 请求参数：
  - 请求体参数（JSON 格式）：
    - content（字符串类型）：要留言的内容。
- 响应结果：
  - 成功（状态码 200）：
    - message（字符串类型）：返回 Successfully left messages，表示留言成功。
  - 失败（状态码 400）：
    - error（字符串类型）：若请求体参数格式不正确，返回具体的数据绑定错误信息；若在数据库插入操作过程中出现错误，返回对应的数据库操作错误信息。
- 请求示例：
  - JSON 格式：
{
    content 这是一条留言示例
}
  - curl 命令格式 （假设服务运行在本地 8080 端口）：
curl -X POST -H Content-Type applicationjson -d '{content这是一条留言示例}' httplocalhost8080leaveM
5、删除留言接口（需登录后访问）
- 接口路径：DELETEDeleteM
- 功能描述：在用户登录状态下，根据用户提交的 user_id 删除数据库 message表中对应的记录。
- 请求参数：
  - 请求体参数（JSON 格式）：
    - user_id（整数类型）：用于指定要删除的消息记录对应的用户 ID，以此来确定要删除的具体消息记录。
- 响应结果：
  - 成功（状态码 200）：
    - message（字符串类型）：返回类似 message Successfully delete messages，告知用户删除操作执行情况。
  - 失败（状态码 400）：
    - error（字符串类型）：若请求体参数格式不正确，返回具体的数据绑定错误信息；若在数据库删除操作过程中出现错误，返回对应的数据库操作错误信息。
- 请求示例：
  - JSON 格式：
{
    user_id 1
}
  - curl 命令格式 （假设服务运行在本地 8080 端口）：
curl -X DELETE -H Content-Type applicationjson -d '{user_id1}' httplocalhost8080deleteM