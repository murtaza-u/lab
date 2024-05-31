# Svelte

## Styles

* Styles are defined using the `<style></style>` block and are scoped to
  the **file**.

```svelte
<style>
	h1 {
		color: rebeccapurple;
	}
</style>

<h1>Hello, Svelte</h1>
```

## Reactivity

### Statements that do trigger DOM redraw

* Assignments

```svelte
count = 10;
```

* Updates

```svelte
count ++;
count = count + 2;
```

### Statements that do **not** trigger DOM redraw

```svelte
myArr.push("element")
```

### Example:

```svelte
<script>
	let obj = {
		count: 0,
		todos: []
	}

	// does trigger DOM redraw
	const inc = () => {
		obj.count ++
	}

	// does *not* trigger DOM redraw
	const addTodo = () => {
		obj.todos.push("Soup")
	}
</script>

<p>TODOS: {JSON.stringify(obj.todos)}</p>
<p>Count: {obj.count}</p>
<button on:click={inc}>+</button>
<button on:click={addTodo}>Add Item</button>
```

## Reactive Declaration

Value of double changes with the value of count.

```svelte
<script>
	let count = 0;

	// let double = 0;
	$: double = count * 2

	const inc = () => {
		count ++
		// double = count * 2
	}
</script>

<p>Count: {count} | Double: {double}</p>
<button on:click={inc}>+</button>
```

## Events

Event listeners can be attached to elements by using the `on:` keyword.

```svelte
<button on:click={() => count ++}>INCREMENT</button>
```

Several properties can be chained to event object by using the `|` seperator.

```svelte
<button on:click|preventDefault|stopPropagation={() => count ++}>INCREMENT</button>
```

## Shorthands

```svelte
<script>
	const href = "https://svelte.dev"
</script>

<!-- <a href="{href}">Click Me</a> -->
<!-- <a href={href}>Click Me</a> -->
<a {href}>Click Me</a>
```
