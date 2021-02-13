#include "mainwindow.hpp"
#include "./ui_mainwindow.h"
#include <QApplication>
#include <QDebug>
#include <QMessageBox>
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    this->initGui();
    this->initThread();

}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::initGui()
{
    this->setDisabled(true);

}

void MainWindow::initThread()
{
    // 实例化子线程并启动
    this->netThread = new NetThread();
    QThread * thread = new QThread();
    this->netThread->moveToThread(thread);
    thread->start();
    // 连接信号槽
    connect(this, &MainWindow::initThreadSignal, this->netThread, &NetThread::initNetThread);
    connect(this, &MainWindow::startLegalCheckSignal, this->netThread, &NetThread::startLegalCheckSlot);
    connect(this->netThread, &NetThread::legalCheckResult, this, [=](int status, QString error){
        QMessageBox box;
        box.setText("状态码:" + QString::number(status) + "\t" + error);
        box.exec();
        if(status != 1){
            QMessageBox isTry;
            isTry.setText("是否重试?");
            isTry.setStandardButtons(QMessageBox::Yes | QMessageBox::No);
            isTry.setDefaultButton(QMessageBox::Yes);
            int isYes = isTry.exec();
            switch (isYes) {
                case QMessageBox::No:{
                    QApplication *app;
                    app->exit();
                }
                default:{
                    emit initThreadSignal();
                }
            }
        }else{
            this->setDisabled(false);
            connect(this->netThread, &NetThread::connectBreakSignal, this, [=](){
                QMessageBox isTry;
                isTry.setText("连接断开!");
                isTry.exec();
                QApplication *app;
                app->exit();
            });
        }
    });
    connect(this->netThread, &NetThread::connectSuccessfully, this, [=](){
        emit startLegalCheckSignal();
    });
    connect(this->netThread, &NetThread::userLoginResult, this, &MainWindow::loginResultSlots);
    connect(this->ui->LoginBtn, &QPushButton::clicked, [=](){
        int userid = ui->idText->text().toInt();
        QString password = ui->pwdText->text();
        emit userLogin(userid, password);

    });
    connect(this, &MainWindow::userLogin, this->netThread, &NetThread::userLogin);
    // 初始化子线程，包括连接服务器，验证合法性
    emit initThreadSignal();
}

void MainWindow::loginResultSlots(int status, bool isSuccess, QString error){
    if(status != 0){
        QMessageBox isTry;
        isTry.setText("错误码：" + QString::number(status) + "\t" + error);
        isTry.exec();
        return;
    }
    if(isSuccess){
        this->token = error;
        emit getFriendList();
    }

}



