#include "crypto.h"
using namespace std;
// 加密
std::string Encrypto(const std::string &publicFileName, const std::string &clear_text){

    BIO *in = BIO_new(BIO_s_file());
    BIO_read_filename(in, publicFileName.c_str());
 
    std::string encrypt_text;
	// BIO *keybio = BIO_new_mem_buf((unsigned char *)pub_key.c_str(), -1);
	RSA* rsa = PEM_read_bio_RSA_PUBKEY(in, nullptr, nullptr, nullptr);
	
	// 注意-----第1种格式的公钥
	// 注意-----第2种格式的公钥（这里以第二种格式为例）
	if (rsa == nullptr) {
		unsigned long err = ERR_get_error(); //获取错误号
		char err_msg[1024] = { 0 };
		ERR_error_string(err, err_msg); // 格式：error:errId:库:函数:原因
		printf("err msg: err:%ld, msg:%s\n", err, err_msg);
		return std::string();
	}
 
	// 获取RSA单次可以处理的数据块的最大长度
	int key_len = RSA_size(rsa);
	int block_len = key_len - 11;    // 因为填充方式为RSA_PKCS1_PADDING, 所以要在key_len基础上减去11
 
	// 申请内存：存贮加密后的密文数据
	char *sub_text = new char[key_len + 1];
	memset(sub_text, 0, key_len + 1);
	int ret = 0;
	int pos = 0;
	std::string sub_str;
	// 对数据进行分段加密（返回值是加密后数据的长度）
	while (pos < clear_text.length()) {
		sub_str = clear_text.substr(pos, block_len);
		memset(sub_text, 0, key_len + 1);
		ret = RSA_public_encrypt(sub_str.length(), (const unsigned char*)sub_str.c_str(), (unsigned char*)sub_text, rsa, RSA_PKCS1_PADDING);
		if (ret >= 0) {
			encrypt_text.append(std::string(sub_text, ret));
		}
		pos += block_len;
	}
	
	// 释放内存  
	BIO_free_all(in);
	RSA_free(rsa);
	delete[] sub_text;
 
	return encrypt_text;

};
// 解密
std::string Decrypto(const std::string &privateFileName, const std::string &cipher_text){
    // std::string strRet;
    // std::cout << "BIO *in" << std::endl;
    BIO *in = BIO_new(BIO_s_file());
    BIO_read_filename(in, privateFileName.c_str());
    std::string decrypt_text;
	RSA* rsa;
 
	rsa = PEM_read_bio_RSAPrivateKey(in, NULL, NULL, NULL);
	if (rsa == nullptr) {
		unsigned long err = ERR_get_error(); //获取错误号
		char err_msg[1024] = { 0 };
		ERR_error_string(err, err_msg); // 格式：error:errId:库:函数:原因
		printf("err msg: err:%ld, msg:%s\n", err, err_msg);
		return std::string();
	}
 
	// 获取RSA单次处理的最大长度
	int key_len = RSA_size(rsa);
	char *sub_text = new char[key_len + 1];
	memset(sub_text, 0, key_len + 1);
	int ret = 0;
	std::string sub_str;
	int pos = 0;
	// 对密文进行分段解密
	while (pos < cipher_text.length()) {
		sub_str = cipher_text.substr(pos, key_len);
		memset(sub_text, 0, key_len + 1);
		ret = RSA_private_decrypt(sub_str.length(), (const unsigned char*)sub_str.c_str(), (unsigned char*)sub_text, rsa, RSA_PKCS1_PADDING);
		if (ret >= 0) {
			decrypt_text.append(std::string(sub_text, ret));
			pos += key_len;
		}
	}
	// 释放内存  
	delete[] sub_text;
	BIO_free_all(in);
	RSA_free(rsa);
 
	return decrypt_text;

};
