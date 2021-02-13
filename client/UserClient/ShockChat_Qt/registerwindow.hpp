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

private:
    Ui::RegisterWindow *ui;
};

#endif // REGISTERWINDOW_HPP
