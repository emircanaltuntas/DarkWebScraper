package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"golang.org/x/net/proxy"
)

// Tor proxy üzerinden HTTP client oluşturma
func createTorClient() *http.Client {
	// Tor Browser genellikle localhost:9150'da çalışır
	torProxy := "127.0.0.1:9150"

	dialer, err := proxy.SOCKS5("tcp", torProxy, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Tor proxy'ye bağlanırken hata oluştu: %v", err)
	}

	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // Zaman aşımı süresi artırıldı
	}

	return client
}

// Masaüstünde "output" klasörü oluşturup dosyayı kaydetme fonksiyonu
func saveToFile(filename, data string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Kullanıcı dizini alınamadı: %v", err)
	}

	// Masaüstü'nde "output" dizini oluşturma
	desktopPath := filepath.Join(usr.HomeDir, "Desktop", "output")
	if _, err := os.Stat(desktopPath); os.IsNotExist(err) {
		err := os.Mkdir(desktopPath, 0755) // "output" klasörünü oluştur
		if err != nil {
			log.Fatalf("output klasörü oluşturulamadı: %v", err)
		}
	}

	// Dosya yolunu masaüstündeki output klasörüne kaydetme
	filePath := filepath.Join(desktopPath, filename)
	err = ioutil.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Dosya yazılamadı: %v", err)
	}

	fmt.Printf("Dosya başarıyla kaydedildi: %s\n", filePath)
}

// Sitenin HTML içeriğini çekme fonksiyonu
func fetchHTML(client *http.Client, url string) string {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("HTTP isteği başarısız oldu: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("HTML okuma hatası: %v", err)
	}

	return string(body)
}

// HTML içinden linkleri çekme fonksiyonu
func extractLinks(html string) []string {
	var links []string
	re := regexp.MustCompile(`href="(http[s]?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		links = append(links, match[1])
	}
	return links
}

// Ekran görüntüsü alma fonksiyonu (Rod kütüphanesi kullanarak)
func takeScreenshot(url, filename string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Kullanıcı dizini alınamadı: %v", err)
	}

	// Masaüstünde "output" dizini oluşturma
	desktopPath := filepath.Join(usr.HomeDir, "Desktop", "output")
	if _, err := os.Stat(desktopPath); os.IsNotExist(err) {
		err := os.Mkdir(desktopPath, 0755) // "output" klasörünü oluştur
		if err != nil {
			log.Fatalf("output klasörü oluşturulamadı: %v", err)
		}
	}

	// Dosya yolunu masaüstündeki output klasörüne kaydetme
	filePath := filepath.Join(desktopPath, filename)

	// Tor proxy kullanarak tarayıcıyı başlatma
	browserLauncher := launcher.New().Proxy("socks5://127.0.0.1:9150").Headless(true).MustLaunch()

	// Rod tarayıcı instance'ı oluşturma
	browser := rod.New().ControlURL(browserLauncher).MustConnect()
	defer browser.MustClose()

	// Sayfayı yükleme ve ekran görüntüsü alma
	page := browser.MustPage(url).MustWaitLoad()
	page.MustScreenshot(filePath)
	fmt.Printf("Ekran görüntüsü başarıyla kaydedildi: %s\n", filePath)
}

func main() {
	// CLI argümanları
	urlFlag := flag.String("url", "", "Bilgileri çekmek istediğiniz web sitesinin URL'si")
	htmlFlag := flag.Bool("html", false, "Sitenin HTML içeriğini çeker")
	linksFlag := flag.Bool("links", false, "Sitedeki linkleri çeker")
	screenshotFlag := flag.Bool("screenshot", false, "Sitenin ekran görüntüsünü alır")
	helpFlag := flag.Bool("h", false, "Yardım menüsünü gösterir")

	flag.Parse()

	if *helpFlag {
		fmt.Println("Kullanım: go run main.go --url=<URL> [--html] [--links] [--screenshot]")
		fmt.Println("Örnek: go run main.go --url=https://örnek.com --html --links --screenshot")
		os.Exit(0)
	}

	// URL argümanı kontrolü
	if *urlFlag == "" {
		fmt.Println("Lütfen bilgilerini çekmek istediğiniz web sitesinin URL'sini girin.")
		os.Exit(1)
	}
	url := *urlFlag

	// Tor proxy ile HTTP client oluşturma
	client := createTorClient()

	// İşlemler
	if *htmlFlag {
		html := fetchHTML(client, url)
		saveToFile("html.txt", html) // Masaüstündeki output klasörüne kaydedecek
		fmt.Println("HTML dosyası başarıyla kaydedildi.")
	}

	if *linksFlag {
		html := fetchHTML(client, url)
		links := extractLinks(html)
		saveToFile("links.txt", fmt.Sprintf("%v", links)) // Masaüstündeki output klasörüne kaydedecek
		fmt.Println("Linkler dosyası başarıyla kaydedildi.")
	}

	if *screenshotFlag {
		takeScreenshot(url, "screenshot.png") // Masaüstündeki output klasörüne kaydedecek
	}
}
