package database

import (
	"go-backend-test/pkg/config"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

// TODO: edit these tests
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
  port: "5433"
  user: "kaliteproject"
  pass: "kaliteproject"
  name: "kaliteproject"
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

	pgDb := NewPgDb(&cfg)
	require.NotNil(t, pgDb)

	db, err := pgDb.ConnectDb()

	require.Nil(t, db)
	require.Error(t, err)
	require.Error(t, pgDb.Ping())

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

	expectedDBPort := "5433"
	if cfg.Database.Port != expectedDBPort {
		t.Errorf("Beklenen veritabanı portu: %s, alınan: %s", expectedDBPort, cfg.Database.Port)
	}

	expectedDBUser := "kaliteproject"
	if cfg.Database.User != expectedDBUser {
		t.Errorf("Beklenen veritabanı kullanıcısı: %s, alınan: %s", expectedDBUser, cfg.Database.User)
	}

	expectedDBPass := "kaliteproject"
	if cfg.Database.Pass != expectedDBPass {
		t.Errorf("Beklenen veritabanı parolası: %s, alınan: %s", expectedDBPass, cfg.Database.Pass)
	}

	expectedDBName := "kaliteproject"
	if cfg.Database.Name != expectedDBName {
		t.Errorf("Beklenen veritabanı adı: %s, alınan: %s", expectedDBName, cfg.Database.Name)
	}
}

// func TestNewPgDb(t *testing.T) {

// 	// Geçici bir YAML dosyası oluştur
// 	tmpfile, err := ioutil.TempFile("", "example")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.Remove(tmpfile.Name())

// 	// Geçici YAML dosyasına örnek yapılandırmayı yaz
// 	yamlData := []byte(`
// httpserver:
//   host: "localhost"
//   port: "8080"
// database:
//   host: "localhost"
//   port: "5432"
//   user: "testuser"
//   pass: "testpass"
//   name: "testdb"
// `)
// 	if _, err := tmpfile.Write(yamlData); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err := tmpfile.Close(); err != nil {
// 		t.Fatal(err)
// 	}

// 	// YAML dosyasından yapılandırmayı yükle
// 	var cfg config.Config
// 	yamlFile, err := ioutil.ReadFile(tmpfile.Name())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
// 		t.Fatalf("YAML dosyasından yapılandırma yüklenirken hata oluştu: %v", err)
// 	}

// 	type args struct {
// 		config config.Config
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 		want *PgDb
// 	}{
// 		{
// 			name: "TestCase Success",
// 			args: args{
// 				config: cfg,
// 			},
// 			want: &PgDb{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewPgDb(&tt.args.config); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewPgDb() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPgDb_ConnectDb(t *testing.T) {
// 	type fields struct {
// 		config config.Config
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    *gorm.DB
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &PgDb{
// 				config: &tt.fields.config,
// 			}
// 			got, err := d.ConnectDb()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("PgDb.ConnectDb() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("PgDb.ConnectDb() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAutoMigrate(t *testing.T) {
// 	type args struct {
// 		db *gorm.DB
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := AutoMigrate(tt.args.db); (err != nil) != tt.wantErr {
// 				t.Errorf("AutoMigrate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestPgDb_Ping(t *testing.T) {
// 	type fields struct {
// 		config config.Config
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &PgDb{
// 				config: &tt.fields.config,
// 			}
// 			if err := d.Ping(); (err != nil) != tt.wantErr {
// 				t.Errorf("PgDb.Ping() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
