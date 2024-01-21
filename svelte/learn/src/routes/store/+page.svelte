<script lang="ts">
	import { count } from './store';
	import Increment from './Increment.svelte';
	import Reset from './Reset.svelte';
</script>

<p>Count: {$count}</p>
<input type="number" placeholder="value of count" bind:value={$count} />
<Increment />
<Reset />
