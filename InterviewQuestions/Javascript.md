## Question 1
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



## Question 2
```javascript
class LivingThing {
  constructor(name) {
    this.name = name
  }
  printName () {
    console.log(this.name)
  }
  greetings () {
    console.log('My name is:')
    this.printName()
  }
}

class Person extends LivingThing {
  constructor(name, printPetName = null) {
    super(name)
    this.printPetName = printPetName
  }
  greetings () {
    console.log('My name is:')
    this.printName()
    if (this.printPetName) {
      console.log('My pets name is:')
      this.printPetName()
    }
  }
}
```
```javascript
const peteThePet = new LivingThing('Pete')
const person = new Person('Jacob', peteThePet.printName)

peteThePet.greetings()
console.log('----------------')
person.greetings()
/* =>
My name is:
Pete
----------------
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
