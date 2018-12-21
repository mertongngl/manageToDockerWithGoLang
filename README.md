# **Go ile Docker Yönetimi**
## - Bu linkten go1.10.2 versiyonu için kurulumu takip ediniz <https://golang.org/dl/> .
## - dockerScripts dizinini $GOPATH/src/ yoluna atınız.
## - main.go script'ini $GOPATH yoluna atınız.
## - Scriptlerdeki paketleri " go get <paket_adı> " şeklinde çekebilirsiniz.
## - main.go script'inde satırları sırasıyla uncomment ederek test edebilirsiniz.
## - pullimagewithauth.go script'inde Docker Hub hesap bilgilerini hardcoded düzenleyiniz.
# ~ **Fonksiyonlar**
## - Background: arkada bir konteynır kaldırıp id'sini döndürüyor.
## - CommitContainer: değişiklik yaptığın konteynırı update ediyor.
## - Hello: ayağa kaldırdığı konteynırın terminalinde hello world yazdırıyor.
## - ListAllImages: makinedeki tüm image'leri listeliyor.
## - PrintLogs: konteynırın loglarını döküyor.
## - PullImage: image çekiyor.
## - PullImageWithAuth: Docker Hub hesap bilgileri ile image çekiyor.
## - StopAll: tüm konteynırları durduruyor.