import { useState } from "react"

function Greeting() {
  const getName = (): string => {
    const names = ["Murtaza", "Bob", "Dave"]
    const idx = Math.floor(Math.random() * names.length)
    return names[idx]
  }

  const [name, setName] = useState(getName())

  const handler = (): void => setName(getName())

  return (
    <>
      <h1>Hello, {name}</h1>
      <button onClick={handler}>Change Me</button>
    </>
  )
}

export default Greeting
