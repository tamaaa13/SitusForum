# SITUS FORUM

Proyek ini merupakan Backend sebuah **Situs Forum** yang menggunakan **Golang**, **GIN**, **MYSQL**, **Docker**

## 🚀 Fitur

- 🔑 Autentikasi menggunakan JWT
- 📝 CRUD **User** dan **Post**
- 📦 Manajemen database menggunakan **migrate**

## 🛠️ Teknologi yang Digunakan

- **Golang**
- **Gin** (Framework)
- **MySQL**
- **Viper** (Konfigurasi)
- **Migrate** (Database Migration)

## 📦 Instalasi dan Menjalankan Proyek

### 1. **Clone repository**

```sh
   git clone https://github.com/tamaaa13/SitusForum.git
```

### 2. **Changes File .env.example To .env**

- (Adjust According To Your Data)

### 3. **Create New File Migration**

```sh
  make migrate-create name=migration_name
```

### 4. **Database Migration**

```sh
     make migrate-up
```

### 5. **Migration Rollback**

```sh
     make migrate-down
```

### 6. **Run Database** (If Use Docker)

```sh
     docker-compose up -d
```

### 7. Run Server

```sh
    go run main.go
```
