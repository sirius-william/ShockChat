# 说明
定义了文本协议，响应体根据需要进行RSA加密
# 内容
<table  width="100%" border="0" cellspacing="1" cellpadding="4" class="tabtop13" align="center">
    <tr align="center" bgcolor="#4472CA">
        <th width="10%">
            <span style="color: #CFDEE7; ">模块</span>
        </th>
        <th width="15%">
            <span style="color: #CFDEE7; ">ID</span>
        </th>
        <th width="12%">
            <span style="color: #CFDEE7; ">发送端</span>
        </th>
        <th width="50%">
            <span style="color: #CFDEE7; ">功能</span>
        </th>
        <th width="10%">
            <span style="color: #CFDEE7; ">proto</span>
        </th>
    </tr>
    <tr align="center" bgcolor="#92B4F4">
        <td rowspan="4">
            <span style="color: #CFDEE7; ">验证连接合法性</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">0x100</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">客户端</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">请求验证连接合法性，携带：无</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">User.proto</span>
        </td>
    </tr>
    <tr align="center" bgcolor="#92B4F4">
        <td>
            <span style="color: #CFDEE7; ">0x101</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">服务端</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">携带：随机盐值</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">User.proto</span>
        </td>
    </tr>
    <tr align="center" bgcolor="#92B4F4">
        <td>
            <span style="color: #CFDEE7; ">0x102</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">客户端</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">携带：指定算法计算后的结果</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">User.proto</span>
        </td>
    </tr>
    <tr align="center" bgcolor="#92B4F4">
        <td>
            <span style="color: #CFDEE7; ">0x103</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">服务端</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">成功或失败响应，失败会包含错误信息</span>
        </td>
        <td>
            <span style="color: #CFDEE7; ">Normal.proto</span>
        </td>
    </tr>
    <tr align="center" bgcolor="#92B4F4">
            <td rowspan="2">
                <span style="color: #CFDEE7; ">用户注册</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">0x200</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">客户端</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">请求注册用户，携带：用户基本信息，包括用户名密码</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">Register.proto</span>
            </td>
        </tr>
    <tr align="center" bgcolor="#92B4F4">
            <td>
                <span style="color: #CFDEE7; ">0x201</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">服务端</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">携带：用户id</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">Register.proto</span>
            </td>
        </tr>
    <tr align="center" bgcolor="#92B4F4">
            <td rowspan="4">
                <span style="color: #CFDEE7; ">用户登录</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">0x300</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">客户端</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">请求登录，携带：userid，密码RSA加密</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">User.proto</span>
            </td>
        </tr>
    <tr align="center" bgcolor="#92B4F4">
            <td>
                <span style="color: #CFDEE7; ">0x301</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">服务端</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">登陆成功，携带：Token</span>
            </td>
            <td>
                <span style="color: #CFDEE7; ">User.proto</span>
            </td>
        </tr>
</table>



