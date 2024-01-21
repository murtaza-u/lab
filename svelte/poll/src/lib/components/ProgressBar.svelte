<script lang="ts">
	import { tweened } from 'svelte/motion';

	export let text = '';
	export let value = 0;
	export let maxValue = 0;
	export let color = '';

	const width = tweened(0);

	$: width.set((value / maxValue) * 100 || 0);
</script>

<div id="container">
	<span>{text}</span>
	<div
		id="progress"
		class="rounded-r-full h-14 relative"
		style="width: {$width}%; background-color: {color};"
	/>
</div>

<style>
	#container span {
		margin: 14px 0 0 14px;
		position: absolute;
		display: inline-block;
		z-index: 99;
	}
</style>
