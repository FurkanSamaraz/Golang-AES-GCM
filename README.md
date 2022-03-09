# Golang-AES-GCM
Golang ile kendi şifreleme algoritmamı yazmak istedik biraz araştırma yaptım ve aes kütüphanesini buldum https://pkg.go.dev/crypto/aes#NewCipher
ve anahtarları kayıt etmemiz gerekiyor bu sıkıcı ve güvensiz bunun yerine 32 bitlik bir anahtar değişkeni yazdım sonra rastgele(rand) işlemine gönderdim sonrasında ise hexadecimale dönüştürdüm sonuç sürekli değişen bir key ve anahtarıda bende ;)

# şifreleme adımları
* Anahtar (32-bytes için AES-256 şifreleme)
* nonce (rastgele sayı)
* şifrelenecek veri
* aesGCM ile şifrele

# şifre çözme
* Anahtar (32-bytes AES-256)
* nonce 
* şifrelenmiş veriden nonce çıkartma
* aesGCM ile çözümle
şifre çözmede nonce(anahtar tekrarlama riski için kullanıldı) değişkenine dikkat edin doğru nonce olmaz ise şifre çözülemez.

# ++ Ornek
# Dosya Şifreleme
projemiz ile bu fonksiyon ile dosya oluşturup şifreleyebiliriz.

func encryptFile(filename string, data []byte, keyString string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

# Dosya Şifre Çözme
Projemiz ile Şifrelediğimiz dosyamızı bu fonksiyon ile çözebiliriz.

func decryptFile(filename string, keyString string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

