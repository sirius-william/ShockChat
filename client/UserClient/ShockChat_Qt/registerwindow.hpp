#ifndef REGISTERWINDOW_HPP
#define REGISTERWINDOW_HPP

#include <QDialog>

namespace Ui {
class RegisterWindow;
}

class RegisterWindow : public QDialog
{
    Q_OBJECT

public:
    explicit RegisterWindow(QWidget *parent = nullptr);
    ~RegisterWindow();
signals:
    void registerSignal(QString username, QString password, QString tel, QString email);
public slots:
    void getRegisterStatus(bool status, QString errorOrUserId);
private:
    Ui::RegisterWindow *ui;
};

#endif // REGISTERWINDOW_HPP
