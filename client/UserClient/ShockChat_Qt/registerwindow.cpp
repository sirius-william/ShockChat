#include "registerwindow.hpp"
#include "ui_registerwindow.h"

#include <QMessageBox>

RegisterWindow::RegisterWindow(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::RegisterWindow)
{
    ui->setupUi(this);


}

RegisterWindow::~RegisterWindow()
{
    delete ui;
}


