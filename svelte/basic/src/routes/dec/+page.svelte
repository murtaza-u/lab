<script>
  let counter = 0;

  function decrement() {
    counter --;
  }
</script>

<p>Count: {counter}</p>
<button on:click={decrement}>Decrement</button>
