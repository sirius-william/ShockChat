#ifndef MESSAGE_HPP
#define MESSAGE_HPP
#include <QByteArray>
#include <QList>

QByteArray packMsg(int id, QByteArray data);
QList<int> unpackMsg(QByteArray head);

#endif // MESSAGE_HPP
