#ifndef MAINWINDOW_HPP
#define MAINWINDOW_HPP

#include <QMainWindow>
#include "protos/Normal.pb.h"
#include "netthread.hpp"
#include "registerwindow.hpp"
QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
    void initGui();
    void initThread();
private:
    Ui::MainWindow *ui;
    NetThread *netThread;
signals:
    void startLegalCheckSignal();
    void initThreadSignal();
    void userLogin(int userid, QString password);
    void getFriendList();
public slots:
    void loginResultSlots(int status, bool isSuccess = true, QString error="");
};
#endif // MAINWINDOW_HPP
