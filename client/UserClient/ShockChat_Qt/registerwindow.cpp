#include "registerwindow.hpp"
#include "ui_registerwindow.h"

#include <QMessageBox>

RegisterWindow::RegisterWindow(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::RegisterWindow)
{
    ui->setupUi(this);
    connect(ui->registerBtn, &QPushButton::clicked, [=](){
        emit registerSignal(ui->usernameText->text(), ui->pwdText->text(), ui->telText->text(), ui->mailText->text());
    });

}

RegisterWindow::~RegisterWindow()
{
    delete ui;
}

void RegisterWindow::getRegisterStatus(bool status, QString errorOrUserId)
{
    QMessageBox messageBox;
    if(status){
        messageBox.setText("register successfully, please remember your id:" + errorOrUserId);
    }else{
        messageBox.setText("register unsuccessfully because:" + errorOrUserId);
    }
    messageBox.exec();
    ui->pwdText->clear();
    ui->usernameText->clear();
    ui->telText->clear();
    ui->mailText->clear();
    this->close();
}
