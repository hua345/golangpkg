const CryptoJS = require('crypto-js');

/**
 * CipherOption, 加密的一些选项:
 *   mode: 加密模式, 可取值(CBC, CFB, CTR, CTRGladman, OFB, ECB), 都在 CryptoJS.mode 对象下
 *   padding: 填充方式, 可取值(Pkcs7, AnsiX923, Iso10126, Iso97971, ZeroPadding, NoPadding), 都在 CryptoJS.pad 对象下
 *   iv: 偏移量, mode === ECB 时, 不需要 iv
 * 返回的是一个加密对象
 */
function AesEncryptCBC(data, key, iv) { // 加密
    key = CryptoJS.enc.Utf8.parse(key);
    iv = CryptoJS.enc.Utf8.parse(iv);
    let encrypted = CryptoJS.AES.encrypt(data, key,
        {
            iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7
        });
    return encrypted.toString();    // 返回的是base64格式的密文
}

function AesDecryptCBC(encrypted, key, iv) { // 解密
    key = CryptoJS.enc.Utf8.parse(key);
    iv = CryptoJS.enc.Utf8.parse(iv);
    let decrypted = CryptoJS.AES.decrypt(encrypted, key,
        {
            iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7
        });
    return decrypted.toString(CryptoJS.enc.Utf8);      //
}

// BASE64
const base64 = {
    en: (data) => CryptoJS.enc.Base64.stringify(data),
    de: (data) => CryptoJS.enc.Base64.parse(data)
};
// SHA256
const sha256 = (data) => {
    return CryptoJS.SHA256(data).toString();
};
// MD5
const md5 = (data) => {
    return CryptoJS.MD5(data).toString();
};


var _ = process.argv.splice(2);

let data, aesKey = '';

if (_[0] != undefined && _[0].length > 0) {
    data = _[0]
} else {
    data = "fangfang";
}
if (_[1] != undefined && _[1].length > 0) {
    aesKey = _[1]
} else {
    aesKey = 'fangfangfangfang'; // 秘钥 16*n:
}
const iv = aesKey.substr(0, 16);

let encrypt = AesEncryptCBC(data, aesKey, iv)
console.log("js AES加密:",encrypt);
console.log("js AES解密:", AesDecryptCBC(encrypt, aesKey, iv))
