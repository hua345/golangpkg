package encrypt

func PasswordEncrypt(password string) string {
	return EncodeMD5(EncodeSHA256(password))
}
