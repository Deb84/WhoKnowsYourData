POST /api/value: ValueJson Object -> add a value
POST /api/values: [] ValueJson Object -> add several value
GET /api/values/{label} -> get a [] ValueJson Object of every label node

GET /api/value/{uuid} -> get value
PATCH /api/value/{uuid} -> edit a value
DELETE /api/value/{uuid} - delete a value

POST /api/relation -> add a relation
POST /api/relations -> add several relations
GET /api/relation/{uuid} -> get a relation
PATCH /api/relation/{uuid} -> edit a relation
DELETE /api/relation/{uuid} -> delete a relation
