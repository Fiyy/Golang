# iFt使用方法

## 工具说明

    PC与Devnet、IDC间的文件传输工具，目前稳定版本适用于Windows、Mac端与Linux服务端的文件传输操作。

## 配置

1. 在Weterm中使用，用Weterm连接服务器

2. 在服务器中执行iFt客户端给出的指令

3. 第一次使用要身份验证，有效期一个月，Username是用户名，还用输入PIN+Token

## 使用指令

1. Linux ————> PC:  
    输入ft put + 文件名或者 ftsz + 文件名
    会弹出一个对话框，选择PC中的存储位置
    可以使用 -s 或 --select 参数指定配置
2. PC ————> Linux：
    输入ft get + 文件名 或者 ftrz + 文件名
    会弹出对话框选择PC中的文件，默认保存在Linux的ft_local文件中
    可以使用 -s 或 --select 参数指定配置
3. 使用ft login可以切换用户
4. ft logout username可以注销用户。
5. 参数表说明：
        -u username：指定用户。

        -d：指定agent所在机器名，不指定则自动获取列表，提示用户选择。

        -o, --overwrite：重名文件替换。

        --rename：重名文件保留。

        -s，--select strings：选择不弹窗目录配置名。

        --skip：跳过重名文件。
