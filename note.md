# Perbaikan Konfigurasi Entity dan ValueObject

## Masalah yang Ditemukan
1. File valueobject mendefinisikan ulang struct yang sama persis dengan entity
2. Valueobject seharusnya hanya berisi struktur untuk transfer data (JSON) dan tidak perlu menduplikasi definisi struct entity
3. Duplikasi ini menyebabkan konflik dan kesulitan maintenance

## File yang Bermasalah
- `valueobject/alamat.go` - mendefinisikan ulang struct Alamat
- `valueobject/identitas.go` - mendefinisikan ulang struct Identitas  
- `valueobject/lansia.go` - mendefinisikan ulang struct Lansia

## Solusi yang Akan Diterapkan
1. Hapus definisi struct duplikat dari valueobject
2. Gunakan alias atau embedded struct dari entity
3. Tetap simpan payload structs untuk API (Insert, Update, Delete, Response)
4. Gunakan helper function untuk konversi entity ke JSON format

## Struktur yang Benar
```go
// valueobject/alamat.go
package valueobject

import "svc-llt-golang/entity"

// Gunakan alias untuk struct utama
type Alamat = entity.Alamat

// Tetap simpan payload structs untuk API
type AlamatPayloadInsert struct {
    Data []Alamat `json:"data" binding:"required"`
    User string
}
// ... dst
```

## Manfaat Perbaikan
1. Eliminasi duplikasi code
2. Single source of truth untuk struktur data
3. Lebih mudah maintenance
4. Menghindari konflik definisi struct