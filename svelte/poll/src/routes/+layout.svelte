<script lang="ts">
	import '../theme.postcss';
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	import { setModeUserPrefers, setModeCurrent } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';

	onMount(() => {
		setModeCurrent(true)
		setModeUserPrefers(true)
	})
</script>

<slot />
