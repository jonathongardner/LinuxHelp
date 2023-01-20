Why does it show 'No one yet' for a few seconds when I fetch in the `componentWillMount`?
```javascript
export class Test extends Component {
  constructor(props) {
    super(props)
    this.state = {
      me: 'No one yet'
    }
  }
  componentWillMount = async () => {
    const response = await fetch(`https://api.github.com/users/jonathongardner`)
    const data = await response.json()
    await this.setState({ me: data.login })
  }
  render () {
    const { me } = this.state
    return (
      <div>
        { me }
      </div>
    );
  }
}
```
