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
