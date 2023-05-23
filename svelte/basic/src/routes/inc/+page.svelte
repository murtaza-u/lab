<script>
  let counter = 0;

  function increment() {
    counter ++;
  }
</script>

<p>Count: {counter}</p>
<button on:click={increment}>Increment</button>
