#include "netthread.hpp"
#include "protos/User.pb.h"
#include "Message.hpp"
#include "protos/token.pb.h"
#include "protos/Register.pb.h"
#include <QCryptographicHash>
NetThread::NetThread(QObject *parent) : QObject(parent) {

}

void NetThread::init()
{
    // 发送请求合法性验证的消息0x100
    QByteArray requestSaltCheck = packMsg(0x100, QByteArray());
    protos::LegalCheckSalt saltProto;
    QString saltMd5;
    this->socket->write(requestSaltCheck);
    if(!this->socket->waitForBytesWritten(TIME_OUT)){
        emit serverError(QString("can not connect to host") + QString(HOST) + QString(":") + QString::number(PORT));
        qDebug() << QString("can not connect to host") + QString(HOST) + QString(":") + QString::number(PORT);
        this->socket->disconnectFromHost();
        return;
    }
    if(this->socket->waitForReadyRead(TIME_OUT)){
        QList<int> dataInfo = unpackMsg(this->socket->read(8));
        int id = dataInfo[1];
        int length = dataInfo[0];
        qDebug() << length;
        QByteArray data = this->socket->read(length);
        if(dataInfo[1] != 0x101){
            emit serverError(QString("message id incorrect, need 0x101, but ") + QString::number(dataInfo[1]));
            qDebug() << QString("message id incorrect, need 0x101, but ") + QString::number(dataInfo[1]);
            this->socket->disconnectFromHost();
            return;
        }
        qDebug() << data;
        saltProto.ParseFromArray(data, data.size());
    }
    // 计算md5
    qDebug() << saltProto.salt().c_str();
    QCryptographicHash md(QCryptographicHash::Md5);
    saltMd5 = QCryptographicHash::hash (QString(saltProto.salt().c_str()).toLatin1(), QCryptographicHash::Md5).toHex();
    qDebug() << saltMd5;
    // 序列化计算结果的proto
    protos::LegalCheckResult saltResultProto;
    saltResultProto.set_result(saltMd5.toStdString());
    QByteArray resultSend = QByteArray(saltResultProto.ByteSizeLong(), 0);
    saltResultProto.SerializePartialToArray(resultSend.data(), saltResultProto.ByteSizeLong());
    // 发送
    socket->write(packMsg(0x102, resultSend));
    if(!this->socket->waitForBytesWritten(TIME_OUT)){
        emit serverError(QString("send result error"));
        qDebug() << QString("send result error");
        this->socket->disconnectFromHost();
        return;
    }
    if(this->socket->waitForReadyRead(TIME_OUT)){
        QList<int> dataInfo = unpackMsg(this->socket->read(8));
        int id = dataInfo[1];
        int length = dataInfo[0];
        QByteArray data = this->socket->read(length);
        if(dataInfo[1] != 0x103){
            emit serverError(QString("message id incorrect, need 0x101, but ") + QString::number(dataInfo[1]));
            qDebug() << QString("message id incorrect, need 0x101, but ") + QString::number(dataInfo[1]);
            this->socket->disconnectFromHost();
            return;
        }
        protos::Status status;
        status.ParseFromArray(data, data.size());
        qDebug() << "status:" << status.status();
        qDebug() << "error:" << status.error().c_str();
        if(status.status() == true){
            emit checkSuccessfulSignal();
        }else{
            qDebug() << status.error().c_str();
            emit sendMsg(QString(status.error().c_str()));
            qDebug()<< QString(status.error().c_str());
            this->socket->disconnectFromHost();
            return;
        }

    }
    // 验证通过后方建立其他消息的信号槽。
    connect(this->socket, &::QAbstractSocket::readyRead, this, [=](){
        QList<int> dataInfo = unpackMsg(this->socket->read(8));
        QByteArray data = this->socket->read(dataInfo[0]);
        switch (dataInfo[1]) {
        case 0x103:{
            protos::Status status;
            status.ParseFromArray(data, data.size());
            emit send0x103ToMain(status.status(), status.error().c_str());
        }
        case 0x301:{
            protos::Token tokenProto;
            tokenProto.ParseFromArray(data, data.size());
            this->token = QString(tokenProto.token().c_str());
            emit getTokenSignal(this->token);
        }
        case 0x201:{
            protos::UserId id;
            id.ParseFromArray(data, data.size());
            emit getIdSignal(QString::number(id.id()));
        }
        }
    });
}

void NetThread::SocketConnect(){
    this->socket = new QTcpSocket();
    this->socket->connectToHost(HOST, PORT);
    if(!this->socket->waitForConnected(TIME_OUT)){
        emit serverError(QString("can not connect to host") + QString(HOST) + QString(":") + QString::number(PORT));
//        this->socket->disconnectFromHost();
        return;
    }
    connect(this->socket, &QAbstractSocket::disconnected, [=](){
        emit socketDisConnectSignal();
    });
}

void NetThread::registerSlot(QString username, QString password, QString tel, QString mail)
{

}

void NetThread::loginSlot(QString id, QString password)
{

}
