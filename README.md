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

