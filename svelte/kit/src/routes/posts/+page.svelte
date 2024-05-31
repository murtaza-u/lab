<script lang="ts">
	import type { PageData } from './$types'

	export let data: PageData
</script>

<h2>Archive</h2>
<ul>
	{#each data.posts as post (post.id)}
		<li><a href="/posts/{post.slug}">{post.title}</a></li>
	{/each}
</ul>

<style>
	h2 {
		text-align: center;
	}

	ul {
		margin-left: 1rem;
	}
</style>
