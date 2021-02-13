#ifndef MESSAGE_HPP
#define MESSAGE_HPP
#include <QByteArray>
#include <QList>
#include "Message.hpp"
QByteArray packMsg(int id, QByteArray data = nullptr){
    QByteArray ba;
    ba.resize(8);
    int length = 0;
    if(data != nullptr){
        length = data.size();
    }

    ba[0] = (uchar)(0x000000ff &length);
    ba[1] = (uchar)((0x0000ff00&data.size())>>8);
    ba[2] = (uchar)((0x00ff0000&data.size())>>16);
    ba[3] = (uchar)((0xff000000&data.size())>>24);

    ba[4] = (uchar)(0x000000ff &id);
    ba[5] = (uchar)((0x0000ff00&id)>>8);
    ba[6] = (uchar)((0x00ff0000&id)>>16);
    ba[7] = (uchar)((0xff000000&id)>>24);
    if(data != nullptr) ba.append(data);
    return ba;
}

QList<int> unpackMsg(QByteArray head){
    int length = head[0] & 0x000000FF;
    length |= ((head[1] << 8) & 0x0000FF00);
    length |= ((head[2] << 16) & 0x00FF0000);
    length |= ((head[3] << 24) & 0xFF000000);
    int id = head[4] & 0x000000FF;
    id |= ((head[5] << 8) & 0x0000FF00);
    id |= ((head[6] << 16) & 0x00FF0000);
    id |= ((head[7] << 24) & 0xFF000000);
    QList<int> res;
    res << id << length;
    return res;
}
#endif // MESSAGE_HPP
