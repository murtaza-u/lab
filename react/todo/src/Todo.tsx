import AddItem from './AddItem'
import './Todo.css'

export type Item = {
  id: string
  checked: boolean
  text: string
}

type TodoProps = {
  items: Item[]
  setItems(items: Item[]): void
}

function Todo({ items, setItems }: TodoProps) {
  return (
    <main>
      <AddItem items={items} setItems={setItems} />
      <TodoList items={items} setItems={setItems} />
    </main>
  )
}

function TodoList({ items, setItems }: TodoProps) {
  const onChangeHandler = (id: string): void => {
    const newItems = items.map(item => {
      if (item.id === id) {
        return { ...item, checked: !item.checked }
      }
      return item
    })
    setItems(newItems);
  }

  const deleteHandler = (id: string): void => {
    const newItems = items.filter(item => item.id !== id)
    setItems(newItems);
  }

  return (
    <ul id="todolist">
      {items.map(item => {
        return <TodoItem
          key={item.id}
          item={item}
          onChangeHandler={onChangeHandler}
          deleteHandler={deleteHandler}
        />
      })}
    </ul>
  )
}

type TodoItemProps = {
  item: Item
  onChangeHandler(id: string): void
  deleteHandler(id: string): void
}

function TodoItem({ item, onChangeHandler, deleteHandler }: TodoItemProps) {
  return (
    <li className="todoitem">
      <button
        onClick={() => deleteHandler(item.id)}
      >Delete</button>
      <input
        type="checkbox"
        checked={item.checked}
        onChange={() => onChangeHandler(item.id)}
      />
      <span
        style={{ textDecoration: item.checked ? "line-through" : "none" }}
      >{item.text}</span>
    </li>
  )
}

export default Todo
