#ifndef NETTHREAD_HPP
#define NETTHREAD_HPP

#include <QMutex>
#include <QObject>
#include <QTcpSocket>
#include <QThread>
#include "protos/Normal.pb.h"
#include "Definations.hpp"
#include "data.hpp"

// 子线程类，用于通信
class NetThread : public QObject
{
    Q_OBJECT
public:
    explicit NetThread(QObject *parent = nullptr);
    ~NetThread();
    QMutex mutex;
private:
    QTcpSocket *socket;
    Data readDataFromSocket();
    QString token;
    int userid;
public slots:
    // 发送者：主线程。请求向服务器验证合法性连接
    void startLegalCheckSlot();
    void initNetThread();
    void userLogin(int userid, QString password);
    void getFriendList();
signals:
    void legalCheckResult(int status, QString error="");
    void userLoginResult(int status, bool isSuccess, QString error="");
    void connectSuccessfully();
    void connectBreakSignal();
};


#endif // NETTHREAD_HPP
