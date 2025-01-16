FROM mongo:6.0

COPY init-mongo.json /init-mongo.json
#El init-mongo.json tiene 12 elementos, de volver a ejecutarlo se agregaran los mismos elementos con diferente ID
#Este script es solo para probar el filtro, podria eliminarse y no afectar la API
CMD mongoimport --host mongodb --username root --password example --authenticationDatabase admin --db productsdb --collection products --type json --file /init-mongo.json --jsonArray
