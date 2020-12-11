```javascript
const person = new Person('Jacob', pet.printName.bind(pet))
// OR
class Pet {
  ...
  printName = () => {
    console.log(this.name)
  }
  // BAD too
  // printName = function printName () {
  //   console.log(this.name)
  // }
}
```
