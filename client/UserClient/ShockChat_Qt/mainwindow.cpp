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
    this->netThread = new NetThread();
    QThread *netQThread = new QThread();
    this->netThread->moveToThread(netQThread);
    netQThread->start();
    this->registerWindow = new RegisterWindow();
    this->registerWindow->setModal(true);
    connect(this, &MainWindow::initNetThreadSignal, this->netThread, &NetThread::init);

    connect(this->netThread, &NetThread::sendMsg, this, &MainWindow::sendMsgSlot);
    connect(this->netThread, &NetThread::serverError, this, &MainWindow::serverErrorSlot);
//    connect(this->netThread, &NetThread::getIdSignal, this, &MainWindow::g);
    connect(this->netThread, &NetThread::getTokenSignal, this, &MainWindow::getTokenSlot);
    connect(this->netThread, &NetThread::loginResSignal, this, &MainWindow::loginResSlot);
    connect(this->netThread, &NetThread::send0x103ToMain, this, &MainWindow::get0x103Msg);
    connect(this, &MainWindow::initNetThreadSocket, this->netThread, &NetThread::SocketConnect, Qt::BlockingQueuedConnection);
    connect(this->netThread, &NetThread::checkSuccessfulSignal, this, [=](){
        QMessageBox messageBox;
        messageBox.setText("legal check successfull");
        messageBox.exec();
    });
    connect(this->netThread, &NetThread::socketDisConnectSignal, this, [=](){
        QMessageBox messageBox;
        messageBox.setText("socket disconnected");
        messageBox.exec();
        QApplication* app;
        app->quit();
    });
    connect(this->ui->RegisterBtn, &QPushButton::clicked, [=](){
        this->registerWindow->show();
    });
    connect(this->registerWindow, &RegisterWindow::registerSignal, this->netThread, &NetThread::registerSlot);
    connect(this->netThread, &NetThread::getIdSignal, this->registerWindow, &RegisterWindow::getRegisterStatus);
    connect(this->ui->LoginBtn, &QPushButton::clicked, this, [=](){
        emit wantToLogin(ui->idText->text(), ui->pwdText->text());
    });
    connect(this, &MainWindow::wantToLogin, this->netThread, &NetThread::loginSlot);
    qDebug() << "main conn end";
    emit initNetThreadSocket();
    emit initNetThreadSignal();
}

MainWindow::~MainWindow()
{
    delete ui;
}


void MainWindow::getTokenSlot(QString _token){
    this->token = _token;
    qDebug()<< this->token;
    QMessageBox message;
    message.setText("login successfully");
    message.exec();
}

void MainWindow::serverErrorSlot(QString err)
{
    QMessageBox messageBox;
    messageBox.setText(err);
    messageBox.exec();
    QApplication* app;
    app->quit();

}
void MainWindow::loginResSlot(bool res){
    QMessageBox messageBox;
    if(res){
        messageBox.setText("legal res check true");
        messageBox.exec();
    }else{
        messageBox.setText("legal res check false");
        messageBox.exec();
        QApplication* app;
        app->quit();
    }
}
void MainWindow::sendMsgSlot(QString err){
    QMessageBox messageBox;
    messageBox.setText(err);
    messageBox.exec();
    if(!this->checked){
        QApplication* app;
        app->quit();
    }
}
void MainWindow::get0x103Msg(bool status, QString err)
{
    QMessageBox messageBox;

    messageBox.exec();
    if(status){
        messageBox.setText("success");
        messageBox.exec();
    }else{
        messageBox.setText(err);
        messageBox.exec();
        if(!this->checked){
            QApplication* app;
            app->quit();
        }
    }
}

void MainWindow::get0x201Msg(QString id)
{

}

void MainWindow::get0x301Msg(QString token)
{

}

