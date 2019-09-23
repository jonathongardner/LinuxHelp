# GraphQL
Example query:
```graphql
query {
  images {
    id
    name
  }
}
```
```bash
curl -H "Content-Type: application/json" -X POST -d @graphql_images.json http://localhost:3000/graphql
```
Example query with params:
```graphql
query {
  images(ids: [2]) {
    id
    name
  }
}
```
```bash
curl -H "Content-Type: application/json" -X POST -d @graphql_image.json http://localhost:3000/graphql
```
Example query with params and variables:
```graphql
query($ids: [ID!]) {
  images(ids: $ids) {
    id
    name
  }
}
```
```bash
curl -H "Content-Type: application/json" -X POST -d @graphql_image_v.json http://localhost:3000/graphql
```
```bash
curl -H "Content-Type: application/json" -X POST -d @graphql_images_v.json http://localhost:3000/graphql
```
