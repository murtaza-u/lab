import { useRef, useState } from "react"
import "./App.css"

function App() {
  const ref = useRef<HTMLInputElement>(null)
  const [color, setColor] = useState("")

  return (
    <div className="container">
      <div
        id="box"
        style={{
          backgroundColor: color
        }}
      ></div>
      <form>
        <input
          type="text"
          ref={ref}
          placeholder="Color"
          onChange={() => setColor(ref.current!.value)}
        />
      </form>
    </div>
  )
}

export default App
