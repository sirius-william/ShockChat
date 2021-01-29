#ifndef MAINWINDOW_HPP
#define MAINWINDOW_HPP

#include <QMainWindow>
#include "protos/Normal.pb.h"
#include "netthread.hpp"
QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
    bool checked;
    NetThread* netThread;

signals:
    void initNetThreadSignal();
    void initNetThreadSocket();
public slots:
    void get0x103Msg(bool status, QString err);
    void get0x201Msg(QString id);
    void get0x301Msg(QString token);
    void sendMsgSlot(QString err);
    void loginResSlot(bool res);
    void getTokenSlot(QString _token);
    void serverErrorSlot(QString err);
private:
    Ui::MainWindow *ui;
    QString token;
};
#endif // MAINWINDOW_HPP
