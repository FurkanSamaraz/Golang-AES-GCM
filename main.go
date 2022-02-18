package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//Güvenliği arttırmak için anahtarı rastgele döngü halinde değiştirmek.
	bytes := make([]byte, 256) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n")
	key := hex.EncodeToString(bytes) //anahtarı bayt cinsinden kodlayın ve gizli olarak saklayın, bir kasaya koyun
	fmt.Printf("key(anahtar) : %s\n", key)

	encrypted := encrypt("Merhaba şifre", key)
	fmt.Printf("encrypted(şifreli) : %s\n", encrypted)

	decrypted := decrypt(encrypted, key)
	fmt.Printf("decrypted(şifre çözüm) : %s\n", decrypted)
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Anahtar hex. olduğundan, kodu bayt'a dönüştürmemiz gerekiyor.
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Anahtarı kullanarak yeni şifre örneği alın.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//GCM oluştur. GCM simetrik anahtar şifreleme blok şifrelemeleri için kullanılır. örnekleri https://golang.hotexamples.com/examples/crypto.cipher/-/NewGCM/golang-newgcm-function-examples.html

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Bir nonce oluşturun. Nonce, GCM'den olmalıdır. anahtar tekrarlama riskine karşılık.
	nonce := make([]byte, aesGCM.NonceSize()) // nonce size yerine byte eklentiside yapılabilir.
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//aesGCM.Seal kullanarak verileri şifreleyin
	//Bu durumda nonce'yi başka bir yere kaydetmek istemediğimiz için onu şifrelenmiş verilere örnek olarak ekliyoruz. Seal'deki ilk nonce argümanı örnektir.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil) //şifrele. örnekler https://golang.hotexamples.com/examples/crypto.cipher/AEAD/Open/golang-aead-open-method-examples.html
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Anahtardan yeni bir Şifre Bloğu oluşturun.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Yeni bir GCM oluşturun.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//nonce boyutunu al.
	nonceSize := aesGCM.NonceSize()

	//Şifrelenmiş verilerden nonce'yi çıkarın.
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Verilerin şifresini çöz
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

/*Golang ile kendi şifreleme algoritmamı yazmak istedik biraz araştırma yaptım ve aes kütüphanesini buldum https://pkg.go.dev/crypto/aes#NewCipher
ve anahtarları kayıt etmemiz gerekiyor bu sıkıcı ve güvensiz bunun yerine 32 bitlik bir anahtar değişkeni yazdım sonra rastgele(rand) işlemine gönderdim sonrasında ise hexadecimale dönüştürdüm sonuç sürekli değişen bir key ve anahtarıda bende ;)
şifreleme adımları
Anahtar (32-bytes için AES-256 şifreleme)
nonce (rastgele sayı)
şifrelenecek veri
şifre çözme
şifre çözmede nonce(anahtar tekrarlama riski için kullanıldı) değişkenine dikkat edin doğru nonce olmaz ise şifre çözülemez.
*/
