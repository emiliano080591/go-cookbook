# Ejercicios de golang

## General
* Leer y escribir archivos
* Serializacion y deserializacion de estructuras
* Refleccion 
## Storage
* Conexion con mysql
* Conexion con redis 
* Conexion con mongodb
## Concurrency
* Concurrencia(async)
## Test
* Mockgen
## Pipeline architecture
* Context package
* Workers
* Basic pipeline
## Data streams
* Implentación de kafka
  * Sarama(library) sync


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
## Bibliotecas para datastreams
```shell
    $ go get gopkg.in/Shopify/sarama.v1
```
## Para correr con Makefile
```shell
$ make all
```