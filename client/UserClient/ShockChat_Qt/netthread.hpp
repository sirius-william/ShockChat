#ifndef NETTHREAD_HPP
#define NETTHREAD_HPP

#include <QObject>
#include <QTcpSocket>
#include <QThread>
#include "protos/Normal.pb.h"
// 子线程类，用于通信
class NetThread : public QObject
{
    Q_OBJECT
public:
    explicit NetThread(QObject *parent = nullptr);
    // 通信socket
    QTcpSocket* socket;


private:
    // 用户token，登录后获取
    QString token;

signals:
    // 注册结果
    void registerResSignal(bool res);
    // 登录结果
    void loginResSignal(bool res);
    // 向主线程发送一些通知，显示在QMessageBox里
    void sendMsg(QString msg);
    // 发送连接服务器失败信号
    void serverError(QString beacuse);
    // 将0x103消息发给主线程进行解析
    void send0x103ToMain(bool status, QString err);
    // 告诉主线程验证合法性通过了
    void checkSuccessfulSignal();
    // 拿到了token
    void getTokenSignal(QString token);
    // 注册用户时拿到了id
    void getIdSignal(QString id);
    // socket断开
    void socketDisConnectSignal();
public slots:
    // 用于用户注册的槽函数
    void registerSlot(QString username, QString password, QString tel, QString mail);
    // 用于用户登录的槽函数
    void loginSlot(QString id, QString password);
    void init();
    void SocketConnect();
};

#endif // NETTHREAD_HPP
