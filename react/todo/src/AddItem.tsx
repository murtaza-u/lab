import { useRef } from 'react'
import { Item } from './Todo'
import './AddItem.css'
import { v4 as uuid } from 'uuid';

type AddItemProps = {
  items: Item[]
  setItems(items: Item[]): void
}

function AddItem({ items, setItems }: AddItemProps) {
  const ref = useRef<HTMLInputElement>(null)

  function handler(e: React.FormEvent) {
    e.preventDefault()
    const txt = ref.current!.value
    ref.current!.value = ""
    const item = {
      id: uuid(),
      text: txt,
      checked: false
    }
    setItems([...items, item])
  }

  return (
    <form id="addItemForm" onSubmit={handler}>
      <input
        type="text"
        placeholder="New Item"
        autoFocus
        required
        ref={ref}
      />
      <button type="submit">Add</button>
    </form>
  )
}

export default AddItem