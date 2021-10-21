# Ejercicios de golang

* Leer y escribir archivos
* Serializacion y deserializacion de estructuras
* Refleccion 
* Conexion con mysql
* Conexion con redis 
* Conexion con mongodb
* Concurrencia(async)
* Mockgen


## Bibliotecas para los storages
```shell
    $ go get github.com/go-sql-driver/mysql
    $ go get gopkg.in/redis.v5
    $ go get gopkg.in/mgo.v2
```
## Bibliotecas para mockear
```shell
    $ go get github.com/golang/mock/mockgen
    $ go get github.com/golang/mock/gomock
```
## Generar el archivo para mockear
```shell
    $ mockgen -source=archivo_para_mockear.go -destination=archivo_destino.go -package=paquete_archivo_a_mockear
```

## Para correr con Makefile
```shell
$ make all
```