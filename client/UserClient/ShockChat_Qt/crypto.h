#include <iostream>
#include <cstring>
#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/err.h>
#include <QString>

// 加密
std::string Encrypto(const std::string &publicFileName, const std::string &data);
// 解密
std::string Decrypto(const std::string &privateFileName, const std::string &data);

// 加盐算法
QString Riddle(QString salt);
