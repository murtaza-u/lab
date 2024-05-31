import "./style.css";

let count = 0;

document.getElementById("app")!.innerHTML = `
  <button id="counter">
    Count: ${count}
  </button>
`;

document.getElementById("counter")!.addEventListener("click", () => {
  document.getElementById("counter")!.innerHTML = `Count: ${++count}`;
});