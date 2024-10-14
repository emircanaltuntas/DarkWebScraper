# DarkWebScraper
**Bu repo Yıldız CTI takım içi görev için oluşturduğum bir DarkWebScraper aracıdır.**
**Görev tanımı ilgili dosyada bulunmaktadır.**

__Tool Hakkında:__

- Bu araçta, dark web'den verileri çekerken güvenliğimizi korumak için Tor Browser kullandım.

- Bu aracı kullanmak için, Tor Browser'ın makinenizde yüklü ve açık olması gerekmektedir.

- Tor Browser'a girmeden önce VPN açmanız önerilir.

- Tor Browser proxy ayarları, araçtaki ayarlarla eşleşmelidir: 

  * Tor Browser'da Proxy Ayarlarını Kontrol Etme:
  * Tor Browser'ın Kullanıcı Ayarlarını Kontrol Edin:
    * Tor Browser'ı açın.
    * Adres çubuğuna about:preferences#network yazıp Enter’a basın.
    * Bağlantı Ayarları (Connection Settings) bölümünü bulun ve bu ayarda hangi proxy portunun kullanıldığını kontrol edin. Genellikle Tor Browser, SOCKS v5 proxy için 127.0.0.1 ve 
  farklı bir port (genellikle 9150) kullanır.

- Hazırsanız terminalde şu komutları çalıştırabilirsiniz:

  * go run main.go -h ile help parametresi çalıştırılabilir.
  * Örnek: go run main.go --url=https://örnek.com --html --links --screenshot şu şekilde kullanılabilir.
  * Bu komutlardan herhangi birini yazdığınızda masaüstünüze "output" adında bir dosya oluşturulacak ve çıktılar bu dosyada yer alacak.
- İyi çalışmalar!

-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

# DarkWebScraper
**This repository is a DarkWebScraper tool that I created for an internal task at Yıldız CTI.** 
**The task description is included in the relevant file.**

__About the Tool:__

- In this tool, I used the Tor Browser to ensure our security while retrieving data from the dark web.

- To use this tool, Tor Browser must be installed and running on your machine.

- It is recommended to activate a VPN before opening Tor Browser.

- The Tor Browser's proxy settings must match those in the tool:

  * Checking Proxy Settings in Tor Browser:
  * Check the User Settings in Tor Browser:
    * Open Tor Browser.
    * In the address bar, type about:preferences#network and press Enter.
    * Find the "Connection Settings" section and verify which proxy port is being used. Typically, Tor Browser uses 127.0.0.1 and a different port (usually 9150) for SOCKS v5 proxy.

- When you're ready, you can run the following commands in the terminal:

  * Use go run main.go -h to display the help parameter.
  * Example: go run main.go --url=https://example.com --html --links --screenshot can be used like this.
  * After running any of these commands, a file named "output" will be created on your desktop, and the results will be stored inside it.

- Happy coding!
