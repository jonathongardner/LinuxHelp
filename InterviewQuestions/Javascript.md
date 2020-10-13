## Question 1
```javascript
class Pet {
  constructor(name) {
    this.name = name
  }
  printName () {
    console.log(this.name)
  }
}

class Person {
  constructor(name, printPetName) {
    this.name = name
    this.printPetName = printPetName
  }
  printName () {
    console.log(this.name)
  }
  greetings () {
    console.log('My name is:')
    this.printName()
    console.log('My pets name is:')
    this.printPetName()
  }
}
```
```javascript
const pet = new Pet('Pete')
const person = new Person('Jacob', pet.printName)
person.greetings()
/* =>
My name is:
Jacob
My pets name is:
Jacob
*/
```
Whats wrong and how to fix this?
<details>
  <summary>Answer</summary>

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
</details>



## Question 2
```javascript
let createdAt = 'NoDate'
fetch(`https://api.github.com/users/jonathongardner`).then(response => response.json()).then(data => {
  // console.log(data)
  createdAt = data['created_at']
})
console.log(`Created At: ${createdAt}`)
/* =>
Created At: NoDate
*/
```
