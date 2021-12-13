
# Master Data Collaction



## Indices

* [Kerusakan](#kerusakan)

  * [Delete kerusakan](#1-delete-kerusakan)
  * [Get Kerusakan](#2-get-kerusakan)
  * [Get List Kerusakan](#3-get-list-kerusakan)
  * [Insert Kerusakan](#4-insert-kerusakan)
  * [Update data kerusakan](#5-update-data-kerusakan)

* [Teknisi](#teknisi)

  * [Delete Teknisi](#1-delete-teknisi)
  * [Get All Teknisi](#2-get-all-teknisi)
  * [Insert Teknisi](#3-insert-teknisi)
  * [Update Antrian](#4-update-antrian)
  * [Update Teknisi](#5-update-teknisi)


--------


## Kerusakan



### 1. Delete kerusakan



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: localhost:7070/v1/kerusakan/id/1
```



### 2. Get Kerusakan



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:7070/v1/kerusakan/id/2
```

***Response***

```js
{
    "code": "success",
    "message": "success",
    "data": {
        "id": 2,
        "jenis_kerusakan": "Mati Total",
        "lama_pengerjaan": "2 (menit)",
        "harga": 300000,
        "version": 1
    }
}
```


### 3. Get List Kerusakan



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:7070/v1/kerusakan
```


***Response***

```js
{
    "code": "success",
    "message": "success",
    "data": [
        {
            "id": 2,
            "jenis_kerusakan": "Mati Total",
            "lama_pengerjaan": "2 (menit)",
            "harga": 300000,
            "version": 1
        }
    ]
}
```



### 4. Insert Kerusakan



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7070/v1/kerusakan/add
```



***Body:***

```js        
{
    "jenis_kerusakan": "Mati Total",
    "lama_pengerjaan": "2 (menit)",
    "harga": 300000
}
```



### 5. Update data kerusakan



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7070/v1/kerusakan/update
```



***Body:***

```js        
{
    "id": 1,
    "jenis_kerusakan": "Boot loop",
    "lama_pengerjaan": "90 (detik)",
    "harga": 250000,
    "version": 1
}
```



## Teknisi



### 1. Delete Teknisi



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: localhost:7070/v1/teknisi/id/1
```



### 2. Get All Teknisi



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:7070/v1/teknisi
```

***Response***

```js
{
    "code": "success",
    "message": "success",
    "data": [
        {
            "id": 2,
            "full_name": "gunawan",
            "specialist": "samsung",
            "platform": "Android",
            "jumlah_antrian": 2,
            "version": 2
        },
        {
            "id": 3,
            "full_name": "Jhon",
            "specialist": "iphone",
            "platform": "ios",
            "jumlah_antrian": 22,
            "version": 15
        },
        {
            "id": 4,
            "full_name": "Jhon",
            "specialist": "iphone",
            "platform": "ios",
            "jumlah_antrian": 23,
            "version": 25
        }
    ]
}
```


### 3. Insert Teknisi



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7070/v1/teknisi/add
```



***Body:***

```js        
{
    "full_name": "gunawan",
    "specialist": "samsung",
    "platform": "Android"
}
```



### 4. Update Antrian



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7070/v1/teknisi/antrian
```



***Body:***

```js        
{
    "id": 2,
    "jumlah_antrian": 2,
    "version": 1
}
```



### 5. Update Teknisi



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7070/v1/teknisi/update
```



***Body:***

```js        
{
    "id": 1,
    "full_name": "boby",
    "specialist": "xiaomi",
    "platform": "Android",
    "version": 1
}
```

