<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';

	export let initial = 0;

	beforeUpdate(() => {
		console.log('before update');
	});

	afterUpdate(() => {
		console.log('after update');
	});
</script>

<p>Count: {initial}</p>
<button on:click={() => initial++}>+</button>
<button on:click={() => initial--}>-</button>
