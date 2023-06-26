# purplestore

## Pkg
- gin   (http fremwork) : github.com/gin-gonic
- sqlx (db interaction) : github.com/jmoiron
- lib pq (db potgres)
- viper 

```
.
├── app.env.sample ( contains application configuration )
├── cmd ( main binary )
├── Makefile ( simplify project commands )
├── docs ( swagger documentation )
├── db
│   └── migrations ( database migrations )
└── internal
    ├── app
    │   ├── controllers ( request response handler )
    │   ├── models ( all about database table )
    │   ├── repository ( database/cache operation )
    │   ├── router ( http router )
    │   ├── schema ( request/response schema )
    │   └── service ( business logic )
    └── pkg ( private lib )
```

```
router --> middleware --> controllers(use schema) --> service --> repository(user model)
```



| Tipe data   | Cakupan bilangan                                |
|-------------|------------------------------------------|
| `uint8`     | 0 ↔ 255           |
| `uint16`    | 0 ↔ 65535  |
| `uint32`    | 0 ↔ 4294967295    |
| `uint32`    | 0 ↔ 4294967295    |
| `uint64`    | 0 ↔ 18446744073709551615    |
| `uint`      | sama dengan uint32 atau uint64 (tergantung nilai)   |
| `byte`    | sama dengan uint8   |
| `int8`    | -128 ↔ 127    |
| `int16`    | -32768 ↔ 32767    |
| `int32`    | -2147483648 ↔ 2147483647    |
| `int64`    | -9223372036854775808 ↔ 9223372036854775807   |
| `int`    | sama dengan int32 atau int64 (tergantung nilai)  |
| `rune`    | sama dengan int32  |


| Tipe data   | Cakupan bilangan                                |
|-------------|------------------------------------------|
| `float32`     | 01.18×10^−38 ↔ 23.4×10^38       |
| `float64`    | 2.23×10^−308 ↔ 1.80×10^308  |
| `complex64`    | complex numbers with float32 real and imaginary parts.    |
| `complex128`   | complex numbers with float64 real and imaginary parts.

