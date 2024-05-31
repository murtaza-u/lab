<script lang="ts">
	import type { PageData } from './$types'

	export let data: PageData
</script>

<h2>{data.title}</h2>
<p>{data.body}</p>

<style>
	h2 {
		text-align: center;
	}

	p {
		margin: 1rem;
	}
</style>
