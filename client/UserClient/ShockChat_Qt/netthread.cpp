#include "netthread.hpp"
#include "Message.hpp"
#include "protos/token.pb.h"
#include "protos/Register.pb.h"
#include "protos/LegalCheck.pb.h"
#include "crypto.h"
#include "protos/UserLogin.pb.h"
#include <QCryptographicHash>

NetThread::NetThread(QObject *parent) : QObject(parent) {

}

NetThread::~NetThread()
{
    if(this->socket->isValid()) this->socket->disconnectFromHost();
    delete this->socket;
    this->socket = nullptr;
}

Data NetThread::readDataFromSocket()
{
    QMutexLocker locker(&this->mutex);
    QByteArray head = this->socket->read(8);
    QList<int> head_i = unpackMsg(head);
    QByteArray data = this->socket->read(head_i[1]);
    Data dataRes = {head_i[0], head_i[1], data};
    return dataRes;
}

void NetThread::startLegalCheckSlot()
{
    if(!this->socket->isValid()){
        emit legalCheckResult(ERROR_SOCKET, "连接不可用");
        if(this->socket != nullptr) {
            delete this->socket;
        }
        this->socket = nullptr;
        return;
    }
    this->socket->write(packMsg(0x100));
    if(!this->socket->waitForBytesWritten(TIME_OUT)){
        emit legalCheckResult(ERROR_SOCKET, "写超时");
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    if(!this->socket->waitForReadyRead(TIME_OUT)){
        emit legalCheckResult(ERROR_SOCKET, "超时");
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    Data saltProtoReceiveFromServer = this->readDataFromSocket();

    if(saltProtoReceiveFromServer.id != 0x101){
        emit legalCheckResult(ERROR_CLIENT, "读到的id错误, 应为0x101, 但为" + QString::number(saltProtoReceiveFromServer.id));
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
    }
    protos::LegalCheckSalt salt;
    salt.ParsePartialFromArray(saltProtoReceiveFromServer.data, saltProtoReceiveFromServer.size);
    QString saltMd5 = Riddle(QString(salt.salt().c_str()));
    protos::LegalCheckResult result;
    result.set_result(saltMd5.toStdString());
    QByteArray saltResultSendToServer = QByteArray(result.ByteSizeLong(), 0);
    result.SerializePartialToArray(saltResultSendToServer.data(), result.ByteSizeLong());
    this->socket->write(packMsg(0x102, saltResultSendToServer));
    if(!this->socket->waitForBytesWritten(TIME_OUT)){
        emit legalCheckResult(ERROR_SOCKET, "发送结果失败");
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    if(!this->socket->waitForReadyRead(TIME_OUT)){
        emit legalCheckResult(ERROR_SOCKET, "读取验证超时");
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    Data saltResultStatusFromServer = this->readDataFromSocket();
    if(saltResultStatusFromServer.id != 0x103){
        emit legalCheckResult(ERROR_CLIENT, "消息id应为0x103, 但却是" + QString::number(saltResultStatusFromServer.id));
        this->socket->disconnectFromHost();
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    protos::LegalCheckStatus status;
    status.ParseFromArray(saltResultStatusFromServer.data, saltResultStatusFromServer.size);
    emit legalCheckResult(status.status(), status.error().c_str());
}

void NetThread::initNetThread()
{
    this->socket = new QTcpSocket();
    this->socket->connectToHost(HOST, PORT);
    if(!this->socket->waitForConnected(TIME_OUT)){
        // 连接超时
        emit legalCheckResult(ERROR_SOCKET, "连接超时");
        delete this->socket;
        this->socket = nullptr;
        return;
    }
    emit connectSuccessfully();

}

void NetThread::userLogin(int userid, QString password)
{
    protos::UserLogin userLogin;
    userLogin.set_id(userid);
    std::string p = Encrypto(PUBLIC_KEY, password.toStdString().c_str());
    userLogin.set_password(p);
    QByteArray userLoginToSend = QByteArray(userLogin.ByteSizeLong(), 0);
    userLogin.SerializePartialToArray(userLoginToSend.data(), userLogin.ByteSizeLong());
    this->socket->write(packMsg(0x300, userLoginToSend));
    if(!this->socket->waitForBytesWritten(TIME_OUT)){
        emit userLoginResult(ERROR_SOCKET, false, "发送消息失败");
        return;
    }
    if(!this->socket->waitForReadyRead(TIME_OUT)){
        emit userLoginResult(ERROR_SOCKET, false, "读超时");
        return;
    }
    Data loginStatusFromServer = this->readDataFromSocket();
    if(loginStatusFromServer.id != 0x301){
        emit userLoginResult(ERROR_SOCKET, false, "读id错误，应为0x301，但为" + QString::number(loginStatusFromServer.id));
        return;
    }
    protos::LoginResult loginResult;
    loginResult.ParseFromArray(loginStatusFromServer.data, loginStatusFromServer.size);
    emit userLoginResult(loginResult.status(), loginResult.issuccess(), loginResult.error().c_str());
}

void NetThread::getFriendList()
{

}


