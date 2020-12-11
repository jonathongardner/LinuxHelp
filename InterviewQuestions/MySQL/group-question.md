Car
```
id - primary key
model_id - foreign_key
color - string
```
Model
```
id - primary key
name - (4Runner, etc.)
make_id - foreign_key (Toyota, Chevy, etc)
type_id - foreign_key (Truck, Car, ...)
```
Is this a valid query?
```sql
SELECT models.name, cars.color FROM models INNER JOIN car on cars.model_id = models.id GROUP BY models.id
```
