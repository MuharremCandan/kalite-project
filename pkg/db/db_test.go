package database

import (
	"go-backend-test/pkg/config"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestLoadConfigFromYAML(t *testing.T) {
	// Geçici bir YAML dosyası oluştur
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Geçici YAML dosyasına örnek yapılandırmayı yaz
	yamlData := []byte(`
httpserver:
  host: "localhost"
  port: "8080"
database:
  host: "localhost"
  port: "5432"
  user: "testuser"
  pass: "testpass"
  name: "testdb"
`)
	if _, err := tmpfile.Write(yamlData); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// YAML dosyasından yapılandırmayı yükle
	var cfg config.Config
	yamlFile, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		t.Fatalf("YAML dosyasından yapılandırma yüklenirken hata oluştu: %v", err)
	}

	// Yapılandırmayı kontrol et
	expectedHost := "localhost"
	if cfg.HttpServer.Host != expectedHost {
		t.Errorf("Beklenen HTTP sunucu hostu: %s, alınan: %s", expectedHost, cfg.HttpServer.Host)
	}

	expectedPort := "8080"
	if cfg.HttpServer.Port != expectedPort {
		t.Errorf("Beklenen HTTP sunucu portu: %s, alınan: %s", expectedPort, cfg.HttpServer.Port)
	}

	expectedDBHost := "localhost"
	if cfg.Database.Host != expectedDBHost {
		t.Errorf("Beklenen veritabanı hostu: %s, alınan: %s", expectedDBHost, cfg.Database.Host)
	}

	expectedDBPort := "5432"
	if cfg.Database.Port != expectedDBPort {
		t.Errorf("Beklenen veritabanı portu: %s, alınan: %s", expectedDBPort, cfg.Database.Port)
	}

	expectedDBUser := "testuser"
	if cfg.Database.User != expectedDBUser {
		t.Errorf("Beklenen veritabanı kullanıcısı: %s, alınan: %s", expectedDBUser, cfg.Database.User)
	}

	expectedDBPass := "testpass"
	if cfg.Database.Pass != expectedDBPass {
		t.Errorf("Beklenen veritabanı parolası: %s, alınan: %s", expectedDBPass, cfg.Database.Pass)
	}

	expectedDBName := "testdb"
	if cfg.Database.Name != expectedDBName {
		t.Errorf("Beklenen veritabanı adı: %s, alınan: %s", expectedDBName, cfg.Database.Name)
	}
}
