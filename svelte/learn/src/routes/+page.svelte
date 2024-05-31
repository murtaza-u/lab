<script lang="ts">
	import Counter from './Counter.svelte';

	export let data;
</script>

<h1>{data.title}</h1>
<p>{data.content}</p>

<Counter initial={10} />
