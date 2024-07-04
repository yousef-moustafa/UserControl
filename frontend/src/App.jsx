export default function App() {
  return (
    <form className="new-item-form">
      <div className="form-row">
        <label htmlFor="item">New User</label>
        <input type="text" id="item"/>
      </div>
      <button className="btn">Add User</button>
    </form>
  )
}