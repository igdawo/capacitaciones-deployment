db = db.getSiblingDB('veterinaria_citiaps');
db.createUser({
  user: 'domi',
  pwd: '123456',
  roles: [{ role: 'readWrite', db: 'veterinaria_citiaps' }],
});
db.createCollection("Dueno");

db.Dueno.insertMany([{
  "_id": ObjectId("000000012f8895c562734f6e"),
  "Nombre": "María",
  "Edad": 25,
  "Sexo": "Mujer",
  "id_perros": []
},
{
  "_id": ObjectId("000000022f8895c562734f71"),
  "Nombre": "Carmen",
  "Edad": 30,
  "Sexo": "Mujer",
  "id_perros": []
},
{
  "_id": ObjectId("000000032f8895c562734f72"),
  "Nombre": "Rosalía",
  "Edad": 19,
  "Sexo": "Mujer",
  "id_perros": []
},
{
  "_id": ObjectId("000000042f8895c562734f73"),
  "Nombre": "Dominique",
  "Edad": 31,
  "Sexo": "Mujer",
  "id_perros": []
},
{
  "_id": 
    ObjectId("000000052f8895c562734f75"),
  "Nombre": "Bastián",
  "Edad": 27,
  "Sexo": "Hombre",
  "id_perros": []
},
{
  "_id": 
    ObjectId("64e26ade9af995b2748f7334"),
  "nombre": "Dominike",
  "edad": 24,
  "sexo": "Mujer",
  "id_perros": []
}]);

db.createCollection("Perro");

db.createCollection("Vacuna");



