```javascript
const getGitUser = () => {
  return fetch('https://api.github.com/users/jonathongardner').then(response => response.json()).then(data => {
    return data
  })
}

const user = getGitUser()
console.log(user.login)
/* =>
undefined
*/
```
Example response:
```json
{
  "login": "jonathongardner",
  ...
}
```
Why is it returning `undefined`?
