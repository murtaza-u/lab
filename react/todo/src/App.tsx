import { useEffect, useState } from "react"
import Todo, { Item } from "./Todo"
import Nav from "./Nav"
import { Route, Routes } from 'react-router-dom'
import About from "./About"
import './App.css'

function fetchItems(): Item[] {
  const key = "items"
  const val = localStorage.getItem(key)
  if (!val) return []
  const item = JSON.parse(val)
  if (!item) return []
  return item
}

function App() {
  const [items, setItems] = useState<Item[]>(fetchItems)

  useEffect(() => {
    localStorage.setItem("items", JSON.stringify(items))
  }, [items])

  return (
    <>
      <Nav />
      <Routes>
        <Route path="/" element={<Todo items={items} setItems={setItems} />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </>
  )
}

export default App
